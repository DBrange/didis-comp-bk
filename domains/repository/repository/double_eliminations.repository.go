package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/double_elimination/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) DoubleEliminationColl() *mongo.Collection {
	return r.doubleEliminationColl
}

func (r *Repository) CreateDoubleElimination(ctx context.Context, doubleEliminationDAO *dao.CreateDoubleEliminationDAOReq) (string, error) {
	doubleEliminationDAO.SetTimeStamp()

	result, err := r.doubleEliminationColl.InsertOne(ctx, doubleEliminationDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for doubleElimination: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error doubleElimination scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting doubleElimination: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}
func (r *Repository) CreateDoubleEliminationEmpty(ctx context.Context) (string, error) {
	var doubleEliminationEmptyDAO dao.CreateDoubleEliminationDAOReq

	doubleEliminationEmptyDAO.Matches = []*primitive.ObjectID{}
	doubleEliminationEmptyDAO.Rounds = []*primitive.ObjectID{}

	doubleEliminationEmptyDAO.SetTimeStamp()

	result, err := r.doubleEliminationColl.InsertOne(ctx, doubleEliminationEmptyDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for doubleElimination: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error doubleElimination scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting doubleElimination: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetDoubleEliminationByID(ctx context.Context, doubleEliminationID string) (*dao.GetDoubleEliminationByIDDAORes, error) {
	var doubleElimination dao.GetDoubleEliminationByIDDAORes

	doubleEliminationOID, err := r.ConvertToObjectID(doubleEliminationID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *doubleEliminationOID}

	err = r.doubleEliminationColl.FindOne(ctx, filter).Decode(&doubleElimination)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for doubleElimination: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the doubleElimination: %w", err)
	}

	return &doubleElimination, nil
}

func (r *Repository) UpdateDoubleEliminationBsonStruct(ctx context.Context, doubleEliminationDAO *dao.UpdateDoubleEliminationDAOReq, update *bson.M, add bool) *bson.M {
	operation := "$pull"
	arrayModifier := "$in"

	if add {
		operation = "$push"
		arrayModifier = "$each"
	}

	setUpdates := bson.M{}

	if doubleEliminationDAO.Matches != nil {
		setUpdates["matches"] = bson.M{arrayModifier: doubleEliminationDAO.Matches}
	}

	if doubleEliminationDAO.Rounds != nil {
		setUpdates["rounds"] = bson.M{arrayModifier: doubleEliminationDAO.Rounds}
	}

	if len(setUpdates) > 0 {
		(*update)[operation] = setUpdates
	}

	return update

}

func (r *Repository) UpdateDoubleElimination(ctx context.Context, doubleEliminationOID *primitive.ObjectID, doubleEliminationInfoDAO *dao.UpdateDoubleEliminationDAOReq, add bool) error {
	doubleEliminationInfoDAO.RenewUpdate()

	filter := bson.M{"_id": doubleEliminationOID}

	update := bson.M{}

	update = *r.UpdateDoubleEliminationBsonStruct(ctx, doubleEliminationInfoDAO, &update, add)
	if update == nil {
		return fmt.Errorf("error updating doubleElimination, nothing to update: %w", nil)
	}

	result, err := r.doubleEliminationColl.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return fmt.Errorf("error updating doubleElimination: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no doubleElimination found with id: %s", customerrors.ErrNotFound, doubleEliminationOID.Hex())
	}

	return nil
}

func (r *Repository) DeleteDoubleElimination(ctx context.Context, doubleEliminationID string) error {
	err := r.SetDeletedAt(ctx, r.doubleEliminationColl, doubleEliminationID, "doubleElimination")
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) AddMatchInDoubleElim(ctx context.Context, doubleElimOID, matchOID *primitive.ObjectID) error {

	filter := bson.M{"_id": doubleElimOID}

	update := bson.M{
		"$push": bson.M{"matches": matchOID},
		"$set":  bson.M{"updated_at": time.Now().UTC()},
	}

	result, err := r.doubleEliminationColl.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return fmt.Errorf("error updating tournament: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no match found in doubleElimination with id: %s", customerrors.ErrNotFound, matchOID.Hex())
	}

	return nil
}

func (r *Repository) GetDoubleElimRoundID(ctx context.Context, doubleEliminationOID *primitive.ObjectID, round models.ROUND) (string, error) {
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"_id": doubleEliminationOID}}},
		bson.D{{Key: "$unwind", Value: "$rounds"}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "rounds",
			"localField":   "rounds",
			"foreignField": "_id",
			"as":           "round",
		}}},
		bson.D{{Key: "$unwind", Value: "$round"}},
		bson.D{{Key: "$match", Value: bson.M{"round.round": round}}},
		bson.D{{Key: "$project", Value: bson.M{"_id": "$round._id"}}},
	}

	cursor, err := r.doubleEliminationColl.Aggregate(ctx, pipeline)
	if err != nil {
		return "", fmt.Errorf("error when aggregating rounds: %w", err)
	}
	defer cursor.Close(ctx)

	var result struct {
		ID *primitive.ObjectID `bson:"_id"`
	}

	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return "", fmt.Errorf("error when decoding round ID: %w", err)
		}

		return result.ID.Hex(), nil
	}

	if err := cursor.Err(); err != nil {
		return "", fmt.Errorf("cursor error: %w", err)
	}

	return "", fmt.Errorf("%w: no round found with round name: %s", customerrors.ErrNotFound, round)
}

func (r *Repository) GetAllDoubleElimRoundIDs(ctx context.Context, doubleEliminationOID *primitive.ObjectID) ([]string, error) {
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"_id": doubleEliminationOID}}},
		bson.D{{Key: "$unwind", Value: "$rounds"}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "rounds",
			"localField":   "rounds",
			"foreignField": "_id",
			"as":           "round",
		}}},
		bson.D{{Key: "$unwind", Value: "$round"}},
		bson.D{{Key: "$project", Value: bson.M{"_id": "$round._id"}}},
	}

	cursor, err := r.doubleEliminationColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error when aggregating rounds: %w", err)
	}
	defer cursor.Close(ctx)

	var roundIDs []string

	for cursor.Next(ctx) {
		var result struct {
			ID *primitive.ObjectID `bson:"_id"`
		}
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("error when decoding round ID: %w", err)
		}
		roundIDs = append(roundIDs, result.ID.Hex())
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return roundIDs, nil
}

func (r *Repository) GetDoubleElimInfoToFinaliseIt(ctx context.Context, doubleElimOID *primitive.ObjectID) (*dao.GetDoubleElimInfoToFinaliseItDAORes, error) {
	var doubleElimInfo *dao.GetDoubleElimInfoToFinaliseItDAORes

	filter := bson.M{"_id": doubleElimOID}

	opts := options.FindOne().SetProjection(bson.M{"total_prize": 1, "points": 1})

	err := r.doubleEliminationColl.FindOne(ctx, filter, opts).Decode(&doubleElimInfo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for doubleElim: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the doubleElim: %w", err)
	}

	return doubleElimInfo, nil
}

func (r *Repository) GetDoubleElimCompetitorChampion(ctx context.Context, doubleElimOID *primitive.ObjectID) (string, error) {
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"_id": doubleElimOID}}},
		bson.D{{Key: "$unwind", Value: "$rounds"}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "rounds",
			"localField":   "rounds",
			"foreignField": "_id",
			"as":           "round",
		}}},
		bson.D{{Key: "$unwind", Value: "$round"}},
		bson.D{{Key: "$match", Value: bson.M{"round.round": "F"}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "matches",
			"localField":   "round._id",
			"foreignField": "round_id",
			"as":           "match",
		}}},
		bson.D{{Key: "$unwind", Value: "$match"}},
		bson.D{{Key: "$project", Value: bson.M{"_id": "$match.winner"}}},
	}

	cursor, err := r.doubleEliminationColl.Aggregate(ctx, pipeline)
	if err != nil {
		return "", fmt.Errorf("error when aggregating match: %w", err)
	}
	defer cursor.Close(ctx)

	var competitor struct {
		ID *primitive.ObjectID `bson:"_id"`
	}

	if cursor.Next(ctx) {
		if err := cursor.Decode(&competitor); err != nil {
			return "", fmt.Errorf("error when decoding round ID: %w", err)
		}

		return competitor.ID.Hex(), nil
	}

	if err := cursor.Err(); err != nil {
		return "", fmt.Errorf("cursor error: %w", err)
	}

	return "", fmt.Errorf("%w: no round found with match winner", customerrors.ErrNotFound)
}
