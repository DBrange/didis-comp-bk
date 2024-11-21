package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	availability_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
	competitor_user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/competitor_user/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) CreateCompetitorUser(ctx context.Context, userOID *primitive.ObjectID, competitorOID *primitive.ObjectID) error {
	competitorUserDAO := &competitor_user_dao.CreateCompetitorUserDAOReq{}
	competitorUserDAO.CompetitorID = *competitorOID
	competitorUserDAO.UserID = *userOID

	competitorUserDAO.SetTimeStamp()

	_, err := r.competitorUserColl.InsertOne(ctx, competitorUserDAO)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return fmt.Errorf("%w: error competitorUser scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return fmt.Errorf("error when inserting competitorUser: %w", err)
	}

	return nil
}

func (r *Repository) GetCompetitorUserByID(ctx context.Context, competitorUserID string) (*competitor_user_dao.GetCompetitorUserByIDDAORes, error) {
	var competitorUser competitor_user_dao.GetCompetitorUserByIDDAORes

	competitorUserOID, err := r.ConvertToObjectID(competitorUserID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *competitorUserOID}

	err = r.competitorUserColl.FindOne(ctx, filter).Decode(&competitorUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for competitorUser: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the competitorUser: %w", err)
	}

	return &competitorUser, nil
}

// func (r *Repository) UpdateCompetitorUser(ctx context.Context, competitorUserID string, competitorUserInfoDAO *competitor_user_dao.UpdateCompetitorUserDAOReq) error {
// 	competitorUserOID, err := r.ConvertToObjectID(competitorUserID)
// 	if err != nil {
// 		return err
// 	}

// 	competitorUserInfoDAO.RenewUpdate()

// 	filter := bson.M{"_id": *competitorUserOID}
// 	update, err := api_assets.StructToBsonMap(competitorUserInfoDAO)
// 	if err != nil {
// 		return err
// 	}

// 	result, err := r.competitorUserColl.UpdateOne(
// 		ctx,
// 		filter,
// 		bson.M{"$set": update},
// 	)
// 	if err != nil {
// 		return fmt.Errorf("error updating competitorUser: %w", err)
// 	}

// 	if result.MatchedCount == 0 {
// 		return fmt.Errorf("%w: no competitorUser found with id: %s", customerrors.ErrNotFound, competitorUserID)
// 	}

// 	return nil
// }

// func (r *Repository) DeleteCompetitorUser(ctx context.Context, competitorUserID string) error {
// 	err := r.SetDeletedAt(ctx, r.competitorUserColl, competitorUserID, "competitorUser")
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (r *Repository) VerifyCompetitorIDInCompetitorUser(ctx context.Context, competitorIDs []*primitive.ObjectID) (bool, error) {
	// Crear el filtro para buscar ambos IDs en la colección
	filter := bson.M{"competitor_id": bson.M{"$in": competitorIDs}}

	// Contar cuántos documentos coinciden con los IDs proporcionados
	count, err := r.userColl.CountDocuments(ctx, filter)
	if err != nil {
		return false, fmt.Errorf("error when counting competitor users: %w", err)
	}

	// Si el número de documentos encontrados es igual al número de IDs proporcionados, ambos existen
	if count == int64(len(competitorIDs)) {
		return true, nil
	}

	// Si no, falta al menos uno de los IDs
	return false, nil
}

func (r *Repository) GetUserCategories(
	ctx context.Context,
	userOID *primitive.ObjectID,
	sport models.SPORT,
	limit int,
	lastOID *primitive.ObjectID,
) ([]*competitor_user_dao.GetUserCategoriesCategoryDAO, error) {
	// a, e :=r.GetUserTournaments(ctx, userOID, sport, limit, lastOID)
	// fmt.Printf("ni idea breo: %+v",a[0])
	// fmt.Printf("ni idea breo: %+v",a[0].Location)
	// fmt.Printf("ni idea breo: %+v",a[0].Organizer)
	// fmt.Printf("ni idea breo: %+v",e)

	// Pipeline de agregación
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"user_id": userOID}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "category_registrations",
			"localField":   "competitor_id",
			"foreignField": "competitor_id",
			"as":           "category_registrations",
		}}},
		bson.D{{Key: "$unwind", Value: "$category_registrations"}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "categories",
			"localField":   "category_registrations.category_id",
			"foreignField": "_id",
			"as":           "category",
		}}},
		bson.D{{Key: "$unwind", Value: "$category"}},
		bson.D{{Key: "$match", Value: bson.M{"category.sport": sport}}},
	}

	// Si `lastOID` no es nil, agregar el filtro por `_id` mayor que `lastOID`
	if lastOID != nil {
		pipeline = append(pipeline, bson.D{
			{Key: "$match", Value: bson.M{
				"category._id": bson.M{"$gt": lastOID},
			}},
		})
	}

	// Aplicar el límite después de filtrar por `sport`
	if limit > 0 {
		pipeline = append(pipeline, bson.D{{Key: "$limit", Value: limit}})
	}

	// Continuación del pipeline con lookup, unwind, etc.
	pipeline = append(pipeline,
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "competitor_users",
			"localField":   "competitor_id",
			"foreignField": "competitor_id",
			"as":           "competitor_users",
		}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "users",
			"localField":   "competitor_users.user_id",
			"foreignField": "_id",
			"as":           "users",
		}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "organizers",
			"localField":   "category.organizer_id",
			"foreignField": "_id",
			"as":           "organizer",
		}}},
		bson.D{{Key: "$unwind", Value: "$organizer"}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "users",
			"localField":   "organizer.user_id",
			"foreignField": "_id",
			"as":           "organizer_user",
		}}},
		bson.D{{Key: "$unwind", Value: "$organizer_user"}},
		bson.D{{Key: "$group", Value: bson.M{
			"_id": "$user_id",
			"categories": bson.M{
				"$push": bson.M{
					"_id":                 "$category._id",
					"name":                "$category.name",
					"sport":               "$category.sport",
					"competitor_type":     "$category.competitor_type",
					"average_score":       "$category.average_score",
					"tournament_quantity": bson.M{"$size": "$category.tournaments"},
					"competitor_data": bson.M{
						"points":          "$category_registrations.points",
						"currentPosition": "$category_registrations.current_position",
						"users": bson.M{
							"$map": bson.M{
								"input": "$users",
								"as":    "user",
								"in": bson.M{
									"_id":        "$$user._id",
									"first_name": "$$user.first_name",
									"last_name":  "$$user.last_name",
								},
							},
						},
					},
					"organizer": bson.M{
						"_id":        "$organizer._id",
						"first_name": "$organizer_user.first_name",
						"last_name":  "$organizer_user.last_name",
					},
				},
			},
		}}},
	)

	// Ejecutar la consulta de agregación
	cursor, err := r.competitorUserColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error when aggregating user categories: %w", err)
	}
	defer cursor.Close(ctx)

	var result []*competitor_user_dao.GetUserCategoriesDAO
	if err := cursor.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("error when decoding user categories: %w", err)
	}

	// Si no se encontraron categorías, devolver un slice vacío
	if len(result) == 0 || len(result[0].Categories) == 0 {
		return []*competitor_user_dao.GetUserCategoriesCategoryDAO{}, nil
	}

	return result[0].Categories, nil
}

func (r *Repository) GetUserCategoriesNumber(
	ctx context.Context,
	userOID *primitive.ObjectID,
	sport models.SPORT,
) (int64, error) {
	// Pipeline de agregación para contar las categorías
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"user_id": userOID}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "category_registrations",
			"localField":   "competitor_id",
			"foreignField": "competitor_id",
			"as":           "category_registrations",
		}}},
		bson.D{{Key: "$unwind", Value: "$category_registrations"}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "categories",
			"localField":   "category_registrations.category_id",
			"foreignField": "_id",
			"as":           "category",
		}}},
		bson.D{{Key: "$unwind", Value: "$category"}},
		bson.D{{Key: "$match", Value: bson.M{"category.sport": sport}}},
		bson.D{{Key: "$group", Value: bson.M{
			"_id":             nil,
			"totalCategories": bson.M{"$sum": 1},
		}}},
	}

	// Ejecutar la consulta de agregación
	cursor, err := r.competitorUserColl.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, fmt.Errorf("error when aggregating user categories count: %w", err)
	}
	defer cursor.Close(ctx)

	// Obtener el resultado
	var result struct {
		TotalCategories int64 `bson:"totalCategories"`
	}
	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return 0, fmt.Errorf("error when decoding user categories count: %w", err)
		}
	} else {
		return 0, nil // No se encontraron categorías
	}

	return result.TotalCategories, nil
}

func (r *Repository) GetUserTournaments(
	ctx context.Context,
	userOID *primitive.ObjectID,
	sport models.SPORT,
	limit int,
	lastOID *primitive.ObjectID, // El último ID desde el que comenzar la siguiente página
) (*competitor_user_dao.GetUserTournamentsDAORes, error) {
	// Paso 1: Buscar en `competitor_users` donde `user_id` sea igual a `userOID`.
	matchStage := bson.D{
		{Key: "$match", Value: bson.M{
			"user_id": userOID,
		}},
	}

	// Paso 2: Hacer un `lookup` en `tournament_registrations`.
	lookupTournamentRegistersStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "tournament_registrations",
			"localField":   "competitor_id",
			"foreignField": "competitor_id",
			"as":           "tournament_registrations",
		}},
	}

	// Paso 3: Desenrollar `tournament_registrations`.
	unwindTournamentRegistersStage := bson.D{
		{Key: "$unwind", Value: bson.M{
			"path": "$tournament_registrations",
		}},
	}

	// Paso 4: Hacer un `lookup` en `tournaments`.
	lookupTournamentsStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "tournaments",
			"localField":   "tournament_registrations.tournament_id",
			"foreignField": "_id",
			"as":           "tournament_details",
		}},
	}

	// Paso 5: Desenrollar `tournament_details`.
	unwindTournamentsStage := bson.D{
		{Key: "$unwind", Value: bson.M{
			"path": "$tournament_details",
		}},
	}

	// Paso 6: Hacer un `lookup` en `locations`.
	lookupLocationStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "locations",
			"localField":   "tournament_details.location_id",
			"foreignField": "_id",
			"as":           "location_details",
		}},
	}

	// Paso 7: Desenrollar `location_details`.
	unwindLocationStage := bson.D{
		{Key: "$unwind", Value: bson.M{
			"path":                       "$location_details",
			"preserveNullAndEmptyArrays": true,
		}},
	}

	// Paso 8: Hacer un `lookup` en `organizers`.
	lookupOrganizerStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "organizers",
			"localField":   "tournament_details.organizer_id",
			"foreignField": "_id",
			"as":           "organizer_details",
		}},
	}

	// Paso 9: Desenrollar `organizer_details`.
	unwindOrganizerStage := bson.D{
		{Key: "$unwind", Value: bson.M{
			"path":                       "$organizer_details",
			"preserveNullAndEmptyArrays": true,
		}},
	}

	// Paso 10: Hacer un `lookup` en `users` para obtener detalles del organizador.
	lookupOrganizerUserStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "users",
			"localField":   "organizer_details.user_id",
			"foreignField": "_id",
			"as":           "organizer_user_details",
		}},
	}

	// Paso 11: Desenrollar `organizer_user_details`.
	unwindOrganizerUserStage := bson.D{
		{Key: "$unwind", Value: bson.M{
			"path":                       "$organizer_user_details",
			"preserveNullAndEmptyArrays": true,
		}},
	}

	// Paso 12: Filtro por deporte, si corresponde.
	var sportFilterStage bson.D
	if sport != "" {
		sportFilterStage = bson.D{
			{Key: "$match", Value: bson.M{
				"tournament_details.sport": sport,
			}},
		}
	}

	// Paso 13: Filtro por `lastOID` si está presente para la paginación.
	var lastOIDFilterStage bson.D
	if lastOID != nil {
		lastOIDFilterStage = bson.D{
			{Key: "$match", Value: bson.M{
				"tournament_details._id": bson.M{
					"$lt": lastOID, // Solo torneos con _id menor que el `lastOID` (paginación hacia atrás)
				},
			}},
		}
	}

	// Paso 14: Ordenar los torneos por la fecha o por `_id` (en orden descendente para obtener los más recientes primero).
	sortStage := bson.D{
		{Key: "$sort", Value: bson.M{
			"tournament_details._id": -1, // Ordenar por _id de más reciente a más antiguo
		}},
	}

	// Paso 15: Limitar el número de resultados.
	limitStage := bson.D{
		{Key: "$limit", Value: limit},
	}

	// Paso 16: Proyección final para mapear a la estructura deseada.
	projectStage := bson.D{
		{Key: "$project", Value: bson.M{
			"_id":           "$tournament_details._id",
			"name":          "$tournament_details.name",
			"start_date":    "$tournament_details.start_date",
			"image":         "$tournament_details.image",
			"finish_date":   "$tournament_details.finish_date",
			"points":        "$tournament_details.points",
			"average_score": "$tournament_details.average_score",
			"total_prize":   "$tournament_details.total_prize",
			"location": bson.M{
				"_id":     "$location_details._id",
				"state":   "$location_details.state",
				"country": "$location_details.country",
				"city":    "$location_details.city",
				"lat":     "$location_details.lat",
				"long":    "$location_details.long",
			},
			"organizer": bson.M{
				"_id":        "$organizer_user_details._id",
				"first_name": "$organizer_user_details.first_name",
				"last_name":  "$organizer_user_details.last_name",
			},
		}},
	}

	// Pipeline para contar el total de torneos
	totalPipeline := mongo.Pipeline{
		matchStage,
		lookupTournamentRegistersStage,
		unwindTournamentRegistersStage,
		lookupTournamentsStage,
		unwindTournamentsStage,
	}

	// Agregar el filtro de deporte si aplica al pipeline de conteo
	if sport != "" {
		totalPipeline = append(totalPipeline, sportFilterStage)
	}

	// Agregar la etapa de conteo al pipeline de conteo
	totalPipeline = append(totalPipeline, bson.D{{Key: "$count", Value: "total"}})

	// Ejecutar la consulta de conteo
	totalCursor, err := r.competitorUserColl.Aggregate(ctx, totalPipeline)
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

	// Pipeline para obtener torneos con paginación
	pipeline := mongo.Pipeline{
		matchStage,
		lookupTournamentRegistersStage,
		unwindTournamentRegistersStage,
		lookupTournamentsStage,
		unwindTournamentsStage,
		lookupLocationStage,
		unwindLocationStage,
		lookupOrganizerStage,
		unwindOrganizerStage,
		lookupOrganizerUserStage,
		unwindOrganizerUserStage,
	}

	// Agregar filtro de deporte y paginación si aplica
	if sport != "" {
		pipeline = append(pipeline, sportFilterStage)
	}
	if lastOID != nil {
		pipeline = append(pipeline, lastOIDFilterStage)
	}

	// Ordenar y limitar los resultados
	pipeline = append(pipeline, sortStage, limitStage, projectStage)

	// Ejecutar la consulta de torneos
	cursor, err := r.competitorUserColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error when aggregating user tournaments: %w", err)
	}
	defer cursor.Close(ctx)

	var tournaments []*competitor_user_dao.GetUserTournamentDAORes
	if err := cursor.All(ctx, &tournaments); err != nil {
		return nil, fmt.Errorf("error when decoding user tournaments: %w", err)
	}

	// Retornar la estructura con los torneos y el total
	result := &competitor_user_dao.GetUserTournamentsDAORes{
		Tournaments: tournaments,
		Total:       total,
	}

	return result, nil
}

func (r *Repository) GetUserAllCompetitorSports(
	ctx context.Context,
	userOID *primitive.ObjectID,
) ([]models.SPORT, error) {

	// Pipeline de agregación
	pipeline := mongo.Pipeline{
		// Filtra los documentos por el user_id
		bson.D{{Key: "$match", Value: bson.M{"user_id": userOID}}},
		// Hace un lookup en la colección de competidores para obtener los detalles del competidor
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "competitors",
			"localField":   "competitor_id",
			"foreignField": "_id",
			"as":           "competitor",
		}}},
		// Expande el array "competitor" para procesar cada elemento individualmente
		bson.D{{Key: "$unwind", Value: "$competitor"}},
		// Agrupa por el campo "sport" para eliminar duplicados
		bson.D{{Key: "$group", Value: bson.M{
			"_id": "$competitor.sport",
		}}},
		// Proyección para obtener el campo "sport" de manera limpia
		bson.D{{Key: "$project", Value: bson.M{
			"sport": "$_id",
			"_id":   0,
		}}},
	}

	// Ejecutar la consulta de agregación
	cursor, err := r.competitorUserColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error when aggregating user sports: %w", err)
	}
	defer cursor.Close(ctx)

	// Decodificar el resultado
	var results []struct {
		Sport models.SPORT `bson:"sport"`
	}

	if err := cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("error when decoding sports: %w", err)
	}

	// Procesar el resultado para devolver un slice de tipos de deportes
	sports := make([]models.SPORT, len(results))
	for i, result := range results {
		sports[i] = result.Sport
	}

	return sports, nil
}

func (r *Repository) GetProfileUserTournaments(
	ctx context.Context,
	userOID *primitive.ObjectID,
	sport models.SPORT,
	limit int,
	lastOID *primitive.ObjectID,
) (*competitor_user_dao.GetProfileUserTournamentsDAORes, error) {
	// Paso 1: Buscar en competitor_users donde user_id sea igual a userOID.
	matchStage := bson.D{
		{Key: "$match", Value: bson.M{
			"user_id": userOID,
		}},
	}

	// Paso 2: Hacer un lookup en tournament_registrations.
	lookupTournamentRegistersStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "tournament_registrations",
			"localField":   "competitor_id",
			"foreignField": "competitor_id",
			"as":           "tournament_registrations",
		}},
	}

	// Paso 3: Desenrollar tournament_registrations.
	unwindTournamentRegistersStage := bson.D{
		{Key: "$unwind", Value: bson.M{
			"path": "$tournament_registrations",
		}},
	}

	// Paso 4: Hacer un lookup en tournaments.
	lookupTournamentsStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "tournaments",
			"localField":   "tournament_registrations.tournament_id",
			"foreignField": "_id",
			"as":           "tournament_details",
		}},
	}

	// Paso 5: Desenrollar tournament_details.
	unwindTournamentsStage := bson.D{
		{Key: "$unwind", Value: bson.M{
			"path": "$tournament_details",
		}},
	}

	// // Paso 6: Hacer un lookup en matches.
	// lookupMatchesStage := bson.D{
	// 	{Key: "$lookup", Value: bson.M{
	// 		"from":         "matches",
	// 		"localField":   "tournament_details.matches",
	// 		"foreignField": "_id",
	// 		"as":           "matches",
	// 	}},
	// }

	// Paso 7: Hacer un lookup en rounds.
	lookupRoundsStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "rounds",
			"localField":   "matches.round_id",
			"foreignField": "_id",
			"as":           "all_rounds",
		}},
	}

	// Paso 8: Hacer un lookup en competitor_matches.
	lookupCompetitorMatchesStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "competitor_matches",
			"localField":   "matches._id",
			"foreignField": "match_id",
			"as":           "competitor_matches",
		}},
	}

	// Paso 9: Hacer un lookup en competitor_users.
	lookupCompetitorUsersStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "competitor_users",
			"localField":   "competitor_matches.competitor_id",
			"foreignField": "competitor_id",
			"as":           "competitor_users",
		}},
	}

	// Paso 10: Hacer un lookup en users.
	lookupUsersStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "users",
			"localField":   "competitor_users.user_id",
			"foreignField": "_id",
			"as":           "users",
		}},
	}

	// Paso 11: Hacer un lookup en guest_competitors.
	lookupGuestCompetitorsStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "guest_competitors",
			"localField":   "competitor_matches.competitor_id",
			"foreignField": "competitor_id",
			"as":           "guest_competitors",
		}},
	}

	// Paso 12: Hacer un lookup en guest_users.
	lookupGuestUsersStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "guest_users",
			"localField":   "guest_competitors.guest_user_id",
			"foreignField": "_id",
			"as":           "guest_users",
		}},
	}

	// Paso 13: Hacer un lookup en locations.
	lookupLocationStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "locations",
			"localField":   "tournament_details.location_id",
			"foreignField": "_id",
			"as":           "location_details",
		}},
	}

	// Paso 14: Desenrollar location_details.
	unwindLocationStage := bson.D{
		{Key: "$unwind", Value: bson.M{
			"path":                       "$location_details",
			"preserveNullAndEmptyArrays": true,
		}},
	}

	// Paso 15: Hacer un lookup en organizers.
	lookupOrganizerStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "organizers",
			"localField":   "tournament_details.organizer_id",
			"foreignField": "_id",
			"as":           "organizer_details",
		}},
	}

	// Paso 16: Desenrollar organizer_details.
	unwindOrganizerStage := bson.D{
		{Key: "$unwind", Value: bson.M{
			"path":                       "$organizer_details",
			"preserveNullAndEmptyArrays": true,
		}},
	}

	// Paso 17: Hacer un lookup en category_registrations.
	lookupCategoryStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "category_registrations",
			"localField":   "competitor_matches.competitor_id",
			"foreignField": "competitor_id",
			"as":           "category_registrations",
		}},
	}

	// Paso 18: Desenrollar category_registration.
	unwindCategoryStage := bson.D{
		{Key: "$unwind", Value: bson.M{
			"path":                       "$category_registration",
			"preserveNullAndEmptyArrays": true,
		}},
	}

	// Paso 19: Hacer un lookup en users para obtener detalles del organizador.
	lookupOrganizerUserStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "users",
			"localField":   "organizer_details.user_id",
			"foreignField": "_id",
			"as":           "organizer_user_details",
		}},
	}

	// Paso 20: Desenrollar organizer_user_details.
	unwindOrganizerUserStage := bson.D{
		{Key: "$unwind", Value: bson.M{
			"path":                       "$organizer_user_details",
			"preserveNullAndEmptyArrays": true,
		}},
	}

	// Nuevo: Filtro para descartar torneos con start_date null
	filterNullStartDateStage := bson.D{
		{Key: "$match", Value: bson.M{
			"tournament_details.start_date": bson.M{
				"$ne": nil,
			},
		}},
	}

	// Filtro por deporte, si corresponde.
	var sportFilterStage bson.D
	if sport != "" {
		sportFilterStage = bson.D{
			{Key: "$match", Value: bson.M{
				"tournament_details.sport": sport,
			}},
		}
	}

	// Filtro por lastOID si está presente para la paginación.
	var lastOIDFilterStage bson.D
	if lastOID != nil {
		// Necesitamos obtener el start_date del torneo correspondiente al lastOID
		var lastTournament struct {
			StartDate *time.Time `bson:"start_date"`
		}
		err := r.tournamentColl.FindOne(ctx, bson.M{"_id": lastOID}).Decode(&lastTournament)
		if err != nil {
			return nil, fmt.Errorf("error finding last tournament: %w", err)
		}

		lastOIDFilterStage = bson.D{
			{Key: "$match", Value: bson.M{
				"$or": []bson.M{
					{
						"tournament_details.start_date": bson.M{
							"$lt": lastTournament.StartDate,
						},
					},
					{
						"$and": []bson.M{
							{
								"tournament_details.start_date": lastTournament.StartDate,
							},
							{
								"tournament_details._id": bson.M{
									"$lt": lastOID,
								},
							},
						},
					},
				},
			}},
		}
	}

	// Ordenar los torneos por start_date y luego por _id
	sortStage := bson.D{
		{Key: "$sort", Value: bson.M{
			"tournament_details.start_date": -1, // Ordenar por fecha de inicio descendente
			"tournament_details._id":        -1, // Desempatar por _id
		}},
	}

	// Limitar el número de resultados.
	limitStage := bson.D{
		{Key: "$limit", Value: limit},
	}

	// Nuevo: Lookup para tournament_groups
	lookupTournamentGroupsStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "tournament_groups",
			"localField":   "tournament_details.groups",
			"foreignField": "_id",
			"as":           "tournament_groups",
		}},
	}

	// Nuevo: Lookup para obtener los matches de los grupos
	lookupGroupMatchesStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "matches",
			"localField":   "tournament_groups.matches",
			"foreignField": "_id",
			"as":           "group_matches",
		}},
	}

	// Modificar el lookupMatchesStage existente para incluir los matches de grupos
	lookupMatchesStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "matches",
			"localField":   "tournament_details.matches",
			"foreignField": "_id",
			"as":           "tournament_matches",
		}},
	}

	// Nuevo: Stage para combinar matches de torneo y grupos
	addFieldsCombineMatchesStage := bson.D{
		{Key: "$addFields", Value: bson.M{
			"matches": bson.M{
				"$concatArrays": []interface{}{
					"$tournament_matches",
					"$group_matches",
				},
			},
		}},
	}

	// Proyección final
	projectStage := bson.D{
		{Key: "$project", Value: bson.M{
			"_id":  "$tournament_details._id",
			"name": "$tournament_details.name",
			"location": bson.M{
				"_id":     "$location_details._id",
				"state":   "$location_details.state",
				"country": "$location_details.country",
				"city":    "$location_details.city",
				"lat":     "$location_details.lat",
				"long":    "$location_details.long",
			},
			"organizer": bson.M{
				"_id":        "$organizer_user_details._id",
				"first_name": "$organizer_user_details.first_name",
				"last_name":  "$organizer_user_details.last_name",
			},
			"matches": bson.M{
				"$filter": bson.M{
					"input": bson.M{
						"$map": bson.M{
							"input": "$matches",
							"as":    "match",
							"in": bson.M{
								"_id":      "$$match._id",
								"result":   "$$match.result",
								"winner":   "$$match.winner",
								"date":     "$$match.date",
								"position": "$$match.position",
								"round": bson.M{
									"$let": bson.M{
										"vars": bson.M{
											"matchRound": bson.M{
												"$filter": bson.M{
													"input": "$all_rounds",
													"as":    "round",
													"cond":  bson.M{"$eq": bson.A{"$$round._id", "$$match.round_id"}},
												},
											},
										},
										"in": bson.M{
											"_id":   bson.M{"$arrayElemAt": bson.A{"$$matchRound._id", 0}},
											"round": bson.M{"$arrayElemAt": bson.A{"$$matchRound.round", 0}},
										},
									},
								},
								"competitors": bson.M{
									"$map": bson.M{
										"input": bson.M{
											"$filter": bson.M{
												"input": "$competitor_matches",
												"as":    "cm",
												"cond":  bson.M{"$eq": bson.A{"$$cm.match_id", "$$match._id"}},
											},
										},
										"as": "cm",
										"in": buildCompetitorProjectionWithoutCurrentPositionAndPosition(),
									},
								},
							},
						},
					},
					"as": "match",
					"cond": bson.M{
						"$or": bson.A{
							bson.M{
								"$anyElementTrue": bson.M{
									"$map": bson.M{
										"input": "$$match.competitors",
										"as":    "competitor",
										"in": bson.M{
											"$anyElementTrue": bson.M{
												"$map": bson.M{
													"input": "$$competitor.users",
													"as":    "user",
													"in": bson.M{
														"$eq": bson.A{"$$user._id", userOID},
													},
												},
											},
										},
									},
								},
							},
							bson.M{
								"$anyElementTrue": bson.M{
									"$map": bson.M{
										"input": "$$match.competitors",
										"as":    "competitor",
										"in": bson.M{
											"$anyElementTrue": bson.M{
												"$map": bson.M{
													"input": "$$competitor.guest_users",
													"as":    "guest_user",
													"in": bson.M{
														"$eq": bson.A{"$$guest_user._id", userOID},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}},
	}

	// Pipeline para contar el total de torneos
	totalPipeline := mongo.Pipeline{
		matchStage,
		lookupTournamentRegistersStage,
		unwindTournamentRegistersStage,
		lookupTournamentsStage,
		unwindTournamentsStage,
		filterNullStartDateStage,    // Agregar filtro de null start_date
		lookupTournamentGroupsStage, // Nuevo
		lookupMatchesStage,          // Modificado
		lookupGroupMatchesStage,     // Nuevo
		addFieldsCombineMatchesStage,
		lookupMatchesStage,
		lookupRoundsStage,
		lookupCompetitorMatchesStage,
		lookupCompetitorUsersStage,
		lookupUsersStage,
		lookupGuestCompetitorsStage,
		lookupGuestUsersStage,
	}

		// Pipeline para obtener torneos con paginación
	pipeline := mongo.Pipeline{
		matchStage,
		lookupTournamentRegistersStage,
		unwindTournamentRegistersStage,
		lookupTournamentsStage,
		unwindTournamentsStage,
		filterNullStartDateStage,
		lookupTournamentGroupsStage,  // Nuevo
		lookupMatchesStage,           // Modificado
		lookupGroupMatchesStage,      // Nuevo
		addFieldsCombineMatchesStage, // Nue
		lookupMatchesStage,
		lookupRoundsStage,
		lookupCompetitorMatchesStage,
		lookupCompetitorUsersStage,
		lookupUsersStage,
		lookupGuestCompetitorsStage,
		lookupGuestUsersStage,
		lookupLocationStage,
		unwindLocationStage,
		lookupOrganizerStage,
		unwindOrganizerStage,
		lookupCategoryStage,
		unwindCategoryStage,
		lookupOrganizerUserStage,
		unwindOrganizerUserStage,
	}

	if sport != "" {
		pipeline = append(pipeline, sportFilterStage)
		totalPipeline = append(totalPipeline, sportFilterStage)
	}
	if lastOID != nil {
		pipeline = append(pipeline, lastOIDFilterStage)
	}

	// Ordenar y limitar los resultados
	pipeline = append(pipeline, sortStage, limitStage, projectStage)
	
	// Agregar la etapa de conteo al pipeline de conteo
	totalPipeline = append(totalPipeline, bson.D{{Key: "$count", Value: "total"}})

	// Ejecutar la consulta de conteo
	totalCursor, err := r.competitorUserColl.Aggregate(ctx, totalPipeline)
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

	// Ejecutar la consulta de torneos
	cursor, err := r.competitorUserColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error when aggregating user tournaments: %w", err)
	}
	defer cursor.Close(ctx)

	var tournaments []*competitor_user_dao.GetProfileUserTournamentDAORes
	if err := cursor.All(ctx, &tournaments); err != nil {
		return nil, fmt.Errorf("error when decoding user tournaments: %w", err)
	}

	// Retornar la estructura con los torneos y el total
	result := &competitor_user_dao.GetProfileUserTournamentsDAORes{
		Tournaments: tournaments,
		Total:       total,
	}

	return result, nil
}

func buildCompetitorProjectionWithoutCurrentPositionAndPosition() bson.M {
	return bson.M{
		"_id":              "$$cm.competitor_id",
		"current_position": nil,
		"users":            buildUsersProjection(),
		"guest_users":      buildGuestUsersProjection(),
	}
}
func buildCompetitorProjectionWithoutPosition(categoryOID *primitive.ObjectID) bson.M {
	return bson.M{
		"_id":              "$$cm.competitor_id",
		"current_position": buildCurrentPositionProjection(categoryOID),
		"users":            buildUsersProjection(),
		"guest_users":      buildGuestUsersProjection(),
	}
}

func (r *Repository) GetProfileCompetitorTournaments(
	ctx context.Context,
	competitorOID, categoryOID *primitive.ObjectID,
	sport models.SPORT,
	limit int,
	lastOID *primitive.ObjectID,
) (*competitor_user_dao.GetProfileUserTournamentsDAORes, error) {
	// Paso 1: Buscar en competitor_users donde user_id sea igual a competitorOID.
	matchStage := bson.D{
		{Key: "$match", Value: bson.M{
			"competitor_id": competitorOID,
		}},
	}

	// Paso 2: Hacer un lookup en tournament_registrations.
	lookupTournamentRegistersStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "tournament_registrations",
			"localField":   "competitor_id",
			"foreignField": "competitor_id",
			"as":           "tournament_registrations",
		}},
	}

	// Paso 3: Desenrollar tournament_registrations.
	unwindTournamentRegistersStage := bson.D{
		{Key: "$unwind", Value: bson.M{
			"path": "$tournament_registrations",
		}},
	}

	// Paso 4: Hacer un lookup en tournaments.
	lookupTournamentsStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "tournaments",
			"localField":   "tournament_registrations.tournament_id",
			"foreignField": "_id",
			"as":           "tournament_details",
		}},
	}

	// Paso 5: Desenrollar tournament_details.
	unwindTournamentsStage := bson.D{
		{Key: "$unwind", Value: bson.M{
			"path": "$tournament_details",
		}},
	}

	// // Paso 6: Hacer un lookup en matches.
	// lookupMatchesStage := bson.D{
	// 	{Key: "$lookup", Value: bson.M{
	// 		"from":         "matches",
	// 		"localField":   "tournament_details.matches",
	// 		"foreignField": "_id",
	// 		"as":           "matches",
	// 	}},
	// }

	// Paso 7: Hacer un lookup en rounds.
	lookupRoundsStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "rounds",
			"localField":   "matches.round_id",
			"foreignField": "_id",
			"as":           "all_rounds",
		}},
	}

	// Paso 8: Hacer un lookup en competitor_matches.
	lookupCompetitorMatchesStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "competitor_matches",
			"localField":   "matches._id",
			"foreignField": "match_id",
			"as":           "competitor_matches",
		}},
	}

	// Paso 9: Hacer un lookup en competitor_users.
	lookupCompetitorUsersStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "competitor_users",
			"localField":   "competitor_matches.competitor_id",
			"foreignField": "competitor_id",
			"as":           "competitor_users",
		}},
	}

	// Paso 10: Hacer un lookup en users.
	lookupUsersStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "users",
			"localField":   "competitor_users.user_id",
			"foreignField": "_id",
			"as":           "users",
		}},
	}

	// Paso 11: Hacer un lookup en guest_competitors.
	lookupGuestCompetitorsStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "guest_competitors",
			"localField":   "competitor_matches.competitor_id",
			"foreignField": "competitor_id",
			"as":           "guest_competitors",
		}},
	}

	// Paso 12: Hacer un lookup en guest_users.
	lookupGuestUsersStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "guest_users",
			"localField":   "guest_competitors.guest_user_id",
			"foreignField": "_id",
			"as":           "guest_users",
		}},
	}

	// Paso 13: Hacer un lookup en locations.
	lookupLocationStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "locations",
			"localField":   "tournament_details.location_id",
			"foreignField": "_id",
			"as":           "location_details",
		}},
	}

	// Paso 14: Desenrollar location_details.
	unwindLocationStage := bson.D{
		{Key: "$unwind", Value: bson.M{
			"path":                       "$location_details",
			"preserveNullAndEmptyArrays": true,
		}},
	}

	// Paso 15: Hacer un lookup en organizers.
	lookupOrganizerStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "organizers",
			"localField":   "tournament_details.organizer_id",
			"foreignField": "_id",
			"as":           "organizer_details",
		}},
	}

	// Paso 16: Desenrollar organizer_details.
	unwindOrganizerStage := bson.D{
		{Key: "$unwind", Value: bson.M{
			"path":                       "$organizer_details",
			"preserveNullAndEmptyArrays": true,
		}},
	}

	// Paso 17: Hacer un lookup en category_registrations.
	lookupCategoryStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "category_registrations",
			"localField":   "competitor_matches.competitor_id",
			"foreignField": "competitor_id",
			"as":           "category_registration",
		}},
	}

	// Paso 19: Hacer un lookup en users para obtener detalles del organizador.
	lookupOrganizerUserStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "users",
			"localField":   "organizer_details.user_id",
			"foreignField": "_id",
			"as":           "organizer_user_details",
		}},
	}

	// Paso 20: Desenrollar organizer_user_details.
	unwindOrganizerUserStage := bson.D{
		{Key: "$unwind", Value: bson.M{
			"path":                       "$organizer_user_details",
			"preserveNullAndEmptyArrays": true,
		}},
	}

	// Nuevo: Filtro para descartar torneos con start_date null
	filterNullStartDateStage := bson.D{
		{Key: "$match", Value: bson.M{
			"tournament_details.start_date": bson.M{
				"$ne": nil,
			},
		}},
	}

	// Filtro por deporte, si corresponde.
	var sportFilterStage bson.D
	if sport != "" {
		sportFilterStage = bson.D{
			{Key: "$match", Value: bson.M{
				"tournament_details.sport": sport,
			}},
		}
	}

	// Filtro por lastOID si está presente para la paginación.
	var lastOIDFilterStage bson.D
	if lastOID != nil {
		// Necesitamos obtener el start_date del torneo correspondiente al lastOID
		var lastTournament struct {
			StartDate *time.Time `bson:"start_date"`
		}
		err := r.tournamentColl.FindOne(ctx, bson.M{"_id": lastOID}).Decode(&lastTournament)
		if err != nil {
			return nil, fmt.Errorf("error finding last tournament: %w", err)
		}

		lastOIDFilterStage = bson.D{
			{Key: "$match", Value: bson.M{
				"$or": []bson.M{
					{
						"tournament_details.start_date": bson.M{
							"$lt": lastTournament.StartDate,
						},
					},
					{
						"$and": []bson.M{
							{
								"tournament_details.start_date": lastTournament.StartDate,
							},
							{
								"tournament_details._id": bson.M{
									"$lt": lastOID,
								},
							},
						},
					},
				},
			}},
		}
	}

	// Ordenar los torneos por start_date y luego por _id
	sortStage := bson.D{
		{Key: "$sort", Value: bson.M{
			"tournament_details.start_date": -1, // Ordenar por fecha de inicio descendente
			"tournament_details._id":        -1, // Desempatar por _id
		}},
	}

	// Limitar el número de resultados.
	limitStage := bson.D{
		{Key: "$limit", Value: limit},
	}
	// Nuevo: Lookup para tournament_groups
	lookupTournamentGroupsStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "tournament_groups",
			"localField":   "tournament_details.groups",
			"foreignField": "_id",
			"as":           "tournament_groups",
		}},
	}

	// Nuevo: Lookup para obtener los matches de los grupos
	lookupGroupMatchesStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "matches",
			"localField":   "tournament_groups.matches",
			"foreignField": "_id",
			"as":           "group_matches",
		}},
	}

	// Modificar el lookupMatchesStage existente para incluir los matches de grupos
	lookupMatchesStage := bson.D{
		{Key: "$lookup", Value: bson.M{
			"from":         "matches",
			"localField":   "tournament_details.matches",
			"foreignField": "_id",
			"as":           "tournament_matches",
		}},
	}

	// Nuevo: Stage para combinar matches de torneo y grupos
	addFieldsCombineMatchesStage := bson.D{
		{Key: "$addFields", Value: bson.M{
			"matches": bson.M{
				"$concatArrays": []interface{}{
					"$tournament_matches",
					"$group_matches",
				},
			},
		}},
	}
	// Proyección final
	projectStage := bson.D{
		{Key: "$project", Value: bson.M{
			"_id":  "$tournament_details._id",
			"name": "$tournament_details.name",
			"location": bson.M{
				"_id":     "$location_details._id",
				"state":   "$location_details.state",
				"country": "$location_details.country",
				"city":    "$location_details.city",
				"lat":     "$location_details.lat",
				"long":    "$location_details.long",
			},
			"organizer": bson.M{
				"_id":        "$organizer_user_details._id",
				"first_name": "$organizer_user_details.first_name",
				"last_name":  "$organizer_user_details.last_name",
			},
			"matches": bson.M{
				"$filter": bson.M{
					"input": bson.M{
						"$map": bson.M{
							"input": "$matches", // Ahora contiene tanto matches de torneo como de grupos
							"as":    "match",
							"in": bson.M{
								"_id":      "$$match._id",
								"result":   "$$match.result",
								"winner":   "$$match.winner",
								"date":     "$$match.date",
								"position": "$$match.position",
								"round": bson.M{
									"$let": bson.M{
										"vars": bson.M{
											"matchRound": bson.M{
												"$filter": bson.M{
													"input": "$all_rounds",
													"as":    "round",
													"cond":  bson.M{"$eq": bson.A{"$$round._id", "$$match.round_id"}},
												},
											},
										},
										"in": bson.M{
											"_id":   bson.M{"$arrayElemAt": bson.A{"$$matchRound._id", 0}},
											"round": bson.M{"$arrayElemAt": bson.A{"$$matchRound.round", 0}},
										},
									},
								},
								"competitors": bson.M{
									"$map": bson.M{
										"input": bson.M{
											"$filter": bson.M{
												"input": "$competitor_matches",
												"as":    "cm",
												"cond":  bson.M{"$eq": bson.A{"$$cm.match_id", "$$match._id"}},
											},
										},
										"as": "cm",
										"in": buildCompetitorProjectionWithoutPosition(categoryOID),
									},
								},
							},
						},
					},
					"as": "match",
					"cond": bson.M{
						"$anyElementTrue": bson.M{
							"$map": bson.M{
								"input": "$$match.competitors",
								"as":    "competitor",
								"in": bson.M{
									"$eq": bson.A{"$$competitor._id", competitorOID},
								},
							},
						},
					},
				},
			},
		}},
	}

	// Pipeline para contar el total de torneos
	pipeline := mongo.Pipeline{
		matchStage,
		lookupTournamentRegistersStage,
		unwindTournamentRegistersStage,
		lookupTournamentsStage,
		unwindTournamentsStage,
		filterNullStartDateStage,
		lookupTournamentGroupsStage,  // Nuevo
		lookupMatchesStage,           // Modificado
		lookupGroupMatchesStage,      // Nuevo
		addFieldsCombineMatchesStage, // Nuevo
		lookupRoundsStage,
		lookupCompetitorMatchesStage,
		lookupCompetitorUsersStage,
		lookupUsersStage,
		lookupGuestCompetitorsStage,
		lookupGuestUsersStage,
		lookupLocationStage,
		unwindLocationStage,
		lookupOrganizerStage,
		unwindOrganizerStage,
		lookupCategoryStage,
		lookupOrganizerUserStage,
		unwindOrganizerUserStage,
	}

	// Actualizar también el pipeline de conteo
	totalPipeline := mongo.Pipeline{
		matchStage,
		lookupTournamentRegistersStage,
		unwindTournamentRegistersStage,
		lookupTournamentsStage,
		unwindTournamentsStage,
		filterNullStartDateStage,
		lookupTournamentGroupsStage,  // Nuevo
		lookupMatchesStage,           // Modificado
		lookupGroupMatchesStage,      // Nuevo
		addFieldsCombineMatchesStage, // Nuevo
		lookupRoundsStage,
		lookupCompetitorMatchesStage,
		lookupCompetitorUsersStage,
		lookupUsersStage,
		lookupGuestCompetitorsStage,
		lookupGuestUsersStage,
	}

	// El resto de la función permanece igual...

	// Agregar filtro de deporte y paginación si aplica
	if sport != "" {
		pipeline = append(pipeline, sportFilterStage)
		totalPipeline = append(totalPipeline, sportFilterStage)
	}
	if lastOID != nil {
		pipeline = append(pipeline, lastOIDFilterStage)
	}

	// Ordenar y limitar los resultados
	pipeline = append(pipeline, sortStage, limitStage, projectStage)

	// Agregar la etapa de conteo al pipeline de conteo
	totalPipeline = append(totalPipeline, bson.D{{Key: "$count", Value: "total"}})

	// Ejecutar la consulta de conteo
	totalCursor, err := r.competitorUserColl.Aggregate(ctx, totalPipeline)
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

	// Ejecutar la consulta de torneos
	cursor, err := r.competitorUserColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error when aggregating user tournaments: %w", err)
	}
	defer cursor.Close(ctx)

	var tournaments []*competitor_user_dao.GetProfileUserTournamentDAORes
	if err := cursor.All(ctx, &tournaments); err != nil {
		return nil, fmt.Errorf("error when decoding user tournaments: %w", err)
	}

	// Retornar la estructura con los torneos y el total
	result := &competitor_user_dao.GetProfileUserTournamentsDAORes{
		Tournaments: tournaments,
		Total:       total,
	}

	return result, nil
}

// func (r *Repository) GetUsersCompetitorAvailability(
// 	ctx context.Context,
// 	userOID *primitive.ObjectID,
// 	day models.DAY,
// 	timeSlot string,
// ) ([]*availability_dao.GetDayTimeSlotDAORes, error) {

// 	pipeline := mongo.Pipeline{
// 		bson.D{{Key: "$match", Value: bson.M{"user_id": userOID}}},
// 		bson.D{{Key: "$lookup", Value: bson.M{
// 			"from":         "availabilities",
// 			"localField":   "competitor_id",
// 			"foreignField": "competitor_id",
// 			"as":           "availability",
// 		}}},
// 		bson.D{{Key: "$unwind", Value: "$availability"}},
// 		bson.D{{Key: "$project", Value: bson.M{
// 			"_id": "$availability._id",
// 			"daily_availabilities": bson.M{"$filter": bson.M{
// 				"input": "$availability.daily_availabilities",
// 				"as":    "daily",
// 				"cond": bson.M{"$eq": bson.A{"$$daily.day", day}},
// 			}},
// 		}}},
// 		bson.D{{Key: "$unwind", Value: bson.M{
// 			"path":                       "$daily_availabilities",
// 			"preserveNullAndEmptyArrays": true,
// 		}}},
// 		bson.D{{Key: "$project", Value: bson.M{
// 			"_id": "$_id",
// 			"daily_availabilities": bson.M{"$cond": bson.M{
// 				"if": bson.M{"$gt": bson.A{"$daily_availabilities", nil}},
// 				"then": []bson.M{
// 					{
// 						"day": "$daily_availabilities.day",
// 						"time_slots": bson.M{"$filter": bson.M{
// 							"input": "$daily_availabilities.time_slots",
// 							"as":    "slot",
// 							"cond": bson.M{"$eq": bson.A{"$$slot.time_slot", timeSlot}},
// 						}},
// 					},
// 				},
// 				"else": bson.A{},
// 			}},
// 		}}},
// 	}

// 	// Ejecutar la consulta de agregación
// 	cursor, err := r.competitorUserColl.Aggregate(ctx, pipeline)
// 	if err != nil {
// 		return nil, fmt.Errorf("error when aggregating user categories: %w", err)
// 	}
// 	defer cursor.Close(ctx)

// 	var result []*availability_dao.GetDayTimeSlotDAORes
// 	if err := cursor.All(ctx, &result); err != nil {
// 		return nil, fmt.Errorf("error when decoding user categories: %w", err)
// 	}

// 	// Si no se encontraron categorías, devolver un slice vacío
// 	if len(result) == 0 {
// 		return []*availability_dao.GetDayTimeSlotDAORes{}, nil
// 	}

// 	return result, nil
// }

func (r *Repository) GetCompetitorIDsFromUser(
	ctx context.Context,
	userOID *primitive.ObjectID,
) ([]*primitive.ObjectID, error) {
	// Definir el filtro para la consulta Find
	filter := bson.M{"user_id": userOID}

	// Definir la proyección para obtener solo el campo "competitor_id"
	projection := bson.M{"competitor_id": 1}

	opts := options.Find().SetProjection(projection)

	// Ejecutar la consulta Find
	cursor, err := r.competitorUserColl.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("error when querying competitor IDs: %w", err)
	}
	defer cursor.Close(ctx)

	// Inicializar el slice para almacenar los resultados
	var competitorIDs []*primitive.ObjectID
	// Iterar sobre los resultados del cursor y extraer el campo "competitor_id"
	for cursor.Next(ctx) {
		var doc struct {
			CompetitorID *primitive.ObjectID `bson:"competitor_id"`
		}
		if err := cursor.Decode(&doc); err != nil {
			return nil, fmt.Errorf("error when decoding result: %w", err)
		}
		// Agregar el "competitor_id" al slice de resultados
		competitorIDs = append(competitorIDs, doc.CompetitorID)
	}

	// Si no se encontraron coincidencias, devolver un slice vacío
	if len(competitorIDs) == 0 {
		return []*primitive.ObjectID{}, nil
	}

	return competitorIDs, nil
}

func (r *Repository) GetUsersAvailability(
	ctx context.Context,
	competitorOID *primitive.ObjectID,
	day models.DAY,
	timeSlot string,
) ([]*availability_dao.GetDayTimeSlotDAORes, error) {
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"competitor_id": competitorOID}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "availabilities",
			"localField":   "user_id",
			"foreignField": "user_id",
			"as":           "availability",
		}}},
		bson.D{{Key: "$unwind", Value: "$availability"}},
		bson.D{{Key: "$project", Value: bson.M{
			"_id": "$availability._id",
			"daily_availabilities": bson.M{"$filter": bson.M{
				"input": "$availability.daily_availabilities",
				"as":    "daily",
				"cond":  bson.M{"$eq": bson.A{"$$daily.day", day}},
			}},
		}}},
		bson.D{{Key: "$unwind", Value: bson.M{
			"path":                       "$daily_availabilities",
			"preserveNullAndEmptyArrays": true,
		}}},
		bson.D{{Key: "$project", Value: bson.M{
			"_id": "$_id",
			"daily_availabilities": bson.M{"$cond": bson.M{
				"if": bson.M{"$gt": bson.A{"$daily_availabilities", nil}},
				"then": []bson.M{
					{
						"day": "$daily_availabilities.day",
						"time_slots": bson.M{"$filter": bson.M{
							"input": "$daily_availabilities.time_slots",
							"as":    "slot",
							"cond":  bson.M{"$eq": bson.A{"$$slot.time_slot", timeSlot}},
						}},
					},
				},
				"else": bson.A{},
			}},
		}}},
	}

	// Ejecutar la consulta de agregación
	cursor, err := r.competitorUserColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error when aggregating user categories: %w", err)
	}
	defer cursor.Close(ctx)

	var result []*availability_dao.GetDayTimeSlotDAORes
	if err := cursor.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("error when decoding user categories: %w", err)
	}

	// Si no se encontraron categorías, devolver un slice vacío
	if len(result) == 0 {
		return []*availability_dao.GetDayTimeSlotDAORes{}, nil
	}

	return result, nil
}
