package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	tournament_models "github.com/DBrange/didis-comp-bk/cmd/api/models/options/tournament"
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
	options *tournament_models.OrganizeTournamentOptions,
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
	tournamentInfoDAO.Availability = tournament_dao.TournamentAvailabilityDAO{}

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

func (r *Repository) UpdateTournamentStartDate(ctx context.Context, tournamentOID *primitive.ObjectID) error {
	filter := bson.M{"_id": tournamentOID}

	currentDate := time.Now().UTC()

	update := bson.M{
		"$set": bson.M{"start_date": currentDate, "updated_at": currentDate},
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

func (r *Repository) VerifyTournamentExists(ctx context.Context, tournamentOID *primitive.ObjectID) error {
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

func (r *Repository) DecrementTotalCompetitorsInTournament(ctx context.Context, tournamentOID *primitive.ObjectID) error {
	filter := bson.M{"_id": tournamentOID}

	update := bson.M{
		"$inc": bson.M{
			"total_competitors": -1,
		},
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

func (r *Repository) GetTournamentTotalCompetitors(ctx context.Context, tournamentOID *primitive.ObjectID) (int, error) {
	var result struct {
		TotalCompetitors int `bson:"total_competitors"`
	}

	filter := bson.M{"_id": tournamentOID}

	opts := options.FindOne().SetProjection(bson.M{"total_competitors": 1})

	err := r.tournamentColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, fmt.Errorf("%w: error when searching for tournament: %s", customerrors.ErrNotFound, err.Error())
		}
		return 0, fmt.Errorf("error when searching for the tournament: %w", err)
	}

	return result.TotalCompetitors, nil
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
func (r *Repository) GetTournamentGroupsIDs(ctx context.Context, tournamentOID *primitive.ObjectID) ([]*primitive.ObjectID, error) {

	filter := bson.M{
		"_id": tournamentOID,
	}

	opts := options.FindOne().SetProjection(bson.M{"groups": 1})

	// Buscar el documento que coincida con el filtro
	var result struct {
		Groups []*primitive.ObjectID `bson:"groups"`
	}
	err := r.tournamentColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for tournament: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the tournament: %w", err)
	}

	// Si llegaste aquí, significa que al menos uno de los groupOIDs está en el documento
	return result.Groups, nil
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

func (r *Repository) GetDoubleElimID(ctx context.Context, tournamentOID *primitive.ObjectID) (string, error) {
	var result struct {
		DoubleElim *primitive.ObjectID `bson:"double_elimination_id"`
	}

	filter := bson.M{"_id": tournamentOID}

	projection := bson.M{"double_elimination_id": 1}

	opts := options.FindOne().SetProjection(projection)

	err := r.tournamentColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", nil // O manejar este caso como prefieras
		}
		return "", err
	}

	return result.DoubleElim.Hex(), nil
}

func (r *Repository) GetRoundID(ctx context.Context, tournamentOID *primitive.ObjectID, roundName models.ROUND) (string, error) {
	var result struct {
		ID primitive.ObjectID `bson:"_id"`
	}

	// Pipeline de agregación
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"_id": tournamentOID}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "rounds",
			"localField":   "rounds",
			"foreignField": "_id",
			"as":           "rounds",
		}}},
		bson.D{{Key: "$unwind", Value: "$rounds"}},
		bson.D{{Key: "$match", Value: bson.M{"rounds.round": roundName}}},
		bson.D{{Key: "$project", Value: bson.M{"_id": "$rounds._id"}}},
	}

	// Ejecutamos la agregación
	cursor, err := r.tournamentColl.Aggregate(ctx, pipeline)
	if err != nil {
		return "", fmt.Errorf("error when searching for the round: %w", err)
	}
	defer cursor.Close(ctx)

	// Verificamos si el cursor tiene documentos
	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return "", fmt.Errorf("error when decoding round ID: %w", err)
		}
	} else {
		return "", fmt.Errorf("%w: no round found with name: %s", customerrors.ErrNotFound, roundName)
	}

	// Convertimos el ObjectID a string
	return result.ID.Hex(), nil
}

func (r *Repository) GetTournamentRoundNames(ctx context.Context, tournamentOID *primitive.ObjectID) ([]models.ROUND, error) {
	pipeline := mongo.Pipeline{
		// Filtramos el documento del torneo por su ID
		bson.D{{Key: "$match", Value: bson.M{"_id": tournamentOID}}},
		// Realizamos el lookup para obtener los detalles de las rondas
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "rounds",
			"localField":   "rounds",
			"foreignField": "_id",
			"as":           "rounds",
		}}},
		// Desenrollamos el array "rounds" para procesar cada elemento
		bson.D{{Key: "$unwind", Value: "$rounds"}},
		// Agrupamos los nombres de las rondas en un array
		bson.D{{Key: "$group", Value: bson.M{
			"_id":    nil,
			"rounds": bson.M{"$addToSet": "$rounds.round"},
		}}},
		// Proyectamos el array de nombres de rondas
		bson.D{{Key: "$project", Value: bson.M{
			"_id":    0,
			"rounds": 1,
		}}},
	}

	cursor, err := r.tournamentColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error when aggregating round names: %w", err)
	}
	defer cursor.Close(ctx)

	var result struct {
		Rounds []models.ROUND `bson:"rounds"`
	}

	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("error when decoding round names: %w", err)
		}

		return result.Rounds, nil
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return nil, fmt.Errorf("no rounds found for tournament ID: %s", tournamentOID.Hex())
}

func (r *Repository) GetCompetitorChampion(ctx context.Context, tournamentOID *primitive.ObjectID) (string, error) {
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"_id": tournamentOID}}},
		bson.D{{Key: "$unwind", Value: "$rounds"}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "rounds",
			"localField":   "rounds",
			"foreignField": "_id",
			"as":           "round",
		}}},
		bson.D{{Key: "$unwind", Value: "$round"}},
		bson.D{{Key: "$match", Value: bson.M{"round.round": "F"}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "matches",
			"localField":   "round._id",
			"foreignField": "round_id",
			"as":           "match",
		}}},
		bson.D{{Key: "$unwind", Value: "$match"}},
		bson.D{{Key: "$project", Value: bson.M{"_id": "$match.winner"}}},
	}

	cursor, err := r.tournamentColl.Aggregate(ctx, pipeline)
	if err != nil {
		return "", fmt.Errorf("error when aggregating match: %w", err)
	}
	defer cursor.Close(ctx)

	var competitor struct {
		ID *primitive.ObjectID `bson:"_id"`
	}

	if cursor.Next(ctx) {
		if err := cursor.Decode(&competitor); err != nil {
			return "", fmt.Errorf("error when decoding round ID: %w", err)
		}

		return competitor.ID.Hex(), nil
	}

	if err := cursor.Err(); err != nil {
		return "", fmt.Errorf("cursor error: %w", err)
	}

	return "", fmt.Errorf("%w: no round found with match winner", customerrors.ErrNotFound)
}

func (r *Repository) GetTournamentAvailavility(ctx context.Context, tournamentOID *primitive.ObjectID) (*tournament_dao.TournamentAvailabilityDAO, error) {
	var result struct {
		Avilability tournament_dao.TournamentAvailabilityDAO `bson:"availability"`
	}

	filter := bson.M{"_id": tournamentOID}

	opts := options.FindOne().SetProjection(bson.M{"availability": 1})

	err := r.tournamentColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for tournament: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the tournament: %w", err)
	}

	return &result.Avilability, nil

}

func (r *Repository) UpdateTournamentAvailability(
	ctx context.Context,
	tournamentOID *primitive.ObjectID,
	availableCourts int,
	averageHours int,
) error {
	// Definir el filtro para ubicar el documento del torneo
	filter := bson.M{"_id": tournamentOID}

	// Crear el conjunto de campos a actualizar
	update := bson.M{
		"$set": bson.M{
			"availability.available_courts": availableCourts,
			"availability.average_hours":    averageHours,
		},
	}

	// Ejecutar la actualización en MongoDB
	_, err := r.tournamentColl.UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: error when updating tournament availability: %s", customerrors.ErrNotFound, err.Error())
		}
		return fmt.Errorf("error when updating tournament availability: %w", err)
	}

	return nil
}

func (r *Repository) GetAllDatesMatchesFromTournament(ctx context.Context, tournamentOID *primitive.ObjectID) ([]time.Time, error) {
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"_id": tournamentOID}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "matches",
			"localField":   "_id",
			"foreignField": "tournament_id",
			"as":           "match",
		}}},
		bson.D{{Key: "$unwind", Value: "$match"}},
		bson.D{{Key: "$match", Value: bson.M{"match.date": bson.M{"$ne": nil}}}}, // Exclude null dates
		bson.D{{Key: "$group", Value: bson.M{
			"_id":   nil,
			"dates": bson.M{"$addToSet": "$match.date"},
		}}},
		bson.D{{Key: "$project", Value: bson.M{
			"_id":   0,
			"dates": 1,
		}}},
	}

	cursor, err := r.tournamentColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error when aggregating match dates: %w", err)
	}
	defer cursor.Close(ctx)

	var result struct {
		Dates []time.Time `bson:"dates"`
	}

	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("error when decoding match dates: %w", err)
		}

		// Return the dates if found
		return result.Dates, nil
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	// If no dates found, return an empty array instead of an error
	return []time.Time{}, nil
}

func (r *Repository) GetTournamentPrimaryInfo(ctx context.Context, tournamentOID *primitive.ObjectID) (*tournament_dao.GetTournamentPrimaryInfoDAORes, error) {
	// Define el pipeline de agregación
	pipeline := mongo.Pipeline{
		// Filtro inicial para encontrar el torneo por su ID
		bson.D{
			{Key: "$match", Value: bson.M{"_id": tournamentOID}},
		},
		// Lookup para obtener la ubicación (location)
		bson.D{
			{Key: "$lookup", Value: bson.M{
				"from":         "locations",
				"localField":   "location_id",
				"foreignField": "_id",
				"as":           "location",
			}},
		},
		// Unwind de location (permitiendo valores nulos)
		bson.D{
			{Key: "$unwind", Value: bson.M{
				"path":                       "$location",
				"preserveNullAndEmptyArrays": true,
			}},
		},
		// Lookup para obtener la ronda (category)
		bson.D{
			{Key: "$lookup", Value: bson.M{
				"from":         "rounds",
				"localField":   "rounds",
				"foreignField": "_id",
				"as":           "rounds",
			}},
		},
		// Unwind de round (permitiendo valores nulos)
		bson.D{
			{Key: "$unwind", Value: bson.M{
				"path":                       "$round",
				"preserveNullAndEmptyArrays": true,
			}},
		},
		// Lookup para obtener la categoría (category)
		bson.D{
			{Key: "$lookup", Value: bson.M{
				"from":         "categories",
				"localField":   "category_id",
				"foreignField": "_id",
				"as":           "category",
			}},
		},
		// Unwind de category (permitiendo valores nulos)
		bson.D{
			{Key: "$unwind", Value: bson.M{
				"path":                       "$category",
				"preserveNullAndEmptyArrays": true,
			}},
		},
		// Lookup para obtener el organizador (organizer)
		bson.D{
			{Key: "$lookup", Value: bson.M{
				"from":         "organizers",
				"localField":   "organizer_id",
				"foreignField": "_id",
				"as":           "organizer",
			}},
		},
		// Unwind de organizer (permitiendo valores nulos)
		bson.D{
			{Key: "$unwind", Value: bson.M{
				"path":                       "$organizer",
				"preserveNullAndEmptyArrays": true,
			}},
		},
		// Lookup para obtener el usuario asociado al organizador
		bson.D{
			{Key: "$lookup", Value: bson.M{
				"from":         "users",
				"localField":   "organizer.user_id",
				"foreignField": "_id",
				"as":           "user",
			}},
		},
		// Unwind de user (permitiendo valores nulos)
		bson.D{
			{Key: "$unwind", Value: bson.M{
				"path":                       "$user",
				"preserveNullAndEmptyArrays": true,
			}},
		},
		// Proyección de los campos que quieres devolver
		bson.D{
			{Key: "$project", Value: bson.M{
				"_id":               "$_id",
				"name":              "$name",
				"finish_date":       "$finish_date",
				"start_date":        "$start_date",
				"points":            "$points",
				"image":             "$image",
				"total_prize":       "$total_prize",
				"total_competitors": "$total_competitors",
				"max_capacity":      "$max_capacity",
				"average_score":     "$average_score",
				"genre":             "$genre",
				"sport":             "$sport",
				"competitor_type":   "$competitor_type",
				"surface":           "$surface",
				"availability": bson.M{
					"available_courts": "$availability.available_courts",
					"average_hours":    "$availability.average_hours",
				},
				"rounds": bson.M{
					"$map": bson.M{
						"input": "$rounds",
						"as":    "round",
						"in": bson.M{
							"_id":   "$$round._id",
							"round": "$$round.round",
						},
					},
				},
				"location": bson.M{
					"_id":     "$location._id",
					"state":   "$location.state",
					"country": "$location.country",
					"city":    "$location.city",
					"lat":     "$location.lat",
					"long":    "$location.long",
				},
				"organizer": bson.M{
					"_id":        "$user._id",
					"first_name": "$user.first_name",
					"last_name":  "$user.last_name",
				},
				"category": bson.M{
					"$cond": bson.M{
						"if":   bson.M{"$eq": bson.A{"$category_id", nil}},
						"then": nil,
						"else": bson.M{
							"_id":  "$category._id",
							"name": "$category.name",
						},
					},
				},
			}},
		},
	}

	// Ejecuta el pipeline de agregación
	cursor, err := r.tournamentColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error during aggregation: %w", err)
	}
	defer cursor.Close(ctx)

	// Decode el primer resultado
	var result tournament_dao.GetTournamentPrimaryInfoDAORes
	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("error decoding cursor result: %w", err)
		}
	}

	// Maneja errores de cursor
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return &result, nil
}

func (r *Repository) GetCategoryIDOfTournament(ctx context.Context, tournamentOID *primitive.ObjectID) (*primitive.ObjectID, error) {
	filter := bson.M{"_id": tournamentOID}

	opts := options.FindOne().SetProjection(bson.M{"category_id": 1})

	var result struct {
		CategoryID *primitive.ObjectID `bson:"category_id"`
	}

	err := r.tournamentColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Retorna nil sin un error si no se encuentra el torneo
			return nil, nil
		}
		return nil, fmt.Errorf("error when searching for the tournament: %w", err)
	}

	return result.CategoryID, nil
}

func (r *Repository) GetTournamentFilters(ctx context.Context, tournamentOID *primitive.ObjectID) (*tournament_dao.GetTournamentFiltersDAORes, error) {
	filter := bson.M{"_id": tournamentOID}

	opts := options.FindOne().SetProjection(bson.M{"surface": 1, "sport": 1, "competitor_type": 1, "max_capacity": 1, "total_competitors": 1, "category_id": 1})

	var result tournament_dao.GetTournamentFiltersDAORes

	err := r.tournamentColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for tournament: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the tournament: %w", err)
	}

	return &result, nil
}
func (r *Repository) GetTournamentMatchesByID(ctx context.Context, tournamentOID *primitive.ObjectID) ([]*primitive.ObjectID, error) {
	filter := bson.M{"_id": tournamentOID}

	opts := options.FindOne().SetProjection(bson.M{"matches": 1})

	var result struct {
		Matches []*primitive.ObjectID `bson:"matches"`
	}

	err := r.tournamentColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for tournament: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the tournament: %w", err)
	}

	return result.Matches, nil
}

func (r *Repository) GetTournamentCompetitorIDsInMatches(ctx context.Context, tournamentOID *primitive.ObjectID) ([]string, error) {

	// Definimos el pipeline de agregación.
	pipeline := mongo.Pipeline{
		// Filtro para el torneo (excluyendo los documentos con deleted_at)
		bson.D{{Key: "$match", Value: bson.M{"_id": tournamentOID, "deleted_at": bson.M{"$exists": false}}}},
		// Realizamos el lookup para obtener los match de "competitor_matches"
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "competitor_matches",
			"localField":   "matches",
			"foreignField": "match_id",
			"as":           "competitor_match",
		}}},
		// Hacemos el unwind para obtener los documentos de la colección "competitor_matches"
		bson.D{{Key: "$unwind", Value: bson.M{"path": "$competitor_match"}}},
		// Filtro para excluir los competidores con ID nulo
		bson.D{{Key: "$match", Value: bson.M{
			"competitor_match.competitor_id": bson.M{"$ne": nil},
		}}},
		// Agrupamos por competitor_id para eliminar duplicados
		bson.D{{Key: "$group", Value: bson.M{
			"_id": "$competitor_match.competitor_id",
		}}},
		// Proyectamos solo el campo "competitor_id"
		bson.D{{Key: "$project", Value: bson.M{
			"competitor_id": "$_id",
		}}},
	}

	// Ejecutamos el query en la colección de torneos
	cursor, err := r.tournamentColl.Aggregate(ctx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: no documents found for tournamentRegistration", customerrors.ErrNotFound)
		}
		return nil, fmt.Errorf("error when searching for tournamentRegistration: %w", err)
	}
	defer cursor.Close(ctx)

	// Extraemos los IDs de los documentos encontrados
	var competitorIDs []string
	for cursor.Next(ctx) {
		var doc struct {
			CompetitorID primitive.ObjectID `bson:"competitor_id"`
		}
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
		competitorIDs = append(competitorIDs, doc.CompetitorID.Hex())
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Devolvemos los IDs de los competidores encontrados
	return competitorIDs, nil
}

