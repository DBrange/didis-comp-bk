package repository

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/tournament_registration/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) CreateTournamentRegistration(ctx context.Context, tournamentRegistrationInfoDAO *dao.CreateTournamentRegistrationDAOReq) error {
	if err := r.VerifyCompetitorAlreadyResgisteredInTournament(ctx, tournamentRegistrationInfoDAO); err != nil {
		return err
	}

	tournamentRegistrationInfoDAO.SetTimeStamp()

	_, err := r.tournamentRegistrationColl.InsertOne(ctx, tournamentRegistrationInfoDAO)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return fmt.Errorf("%w: error tournamentRegistration scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return fmt.Errorf("error when inserting tournamentRegistration: %w", err)
	}

	return nil
}

func (r *Repository) GetTournamentRegistrationByID(ctx context.Context, tournamentRegistrationID string) (*dao.GetTournamentRegistrationByIDDAORes, error) {
	var tournamentRegistration dao.GetTournamentRegistrationByIDDAORes

	tournamentRegistrationOID, err := r.ConvertToObjectID(tournamentRegistrationID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *tournamentRegistrationOID}

	err = r.tournamentRegistrationColl.FindOne(ctx, filter).Decode(&tournamentRegistration)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for tournamentRegistration: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the tournamentRegistration: %w", err)
	}

	return &tournamentRegistration, nil
}

func (r *Repository) DeleteTournamentRegistration(ctx context.Context, tournamentRegistrationID string) error {
	err := r.SetDeletedAt(ctx, r.tournamentRegistrationColl, tournamentRegistrationID, "tournamentRegistration")
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) VerifyCompetitorAlreadyResgisteredInTournament(ctx context.Context, tournamentRegistrationInfoDAO *dao.CreateTournamentRegistrationDAOReq) error {
	filter := bson.M{
		"tournament_id": tournamentRegistrationInfoDAO.TournamentID,
		"competitor_id": tournamentRegistrationInfoDAO.CompetitorID,
		// "$or": []bson.M{
		// 	{"deleted_at": bson.M{"$exists": false}},
		// 	{"deleted_at": nil},
		// },
	}

	var documentFinded *dao.CreateTournamentRegistrationDAOReq

	err := r.tournamentRegistrationColl.FindOne(ctx, filter).Decode(&documentFinded)
	if err == nil {
		return fmt.Errorf("error relation in tournamentRegistration already exists: %w", err)
	}

	return nil
}

func (r *Repository) GetAllCompetitorInTournament(ctx context.Context, tournamentOID *primitive.ObjectID, limit, page int) ([]dao.GetTournamentRegistrationByIDDAORes, error) {
	filter := bson.M{"tournament_id": *tournamentOID}

	skip := (page - 1) * limit

	opts := options.Find().
		SetLimit(int64(limit)).
		SetSkip(int64(skip))

	cursor, err := r.tournamentRegistrationColl.Find(ctx, filter, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for tournamentRegistration: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the tournamentRegistration: %w", err)
	}

	defer cursor.Close(ctx)

	var tournamentRegistrations []dao.GetTournamentRegistrationByIDDAORes

	err = cursor.All(ctx, &tournamentRegistrations)
	if err != nil {
		return nil, fmt.Errorf("error when decoding tournamentRegistrations: %w", err)
	}

	if len(tournamentRegistrations) == 0 {
		return nil, fmt.Errorf("%w: no tournament registrations found for tournament ID: %s on page %d", customerrors.ErrNotFound, tournamentOID.Hex(), page)
	}

	return tournamentRegistrations, nil
}

func (r *Repository) GetCompetitorsInTournament(ctx context.Context, tournamentOID, categoryOID, lastOID *primitive.ObjectID, limit int) ([]*dao.GetCompetitorsInTournamentDAORes, error) {
	pipeline := r.getCompetitorsInTournamentBuildBasePipeline(tournamentOID)

	if categoryOID != nil {
		pipeline = r.getCompetitorsInTournamentAppendCategoryFilter(pipeline, categoryOID)
	}

	pipeline = r.getCompetitorsInTournamentAppendFinalStages(pipeline, categoryOID != nil)
	pipeline = r.getCompetitorsInTournamentAppendSortAndPagination(pipeline, lastOID, limit)

	return r.getCompetitorsInTournamentExecuteAggregation(ctx, pipeline)
}

func (r *Repository) getCompetitorsInTournamentBuildBasePipeline(tournamentOID *primitive.ObjectID) mongo.Pipeline {
	return mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"tournament_id": tournamentOID}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "competitor_users",
			"localField":   "competitor_id",
			"foreignField": "competitor_id",
			"as":           "competitor_user",
		}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "guest_competitors",
			"localField":   "competitor_id",
			"foreignField": "competitor_id",
			"as":           "guest_competitor",
		}}},
		bson.D{{Key: "$addFields", Value: bson.M{
			"is_guest": bson.M{"$gt": bson.A{bson.M{"$size": "$guest_competitor"}, 0}},
		}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "users",
			"localField":   "competitor_user.user_id",
			"foreignField": "_id",
			"as":           "user",
		}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "guest_users",
			"localField":   "guest_competitor.guest_user_id",
			"foreignField": "_id",
			"as":           "guest_user",
		}}},
	}
}

func (r *Repository) getCompetitorsInTournamentAppendCategoryFilter(pipeline mongo.Pipeline, categoryOID *primitive.ObjectID) mongo.Pipeline {
	return append(pipeline,
		bson.D{{Key: "$lookup", Value: bson.M{
			"from": "category_registrations",
			"let": bson.M{
				"competitor_id": "$competitor_id",
			},
			"pipeline": bson.A{
				bson.D{{Key: "$match", Value: bson.M{
					"$expr": bson.M{
						"$and": bson.A{
							bson.M{"$eq": bson.A{"$competitor_id", "$$competitor_id"}},
							bson.M{"$eq": bson.A{"$category_id", categoryOID}},
						},
					},
				}}},
			},
			"as": "category_registration",
		}}},
		bson.D{{Key: "$unwind", Value: bson.M{"path": "$category_registration", "preserveNullAndEmptyArrays": true}}},
	)
}

func (r *Repository) getCompetitorsInTournamentAppendSortAndPagination(pipeline mongo.Pipeline, lastOID *primitive.ObjectID, limit int) mongo.Pipeline {
	// Agregar etapa de ordenamiento
	sortStage := bson.D{{Key: "$sort", Value: bson.D{
		{Key: "_id", Value: 1}, // Siempre ordenar por _id para asegurar un orden consistente
	}}}
	pipeline = append(pipeline, sortStage)

	// Aplicar paginación
	if lastOID != nil {
		matchStage := bson.D{{Key: "$match", Value: bson.M{
			"_id": bson.M{"$gt": lastOID},
		}}}
		pipeline = append(pipeline, matchStage)
	}

	// Aplicar límite
	limitStage := bson.D{{Key: "$limit", Value: limit}}
	pipeline = append(pipeline, limitStage)

	return pipeline
}

func (r *Repository) getCompetitorsInTournamentAppendFinalStages(pipeline mongo.Pipeline, includeCurrentPosition bool) mongo.Pipeline {
	groupStage := bson.D{{Key: "$group", Value: bson.M{
		"_id":      "$competitor_id",
		"is_guest": bson.M{"$first": "$is_guest"},
		"users": bson.M{
			"$push": bson.M{
				"$cond": bson.A{
					bson.M{"$eq": bson.A{"$is_guest", false}},
					"$user",
					bson.M{"$literal": bson.A{}},
				},
			},
		},
		"guest_users": bson.M{
			"$push": bson.M{
				"$cond": bson.A{
					bson.M{"$eq": bson.A{"$is_guest", true}},
					"$guest_user",
					bson.M{"$literal": bson.A{}},
				},
			},
		},
	}}}

	projectStage := bson.D{{Key: "$project", Value: bson.M{
		"_id": "$_id",
		"users": bson.M{
			"$reduce": bson.M{
				"input":        "$users",
				"initialValue": bson.A{},
				"in": bson.M{
					"$concatArrays": bson.A{"$$value", "$$this"},
				},
			},
		},
		"guest_users": bson.M{
			"$reduce": bson.M{
				"input":        "$guest_users",
				"initialValue": bson.A{},
				"in": bson.M{
					"$concatArrays": bson.A{"$$value", "$$this"},
				},
			},
		},
	}}}

	if includeCurrentPosition {
		groupStage[0].Value.(bson.M)["current_position"] = bson.M{"$first": "$category_registration.current_position"}
		projectStage[0].Value.(bson.M)["current_position"] = "$current_position"
	}

	return append(pipeline, groupStage, projectStage)
}

func (r *Repository) getCompetitorsInTournamentExecuteAggregation(ctx context.Context, pipeline mongo.Pipeline) ([]*dao.GetCompetitorsInTournamentDAORes, error) {
	var competitorsDAO []*dao.GetCompetitorsInTournamentDAORes

	cursor, err := r.tournamentRegistrationColl.Aggregate(ctx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for tournamentRegistration: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the tournamentRegistration: %w", err)
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &competitorsDAO); err != nil {
		return nil, fmt.Errorf("error when decoding tournamentRegistration: %w", err)
	}

	return competitorsDAO, nil
}
