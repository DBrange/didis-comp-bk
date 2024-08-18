package repository

import (
	"context"
	"fmt"
	"time"

	api_utils "github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/tournament_group/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) TournamentGroupColl() *mongo.Collection {
	return r.tournamentGroupColl
}

func (r *Repository) CreateTournamentGroup(ctx context.Context, tournamentOID *primitive.ObjectID, position int) (string, error) {
	tournamentGroupDAO := dao.CreateTournamentGroupDAOReq{
		TournamentID: *tournamentOID,
		Matches:      []*primitive.ObjectID{},
		Competitors:  []*dao.TournamentGroupCompetitorDAOReq{},
		Position:     position,
	}

	tournamentGroupDAO.SetTimeStamp()

	result, err := r.tournamentGroupColl.InsertOne(ctx, tournamentGroupDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for tournamentGroup: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error tournamentGroup scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting tournamentGroup: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetTournamentGroupByID(ctx context.Context, tournamentGroupOID *primitive.ObjectID) (*dao.GetTournamentGroupDAORes, error) {
	var tournamentGroup dao.GetTournamentGroupDAORes

	filter := bson.M{"_id": *tournamentGroupOID}

	err := r.tournamentGroupColl.FindOne(ctx, filter).Decode(&tournamentGroup)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for tournamentGroup: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the tournamentGroup: %w", err)
	}

	return &tournamentGroup, nil
}

func (r *Repository) UpdateTournamentGroup(ctx context.Context, tournamentGroupID string, tournamentGroupInfoDAO *dao.UpdateTournamentGroupDAOReq) error {
	tournamentGroupOID, err := r.ConvertToObjectID(tournamentGroupID)
	if err != nil {
		return err
	}

	tournamentGroupInfoDAO.RenewUpdate()

	filter := bson.M{"_id": *tournamentGroupOID}
	update, err := api_utils.StructToBsonMap(tournamentGroupInfoDAO)
	if err != nil {
		return err
	}

	result, err := r.tournamentGroupColl.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": update},
	)
	if err != nil {
		return fmt.Errorf("error updating tournamentGroup: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no tournamentGroup found with id: %s", customerrors.ErrNotFound, tournamentGroupID)
	}

	return nil
}

func (r *Repository) DeleteTournamentGroup(ctx context.Context, tournamentGroupID string) error {
	err := r.SetDeletedAt(ctx, r.tournamentGroupColl, tournamentGroupID, "tournamentGroup")
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) AddCompetitorInGroup(ctx context.Context, groupOID, competitorOID *primitive.ObjectID) error {
	filter := bson.M{"_id": groupOID}
	competitor := &dao.TournamentGroupCompetitorDAOReq{
		CompetitorID: competitorOID,
	}

	update := bson.M{"$push": bson.M{"competitors": competitor}}

	result, err := r.tournamentGroupColl.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return fmt.Errorf("error updating tournamentGroup: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no tournamentGroup found with id: %s", customerrors.ErrNotFound, groupOID.Hex())
	}

	return nil
}

func (r *Repository) AddCompetitorsToTournamentGroups(ctx context.Context, tournamentOID *primitive.ObjectID, groupDTOs []*dao.AddCompetitorsToTournamentGroupsDAOReq) error {
	var operations []mongo.WriteModel

	for _, group := range groupDTOs {
		competitors := make([]*dao.TournamentGroupCompetitorDAOReq, len(group.Competitors))
		for i, competitorID := range group.Competitors {
			competitors[i] = &dao.TournamentGroupCompetitorDAOReq{
				CompetitorID: competitorID,
			}
		}

		filter := bson.M{"_id": group.GroupID, "tournament_id": tournamentOID}
		update := bson.M{
			"$set": bson.M{
				"competitors": competitors,
			},
		}

		operation := mongo.NewUpdateOneModel().
			SetFilter(filter).
			SetUpdate(update).
			SetUpsert(false)

		operations = append(operations, operation)
	}

	_, err := r.tournamentGroupColl.BulkWrite(ctx, operations, options.BulkWrite().SetOrdered(false))
	if err != nil {
		return fmt.Errorf("error updating tournament groups: %v", err)
	}

	return nil
}

func (r *Repository) AddMultipleMatchesInTournamentGroup(ctx context.Context, groupOID, tournamentOID *primitive.ObjectID, matchOIDs []*primitive.ObjectID) error {

	var operations []mongo.WriteModel

	filter := bson.M{"_id": groupOID, "tournament_id": tournamentOID}
	update := bson.M{
		"$set": bson.M{
			"matches":    matchOIDs,
			"updated_at": time.Now().UTC(),
		},
	}

	operation := mongo.NewUpdateOneModel().
		SetFilter(filter).
		SetUpdate(update).
		SetUpsert(false)

	operations = append(operations, operation)

	_, err := r.tournamentGroupColl.BulkWrite(ctx, operations, options.BulkWrite().SetOrdered(false))
	if err != nil {
		return fmt.Errorf("error updating tournament groups: %v", err)
	}

	return nil
}

func (r *Repository) AddMatchInTournamentGroup(ctx context.Context, groupOID, tournamentOID, matchOID *primitive.ObjectID) error {
	filter := bson.M{"_id": groupOID, "tournament_id": tournamentOID}

	update := bson.M{
		"$push": bson.M{"matches": matchOID},
	}

	result, err := r.tournamentGroupColl.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return fmt.Errorf("error updating tournamentGroup: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no tournamentGroup found with id: %s", customerrors.ErrNotFound, groupOID.Hex())
	}

	return nil

}

func (r *Repository) GetTournamentGroupMatches(ctx context.Context, groupID *primitive.ObjectID) ([]string, []string, error) {
	filter := bson.M{"_id": groupID}

	projection := bson.M{"matches": 1, "competitors": 1}

	opts := options.FindOne().SetProjection(projection)

	var result struct {
		Matches     []*primitive.ObjectID                  `bson:"matches"`
		Competitors []*dao.TournamentGroupCompetitorDAOReq `bson:"competitors"`
	}

	err := r.tournamentGroupColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil, fmt.Errorf("%w: error when searching for tournamentGroup: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, nil, fmt.Errorf("error when searching for the tournamentGroup: %w", err)
	}

	matchesStr := make([]string, len(result.Matches))
	competitorStr := make([]string, len(result.Competitors))
	for i, match := range result.Matches {
		matchesStr[i] = match.Hex()
	}

	for i, comeptitor := range result.Competitors {
		competitorStr[i] = comeptitor.CompetitorID.Hex()
	}

	return matchesStr, competitorStr, nil
}

func (r *Repository) GetTournamentGroupMatchesByPosition(ctx context.Context, position int, tournamentOID *primitive.ObjectID) ([]string, []string, error) {
	filter := bson.M{"position": position, "tournament_id": tournamentOID}

	projection := bson.M{"matches": 1, "competitors": 1}

	opts := options.FindOne().SetProjection(projection)

	var result struct {
		Matches     []*primitive.ObjectID                  `bson:"matches"`
		Competitors []*dao.TournamentGroupCompetitorDAOReq `bson:"competitors"`
	}

	err := r.tournamentGroupColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil, fmt.Errorf("%w: error when searching for tournamentGroup: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, nil, fmt.Errorf("error when searching for the tournamentGroup: %w", err)
	}

	matchesStr := make([]string, len(result.Matches))
	competitorStr := make([]string, len(result.Competitors))
	for i, match := range result.Matches {
		matchesStr[i] = match.Hex()
	}

	for i, comeptitor := range result.Competitors {
		competitorStr[i] = comeptitor.CompetitorID.Hex()
	}

	return matchesStr, competitorStr, nil
}

func (r *Repository) UpdateGroupPositions(ctx context.Context, groupOID *primitive.ObjectID, position int) error {
	filter := bson.M{"_id": groupOID}

	update := bson.M{"$set": bson.M{"position": position}}

	result, err := r.tournamentGroupColl.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return fmt.Errorf("error updating group: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no group found with id: %s", customerrors.ErrNotFound, groupOID.Hex())
	}

	return nil
}

func (r *Repository) DeleteGroupByPosition(ctx context.Context, position int, tournamentOID *primitive.ObjectID) error {
	filter := bson.M{"position": position, "tournament_id": tournamentOID}

	result, err := r.tournamentGroupColl.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("%w: error deleting group: %s", customerrors.ErrDeleted, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("%w: no group found with position: %d", customerrors.ErrNotFound, position)
	}

	return nil

}