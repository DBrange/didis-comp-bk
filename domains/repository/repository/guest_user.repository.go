package repository

import (
	"context"
	"fmt"

	api_assets "github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/guest_user/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateGuestUser(ctx context.Context, guestUserInfoDAO *dao.CreateGuestUserDAOReq) (string, error) {
	guestUserInfoDAO.SetTimeStamp()

	result, err := r.guestUserColl.InsertOne(ctx, guestUserInfoDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for guestUser: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error guestUser scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting guestUser: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetGuestUserByID(ctx context.Context, guestUserID string) (*dao.GetGuestUserByIDDAORes, error) {
	var guestUser dao.GetGuestUserByIDDAORes

	guestUserOID, err := r.ConvertToObjectID(guestUserID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *guestUserOID}

	err = r.guestUserColl.FindOne(ctx, filter).Decode(&guestUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for guestUser: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the guestUser: %w", err)
	}

	return &guestUser, nil
}

func (r *Repository) UpdateGuestUser(ctx context.Context, guestUserID string, guestUserInfoDAO *dao.UpdateGuestUserDAOReq) error {
	guestUserOID, err := r.ConvertToObjectID(guestUserID)
	if err != nil {
		return err
	}

	guestUserInfoDAO.RenewUpdate()

	filter := bson.M{"_id": *guestUserOID}
	update, err := api_assets.StructToBsonMap(guestUserInfoDAO)
	if err != nil {
		return err
	}

	result, err := r.guestUserColl.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": update},
	)
	if err != nil {
		return fmt.Errorf("error updating guestUser: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no guestUser found with id: %s", customerrors.ErrNotFound, guestUserID)
	}

	return nil
}

func (r *Repository) DeleteGuestUser(ctx context.Context, guestUserID string) error {
	err := r.SetDeletedAt(ctx, r.guestUserColl, guestUserID, "guestUser")
	if err != nil {
		return err
	}

	return nil
}
