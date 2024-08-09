package repository

import (
	"context"
	"fmt"
	"time"

	models "github.com/DBrange/didis-comp-bk/cmd/api/models/options/tournament"
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

func (r *Repository) UpdateTournamentInfo(ctx context.Context, tournamentID string, tournamentDAO *tournament_dao.UpdateTournamentInfoDAOReq) (string, error) {
	tournamentOID, err := r.ConvertToObjectID(tournamentID)
	if err != nil {
		return "", err
	}

	filter := bson.M{"_id": tournamentOID}

	projection := bson.M{"location_id": 1}

	opts := options.FindOneAndUpdate().SetProjection(projection)

	var updatedDocument struct {
		LocationID string `bson:"location_id"`
	}

	err = r.tournamentColl.FindOneAndUpdate(
		ctx,
		filter,
		bson.M{"$set": tournamentDAO},
		opts,
	).Decode(&updatedDocument)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", fmt.Errorf("%w: error when searching for the user: %s", customerrors.ErrNotFound, err.Error())
		}
		return "", fmt.Errorf("error when searching for the user: %w", err)
	}

	return updatedDocument.LocationID, nil
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

func (r *Repository) VerifyTournamentsExists(ctx context.Context, tournamentOID *primitive.ObjectID) error {
	var result struct{}
fmt.Printf("asd %v",tournamentOID)
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
