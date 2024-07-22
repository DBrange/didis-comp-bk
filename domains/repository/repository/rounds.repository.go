package repository

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/round/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateRound(ctx context.Context, round *dao.CreateRoundDAOReq) (string, error) {
	result, err := r.roundColl.InsertOne(ctx, round)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error round scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting round: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetRoundByID(ctx context.Context, roundID string) (*dao.GetRoundByIDDAORes, error) {
	var round dao.GetRoundByIDDAORes

	roundOID, err := r.ConvertToObjectID(roundID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *roundOID}

	err = r.roundColl.FindOne(ctx, filter).Decode(&round)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for round: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the round: %w", err)
	}

	return &round, nil
}
