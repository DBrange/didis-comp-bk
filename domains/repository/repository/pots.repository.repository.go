package repository

import (
	"context"
	"fmt"

	api_assets "github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/pot/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) PotColl() *mongo.Collection {
	return r.potColl
}

func (r *Repository) CreatePot(ctx context.Context, tournamentOID *primitive.ObjectID, position int) (string, error) {
	potInfoDAO := dao.CreatePotDAOReq{
		TournamentID: *tournamentOID,
		Competitors:  []primitive.ObjectID{},
		Position:     position,
	}

	potInfoDAO.SetTimeStamp()

	result, err := r.potColl.InsertOne(ctx, potInfoDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for pot: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error pot scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting pot: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetPotByID(ctx context.Context, potID string) (*dao.GetPotByIDDAORes, error) {
	var pot dao.GetPotByIDDAORes

	potOID, err := r.ConvertToObjectID(potID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *potOID}

	err = r.potColl.FindOne(ctx, filter).Decode(&pot)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for pot: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the pot: %w", err)
	}

	return &pot, nil
}

func (r *Repository) UpdatePot(ctx context.Context, potID string, potInfoDAO *dao.UpdatePotDAOReq) error {
	potOID, err := r.ConvertToObjectID(potID)
	if err != nil {
		return err
	}

	potInfoDAO.RenewUpdate()

	filter := bson.M{"_id": *potOID}
	update, err := api_assets.StructToBsonMap(potInfoDAO)
	if err != nil {
		return err
	}

	result, err := r.potColl.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": update},
	)
	if err != nil {
		return fmt.Errorf("error updating pot: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no pot found with id: %s", customerrors.ErrNotFound, potID)
	}

	return nil
}

func (r *Repository) DeletePot(ctx context.Context, potID string) error {
	err := r.SetDeletedAt(ctx, r.potColl, potID, "pot")
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) AddCompetitorInPot(ctx context.Context, potOID, competitorOID *primitive.ObjectID) error {
	filter := bson.M{"_id": potOID}

	update := bson.M{"$push": bson.M{"competitors": competitorOID}}

	result, err := r.potColl.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return fmt.Errorf("error updating pot: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no pot found with id: %s", customerrors.ErrNotFound, potOID.Hex())
	}

	return nil
}

func (r *Repository) RemoveCompetitorOfPot(ctx context.Context, potOID, competitorOID *primitive.ObjectID) error {
	filter := bson.M{"_id": potOID}

	update := bson.M{"$pull": bson.M{"competitors": competitorOID}}

	_, err := r.potColl.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return fmt.Errorf("error updating pot: %w", err)
	}

	return nil
}

func (r *Repository) SetCompetitorsInPots(ctx context.Context, tournamentOID *primitive.ObjectID, potDTOs []*dao.SetPotCompetitorDAOReq) error {
	var operations []mongo.WriteModel

	for _, pot := range potDTOs {
		filter := bson.M{"_id": pot.PotID, "tournament_id": tournamentOID}
		update := bson.M{
			"$set": bson.M{
				"competitors": pot.Competitors,
			},
		}

		operation := mongo.NewUpdateOneModel().
			SetFilter(filter).
			SetUpdate(update).
			SetUpsert(false)

		operations = append(operations, operation)
	}

	_, err := r.potColl.BulkWrite(ctx, operations, options.BulkWrite().SetOrdered(false))
	if err != nil {
		return fmt.Errorf("error updating tournament pots: %v", err)
	}

	return nil
}

func (r *Repository) UpdatePotPositions(ctx context.Context, potOID *primitive.ObjectID, position int) error {
	filter := bson.M{"_id": potOID}

	update := bson.M{"$set":bson.M{"position": position}}

	result, err := r.potColl.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return fmt.Errorf("error updating pot: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no pot found with id: %s", customerrors.ErrNotFound, potOID.Hex())
	}

	return nil
}

func (r *Repository) DeletePotByPosition(ctx context.Context, position int, tournamentOID *primitive.ObjectID) error {
	filter := bson.M{"position": position, "tournament_id": tournamentOID}

	result, err := r.potColl.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("%w: error deleting pot: %s", customerrors.ErrDeleted, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("%w: no pot found with position: %d", customerrors.ErrNotFound, position)
	}

	return nil

}
