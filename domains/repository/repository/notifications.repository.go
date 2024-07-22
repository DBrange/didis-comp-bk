package repository

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/notification/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateNotification(ctx context.Context, notificationInfoDAO *dao.CreateNotificationDAOReq) (string, error) {
	notificationInfoDAO.SetTimeStamp()

	result, err := r.notificationColl.InsertOne(ctx, notificationInfoDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for notification: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error notification scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting notification: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetNotificationByID(ctx context.Context, notificationID string) (*dao.GetNotificationByIDDAORes, error) {
	var notification dao.GetNotificationByIDDAORes

	notificationOID, err := r.ConvertToObjectID(notificationID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *notificationOID}

	err = r.notificationColl.FindOne(ctx, filter).Decode(&notification)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for notification: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the notification: %w", err)
	}

	return &notification, nil
}

// func (r *Repository) UpdateNotification(ctx context.Context, notificationID string, notificationInfoDAO *dao.UpdateNotificationDAOReq) error {
// 	notificationOID, err := r.ConvertToObjectID(notificationID)
// 	if err != nil {
// 		return err
// 	}

// 	notificationInfoDAO.RenewUpdate()

// 	filter := bson.M{"_id": *notificationOID}
// 	update, err := api_assets.StructToBsonMap(notificationInfoDAO)
// 	if err != nil {
// 		return err
// 	}

// 	result, err := r.notificationColl.UpdateOne(
// 		ctx,
// 		filter,
// 		bson.M{"$set": update},
// 	)
// 	if err != nil {
// 		return fmt.Errorf("error updating notification: %w", err)
// 	}

// 	if result.MatchedCount == 0 {
// 		return fmt.Errorf("%w: no notification found with id: %s", customerrors.ErrNotFound, notificationID)
// 	}

// 	return nil
// }

func (r *Repository) DeleteNotification(ctx context.Context, notificationID string) error {
	err := r.setDeletedAt(ctx, r.notificationColl, notificationID, "notification")
	if err != nil {
		return err
	}

	return nil
}
