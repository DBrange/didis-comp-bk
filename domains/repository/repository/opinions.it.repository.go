package repository

import (
	"context"
	"fmt"

	api_assets "github.com/DBrange/didis-comp-bk/cmd/api/assets"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/opinion/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateOpinion(ctx context.Context, opinionInfoDAO *dao.CreateOpinionDAOReq) (string, error) {
	opinionInfoDAO.SetTimeStamp()

	result, err := r.opinionColl.InsertOne(ctx, opinionInfoDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for opinion: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error opinion scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting opinion: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetOpinionByID(ctx context.Context, opinionID string) (*dao.GetOpinionByIDDAORes, error) {
	var opinion dao.GetOpinionByIDDAORes

	opinionOID, err := r.ConvertToObjectID(opinionID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *opinionOID}

	err = r.opinionColl.FindOne(ctx, filter).Decode(&opinion)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for opinion: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the opinion: %w", err)
	}

	return &opinion, nil
}

func (r *Repository) UpdateOpinion(ctx context.Context, opinionID string, opinionInfoDAO *dao.UpdateOpinionDAOReq) error {
	opinionOID, err := r.ConvertToObjectID(opinionID)
	if err != nil {
		return err
	}

	opinionInfoDAO.RenewUpdate()

	filter := bson.M{"_id": *opinionOID}
	update, err := api_assets.StructToBsonMap(opinionInfoDAO)
	if err != nil {
		return err
	}

	result, err := r.opinionColl.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": update},
	)
	if err != nil {
		return fmt.Errorf("error updating opinion: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no opinion found with id: %s", customerrors.ErrNotFound, opinionID)
	}

	return nil
}

func (r *Repository) DeleteOpinion(ctx context.Context, opinionID string) error {
	err := r.setDeletedAt(ctx, r.opinionColl, opinionID, "opinion")
	if err != nil {
		return err
	}

	return nil
}
