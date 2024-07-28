package repository

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/guest_competitor/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateGuestCompetitor(ctx context.Context, guestCompetitorInfoDAO *dao.CreateGuestCompetitorDAOReq) (string, error) {
	guestCompetitorInfoDAO.SetTimeStamp()

	result, err := r.guestCompetitorColl.InsertOne(ctx, guestCompetitorInfoDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for guestCompetitor: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error guestCompetitor scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting guestCompetitor: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetGuestCompetitorByID(ctx context.Context, guestCompetitorID string) (*dao.GetGuestCompetitorByIDDAORes, error) {
	var guestCompetitor dao.GetGuestCompetitorByIDDAORes

	guestCompetitorOID, err := r.ConvertToObjectID(guestCompetitorID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *guestCompetitorOID}

	err = r.guestCompetitorColl.FindOne(ctx, filter).Decode(&guestCompetitor)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for guestCompetitor: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the guestCompetitor: %w", err)
	}

	return &guestCompetitor, nil
}

// func (r *Repository) UpdateGuestCompetitor(ctx context.Context, guestCompetitorID string, guestCompetitorInfoDAO *dao.UpdateGuestCompetitorDAOReq) error {
// 	guestCompetitorOID, err := r.ConvertToObjectID(guestCompetitorID)
// 	if err != nil {
// 		return err
// 	}

// 	guestCompetitorInfoDAO.RenewUpdate()

// 	filter := bson.M{"_id": *guestCompetitorOID}
// 	update, err := api_assets.StructToBsonMap(guestCompetitorInfoDAO)
// 	if err != nil {
// 		return err
// 	}

// 	result, err := r.guestCompetitorColl.UpdateOne(
// 		ctx,
// 		filter,
// 		bson.M{"$set": update},
// 	)
// 	if err != nil {
// 		return fmt.Errorf("error updating guestCompetitor: %w", err)
// 	}

// 	if result.MatchedCount == 0 {
// 		return fmt.Errorf("%w: no guestCompetitor found with id: %s", customerrors.ErrNotFound, guestCompetitorID)
// 	}

// 	return nil
// }

func (r *Repository) DeleteGuestCompetitor(ctx context.Context, guestCompetitorID string) error {
	err := r.SetDeletedAt(ctx, r.guestCompetitorColl, guestCompetitorID, "guestCompetitor")
	if err != nil {
		return err
	}

	return nil
}
