package repository

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/organizer/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	organizer.Categories = []primitive.ObjectID{}
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

func (r *Repository) VerifyOrganizerExists(ctx context.Context, organizerID string) error {
	var result struct{}

	organizerOID, err := r.ConvertToObjectID(organizerID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": organizerOID}

	opts := options.FindOne().SetProjection(bson.M{"_id": 1})

	err = r.organizerColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: error when searching for organizer: %s", customerrors.ErrNotFound, err.Error())
		}
		return fmt.Errorf("error when searching for the organizer: %w", err)
	}

	return nil
}

func (r *Repository) AddCategoryInOrganizer(ctx context.Context, organizerOID, categoryOID *primitive.ObjectID) error {
	filter := bson.M{"_id": organizerOID}

	update := bson.M{
		"$push": bson.M{"categories": categoryOID},
	}

	result, err := r.organizerColl.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error updating organizer: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no opinion found with id: %s", customerrors.ErrNotFound, organizerOID.Hex())
	}

	return nil
}

func (r *Repository) GetCategoriesFromOrganizer(ctx context.Context, organizerOID *primitive.ObjectID, sport models.SPORT, competitorType models.COMPETITOR_TYPE) ([]dao.GetCategoriesFromOrganizerDAORes, error) {
	pipeline := mongo.Pipeline{
		// Match the organizer by ID
		bson.D{{Key: "$match", Value: bson.M{"_id": organizerOID}}},

		// Unwind the array of category IDs
		bson.D{{Key: "$unwind", Value: "$categories"}},

		// Lookup the categories using the unwound category IDs
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "categories",
			"localField":   "categories",
			"foreignField": "_id",
			"as":           "category",
		}}},

		// Unwind the array of categories (should only be one per unwound ID)
		bson.D{{Key: "$unwind", Value: "$category"}},

		// Match the categories by sport and competitor type
		bson.D{{Key: "$match", Value: bson.M{
			"category.sport":           sport,
			"category.competitor_type": competitorType,
		}}},

		// Lookup the category registrations using the category IDs
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "category_registrations",
			"localField":   "category._id",
			"foreignField": "category_id",
			"as":           "category_registration",
		}}},

		// Unwind the array of category registrations
		bson.D{{Key: "$unwind", Value: "$category_registration"}},

		// Lookup the competitor users using the competitor IDs
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "competitor_users",
			"localField":   "category_registration.competitor_id",
			"foreignField": "competitor_id",
			"as":           "competitor_user",
		}}},

		// Unwind the array of competitor users
		bson.D{{Key: "$unwind", Value: "$competitor_user"}},

		// Lookup the users using the user IDs
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "users",
			"localField":   "competitor_user.user_id",
			"foreignField": "_id",
			"as":           "user",
		}}},

		// Unwind the array of users
		bson.D{{Key: "$unwind", Value: "$user"}},

		// Group by category and competitor, and nest the users
		bson.D{{Key: "$group", Value: bson.M{
			"_id": bson.M{
				"category_id":   "$category._id",
				"competitor_id": "$competitor_user.competitor_id",
			},
			"category_id":          bson.M{"$first": "$category._id"},
			"competitor_id":        bson.M{"$first": "$competitor_user.competitor_id"},
			"total_participants":   bson.M{"$first": "$category.total_participants"},
			"points":               bson.M{"$first": "$category_registration.points"},
			"current_position":     bson.M{"$first": "$category_registration.current_position"},
			"registered_positions": bson.M{"$first": "$category_registration.registered_positions"},
			"users": bson.M{
				"$push": bson.M{
					"user_id":    "$user._id",
					"first_name": "$user.first_name",
					"last_name":  "$user.last_name",
					"image":      "$user.image",
				},
			},
		}}},

		// Group by category and nest the competitors
		bson.D{{Key: "$group", Value: bson.M{
			"_id":                "$category_id",
			"category_id":        bson.M{"$first": "$category_id"},
			"total_participants": bson.M{"$first": "$total_participants"},
			"competitors": bson.M{
				"$push": bson.M{
					"competitor_id":        "$competitor_id",
					"points":               "$points",
					"current_position":     "$current_position",
					"registered_positions": "$registered_positions",
					"users":                "$users",
				},
			},
		}}},

		// Project the final structure
		bson.D{{Key: "$project", Value: bson.M{
			"category_id":        "$category_id",
			"competitors":        "$competitors",
			"total_participants": "$total_participants",
		}}},
	}

	// Execute the aggregation pipeline
	cursor, err := r.organizerColl.Aggregate(ctx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for categoryRegistration: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the categoryRegistration: %w", err)
	}

	defer cursor.Close(ctx)

	var categoriesDAO []dao.GetCategoriesFromOrganizerDAORes
	if err = cursor.All(ctx, &categoriesDAO); err != nil {
		return nil, fmt.Errorf("error when decoding categoryRegistration: %w", err)
	}

	return categoriesDAO, nil
}
