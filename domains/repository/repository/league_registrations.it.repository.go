package repository

import (
	"context"
	"fmt"

	api_assets "github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/league_registration/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateLeagueRegistration(ctx context.Context, leagueRegistrationInfoDAO *dao.CreateLeagueRegistrationDAOReq) (string, error) {
	leagueRegistrationInfoDAO.SetTimeStamp()

	result, err := r.leagueRegistrationColl.InsertOne(ctx, leagueRegistrationInfoDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for leagueRegistration: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error leagueRegistration scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting leagueRegistration: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetLeagueRegistrationByID(ctx context.Context, leagueRegistrationID string) (*dao.GetLeagueRegistrationByIDDAORes, error) {
	var leagueRegistration dao.GetLeagueRegistrationByIDDAORes

	leagueRegistrationOID, err := r.ConvertToObjectID(leagueRegistrationID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *leagueRegistrationOID}

	err = r.leagueRegistrationColl.FindOne(ctx, filter).Decode(&leagueRegistration)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for leagueRegistration: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the leagueRegistration: %w", err)
	}

	return &leagueRegistration, nil
}

func (r *Repository) UpdateLeagueRegistration(ctx context.Context, leagueRegistrationID string, leagueRegistrationInfoDAO *dao.UpdateLeagueRegistrationDAOReq) error {
	leagueRegistrationOID, err := r.ConvertToObjectID(leagueRegistrationID)
	if err != nil {
		return err
	}

	leagueRegistrationInfoDAO.RenewUpdate()

	filter := bson.M{"_id": *leagueRegistrationOID}
	update, err := api_assets.StructToBsonMap(leagueRegistrationInfoDAO)
	if err != nil {
		return err
	}

	result, err := r.leagueRegistrationColl.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": update},
	)
	if err != nil {
		return fmt.Errorf("error updating leagueRegistration: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no leagueRegistration found with id: %s", customerrors.ErrNotFound, leagueRegistrationID)
	}

	return nil
}

func (r *Repository) DeleteLeagueRegistration(ctx context.Context, leagueRegistrationID string) error {
	err := r.SetDeletedAt(ctx, r.leagueRegistrationColl, leagueRegistrationID, "leagueRegistration")
	if err != nil {
		return err
	}

	return nil
}
