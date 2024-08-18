package repository

import (
	"context"
	"fmt"
	"time"

	models "github.com/DBrange/didis-comp-bk/cmd/api/models/options/tournament"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/pot/dao"
	tournament_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/tournament/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) CreateTournament(
	ctx context.Context,
	tournamentInfoDAO *tournament_dao.CreateTournamentDAOReq,
	locationID string,
	options *models.OrganizeTournamentOptions,
	categoryID *string,
	organizerID string,
) (string, error) {
	locationOID, err := r.ConvertToObjectID(locationID)
	if err != nil {
		return "", err
	}

	tournamentInfoDAO.LocationID = *locationOID

	organizerOID, err := r.ConvertToObjectID(organizerID)
	if err != nil {
		return "", err
	}

	tournamentInfoDAO.OrganizerID = *organizerOID

	if categoryID != nil {
		categoryOID, err := r.ConvertToObjectID(*categoryID)
		if err != nil {
			return "", err
		}

		tournamentInfoDAO.CategoryID = categoryOID
	}

	tournamentInfoDAO.Rounds = []primitive.ObjectID{}
	tournamentInfoDAO.Matches = []primitive.ObjectID{}
	tournamentInfoDAO.Pots = []primitive.ObjectID{}
	tournamentInfoDAO.Groups = []primitive.ObjectID{}

	tournamentInfoDAO.SetTimeStamp()

	result, err := r.tournamentColl.InsertOne(ctx, tournamentInfoDAO)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error tournament scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}
		return "", fmt.Errorf("error when inserting tournament: %w", err)
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *Repository) UpdateTournamentInfo(ctx context.Context, tournamentOID *primitive.ObjectID, tournamentDAO *tournament_dao.UpdateTournamentInfoDAOReq) error {
	filter := bson.M{"_id": tournamentOID}

	tournamentDAO.RenewUpdate()

	result, err := r.tournamentColl.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": tournamentDAO},
	)
	if err != nil {
		return fmt.Errorf("error updating tournament: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no tournament found with id: %s", customerrors.ErrNotFound, tournamentOID.Hex())
	}

	return nil
}

func (r *Repository) UpdateTournamentFinishDate(ctx context.Context, tournamentOID *primitive.ObjectID) error {
	filter := bson.M{"_id": tournamentOID}

	currentDate := time.Now().UTC()

	update := bson.M{
		"$set": bson.M{"finish_date": currentDate, "updated_at": currentDate},
	}

	result, err := r.tournamentColl.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return fmt.Errorf("error updating tournament: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no tournament found with id: %s", customerrors.ErrNotFound, tournamentOID.Hex())
	}

	return nil
}

// func (r *Repository) RemoveTournamentOptionsBsonStruct(ctx context.Context, tournamentDAO *tournament_dao.UpdateTournamentOptionsDAOReq, update *bson.M) *bson.M {
// 	setUpdates := bson.M{}

// 	if tournamentDAO.Pots != nil {
// 		setUpdates["pots"] = bson.M{"$in": tournamentDAO.Pots}
// 	}

// 	if tournamentDAO.Groups != nil {
// 		setUpdates["groups"] = bson.M{"$in": tournamentDAO.Groups}
// 	}

// 	if len(setUpdates) > 0 {
// 		(*update)["$pull"] = setUpdates
// 	}

// 	return update
// }

// func (r *Repository) AddTournamentOptionsBsonStruct(ctx context.Context, tournamentDAO *tournament_dao.UpdateTournamentOptionsDAOReq, update *bson.M) *bson.M {
// 	setUpdates := bson.M{}

// 	if tournamentDAO.Pots != nil {
// 		setUpdates["pots"] = bson.M{"$each": tournamentDAO.Pots}
// 	}

// 	if tournamentDAO.Groups != nil {
// 		setUpdates["groups"] = bson.M{"$each": tournamentDAO.Groups}
// 	}

// 	if len(setUpdates) > 0 {
// 		(*update)["$push"] = setUpdates
// 	}

// 	if tournamentDAO.DoubleEliminationID != nil {
// 		(*update)["$set"] = bson.M{"double_elimination_id": tournamentDAO.DoubleEliminationID}
// 	}

// 	return update

// }

func (r *Repository) UpdateTournamentRelationsBsonStruct(ctx context.Context, tournamentDAO *tournament_dao.UpdateTournamentOptionsDAOReq, update *bson.M, add bool) *bson.M {
	operation := "$pull"
	arrayModifier := "$in"
	if add {
		operation = "$push"
		arrayModifier = "$each"
	}

	setUpdates := bson.M{}
	if tournamentDAO.Pots != nil {
		setUpdates["pots"] = bson.M{arrayModifier: tournamentDAO.Pots}
	}

	if tournamentDAO.Groups != nil {
		setUpdates["groups"] = bson.M{arrayModifier: tournamentDAO.Groups}
	}

	if tournamentDAO.Matches != nil {
		setUpdates["matches"] = bson.M{arrayModifier: tournamentDAO.Matches}
	}

	if tournamentDAO.Rounds != nil {
		setUpdates["rounds"] = bson.M{arrayModifier: tournamentDAO.Rounds}
	}

	if len(setUpdates) > 0 {
		(*update)[operation] = setUpdates
	}

	// Manejar DoubleEliminationID si es necesario
	if add && tournamentDAO.DoubleEliminationID != nil {
		(*update)["$set"] = bson.M{"double_elimination_id": tournamentDAO.DoubleEliminationID}
	}

	return update
}

func (r *Repository) UpdateTournamentRelations(ctx context.Context, tournamentOID *primitive.ObjectID, tournamentDAO *tournament_dao.UpdateTournamentOptionsDAOReq, add bool) error {
	tournamentDAO.RenewUpdate()

	filter := bson.M{"_id": tournamentOID}

	update := bson.M{}

	update = *r.UpdateTournamentRelationsBsonStruct(ctx, tournamentDAO, &update, add)

	if update == nil {
		return fmt.Errorf("error updating tournament, nothing to update: %w", nil)
	}

	result, err := r.tournamentColl.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return fmt.Errorf("error updating tournament: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no tournament found with id: %s", customerrors.ErrNotFound, tournamentOID.Hex())
	}

	return nil
}

func (r *Repository) AddCategoryInTournament(ctx context.Context, tournamentID string, categoryID string) error {
	tournamentOID, err := r.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	categoryOID, err := r.ConvertToObjectID(categoryID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": *tournamentOID}

	update := bson.M{"category_id": *categoryOID}

	currentDate := time.Now().UTC()
	update["updated_at"] = currentDate

	result, err := r.tournamentColl.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": update},
	)
	if err != nil {
		return fmt.Errorf("error updating tournament: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no tournament found with id: %s", customerrors.ErrNotFound, categoryID)
	}

	return nil
}
func (r *Repository) AddMatchInTournament(ctx context.Context, tournamentOID, matchOID *primitive.ObjectID) error {

	filter := bson.M{"_id": tournamentOID}

	update := bson.M{
		"$push": bson.M{"matches": matchOID},
		"$set":  bson.M{"updated_at": time.Now().UTC()},
	}

	result, err := r.tournamentColl.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return fmt.Errorf("error updating tournament: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no tournament found with id: %s", customerrors.ErrNotFound, matchOID.Hex())
	}

	return nil
}

func (r *Repository) AddMultipleMatchesInTournament(ctx context.Context, tournamentOID *primitive.ObjectID, matchOIDs []*primitive.ObjectID) error {

	var operations []mongo.WriteModel

	filter := bson.M{"_id": tournamentOID}
	update := bson.M{
		"$set": bson.M{
			"matches":    matchOIDs,
			"updated_at": time.Now().UTC(),
		},
	}

	operation := mongo.NewUpdateOneModel().
		SetFilter(filter).
		SetUpdate(update).
		SetUpsert(false)

	operations = append(operations, operation)

	_, err := r.tournamentGroupColl.BulkWrite(ctx, operations, options.BulkWrite().SetOrdered(false))
	if err != nil {
		return fmt.Errorf("error updating tournament groups: %v", err)
	}

	return nil
}

func (r *Repository) VerifyTournamentsExists(ctx context.Context, tournamentOID *primitive.ObjectID) error {
	var result struct{}

	filter := bson.M{"_id": tournamentOID}

	opts := options.FindOne().SetProjection(bson.M{"_id": 1})

	err := r.tournamentColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: error when searching for tournament: %s", customerrors.ErrNotFound, err.Error())
		}
		return fmt.Errorf("error when searching for the tournament: %w", err)
	}

	return nil
}

func (r *Repository) VerifyTournamentsCapacity(ctx context.Context, tournamentOID *primitive.ObjectID) (bool, error) {

	var result struct {
		TotalCompetitors int `bson:"total_competitors"`
		MaxCapacity      int `bson:"max_capacity"`
	}

	filter := bson.M{"_id": tournamentOID}

	err := r.tournamentColl.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil // O manejar este caso como prefieras
		}
		return false, err
	}
	fmt.Printf("totales %d y maximo %d", result.TotalCompetitors, result.MaxCapacity)
	available := result.TotalCompetitors < result.MaxCapacity

	return available, nil

}

func (r *Repository) IncrementTotalCompetitorsInTournament(ctx context.Context, tournamentOID *primitive.ObjectID) error {
	filter := bson.M{"_id": tournamentOID}

	update := bson.M{
		"$inc": bson.M{
			"total_competitors": 1,
		},
	}

	result, err := r.tournamentColl.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return fmt.Errorf("error updating category: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no tournament found with id: %s", customerrors.ErrNotFound, tournamentOID.Hex())
	}

	return nil
}

func (r *Repository) GetTournamentInfoToFinaliseIt(ctx context.Context, tournamentOID *primitive.ObjectID) (*tournament_dao.GetTournamentInfoToFinaliseItDAORes, error) {
	var tournamentInfo *tournament_dao.GetTournamentInfoToFinaliseItDAORes

	filter := bson.M{"_id": tournamentOID}

	opts := options.FindOne().SetProjection(bson.M{"category_id": 1, "total_prize": 1, "points": 1})

	err := r.tournamentColl.FindOne(ctx, filter, opts).Decode(&tournamentInfo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for tournament: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the tournament: %w", err)
	}

	return tournamentInfo, nil
}

func (r *Repository) VerifyTournamentsAlreadyFinished(ctx context.Context, tournamentOID *primitive.ObjectID) error {
	var result struct {
		FinishDate *time.Time `bson:"finish_date"`
	}

	filter := bson.M{"_id": tournamentOID}

	opts := options.FindOne().SetProjection(bson.M{"finish_date": 1})

	err := r.tournamentColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: error when searching for tournament: %s", customerrors.ErrNotFound, err.Error())
		}
		return fmt.Errorf("error when searching for the tournament: %w", err)
	}

	// Verifica si el campo finish_date tiene un valor
	if result.FinishDate != nil {
		return fmt.Errorf("tournament has already finished")
	}

	return nil
}

func (r *Repository) VerifyMultipleGroupsInTournament(ctx context.Context, tournamentOID *primitive.ObjectID, groupOIDs []*primitive.ObjectID) error {
	if len(groupOIDs) == 0 {
		return nil // No hay grupos para verificar
	}
	fmt.Printf("el id es %v", groupOIDs)
	// Ajustar el filtro para que se asegure de que todos los groupOIDs estén en el array groups
	filter := bson.M{
		"_id":    tournamentOID,
		"groups": bson.M{"$all": groupOIDs},
	}

	opts := options.FindOne().SetProjection(bson.M{"groups": 1})

	// Buscar el documento que coincida con el filtro
	var result struct {
		Groups []*primitive.ObjectID `bson:"groups"`
	}
	err := r.tournamentColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: error when searching for tournament: %s", customerrors.ErrNotFound, err.Error())
		}
		return fmt.Errorf("error when searching for the tournament: %w", err)
	}

	// Verificar si se encontraron todos los groupOIDs
	if len(result.Groups) != len(groupOIDs) {
		missingIDs := r.getMissingIDs(groupOIDs, result.Groups)
		return fmt.Errorf("%w: los siguientes grupos no fueron encontrados en el torneo: %v", customerrors.ErrNotFound, missingIDs)
	}

	return nil
}

func (r *Repository) VerifyMultipleTournamentPot(ctx context.Context, tournamentOID *primitive.ObjectID, potOIDs []*primitive.ObjectID) error {
	if len(potOIDs) == 0 {
		return nil // No hay grupos para verificar
	}

	// Ajustar el filtro para que se asegure de que todos los groupOIDs estén en el array groups
	filter := bson.M{
		"_id":  tournamentOID,
		"pots": bson.M{"$all": potOIDs},
	}

	opts := options.FindOne().SetProjection(bson.M{"pots": 1})

	// Buscar el documento que coincida con el filtro
	var result struct {
		Pots []*primitive.ObjectID `bson:"pots"`
	}
	err := r.tournamentColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: error when searching for tournament: %s", customerrors.ErrNotFound, err.Error())
		}
		return fmt.Errorf("error when searching for the tournament: %w", err)
	}

	// Verificar si se encontraron todos los groupOIDs
	if len(result.Pots) != len(potOIDs) {
		missingIDs := r.getMissingIDs(potOIDs, result.Pots)
		return fmt.Errorf("%w: los siguientes grupos no fueron encontrados en el torneo: %v", customerrors.ErrNotFound, missingIDs)
	}

	return nil
}

func (r *Repository) VerifyTournamentGroupInTournament(ctx context.Context, tournamentOID *primitive.ObjectID, groupOIDs []*primitive.ObjectID) error {
	if len(groupOIDs) == 0 {
		return nil // No hay grupos para verificar
	}

	// Ajustar el filtro para que se asegure de que al menos uno de los groupOIDs esté en el array groups
	filter := bson.M{
		"_id":    tournamentOID,
		"groups": bson.M{"$in": groupOIDs},
	}

	opts := options.FindOne().SetProjection(bson.M{"groups": 1})

	// Buscar el documento que coincida con el filtro
	var result struct {
		Groups []*primitive.ObjectID `bson:"groups"`
	}
	err := r.tournamentColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: error when searching for tournament: %s", customerrors.ErrNotFound, err.Error())
		}
		return fmt.Errorf("error when searching for the tournament: %w", err)
	}

	// Si llegaste aquí, significa que al menos uno de los groupOIDs está en el documento
	return nil
}

func (r *Repository) RemoveMultipleTournamentMatches(ctx context.Context, tournamentOID *primitive.ObjectID, matchesToRemoveOIDs []*primitive.ObjectID) error {
	filter := bson.M{"_id": tournamentOID}

	update := bson.M{
		"$pull": bson.M{
			"matches": bson.M{
				"$in": matchesToRemoveOIDs,
			},
		},
	}

	result, err := r.tournamentColl.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error updating tournament: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no tournament found with id: %s", customerrors.ErrNotFound, tournamentOID.Hex())
	}

	return nil

}

func (r *Repository) VerifyTournamentPot(ctx context.Context, tournamentOID, potOID *primitive.ObjectID) error {
	potOIDSlice := []*primitive.ObjectID{potOID}

	filter := bson.M{
		"_id":  tournamentOID,
		"pots": bson.M{"$in": potOIDSlice},
	}

	opts := options.FindOne().SetProjection(bson.M{"pots": 1})

	// Buscar el documento que coincida con el filtro
	var result struct {
		Pots []*primitive.ObjectID `bson:"pots"`
	}
	err := r.tournamentColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: error when searching for tournament: %s", customerrors.ErrNotFound, err.Error())
		}
		return fmt.Errorf("error when searching for the tournament: %w", err)
	}

	// Si llegaste aquí, significa que al menos uno de los groupOIDs está en el documento
	return nil
}

func (r *Repository) AddPotInTournament(ctx context.Context, tournamentOID, potOID *primitive.ObjectID) error {
	filter := bson.M{"_id": tournamentOID}

	update := bson.M{"$push": bson.M{"pots": potOID}}

	result, err := r.tournamentColl.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return fmt.Errorf("error updating tournament: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no tournament found with id: %s", customerrors.ErrNotFound, tournamentOID.Hex())
	}

	return nil
}

func (r *Repository) AddGroupInTournament(ctx context.Context, tournamentOID, groupOID *primitive.ObjectID) error {
	filter := bson.M{"_id": tournamentOID}

	update := bson.M{"$push": bson.M{"groups": groupOID}}

	result, err := r.tournamentColl.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return fmt.Errorf("error updating tournament: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no tournament found with id: %s", customerrors.ErrNotFound, tournamentOID.Hex())
	}

	return nil
}

func (r *Repository) GetTournamentPotPositions(ctx context.Context, tournamentOID *primitive.ObjectID) ([]*dao.PotOrGroupPositionDAORes, error) {
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"_id": tournamentOID}}},
		bson.D{{Key: "$unwind", Value: "$pots"}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "pots",
			"localField":   "pots",
			"foreignField": "_id",
			"as":           "pot",
		}}},
		bson.D{{Key: "$unwind", Value: "$pot"}},
		bson.D{{Key: "$project", Value: bson.M{
			"_id":      "$pot._id",
			"position": "$pot.position",
		}}},
	}

	cursor, err := r.tournamentColl.Aggregate(ctx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for tournament: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the tournament: %w", err)
	}
	defer cursor.Close(ctx)

	var pots []*dao.PotOrGroupPositionDAORes

	if err := cursor.All(ctx, &pots); err != nil {
		return nil, fmt.Errorf("error when decoding tournament: %w", err)
	}

	return pots, nil
}

func (r *Repository) GetTournamentGroupPositions(ctx context.Context, tournamentOID *primitive.ObjectID) ([]*dao.PotOrGroupPositionDAORes, error) {
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"_id": tournamentOID}}},
		bson.D{{Key: "$unwind", Value: "$groups"}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "tournament_groups",
			"localField":   "groups",
			"foreignField": "_id",
			"as":           "group",
		}}},
		bson.D{{Key: "$unwind", Value: "$group"}},
		bson.D{{Key: "$project", Value: bson.M{
			"_id":      "$group._id",
			"position": "$group.position",
		}}},
	}

	cursor, err := r.tournamentColl.Aggregate(ctx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for tournament: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the tournament: %w", err)
	}
	defer cursor.Close(ctx)

	var groups []*dao.PotOrGroupPositionDAORes

	if err := cursor.All(ctx, &groups); err != nil {
		return nil, fmt.Errorf("error when decoding tournament: %w", err)
	}

fmt.Printf("estas son las posiciones %v", groups)
	return groups, nil
}

func (r *Repository) SetPotsInTournament(ctx context.Context, tournamentOID *primitive.ObjectID, potsIDs []*primitive.ObjectID) error {
	var operations []mongo.WriteModel

	filter := bson.M{"_id": tournamentOID}
	update := bson.M{
		"$set": bson.M{
			"competitors": potsIDs,
		},
	}

	operation := mongo.NewUpdateOneModel().
		SetFilter(filter).
		SetUpdate(update).
		SetUpsert(false)

	operations = append(operations, operation)

	_, err := r.tournamentColl.BulkWrite(ctx, operations, options.BulkWrite().SetOrdered(false))
	if err != nil {
		return fmt.Errorf("error updating tournament pots: %v", err)
	}

	return nil
}

func (r *Repository) RemovePotToTournament(ctx context.Context, tournamentOID *primitive.ObjectID, position int) error {
	// Primero buscamos el pot que tiene la posición específica
	filterToFind := bson.M{"position": position, "tournament_id": tournamentOID}

	var pot struct {
		ID primitive.ObjectID `bson:"_id"`
	}

	err := r.potColl.FindOne(ctx, filterToFind).Decode(&pot)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: no pot found with position: %d", customerrors.ErrNotFound, position)
		}
		return fmt.Errorf("error finding pot by position: %w", err)
	}

	// Ahora removemos el pot con ese ID del array de pots en el tournament
	filterToUpdate := bson.M{"_id": tournamentOID}

	update := bson.M{
		"$pull": bson.M{
			"pots": pot.ID,
		},
	}

	result, err := r.tournamentColl.UpdateOne(ctx, filterToUpdate, update)
	if err != nil {
		return fmt.Errorf("error removing pot from tournament: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no tournament found with id: %s", customerrors.ErrNotFound, tournamentOID.Hex())
	}

	return nil
}
func (r *Repository) RemoveGroupToTournament(ctx context.Context, tournamentOID *primitive.ObjectID, position int) error {
	// Primero buscamos el pot que tiene la posición específica
	filterToFind := bson.M{"position": position, "tournament_id": tournamentOID}

	var group struct {
		ID primitive.ObjectID `bson:"_id"`
	}

	err := r.tournamentGroupColl.FindOne(ctx, filterToFind).Decode(&group)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: no group found with position: %d", customerrors.ErrNotFound, position)
		}
		return fmt.Errorf("error finding group by position: %w", err)
	}

	// Ahora removemos el group con ese ID del array de groups en el tournament
	filterToUpdate := bson.M{"_id": tournamentOID}

	update := bson.M{
		"$pull": bson.M{
			"groups": group.ID,
		},
	}

	result, err := r.tournamentColl.UpdateOne(ctx, filterToUpdate, update)
	if err != nil {
		return fmt.Errorf("error removing group from tournament: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no tournament found with id: %s", customerrors.ErrNotFound, tournamentOID.Hex())
	}

	return nil
}

func (r *Repository) VerifyNumberPotsInTournament(ctx context.Context, tournamentOID *primitive.ObjectID, position int) error {
	// Primero obtenemos el número de pots en el torneo.
	var tournament struct {
		Pots []primitive.ObjectID `bson:"pots"`
	}

	filter := bson.M{"_id": tournamentOID}

	projection := bson.M{"pots": 1}

	opts := options.FindOne().SetProjection(projection)

	err := r.tournamentColl.FindOne(ctx, filter, opts).Decode(&tournament)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: no tournament found with id: %s", customerrors.ErrNotFound, tournamentOID.Hex())
		}
		return fmt.Errorf("error when retrieving tournament: %w", err)
	}

	// Obtenemos el número de pots.
	numPots := len(tournament.Pots)

	// Verificamos si la posición es válida.
	if position <= numPots {
		return fmt.Errorf("invalid position: %d is lesser than or equal to the number of pots in the tournament (%d)", position, numPots)
	}

	// Si la posición es válida, no retornamos ningún error.
	return nil
}

func (r *Repository) VerifyNumberGroupsInTournament(ctx context.Context, tournamentOID *primitive.ObjectID, position int) error {
	// Primero obtenemos el número de groups en el torneo.
	var tournament struct {
		Groups []primitive.ObjectID `bson:"groups"`
	}

	filter := bson.M{"_id": tournamentOID}

	projection := bson.M{"groups": 1}

	opts := options.FindOne().SetProjection(projection)

	err := r.tournamentColl.FindOne(ctx, filter, opts).Decode(&tournament)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: no tournament found with id: %s", customerrors.ErrNotFound, tournamentOID.Hex())
		}
		return fmt.Errorf("error when retrieving tournament: %w", err)
	}

	// Obtenemos el número de groups.
	numGroups := len(tournament.Groups)

	// Verificamos si la posición es válida.
	if position <= numGroups {
		return fmt.Errorf("invalid position: %d is lesser than or equal to the number of pots in the tournament (%d)", position, numGroups)
	}

	// Si la posición es válida, no retornamos ningún error.
	return nil
}

// func (r *Repository) AddCompetitorInTournament(ctx context.Context, tournamentID, competitorID string) ( error) {
// 	tournamentOID, err := r.ConvertToObjectID(tournamentID)
// 	if err != nil {
// 		return  err
// 	}

// 	competitorOID, err := r.ConvertToObjectID(competitorID)
// 	if err != nil {
// 		return  err
// 	}

// 	filter := bson.M{"_id": tournamentOID}

// 	update := bson.M{
// 		"&push": bson.M{
// 			"competitors": bson.M{"$each": competitorOID},
// 		},
// 	}

// 	result, err := r.tournamentColl.UpdateOne(
// 		ctx,
// 		filter,
// 		update,
// 	)
// 	if err != nil {
// 		return fmt.Errorf("error updating tournament: %w", err)
// 	}

// 	if result.MatchedCount == 0 {
// 		return fmt.Errorf("%w: no tournament found with id: %s", customerrors.ErrNotFound, tournamentID)
// 	}

// 	return  nil
// }
