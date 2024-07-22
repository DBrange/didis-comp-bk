package repository

import (
	"context"
	"fmt"

	api_assets "github.com/DBrange/didis-comp-bk/cmd/api/assets"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/team/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateTeam(ctx context.Context, teamInfoDAO *dao.CreateTeamDAOReq) (string, error) {
	teamInfoDAO.SetTimeStamp()

	result, err := r.teamColl.InsertOne(ctx, teamInfoDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for team: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error team scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting team: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetTeamByID(ctx context.Context, teamID string) (*dao.GetTeamByIDDAORes, error) {
	var team dao.GetTeamByIDDAORes

	teamOID, err := r.ConvertToObjectID(teamID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *teamOID}

	err = r.teamColl.FindOne(ctx, filter).Decode(&team)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for team: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the team: %w", err)
	}

	return &team, nil
}

func (r *Repository) UpdateTeam(ctx context.Context, teamID string, teamInfoDAO *dao.UpdateTeamDAOReq) error {
	teamOID, err := r.ConvertToObjectID(teamID)
	if err != nil {
		return err
	}

	teamInfoDAO.RenewUpdate()

	filter := bson.M{"_id": *teamOID}
	update, err := api_assets.StructToBsonMap(teamInfoDAO)
	if err != nil {
		return err
	}

	result, err := r.teamColl.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": update},
	)
	if err != nil {
		return fmt.Errorf("error updating team: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no team found with id: %s", customerrors.ErrNotFound, teamID)
	}

	return nil
}

func (r *Repository) DeleteTeam(ctx context.Context, teamID string) error {
	err := r.setDeletedAt(ctx, r.teamColl, teamID, "team")
	if err != nil {
		return err
	}

	return nil
}
