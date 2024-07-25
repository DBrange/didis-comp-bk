package repository

import (
	"context"
	"fmt"
	"time"

	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	tournament_opts "github.com/DBrange/didis-comp-bk/domains/repository/models/tournament"
	tournament_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/tournament/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) OrganizeTournament(
	ctx context.Context,
	tournamentInfoDAO *tournament_dao.CreateTournamentDAOReq,
	locationInfoDAO *location_dao.CreateLocationDAOReq,
	options *tournament_opts.OrganizeTournamentOptions,
	leagueID *string,
	organizerID string,
) error {

	locationID, err := r.CreateLocation(ctx, locationInfoDAO)
	if err != nil {
		return err
	}

	// if err = r.LeagueExists(ctx, leagueID); err != nil{
	// 	return nil
	// }

	if err := r.OrganizerExists(ctx, organizerID); err != nil {
		return err
	}

	if err := r.CreateTournament(
		ctx,
		tournamentInfoDAO,
		locationID,
		options,
		leagueID,
		organizerID,
	); err != nil {
		return err
	}

	return nil
}

func (r *Repository) CreateTournament(
	ctx context.Context,
	tournamentInfoDAO *tournament_dao.CreateTournamentDAOReq,
	locationID string,
	options *tournament_opts.OrganizeTournamentOptions,
	leagueID *string,
	organizerID string,
) error {
	locationOID, err := r.ConvertToObjectID(locationID)
	if err != nil {
		return err
	}

	tournamentInfoDAO.LocationID = *locationOID

	organizerOID, err := r.ConvertToObjectID(organizerID)
	if err != nil {
		return err
	}

	tournamentInfoDAO.OrganizerID = *organizerOID

	if leagueID != nil {
		leagueOID, err := r.ConvertToObjectID(*leagueID)
		if err != nil {
			return err
		}

		tournamentInfoDAO.LeagueID = leagueOID
	}

	tournamentInfoDAO.Rounds = []primitive.ObjectID{}
	tournamentInfoDAO.Matches = []primitive.ObjectID{}
	tournamentInfoDAO.Pots = []primitive.ObjectID{}
	tournamentInfoDAO.Groups = []primitive.ObjectID{}

	tournamentInfoDAO.SetTimeStamp()

	fmt.Printf("este es torneo: %v", tournamentInfoDAO)

	_, err = r.tournamentColl.InsertOne(ctx, tournamentInfoDAO)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return fmt.Errorf("%w: error tournament scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}
		return fmt.Errorf("error when inserting tournament: %w", err)
	}

	return nil
}

func (r *Repository) AddLeagueInTournament(ctx context.Context, tournamentOID *primitive.ObjectID, leagueOID *primitive.ObjectID) error {
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
		return fmt.Errorf("%w: no tournament found with id: %s", customerrors.ErrNotFound, leagueOID.Hex())
	}

	return nil
}
