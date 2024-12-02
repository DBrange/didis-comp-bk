package repository

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
	competitor_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/competitor/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) CreateCompetitor(ctx context.Context, sport models.SPORT, competitorType models.COMPETITOR_TYPE, OID *primitive.ObjectID) (string, error) {
	competitorDAO := &competitor_dao.CreateCompetitorDAOReq{}

	switch competitorType {
	case models.COMPETITOR_TYPE_SINGLE:
		competitorDAO.SingleID = OID

	case models.COMPETITOR_TYPE_DOUBLE:
		competitorDAO.DoubleID = OID

	case models.COMPETITOR_TYPE_TEAM:
		competitorDAO.TeamID = OID
	}

	competitorDAO.Sport = sport

	competitorDAO.SetTimeStamp()

	result, err := r.competitorColl.InsertOne(ctx, competitorDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for competitor: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error competitor scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting competitor: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetCompetitorByID(ctx context.Context, competitorID string) (*competitor_dao.GetCompetitorByIDDAORes, error) {
	var competitor competitor_dao.GetCompetitorByIDDAORes

	competitorOID, err := r.ConvertToObjectID(competitorID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *competitorOID}

	err = r.competitorColl.FindOne(ctx, filter).Decode(&competitor)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for competitor: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the competitor: %w", err)
	}

	return &competitor, nil
}

// func (r *Repository) UpdateCompetitor(ctx context.Context, competitorID string, competitorInfoDAO *competitor_dao.UpdateCompetitorDAOReq) error {
// 	competitorOID, err := r.ConvertToObjectID(competitorID)
// 	if err != nil {
// 		return err
// 	}

// 	competitorInfoDAO.RenewUpdate()

// 	filter := bson.M{"_id": *competitorOID}
// 	update, err := api_assets.StructToBsonMap(competitorInfoDAO)
// 	if err != nil {
// 		return err
// 	}

// 	result, err := r.competitorColl.UpdateOne(
// 		ctx,
// 		filter,
// 		bson.M{"$set": update},
// 	)
// 	if err != nil {
// 		return fmt.Errorf("error updating competitor: %w", err)
// 	}

// 	if result.MatchedCount == 0 {
// 		return fmt.Errorf("%w: no competitor found with id: %s", customerrors.ErrNotFound, competitorID)
// 	}

// 	return nil
// }

func (r *Repository) DeleteCompetitor(ctx context.Context, competitorID string) error {
	err := r.SetDeletedAt(ctx, r.competitorColl, competitorID, "competitor")
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) VerifyCompetitorExists(ctx context.Context, competitorOID *primitive.ObjectID) error {
	var result struct{}

	filter := bson.M{"_id": competitorOID}

	opts := options.FindOne().SetProjection(bson.M{"_id": 1})

	err := r.competitorColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: error when searching for competitor: %s", customerrors.ErrNotFound, err.Error())
		}
		return fmt.Errorf("error when searching for the competitor: %w", err)
	}

	return nil
}

func (r *Repository) VerifyMultipleCompetitorsExists(ctx context.Context, competitorOIDs []*primitive.ObjectID) error {
	if len(competitorOIDs) == 0 {
		return nil // No hay competidores para verificar
	}

	filter := bson.M{"_id": bson.M{"$in": competitorOIDs}}
	opts := options.Find().SetProjection(bson.M{"_id": 1})

	cursor, err := r.competitorColl.Find(ctx, filter, opts)
	if err != nil {
		return fmt.Errorf("error when searching for competitors: %w", err)
	}
	defer cursor.Close(ctx)

	var foundIDs []*primitive.ObjectID
	for cursor.Next(ctx) {
		var result struct {
			ID *primitive.ObjectID `bson:"_id"`
		}
		if err := cursor.Decode(&result); err != nil {
			return fmt.Errorf("error decoding competitor: %w", err)
		}
		foundIDs = append(foundIDs, result.ID)
	}

	if err := cursor.Err(); err != nil {
		return fmt.Errorf("error iterating cursor: %w", err)
	}

	if len(foundIDs) != len(competitorOIDs) {
		missingIDs := r.getMissingIDs(competitorOIDs, foundIDs)
		return fmt.Errorf("%w: the following competitors were not found: %v", customerrors.ErrNotFound, missingIDs)
	}

	return nil
}
func (r *Repository) GetMultipleAvailabilitiesByCompetitor(ctx context.Context, competitorOIDs []*primitive.ObjectID) ([][]*dao.GetDailyAvailabilityByIDDAORes, error) {
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"competitor_id": bson.M{"$in": competitorOIDs}}}},
		bson.D{{Key: "$project", Value: bson.M{
			"daily_availabilities": 1,
		}}},
	}

	cursor, err := r.availabilityColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error when aggregating availabilities: %w", err)
	}
	defer cursor.Close(ctx)

	var result [][]*dao.GetDailyAvailabilityByIDDAORes

	for cursor.Next(ctx) {
		var dailyAvailability dao.GetAvailabilityByIDDAORes
		if err := cursor.Decode(&dailyAvailability); err != nil {
			return nil, fmt.Errorf("error when decoding availabilities: %w", err)
		}
		result = append(result, dailyAvailability.DailyAvailabilities)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return result, nil
}


// func (r *Repository) CreateCompetitorType(ctx context.Context, competitorType models.COMPETITOR_TYPE, userID string) (string, error) {
// 	type createTypeCompetitor func(ctx context.Context) (string, error)

// 	createMap := map[models.COMPETITOR_TYPE]createTypeCompetitor{
// 		models.COMPETITOR_TYPE_SINGLE: func(ctx context.Context) (string, error) {
// 			singleDAO := &single_dao.CreateSingleDAOReq{}
// 			return r.CreateSingle(ctx, singleDAO)
// 		},
// 		models.COMPETITOR_TYPE_DOUBLE: func(ctx context.Context) (string, error) {
// 			doubleDAO := &double_dao.CreateDoubleDAOReq{}
// 			return r.CreateDouble(ctx, doubleDAO)
// 		},
// 		models.COMPETITOR_TYPE_TEAM: func(ctx context.Context) (string, error) {
// 			teamDAO := &team_dao.CreateTeamDAOReq{}
// 			// teamDAO.Admins = []
// 			return r.CreateTeam(ctx, teamDAO)
// 		},
// 	}

// 	create, ok := createMap[competitorType]
// 	if !ok {
// 		return "", fmt.Errorf("error competitor type no exists: %w", customerrors.ErrNotFound)
// 	}

// 	return create(ctx)
// }
