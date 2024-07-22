package repository

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/tournament_registration/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateTournamentRegistration(ctx context.Context, tournamentRegistrationInfoDAO *dao.CreateTournamentRegistrationDAOReq) (string, error) {
	tournamentRegistrationInfoDAO.SetTimeStamp()

	result, err := r.tournamentRegistrationColl.InsertOne(ctx, tournamentRegistrationInfoDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for tournamentRegistration: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error tournamentRegistration scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting tournamentRegistration: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetTournamentRegistrationByID(ctx context.Context, tournamentRegistrationID string) (*dao.GetTournamentRegistrationByIDDAORes, error) {
	var tournamentRegistration dao.GetTournamentRegistrationByIDDAORes

	tournamentRegistrationOID, err := r.ConvertToObjectID(tournamentRegistrationID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *tournamentRegistrationOID}

	err = r.tournamentRegistrationColl.FindOne(ctx, filter).Decode(&tournamentRegistration)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for tournamentRegistration: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the tournamentRegistration: %w", err)
	}

	return &tournamentRegistration, nil
}

// func (r *Repository) UpdateTournamentRegistration(ctx context.Context, tournamentRegistrationID string, tournamentRegistrationInfoDAO *dao.UpdateTournamentRegistrationDAOReq) error {
// 	tournamentRegistrationOID, err := r.ConvertToObjectID(tournamentRegistrationID)
// 	if err != nil {
// 		return err
// 	}

// 	tournamentRegistrationInfoDAO.RenewUpdate()

// 	filter := bson.M{"_id": *tournamentRegistrationOID}
// 	update, err := api_assets.StructToBsonMap(tournamentRegistrationInfoDAO)
// 	if err != nil {
// 		return err
// 	}

// 	result, err := r.tournamentRegistrationColl.UpdateOne(
// 		ctx,
// 		filter,
// 		bson.M{"$set": update},
// 	)
// 	if err != nil {
// 		return fmt.Errorf("error updating tournamentRegistration: %w", err)
// 	}

// 	if result.MatchedCount == 0 {
// 		return fmt.Errorf("%w: no tournamentRegistration found with id: %s", customerrors.ErrNotFound, tournamentRegistrationID)
// 	}

// 	return nil
// }

func (r *Repository) DeleteTournamentRegistration(ctx context.Context, tournamentRegistrationID string) error {
	err := r.setDeletedAt(ctx, r.tournamentRegistrationColl, tournamentRegistrationID, "tournamentRegistration")
	if err != nil {
		return err
	}

	return nil
}
