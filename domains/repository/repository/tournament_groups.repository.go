package repository

import (
	"context"
	"fmt"

	api_utils "github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/tournament_group/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) TournamentGroupColl() *mongo.Collection {
	return r.tournamentGroupColl
}

func (r *Repository) CreateTournamentGroup(ctx context.Context, tournamentID string) (string, error) {
	tournamenOID, err := r.ConvertToObjectID(tournamentID)
	if err != nil {
		return "", err
	}

	var tournamentGroupDAO dao.CreateTournamentGroupDAOReq

	tournamentGroupDAO.TournamentID = *tournamenOID

	tournamentGroupDAO.Matches = []primitive.ObjectID{}
	tournamentGroupDAO.Competitors = []primitive.ObjectID{}

	tournamentGroupDAO.SetTimeStamp()

	result, err := r.tournamentGroupColl.InsertOne(ctx, tournamentGroupDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for tournamentGroup: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error tournamentGroup scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting tournamentGroup: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetTournamentGroupByID(ctx context.Context, tournamentGroupID string) (*dao.GetTournamentGroupByIDDAORes, error) {
	var tournamentGroup dao.GetTournamentGroupByIDDAORes

	tournamentGroupOID, err := r.ConvertToObjectID(tournamentGroupID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *tournamentGroupOID}

	err = r.tournamentGroupColl.FindOne(ctx, filter).Decode(&tournamentGroup)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for tournamentGroup: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the tournamentGroup: %w", err)
	}

	return &tournamentGroup, nil
}

func (r *Repository) UpdateTournamentGroup(ctx context.Context, tournamentGroupID string, tournamentGroupInfoDAO *dao.UpdateTournamentGroupDAOReq) error {
	tournamentGroupOID, err := r.ConvertToObjectID(tournamentGroupID)
	if err != nil {
		return err
	}

	tournamentGroupInfoDAO.RenewUpdate()

	filter := bson.M{"_id": *tournamentGroupOID}
	update, err := api_utils.StructToBsonMap(tournamentGroupInfoDAO)
	if err != nil {
		return err
	}

	result, err := r.tournamentGroupColl.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": update},
	)
	if err != nil {
		return fmt.Errorf("error updating tournamentGroup: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no tournamentGroup found with id: %s", customerrors.ErrNotFound, tournamentGroupID)
	}

	return nil
}

func (r *Repository) DeleteTournamentGroup(ctx context.Context, tournamentGroupID string) error {
	err := r.SetDeletedAt(ctx, r.tournamentGroupColl, tournamentGroupID, "tournamentGroup")
	if err != nil {
		return err
	}

	return nil
}
