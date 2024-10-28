package repository

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	competitor_user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/competitor_user/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
			"total_prize": "$tournament_details.total_prize",
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
			"_id":   "$competitor.sport",
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
