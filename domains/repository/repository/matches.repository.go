package repository

import (
	"context"
	"fmt"

	api_assets "github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/match/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateMatch(ctx context.Context, match *dao.CreateMatchDAOReq) (string, error) {
	match.SetTimeStamp()

	result, err := r.matchColl.InsertOne(ctx, match)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error match scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting match: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetMatchByID(ctx context.Context, matchID string) (*dao.GetMatchByIDDAORes, error) {
	var match dao.GetMatchByIDDAORes

	matchOID, err := r.ConvertToObjectID(matchID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *matchOID}

	err = r.matchColl.FindOne(ctx, filter).Decode(&match)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for match: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the match: %w", err)
	}

	return &match, nil
}

func (r *Repository) UpdateMatch(ctx context.Context, matchID string, matchInfoDAO *dao.UpdateMatchDAOReq) error {
	matchOID, err := r.ConvertToObjectID(matchID)
	if err != nil {
		return err
	}

	matchInfoDAO.RenewUpdate()

	filter := bson.M{"_id": *matchOID}
	update, err := api_assets.StructToBsonMap(matchInfoDAO)
	if err != nil {
		return err
	}

	result, err := r.matchColl.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": update},
	)
	if err != nil {
		return fmt.Errorf("error updating match: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no match found with id: %s", customerrors.ErrNotFound, matchID)
	}

	return nil
}

func (r *Repository) DeleteMatch(ctx context.Context, matchID string) error {
	err := r.SetDeletedAt(ctx, r.matchColl, matchID, "match")
	if err != nil {
		return err
	}

	return nil
}
