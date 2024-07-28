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
	leagueID *string,
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

	if leagueID != nil {
		leagueOID, err := r.ConvertToObjectID(*leagueID)
		if err != nil {
			return "", err
		}

		tournamentInfoDAO.LeagueID = leagueOID
	}

	tournamentInfoDAO.Rounds = []primitive.ObjectID{}
	tournamentInfoDAO.Matches = []primitive.ObjectID{}
	tournamentInfoDAO.Pots = []primitive.ObjectID{}
	tournamentInfoDAO.Groups = []primitive.ObjectID{}

	tournamentInfoDAO.SetTimeStamp()

	fmt.Printf("este es torneo: %v", tournamentInfoDAO)

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

func (r *Repository) RemoveTournamentOptionsBsonStruct(ctx context.Context, tournamentDAO *tournament_dao.UpdateTournamentOptionsDAOReq, update *bson.M) *bson.M {
	setUpdates := bson.M{}

	if tournamentDAO.Pots != nil {
		setUpdates["pots"] = bson.M{"$in": tournamentDAO.Pots}
	}

	if tournamentDAO.Groups != nil {
		setUpdates["groups"] = bson.M{"$in": tournamentDAO.Groups}
	}
	if tournamentDAO.Rounds != nil {
		setUpdates["rounds"] = bson.M{"$in": tournamentDAO.Rounds}
	}

	if len(setUpdates) > 0 {
		(*update)["$pull"] = setUpdates
	}

	return update
}

func (r *Repository) AddTournamentOptionsBsonStruct(ctx context.Context, tournamentDAO *tournament_dao.UpdateTournamentOptionsDAOReq, update *bson.M) *bson.M {
	setUpdates := bson.M{}

	if tournamentDAO.Pots != nil {
		setUpdates["pots"] = bson.M{"$each": tournamentDAO.Pots}
	}

	if tournamentDAO.Groups != nil {
		setUpdates["groups"] = bson.M{"$each": tournamentDAO.Groups}
	}

	if tournamentDAO.Rounds != nil {
		setUpdates["rounds"] = bson.M{"$each": tournamentDAO.Rounds}
	}

	if len(setUpdates) > 0 {
		(*update)["$push"] = setUpdates
	}

	if tournamentDAO.DoubleEliminationID != nil {
		(*update)["$set"] = bson.M{"double_elimination_id": tournamentDAO.DoubleEliminationID}
	}

	return update

}

func (r *Repository) UpdateTournamentOptions(ctx context.Context, tournamentID string, tournamentDAO *tournament_dao.UpdateTournamentOptionsDAOReq, add bool) error {
	tournamentOID, err := r.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}
	tournamentDAO.RenewUpdate()

	filter := bson.M{"_id": tournamentOID}

	update := bson.M{}

	if add {
		update = *r.AddTournamentOptionsBsonStruct(ctx, tournamentDAO, &update)
	} else {
		// Remove
		update = *r.RemoveTournamentOptionsBsonStruct(ctx, tournamentDAO, &update)
	}
	if update == nil {
		return fmt.Errorf("error updating tournament, nothing to update: %w", err)
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
		return fmt.Errorf("%w: no tournament found with id: %s", customerrors.ErrNotFound, tournamentID)
	}

	return nil
}

func (r *Repository) AddLeagueInTournament(ctx context.Context, tournamentID string, leagueID string) error {
	tournamentOID, err := r.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	leagueOID, err := r.ConvertToObjectID(leagueID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": *tournamentOID}

	update := bson.M{"league_id": *leagueOID}

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
		return fmt.Errorf("%w: no tournament found with id: %s", customerrors.ErrNotFound, leagueID)
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
