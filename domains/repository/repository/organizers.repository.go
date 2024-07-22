package repository

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/organizer/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) CreateOrganizer(ctx context.Context, userID string) error {
	userOID, err := r.ConvertToObjectID(userID)
	if err != nil {
		return err
	}

	var organizer dao.CreateOrganizerDAOReq
	organizer.UserID = *userOID
	organizer.SetTimeStamp()

	_, err = r.organizerColl.InsertOne(ctx, organizer)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return fmt.Errorf("%w: error organizer scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return fmt.Errorf("error when inserting organizer: %w", err)
	}

	return nil
}

// func (r *Repository) GetOrganizerByID(ctx context.Context, organizerID string) (any, error) {
// 	var organizer dao.GetLocationByIDDAORes

// 	organizerOID, err := r.ConvertToObjectID(organizerID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	filter := bson.M{"_id": organizerOID}

// 	err = r.organizerColl.FindOne(ctx, filter).Decode(&organizer)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil, fmt.Errorf("%w: error when searching for organizer: %s", customerrors.ErrNotFound, err.Error())
// 		}
// 		return nil, fmt.Errorf("error when searching for the organizer: %w", err)
// 	}

// 	return &organizer, nil
// }

func (r *Repository) OrganizerExists(ctx context.Context, organizerID string) ( error) {
	var result struct{}

	organizerOID, err := r.ConvertToObjectID(organizerID)
	if err != nil {
		return  err
	}

	filter := bson.M{"_id": organizerOID}

	opts := options.FindOne().SetProjection(bson.M{"_id": 1})

	err = r.organizerColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return  fmt.Errorf("%w: error when searching for organizer: %s", customerrors.ErrNotFound, err.Error())
		}
		return  fmt.Errorf("error when searching for the organizer: %w", err)
	}

	return  nil
}
