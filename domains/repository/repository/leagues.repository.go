package repository

import (
	"context"
	"fmt"

	api_assets "github.com/DBrange/didis-comp-bk/cmd/api/assets"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/league/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateLeague(ctx context.Context, leagueInfoDAO *dao.CreateLeagueDAOReq) (string, error) {
	leagueInfoDAO.SetTimeStamp()

	result, err := r.leagueColl.InsertOne(ctx, leagueInfoDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for league: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error league scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting league: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetLeagueByID(ctx context.Context, leagueID string) (*dao.GetLeagueByIDDAORes, error) {
	var league dao.GetLeagueByIDDAORes

	leagueOID, err := r.ConvertToObjectID(leagueID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *leagueOID}

	err = r.leagueColl.FindOne(ctx, filter).Decode(&league)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for league: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the league: %w", err)
	}

	return &league, nil
}

func (r *Repository) UpdateLeague(ctx context.Context, leagueID string, leagueInfoDAO *dao.UpdateLeagueDAOReq) error {
	leagueOID, err := r.ConvertToObjectID(leagueID)
	if err != nil {
		return err
	}

	leagueInfoDAO.RenewUpdate()

	filter := bson.M{"_id": *leagueOID}
	update, err := api_assets.StructToBsonMap(leagueInfoDAO)
	if err != nil {
		return err
	}

	result, err := r.leagueColl.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": update},
	)
	if err != nil {
		return fmt.Errorf("error updating league: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no league found with id: %s", customerrors.ErrNotFound, leagueID)
	}

	return nil
}

func (r *Repository) DeleteLeague(ctx context.Context, leagueID string) error {
	err := r.setDeletedAt(ctx, r.leagueColl, leagueID, "league")
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) OrganizeLeague(ctx context.Context, leagueInfoDAO any) error{
	return nil
}
