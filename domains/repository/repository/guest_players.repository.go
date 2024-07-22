package repository

import (
	"context"
	"fmt"

	api_assets "github.com/DBrange/didis-comp-bk/cmd/api/assets"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/guest_player/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateGuestPlayer(ctx context.Context, guestPlayerInfoDAO *dao.CreateGuestPlayerDAOReq) (string, error) {
	guestPlayerInfoDAO.SetTimeStamp()

	result, err := r.guestPlayerColl.InsertOne(ctx, guestPlayerInfoDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for guestPlayer: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error guestPlayer scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting guestPlayer: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetGuestPlayerByID(ctx context.Context, guestPlayerID string) (*dao.GetGuestPlayerByIDDAORes, error) {
	var guestPlayer dao.GetGuestPlayerByIDDAORes

	guestPlayerOID, err := r.ConvertToObjectID(guestPlayerID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *guestPlayerOID}

	err = r.guestPlayerColl.FindOne(ctx, filter).Decode(&guestPlayer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for guestPlayer: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the guestPlayer: %w", err)
	}

	return &guestPlayer, nil
}

func (r *Repository) UpdateGuestPlayer(ctx context.Context, guestPlayerID string, guestPlayerInfoDAO *dao.UpdateGuestPlayerDAOReq) error {
	guestPlayerOID, err := r.ConvertToObjectID(guestPlayerID)
	if err != nil {
		return err
	}

	guestPlayerInfoDAO.RenewUpdate()

	filter := bson.M{"_id": *guestPlayerOID}
	update, err := api_assets.StructToBsonMap(guestPlayerInfoDAO)
	if err != nil {
		return err
	}

	result, err := r.guestPlayerColl.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": update},
	)
	if err != nil {
		return fmt.Errorf("error updating guestPlayer: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no guestPlayer found with id: %s", customerrors.ErrNotFound, guestPlayerID)
	}

	return nil
}

func (r *Repository) DeleteGuestPlayer(ctx context.Context, guestPlayerID string) error {
	err := r.setDeletedAt(ctx, r.guestPlayerColl, guestPlayerID, "guestPlayer")
	if err != nil {
		return err
	}

	return nil
}
