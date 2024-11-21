package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	api_assets "github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/category/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) CreateCategory(ctx context.Context, organizerID string, categoryInfoDAO *dao.CreateCategoryDAOReq) (string, error) {
	organizerOID, err := r.ConvertToObjectID(organizerID)
	if err != nil {
		return "", err
	}

	categoryInfoDAO.OrganizerID = *organizerOID
	categoryInfoDAO.Tournaments = []primitive.ObjectID{}

	categoryInfoDAO.SetTimeStamp()

	result, err := r.categoryColl.InsertOne(ctx, categoryInfoDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for category: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error category scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting category: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetCategoryByID(ctx context.Context, categoryID string) (*dao.GetCategoryByIDDAORes, error) {
	var category dao.GetCategoryByIDDAORes

	categoryOID, err := r.ConvertToObjectID(categoryID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *categoryOID}

	err = r.categoryColl.FindOne(ctx, filter).Decode(&category)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for category: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the category: %w", err)
	}

	return &category, nil
}

func (r *Repository) UpdateCategory(ctx context.Context, categoryOID *primitive.ObjectID, categoryInfoDAO *dao.UpdateCategoryDAOReq) error {
	categoryInfoDAO.RenewUpdate()

	filter := bson.M{"_id": *categoryOID}
	update, err := api_assets.StructToBsonMap(categoryInfoDAO)
	if err != nil {
		return err
	}

	result, err := r.categoryColl.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": update},
	)
	if err != nil {
		return fmt.Errorf("error updating category: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no category found with id: %s", customerrors.ErrNotFound, categoryOID.Hex())
	}

	return nil
}

func (r *Repository) DeleteCategory(ctx context.Context, categoryID string) error {
	err := r.SetDeletedAt(ctx, r.categoryColl, categoryID, "category")
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) OrganizeCategory(ctx context.Context, organizerID string, categoryInfoDAO *dao.CreateCategoryDAOReq) error {
	if err := r.VerifyOrganizerExists(ctx, organizerID); err != nil {
		return err
	}

	if _, err := r.CreateCategory(ctx, organizerID, categoryInfoDAO); err != nil {
		return err
	}

	return nil
}

func (r *Repository) AddTournamentInCategory(ctx context.Context, categoryID string, tournamentID string) error {
	categoryOID, err := r.ConvertToObjectID(categoryID)
	if err != nil {
		return err
	}

	tournamentOID, err := r.ConvertToObjectID(tournamentID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": *categoryOID}

	update := bson.M{"tournaments": tournamentOID}

	currentDate := time.Now().UTC()
	updatedAt := bson.M{"updated_at": currentDate}

	result, err := r.categoryColl.UpdateOne(
		ctx,
		filter,
		bson.M{
			"$addToSet": update,
			"$set":      updatedAt,
		},
	)
	if err != nil {
		return fmt.Errorf("error updating category: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no category found with id: %s", customerrors.ErrNotFound, categoryID)
	}

	return r.AddCategoryInTournament(ctx, tournamentID, categoryID)

}

func (r *Repository) VerifyCategoryExists(ctx context.Context, categoryID string) error {
	var result struct{}

	categoryOID, err := r.ConvertToObjectID(categoryID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": categoryOID}

	opts := options.FindOne().SetProjection(bson.M{"_id": 1})

	err = r.categoryColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: error when searching for category: %s", customerrors.ErrNotFound, err.Error())
		}
		return fmt.Errorf("error when searching for the category: %w", err)
	}

	return nil
}

func (r *Repository) GetCategoryInfoByID(ctx context.Context, categoryOID *primitive.ObjectID) (*dao.GetCategoryInfoByIDDAORes, error) {
	pipeline := mongo.Pipeline{
		bson.D{{
			Key: "$match", Value: bson.M{
				"_id": categoryOID,
			},
		}},
		bson.D{{
			Key: "$lookup", Value: bson.M{
				"from":         "organizers",
				"localField":   "organizer_id",
				"foreignField": "_id",
				"as":           "organizer",
			},
		}},
		bson.D{{Key: "$unwind", Value: "$organizer"}},
		bson.D{{
			Key: "$lookup", Value: bson.M{
				"from":         "users",
				"localField":   "organizer.user_id",
				"foreignField": "_id",
				"as":           "user",
			},
		}},
		bson.D{{Key: "$unwind", Value: "$user"}},
		bson.D{{
			Key: "$project", Value: bson.M{
				"_id":                1,
				"name":               1,
				"genre":              1,
				"total_participants": 1,
				"range_movement":     1,
				"sport":              1,
				"competitor_type":    1,
				"organizer": bson.M{
					"_id":        "$organizer._id",
					"first_name": "$user.first_name",
					"last_name":  "$user.last_name",
				},
			},
		}},
	}

	cursor, err := r.categoryColl.Aggregate(ctx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for category: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the category: %w", err)
	}

	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		var categoryInfoDAO dao.GetCategoryInfoByIDDAORes
		if err := cursor.Decode(&categoryInfoDAO); err != nil {
			return nil, fmt.Errorf("error when decoding category: %w", err)
		}

		return &categoryInfoDAO, nil
	}

	return nil, fmt.Errorf("%w: category not found", customerrors.ErrNotFound)
}

func (r *Repository) IncrementTotalParticipants(ctx context.Context, categoryOID *primitive.ObjectID) error {
	filter := bson.M{"_id": categoryOID}

	update := bson.M{
		"$inc": bson.M{
			"total_participants": 1,
		},
	}

	result, err := r.categoryColl.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return fmt.Errorf("error updating category: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no category found with id: %s", customerrors.ErrNotFound, categoryOID.Hex())
	}

	return nil
}

func (r *Repository) DecrementTotalParticipants(ctx context.Context, categoryOID *primitive.ObjectID) error {
	filter := bson.M{"_id": categoryOID}

	update := bson.M{
		"$inc": bson.M{
			"total_participants": -1,
		},
	}

	result, err := r.categoryColl.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return fmt.Errorf("error updating category: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no category found with id: %s", customerrors.ErrNotFound, categoryOID.Hex())
	}

	return nil
}

func (r *Repository) GetTournamentsFromCategory(ctx context.Context, categoryOID *primitive.ObjectID, sport models.SPORT, competitorType models.COMPETITOR_TYPE, limit int, lastOID *primitive.ObjectID) ([]*dao.GetTournamentsFromCategoryDAORes, error) {
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"_id": categoryOID}}},
		bson.D{{
			Key: "$lookup",
			Value: bson.M{
				"from":         "tournaments",
				"localField":   "tournaments",
				"foreignField": "_id",
				"as":           "tournament",
			},
		}},
		bson.D{{Key: "$unwind", Value: "$tournament"}},
		bson.D{{Key: "$match", Value: bson.M{
			"tournament.sport":           sport,
			"tournament.competitor_type": competitorType,
		}}},
		bson.D{{
			Key: "$lookup",
			Value: bson.M{
				"from":         "locations",
				"localField":   "tournament.location_id",
				"foreignField": "_id",
				"as":           "location",
			},
		}},
		bson.D{{Key: "$unwind", Value: "$location"}},
	}

	if lastOID != nil {
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.M{
			"tournament._id": bson.M{"$gt": lastOID},
		}}})
	}

	pipeline = append(pipeline,
		bson.D{{Key: "$limit", Value: limit}},
		bson.D{{Key: "$group", Value: bson.M{
			"_id": "$_id",
			"tournaments": bson.M{
				"$push": bson.M{
					"_id":         "$tournament._id",
					"name":        "$tournament.name",
					"points":      "$tournament.points",
					"location": bson.M{
						"state":"$location.state",
						"country":"$location.country",
						"city":"$location.city",
						"lat":"$location.lat",
						"long":"$location.long",
					},
					"total_prize": "$tournament.total_prize",
					"average_score": "$tournament.average_score",
					"start_date":  "$tournament.start_date",
					"finish_date": "$tournament.finish_date",
				},
			},
		}}},
		bson.D{{Key: "$project", Value: bson.M{
			"_id":         0,
			"tournaments": 1,
		}}},
	)

	cursor, err := r.categoryColl.Aggregate(ctx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for category: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the category: %w", err)
	}
	defer cursor.Close(ctx)

	var result []struct {
		Tournaments []*dao.GetTournamentsFromCategoryDAORes `bson:"tournaments"`
	}
	if err = cursor.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("error when decoding category: %w", err)
	}

	if len(result) == 0 {
		return []*dao.GetTournamentsFromCategoryDAORes{}, nil
	}

	return result[0].Tournaments, nil
}

func (r *Repository) GetTournamentsByNameFromCategory(ctx context.Context, categoryOID *primitive.ObjectID, sport models.SPORT, competitorType models.COMPETITOR_TYPE, tournamentName string) ([]*dao.GetTournamentsFromCategoryDAORes, error) {
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"_id": categoryOID}}},
		bson.D{{
			Key: "$lookup",
			Value: bson.M{
				"from":         "tournaments",
				"localField":   "tournaments",
				"foreignField": "_id",
				"as":           "tournament",
			},
		}},
		bson.D{{Key: "$unwind", Value: "$tournament"}},
		bson.D{{Key: "$match", Value: bson.M{
			"tournament.sport":           sport,
			"tournament.competitor_type": competitorType,
			"tournament.name":            bson.M{"$regex": "^" + tournamentName, "$options": "i"}, // Búsqueda que empieza con el nombre (insensible a mayúsculas)
		}}},
		bson.D{{
			Key: "$lookup",
			Value: bson.M{
				"from":         "locations",
				"localField":   "tournament.location_id",
				"foreignField": "_id",
				"as":           "location",
			},
		}},
		bson.D{{Key: "$unwind", Value: "$location"}},
		bson.D{{Key: "$limit", Value: 10}}, // Limitar a 10 resultados
		bson.D{{Key: "$group", Value: bson.M{
			"_id": "$_id",
			"tournaments": bson.M{
				"$push": bson.M{
					"_id":          "$tournament._id",
					"name":         "$tournament.name",
					"points":       "$tournament.points",
					"location": bson.M{
						"state":   "$location.state",
						"country": "$location.country",
						"city":    "$location.city",
						"lat":     "$location.lat",
						"long":    "$location.long",
					},
					"total_prize":    "$tournament.total_prize",
					"average_score":  "$tournament.average_score",
					"start_date":     "$tournament.start_date",
					"finish_date":    "$tournament.finish_date",
				},
			},
		}}},
		bson.D{{Key: "$project", Value: bson.M{
			"_id":         0,
			"tournaments": 1,
		}}},
	}

	cursor, err := r.categoryColl.Aggregate(ctx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for category: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the category: %w", err)
	}
	defer cursor.Close(ctx)

	var result []struct {
		Tournaments []*dao.GetTournamentsFromCategoryDAORes `bson:"tournaments"`
	}
	if err = cursor.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("error when decoding category: %w", err)
	}

	if len(result) == 0 {
		return []*dao.GetTournamentsFromCategoryDAORes{}, nil
	}

	return result[0].Tournaments, nil
}


func (r *Repository) GetTournamentsFromCategoryNumber(ctx context.Context, categoryOID *primitive.ObjectID, sport models.SPORT, competitorType models.COMPETITOR_TYPE) (int, error) {
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"_id": categoryOID}}},
		bson.D{{
			Key: "$lookup",
			Value: bson.M{
				"from":         "tournaments",
				"localField":   "tournaments",
				"foreignField": "_id",
				"as":           "tournament",
			},
		}},
		bson.D{{Key: "$unwind", Value: "$tournament"}},
		bson.D{{Key: "$match", Value: bson.M{
			"tournament.sport":           sport,
			"tournament.competitor_type": competitorType,
		}}},
	}

	// Añadimos la etapa de $count para obtener la cantidad de documentos que cumplen los filtros
	pipeline = append(pipeline, bson.D{{Key: "$count", Value: "tournamentCount"}})

	cursor, err := r.categoryColl.Aggregate(ctx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, fmt.Errorf("%w: error when searching for category: %s", customerrors.ErrNotFound, err.Error())
		}
		return 0, fmt.Errorf("error when searching for the category: %w", err)
	}
	defer cursor.Close(ctx)

	var result []struct {
		TournamentCount int `bson:"tournamentCount"`
	}
	if err = cursor.All(ctx, &result); err != nil {
		return 0, fmt.Errorf("error when decoding category count: %w", err)
	}

	if len(result) == 0 {
		return 0, nil
	}

	return result[0].TournamentCount, nil
}


func (r *Repository) GetCompetitorTournamentsInCategory(ctx context.Context, categoryOID, competitorOID, lastOID *primitive.ObjectID, limit int) ([]*dao.GetTournamentsFromCategoryDAORes, error) {
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"_id": categoryOID}}},
		bson.D{{
			Key: "$lookup", Value: bson.M{
				"from":         "tournaments",
				"localField":   "tournaments",
				"foreignField": "_id",
				"as":           "tournament",
			},
		}},
		bson.D{{Key: "$unwind", Value: "$tournament"}},
		bson.D{{
			Key: "$lookup", Value: bson.M{
				"from":         "tournament_registrations",
				"localField":   "tournament._id",
				"foreignField": "tournament_id",
				"as":           "tournament_registration",
			},
		}},
		bson.D{{Key: "$unwind", Value: "$tournament_registration"}},
		bson.D{{Key: "$match", Value: bson.M{
			"tournament_registration.competitor_id": competitorOID,
		}}},
	}

	if lastOID != nil {
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.M{
			"tournament._id": bson.M{"$gt": lastOID},
		}}})
	}

	pipeline = append(pipeline,
		bson.D{{Key: "$limit", Value: limit}},
		bson.D{{Key: "$group", Value: bson.M{
			"_id": "$_id",
			"tournaments": bson.M{
				"$push": bson.M{
					"_id":         "$tournament._id",
					"name":        "$tournament.name",
					"points":      "$tournament.points",
					"start_date":  "$tournament.start_date",
					"finish_date": "$tournament.finish_date",
				},
			},
		}}},
		bson.D{{Key: "$project", Value: bson.M{
			"_id":         0,
			"tournaments": 1,
		}}},
	)

	cursor, err := r.categoryColl.Aggregate(ctx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for category: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the category: %w", err)
	}
	defer cursor.Close(ctx)

	var result []struct {
		Tournaments []*dao.GetTournamentsFromCategoryDAORes `bson:"tournaments"`
	}

	if err = cursor.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("error when decoding category: %w", err)
	}

	if len(result) == 0 {
		return []*dao.GetTournamentsFromCategoryDAORes{}, nil
	}

	return result[0].Tournaments, nil

}
