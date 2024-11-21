package repository

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	competitor_user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/competitor_user/dao"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/organizer/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) CreateOrganizer(ctx context.Context, userID string) error {
	userOID, err := r.ConvertToObjectID(userID)
	if err != nil {
		return err
	}

	var organizer dao.CreateOrganizerDAOReq
	organizer.UserID = *userOID
	organizer.Categories = []primitive.ObjectID{}
	organizer.Tournaments = []primitive.ObjectID{}
	organizer.SetTimeStamp()

	_, err = r.organizerColl.InsertOne(ctx, organizer)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return fmt.Errorf("%w: error organizer scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return fmt.Errorf("error when inserting organizer: %w", err)
	}

	return nil
}

// func (r *Repository) GetOrganizerByID(ctx context.Context, organizerID string) (any, error) {
// 	var organizer dao.GetLocationByIDDAORes

// 	organizerOID, err := r.ConvertToObjectID(organizerID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	filter := bson.M{"_id": organizerOID}

// 	err = r.organizerColl.FindOne(ctx, filter).Decode(&organizer)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil, fmt.Errorf("%w: error when searching for organizer: %s", customerrors.ErrNotFound, err.Error())
// 		}
// 		return nil, fmt.Errorf("error when searching for the organizer: %w", err)
// 	}

// 	return &organizer, nil
// }

func (r *Repository) VerifyOrganizerExists(ctx context.Context, organizerID string) error {
	var result struct{}

	organizerOID, err := r.ConvertToObjectID(organizerID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": organizerOID}

	opts := options.FindOne().SetProjection(bson.M{"_id": 1})

	err = r.organizerColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: error when searching for organizer: %s", customerrors.ErrNotFound, err.Error())
		}
		return fmt.Errorf("error when searching for the organizer: %w", err)
	}

	return nil
}

func (r *Repository) GetOrganizerData(ctx context.Context, userOID *primitive.ObjectID) (*dao.GetOrganizerDataDAORes, error) {
	filter := bson.M{"user_id": userOID}

	var result dao.GetOrganizerDataDAORes
	err := r.organizerColl.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for organizer: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the organizer: %w", err)
	}

	return &result, nil
}

func (r *Repository) AddCategoryInOrganizer(ctx context.Context, organizerOID, categoryOID *primitive.ObjectID) error {
	filter := bson.M{"_id": organizerOID}

	update := bson.M{
		"$push": bson.M{"categories": categoryOID},
	}

	result, err := r.organizerColl.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error updating organizer: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no opinion found with id: %s", customerrors.ErrNotFound, organizerOID.Hex())
	}

	return nil
}

func (r *Repository) AddTournamentInOrganizer(ctx context.Context, organizerOID, tournamentOID *primitive.ObjectID) error {
	filter := bson.M{"_id": organizerOID}

	update := bson.M{
		"$push": bson.M{"tournaments": tournamentOID},
	}

	result, err := r.organizerColl.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error updating organizer: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no opinion found with id: %s", customerrors.ErrNotFound, organizerOID.Hex())
	}

	return nil
}

func (r *Repository) GetSportsFromOrganizerCategories(ctx context.Context, organizerOID *primitive.ObjectID) ([]models.SPORT, error) {
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"_id": organizerOID}}},
		bson.D{{Key: "$unwind", Value: "$categories"}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "categories",
			"localField":   "categories",
			"foreignField": "_id",
			"as":           "category",
		}}},
		bson.D{{Key: "$unwind", Value: "$category"}},

		// Agrupar por sport para eliminar duplicados
		bson.D{{Key: "$group", Value: bson.M{
			"_id": "$category.sport",
		}}},

		// Proyección para devolver sólo el campo sport
		bson.D{{Key: "$project", Value: bson.M{
			"sport": "$_id",
			"_id":   0,
		}}},
	}

	cursor, err := r.organizerColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error when aggregating sports from organizer categories: %w", err)
	}
	defer cursor.Close(ctx)

	// Decodificar los resultados y convertir a slice de models.SPORT
	var sportDocs []struct {
		Sport models.SPORT `bson:"sport"`
	}
	if err = cursor.All(ctx, &sportDocs); err != nil {
		return nil, fmt.Errorf("error when decoding sports from organizer categories: %w", err)
	}

	// Extraer el campo Sport en un slice de models.SPORT
	sports := make([]models.SPORT, len(sportDocs))
	for i, doc := range sportDocs {
		sports[i] = doc.Sport
	}

	return sports, nil
}


func (r *Repository) GetCategoriesFromOrganizer(ctx context.Context, organizerOID *primitive.ObjectID, sport models.SPORT, competitorType *models.COMPETITOR_TYPE) ([]dao.GetCategoriesFromOrganizerDAORes, error) {
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"_id": organizerOID}}},
		bson.D{{Key: "$unwind", Value: "$categories"}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "categories",
			"localField":   "categories",
			"foreignField": "_id",
			"as":           "category",
		}}},
		bson.D{{Key: "$unwind", Value: bson.M{"path": "$category", "preserveNullAndEmptyArrays": true}}},
	}

	// Filtro por `sport` y `competitorType` si este último no es nil
	categoryMatch := bson.M{"category.sport": sport}
	if competitorType != nil {
		categoryMatch["category.competitor_type"] = *competitorType
	}

	// Añadir el filtro al pipeline
	pipeline = append(pipeline, bson.D{{Key: "$match", Value: categoryMatch}})

	// Continuar con el resto del pipeline
	pipeline = append(pipeline,
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "category_registrations",
			"localField":   "category._id",
			"foreignField": "category_id",
			"as":           "category_registration",
		}}},
		bson.D{{Key: "$unwind", Value: bson.M{"path": "$category_registration", "preserveNullAndEmptyArrays": true}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "competitor_users",
			"localField":   "category_registration.competitor_id",
			"foreignField": "competitor_id",
			"as":           "competitor_user",
		}}},
		bson.D{{Key: "$unwind", Value: bson.M{"path": "$competitor_user", "preserveNullAndEmptyArrays": true}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "users",
			"localField":   "competitor_user.user_id",
			"foreignField": "_id",
			"as":           "user",
		}}},
		bson.D{{Key: "$unwind", Value: bson.M{"path": "$user", "preserveNullAndEmptyArrays": true}}},
		bson.D{{Key: "$group", Value: bson.M{
			"_id": bson.M{
				"category_id":   "$category._id",
				"competitor_id": "$competitor_user.competitor_id",
			},
			"category_id":          bson.M{"$first": "$category._id"},
			"name":                 bson.M{"$first": "$category.name"},
			"competitor_id":        bson.M{"$first": "$competitor_user.competitor_id"},
			"total_participants":   bson.M{"$first": "$category.total_participants"},
			"points":               bson.M{"$first": "$category_registration.points"},
			"current_position":     bson.M{"$first": "$category_registration.current_position"},
			"registered_positions": bson.M{"$first": "$category_registration.registered_positions"},
			"users": bson.M{
				"$push": bson.M{
					"_id":        "$user._id",
					"first_name": "$user.first_name",
					"last_name":  "$user.last_name",
					"image":      "$user.image",
				},
			},
		}}},
		bson.D{{Key: "$sort", Value: bson.M{"current_position": 1}}},
		bson.D{{Key: "$group", Value: bson.M{
			"_id":                "$category_id",
			"category_id":        bson.M{"$first": "$category_id"},
			"total_participants": bson.M{"$first": "$total_participants"},
			"name":               bson.M{"$first": "$name"},
			"competitors": bson.M{
				"$push": bson.M{
					"$cond": bson.M{
						"if": bson.M{"$ne": []interface{}{"$competitor_id", nil}},
						"then": bson.M{
							"_id":                  "$competitor_id",
							"points":               "$points",
							"current_position":     "$current_position",
							"registered_positions": "$registered_positions",
							"users":                "$users",
						},
						"else": nil,
					},
				},
			},
		}}},
		bson.D{{Key: "$project", Value: bson.M{
			"_id":   "$category_id",
			"total": "$total_participants",
			"name":  "$name",
			"competitors": bson.M{
				"$filter": bson.M{
					"input": "$competitors",
					"as":    "competitor",
					"cond":  bson.M{"$ne": []interface{}{"$$competitor", nil}},
				},
			},
		}}},
	)

	// Ejecutar el pipeline de agregación
	cursor, err := r.organizerColl.Aggregate(ctx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for categoryRegistration: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the categoryRegistration: %w", err)
	}

	defer cursor.Close(ctx)

	var categoriesDAO []dao.GetCategoriesFromOrganizerDAORes
	if err = cursor.All(ctx, &categoriesDAO); err != nil {
		return nil, fmt.Errorf("error when decoding categoryRegistration: %w", err)
	}

	return categoriesDAO, nil
}



func (r *Repository) GetTournamentSportsInOrganizer(ctx context.Context, organizerOID *primitive.ObjectID) ([]models.SPORT, error) {
	// Pipeline para obtener deportes únicos de los torneos del organizador
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{"_id": organizerOID}}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "tournaments",
			"localField":   "tournaments",
			"foreignField": "_id",
			"as":           "tournaments",
		}}},
		{{Key: "$unwind", Value: bson.M{"path": "$tournaments", "preserveNullAndEmptyArrays": true}}},
		{{Key: "$group", Value: bson.M{
			"_id":    nil,
			"sports": bson.M{"$addToSet": "$tournaments.sport"}, // Usamos $addToSet para evitar duplicados
		}}},
		{{Key: "$project", Value: bson.M{"sports": 1, "_id": 0}}},
	}

	cursor, err := r.organizerColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error when aggregating unique sports: %w", err)
	}
	defer cursor.Close(ctx)

	// Decodificamos el resultado para obtener los deportes únicos
	var result []bson.M
	if err := cursor.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("error when decoding unique sports: %w", err)
	}

	// Procesamos el resultado y convertimos cada deporte en el tipo SPORT
	var uniqueSports []models.SPORT // Siempre inicializado como un slice vacío
	if len(result) > 0 {
		for _, sport := range result[0]["sports"].(bson.A) {
			if sportStr, ok := sport.(string); ok {
				if parsedSport, err := models.ParseSport(sportStr); err == nil {
					uniqueSports = append(uniqueSports, parsedSport)
				}
			}
		}
	}
fmt.Printf("nose: %+v", uniqueSports)
	// Retornamos el slice, ya sea con datos o vacío
	return uniqueSports, nil
}



func (r *Repository) GetTournamentsInOrganizer(
	ctx context.Context,
	organizerOID *primitive.ObjectID,
	sport models.SPORT,
	limit int,
	lastOID *primitive.ObjectID,
) (*competitor_user_dao.GetUserTournamentsDAORes, error) {
	// Pipeline para contar el total de torneos
	totalPipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{"_id": organizerOID}}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "tournaments",
			"localField":   "tournaments",
			"foreignField": "_id",
			"as":           "tournaments",
		}}},
		{{Key: "$unwind", Value: "$tournaments"}},
		{{Key: "$match", Value: bson.M{"tournaments.sport": sport}}},
		{{Key: "$count", Value: "total"}},
	}

	totalCursor, err := r.organizerColl.Aggregate(ctx, totalPipeline)
	if err != nil {
		return nil, fmt.Errorf("error when aggregating total tournaments: %w", err)
	}
	defer totalCursor.Close(ctx)

	var totalResult []bson.M
	if err := totalCursor.All(ctx, &totalResult); err != nil {
		return nil, fmt.Errorf("error when decoding total tournaments: %w", err)
	}

	total := 0
	if len(totalResult) > 0 {
		total = int(totalResult[0]["total"].(int32))
	}

	// Pipeline para obtener los torneos con detalles
	// Pipeline para obtener los torneos con detalles
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{"_id": organizerOID}}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "tournaments",
			"localField":   "tournaments",
			"foreignField": "_id",
			"as":           "tournaments",
		}}},
		{{Key: "$unwind", Value: "$tournaments"}},
		{{Key: "$match", Value: bson.M{"tournaments.sport": sport}}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "locations",
			"localField":   "tournaments.location_id",
			"foreignField": "_id",
			"as":           "location",
		}}},
		{{Key: "$unwind", Value: "$location"}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "users",
			"localField":   "user_id",
			"foreignField": "_id",
			"as":           "user",
		}}},
		{{Key: "$unwind", Value: "$user"}},

		{{Key: "$sort", Value: bson.M{"tournaments._id": -1}}}, // Ordena por fecha de creación descendente (los más recientes primero)

		{{Key: "$project", Value: bson.M{
			"_id":           "$tournaments._id",
			"name":          "$tournaments.name",
			"start_date":    "$tournaments.start_date",
			"image":         "$tournaments.image",
			"finish_date":   "$tournaments.finish_date",
			"points":        "$tournaments.points",
			"average_score": "$tournaments.average_score",
			"total_prize":   "$tournaments.total_prize",
			"location": bson.M{
				"_id":     "$location._id",
				"state":   "$location.state",
				"country": "$location.country",
				"city":    "$location.city",
				"lat":     "$location.lat",
				"long":    "$location.long",
			},
			"organizer": bson.M{
				"_id":        "$user._id",
				"first_name": "$user.first_name",
				"last_name":  "$user.last_name",
			},
		}}},
	}

	// Si tienes lastOID, filtra por este campo antes de aplicar el sort
	if lastOID != nil {
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.M{"tournaments._id": bson.M{"$lt": lastOID}}}})
	}

	// Limita los resultados
	pipeline = append(pipeline, bson.D{{Key: "$limit", Value: limit}})

	cursor, err := r.organizerColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error when aggregating organizer tournaments: %w", err)
	}
	defer cursor.Close(ctx)

	var tournaments []*competitor_user_dao.GetUserTournamentDAORes
	if err := cursor.All(ctx, &tournaments); err != nil {
		return nil, fmt.Errorf("error when decoding organizer tournaments: %w", err)
	}

	result := &competitor_user_dao.GetUserTournamentsDAORes{
		Tournaments: tournaments,
		Total:       total,
	}
	for i, v := range result.Tournaments {
		fmt.Printf("estooo %v %+v", i, v.ID)
	}
	return result, nil
}

func (r *Repository) GetOrganizerIDByUserID(ctx context.Context, userOID *primitive.ObjectID) (*string, error) {
	filter := bson.M{
		"user_id": userOID,
	}

	var result struct {
		ID *primitive.ObjectID `bson:"_id"`
	}

	err := r.organizerColl.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Retornar nil en lugar de una cadena vacía si no se encuentra el documento
			return nil, nil
		}
		return nil, fmt.Errorf("error when searching for the user: %w", err)
	}

	// Convertir el ObjectID en string y retornar el puntero
	organizerID := result.ID.Hex()
	return &organizerID, nil
}
