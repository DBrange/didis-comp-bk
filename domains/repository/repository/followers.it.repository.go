package repository

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/follower/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) CreateFollower(ctx context.Context, followerDAO *dao.CreateFollowerDAOReq) error {
	followerDAO.SetTimeStamp()

	_, err := r.followerColl.InsertOne(ctx, followerDAO)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return fmt.Errorf("%w: error follower scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return fmt.Errorf("error when inserting follower: %w", err)
	}

	return nil
}

func (r *Repository) VerifyFollowerExistsRelation(ctx context.Context, followerDAO *dao.CreateFollowerDAOReq) error {
	filter := bson.M{"from": followerDAO.From}

	if followerDAO.ToUser != nil {
		filter["to_user"] = followerDAO.ToUser
	} else {
		filter["to_competitor"] = followerDAO.ToCompetitor
	}

	projection := bson.M{"_id": 1}

	opts := options.FindOne().SetProjection(projection)

	var documentFinded struct{}

	if err := r.followerColl.FindOne(ctx, filter, opts).Decode(&documentFinded); err == nil {
		return fmt.Errorf("%w: follower already exists", customerrors.ErrAuthorizationHeader)
	} else if err != mongo.ErrNoDocuments {
		return fmt.Errorf("error when checking for existing follower: %w", err)
	}

	return nil
}

func (r *Repository) GetFollowerByID(ctx context.Context, followerID string) (*dao.GetFollowerByIDDAORes, error) {
	var follower dao.GetFollowerByIDDAORes

	followerOID, err := r.ConvertToObjectID(followerID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *followerOID}

	err = r.followerColl.FindOne(ctx, filter).Decode(&follower)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for follower: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the follower: %w", err)
	}

	return &follower, nil
}

func (r *Repository) DeleteFollower(ctx context.Context, followerID string) error {
	err := r.SetDeletedAt(ctx, r.followerColl, followerID, "follower")
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetCompetitorsFollowed(ctx context.Context, userOID *primitive.ObjectID, name string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) ([]*dao.GetCompetitorFollowedDAORes, error) {
	var followers []*dao.GetCompetitorFollowedDAORes

	if name == "" {
		return followers, nil
	}

	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{
			"from": userOID,
		}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "users",
			"localField":   "to_user",
			"foreignField": "_id",
			"as":           "user",
		}}},
		bson.D{{Key: "$unwind", Value: "$user"}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "competitor_users",
			"localField":   "user._id",
			"foreignField": "user_id",
			"as":           "competitor_user",
		}}},
		bson.D{{Key: "$unwind", Value: "$competitor_user"}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "competitor_users",
			"localField":   "user._id",
			"foreignField": "user_id",
			"as":           "competitor_user",
		}}},
		bson.D{{Key: "$unwind", Value: "$competitor_user"}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "competitors",
			"localField":   "competitor_user.competitor_id",
			"foreignField": "_id",
			"as":           "competitor",
		}}},
		bson.D{{Key: "$unwind", Value: "$competitor"}},
		bson.D{{Key: "$match", Value: bson.M{
			"competitor.sport": sport,
		}}},
		bson.D{{Key: "$limit", Value: 10}},
		bson.D{{Key: "$group", Value: bson.M{
			"_id": "$competitor_user.competitor_id",
		}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "competitor_users",
			"localField":   "_id",
			"foreignField": "competitor_id",
			"as":           "all_competitor_users",
		}}},
		bson.D{{Key: "$unwind", Value: "$all_competitor_users"}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "users",
			"localField":   "all_competitor_users.user_id",
			"foreignField": "_id",
			"as":           "all_users",
		}}},
		bson.D{{Key: "$unwind", Value: "$all_users"}},
		bson.D{{Key: "$group", Value: bson.M{
			"_id": "$_id",
			"users": bson.M{
				"$push": bson.M{
					"_id":        "$all_users._id",
					"first_name": "$all_users.first_name",
					"last_name":  "$all_users.last_name",
					"image":      "$all_users.image",
				},
			},
		}}},
		bson.D{{Key: "$project", Value: bson.M{
			"_id":   "$_id",
			"users": 1,
		}}},
	}

	pipeline = r.agetParticipantsOfCategoryNameFilter(pipeline, name)
	pipeline = r.singlesOrDoublesCategoryFilter(pipeline, competitorType)

	cursor, err := r.followerColl.Aggregate(ctx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for follower: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the follower: %w", err)
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &followers); err != nil {
		return nil, fmt.Errorf("error when decoding follower: %w", err)
	}
	for _, f := range followers {

		fmt.Printf("asas %+v", f)
	}
	return followers, nil
}
