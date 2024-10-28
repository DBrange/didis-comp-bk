package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	api_assets "github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/category_registration/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) CategoryRegistrationColl() *mongo.Collection {
	return r.categoryRegistrationColl
}

func (r *Repository) CreateCategoryRegistration(ctx context.Context, categoryRegistrationDAO *dao.CreateCategoryRegistrationDAOReq) error {
	categoryRegistrationDAO.SetTimeStamp()

	categoryRegistrationDAO.RegisteredPositions = []dao.RegistedPositionDAORes{}

	_, err := r.categoryRegistrationColl.InsertOne(ctx, categoryRegistrationDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("%w: error duplicate key for categoryRegistration: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return fmt.Errorf("%w: error categoryRegistration scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return fmt.Errorf("error when inserting categoryRegistration: %w", err)
	}

	return nil
}

func (r *Repository) VerifyCategoryExistsRelation(ctx context.Context, categoryRegistrationDAO *dao.CreateCategoryRegistrationDAOReq) error {
	filter := bson.M{
		"category_id":   categoryRegistrationDAO.CategoryID,
		"competitor_id": categoryRegistrationDAO.CompetitorID,
	}

	projection := bson.M{"_id": 1}

	opts := options.FindOne().SetProjection(projection)

	var documentFinded struct{}

	if err := r.categoryRegistrationColl.FindOne(ctx, filter, opts).Decode(&documentFinded); err == nil {
		return fmt.Errorf("%w: categoryRegistration already exists", customerrors.ErrAlreadyExits)
	} else if err != mongo.ErrNoDocuments {
		return fmt.Errorf("error when checking for existing categoryRegistration: %w", err)
	}

	return nil
}

func (r *Repository) GetCategoryRegistrationByID(ctx context.Context, categoryRegistrationID string) (*dao.GetCategoryRegistrationByIDDAORes, error) {
	var categoryRegistration dao.GetCategoryRegistrationByIDDAORes

	categoryRegistrationOID, err := r.ConvertToObjectID(categoryRegistrationID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *categoryRegistrationOID}

	err = r.categoryRegistrationColl.FindOne(ctx, filter).Decode(&categoryRegistration)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for categoryRegistration: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the categoryRegistration: %w", err)
	}

	return &categoryRegistration, nil
}

func (r *Repository) UpdateCategoryRegistration(ctx context.Context, categoryRegistrationID string, categoryRegistrationInfoDAO *dao.UpdateCategoryRegistrationDAOReq) error {
	categoryRegistrationOID, err := r.ConvertToObjectID(categoryRegistrationID)
	if err != nil {
		return err
	}

	categoryRegistrationInfoDAO.RenewUpdate()

	filter := bson.M{"_id": *categoryRegistrationOID}
	update, err := api_assets.StructToBsonMap(categoryRegistrationInfoDAO)
	if err != nil {
		return err
	}

	result, err := r.categoryRegistrationColl.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": update},
	)
	if err != nil {
		return fmt.Errorf("error updating categoryRegistration: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no categoryRegistration found with id: %s", customerrors.ErrNotFound, categoryRegistrationID)
	}

	return nil
}

func (r *Repository) DeleteCategoryRegistration(ctx context.Context, categoryRegistrationID string) error {
	err := r.SetDeletedAt(ctx, r.categoryRegistrationColl, categoryRegistrationID, "categoryRegistration")
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) PermaDeleteCategoryRegistration(ctx context.Context, mc *mongo.Collection, categoryRegistrationID string) error {
	return r.DeleteByID(ctx, mc, categoryRegistrationID, "categoryRegistration")
}

func (r *Repository) UpdateCompetitorPoints(ctx context.Context, categoryOID, competitorOID *primitive.ObjectID, points int) error {
	filter := bson.M{"category_id": categoryOID, "competitor_id": competitorOID}

	update := bson.M{"$set": bson.M{
		"points": points,
	}}

	result, err := r.categoryRegistrationColl.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return fmt.Errorf("error updating categoryRegistration: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no categoryRegistration found with category_id: %s or competitor_id %s", customerrors.ErrNotFound, categoryOID.Hex(), competitorOID.Hex())
	}

	return nil
}



func (r *Repository) GetProfileInfoInCategory(ctx context.Context, categoryOID, competitorOID *primitive.ObjectID) (*dao.GetProfileInfoInCategoryDAORes, error) {
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"category_id": categoryOID, "competitor_id": competitorOID}}},

		// Lookup competitor users and guest competitors
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "competitor_users",
			"localField":   "competitor_id",
			"foreignField": "competitor_id",
			"as":           "competitor_users",
		}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "guest_competitors",
			"localField":   "competitor_id",
			"foreignField": "competitor_id",
			"as":           "guest_competitors",
		}}},

		// Determine if it's a guest
		bson.D{{Key: "$addFields", Value: bson.M{
			"is_guest": bson.M{"$gt": bson.A{bson.M{"$size": "$guest_competitors"}, 0}},
		}}},

		// Lookup regular users and guest users
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "users",
			"localField":   "competitor_users.user_id",
			"foreignField": "_id",
			"as":           "users",
		}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "guest_users",
			"localField":   "guest_competitors.guest_user_id",
			"foreignField": "_id",
			"as":           "guest_users",
		}}},

		// Lookup competitor stats
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "competitor_stats",
			"localField":   "competitor_id",
			"foreignField": "competitor_id",
			"as":           "competitor_stats",
		}}},
		bson.D{{Key: "$unwind", Value: bson.M{"path": "$competitor_stats", "preserveNullAndEmptyArrays": true}}}, // Preserva la estructura si no hay stats

		// Unwind user and guest user arrays to simplify processing
		bson.D{{Key: "$unwind", Value: bson.M{"path": "$users", "preserveNullAndEmptyArrays": true}}},
		bson.D{{Key: "$unwind", Value: bson.M{"path": "$guest_users", "preserveNullAndEmptyArrays": true}}},

		// Group by competitor ID and separate users and guest users based on is_guest
		bson.D{{Key: "$group", Value: bson.M{
			"_id":      "$competitor_id",
			"is_guest": bson.M{"$first": "$is_guest"},
			"users": bson.M{
				"$push": bson.M{
					"$cond": bson.A{
						bson.M{"$eq": bson.A{"$is_guest", false}},
						bson.M{
							"_id":        "$users._id",
							"first_name": "$users.first_name",
							"last_name":  "$users.last_name",
							"image":      "$users.image",
						},
						bson.M{"$literal": bson.A{}}, // Empty array if it's not a regular user
					},
				},
			},
			"guest_users": bson.M{
				"$push": bson.M{
					"$cond": bson.A{
						bson.M{"$eq": bson.A{"$is_guest", true}},
						bson.M{
							"_id":        "$guest_users._id",
							"first_name": "$guest_users.first_name",
							"last_name":  "$guest_users.last_name",
							"image":      "$guest_users.image",
						},
						bson.M{"$literal": bson.A{}}, // Empty array if it's not a guest user
					},
				},
			},
			"points":               bson.M{"$first": "$points"},
			"current_position":     bson.M{"$first": "$current_position"},
			"registered_positions": bson.M{"$first": "$registered_positions"},
			"competitor_stats": bson.M{"$first": bson.M{
				"_id":             "$competitor_stats._id",
				"total_wins":      "$competitor_stats.total_wins",
				"total_losses":    "$competitor_stats.total_losses",
				"money_earned":    "$competitor_stats.money_earned",
				"tournaments_won": "$competitor_stats.tournaments_won",
			}},
		}}},

		// Ensure correct output structure
		bson.D{{Key: "$project", Value: bson.M{
			"_id":              "$_id",
			"points":           "$points",
			"current_position": "$current_position",
			"users": bson.M{
				"$cond": bson.M{
					"if":   bson.M{"$eq": bson.A{"$is_guest", false}},
					"then": "$users",
					"else": bson.A{}, // Empty array if it's a guest user
				},
			},
			"guest_users": bson.M{
				"$cond": bson.M{
					"if":   bson.M{"$eq": bson.A{"$is_guest", true}},
					"then": "$guest_users",
					"else": bson.A{}, // Empty array if it's a regular user
				},
			},
			"competitor_stats": "$competitor_stats",
		}}},
	}

	var competitorInfo dao.GetProfileInfoInCategoryDAORes

	cursor, err := r.categoryRegistrationColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error when searching for the categoryRegistration: %w", err)
	}
	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		if err := cursor.Decode(&competitorInfo); err != nil {
			return nil, fmt.Errorf("error when decoding categoryRegistration: %w", err)
		}
	} else {
		return nil, fmt.Errorf("%w: no categoryRegistration found for the given IDs", customerrors.ErrNotFound)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return &competitorInfo, nil
}

func (r *Repository) GetParticipantsOfCategory(ctx context.Context, categoryOID *primitive.ObjectID, sport models.SPORT, competitorType models.COMPETITOR_TYPE, limit int, lastOID *primitive.ObjectID) ([]*dao.GetCompetitorsOfCategoryDAORes, error) {
	var categoryRegistration []*dao.GetCompetitorsOfCategoryDAORes

	pipeline := r.getParticipantsOfCategoryBasePipeline(categoryOID, sport)
	pipeline = r.getParticipantsOfCategoryFinalStages(pipeline)
	pipeline = r.singlesOrDoublesCategoryFilter(pipeline, competitorType)
	pipeline = r.getParticipantsOfCategorySortAndPagination(pipeline, lastOID, limit)

	cursor, err := r.categoryRegistrationColl.Aggregate(ctx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for categoryRegistration: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the categoryRegistration: %w", err)
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &categoryRegistration); err != nil {
		return nil, fmt.Errorf("error when decoding categoryRegistration: %w", err)
	}

	return categoryRegistration, nil
}

func (r *Repository) GetCompetitorsOfCategoryByName(ctx context.Context, categoryOID *primitive.ObjectID, sport models.SPORT, competitorType models.COMPETITOR_TYPE, name string) ([]*dao.GetCompetitorsOfCategoryDAORes, error) {
	var categoryRegistration []*dao.GetCompetitorsOfCategoryDAORes

	// Si el nombre está vacío, retornar un slice vacío
	if strings.TrimSpace(name) == "" {
		return categoryRegistration, nil
	}

	// // Extracción de partes del nombre para la búsqueda
	// nameParts := strings.Fields(name)
	// var firstNameQuery, lastNameQuery string

	// if len(nameParts) > 0 {
	// 	firstNameQuery = nameParts[0]
	// }

	// if len(nameParts) > 1 {
	// 	lastNameQuery = strings.Join(nameParts[1:], " ")
	// }

	pipeline := r.getParticipantsOfCategoryBasePipeline(categoryOID, sport)
	pipeline = r.getParticipantsOfCategoryFinalStages(pipeline)
	pipeline = r.agetParticipantsOfCategoryNameFilter(pipeline, name, true)
	pipeline = r.singlesOrDoublesCategoryFilter(pipeline, competitorType)
	// pipeline = r.getParticipantsOfCategorySortAndPagination(pipeline, lastOID, limit)
	// Agregar etapa de ordenamiento
	sortStage := bson.D{{Key: "$sort", Value: bson.D{
		{Key: "_id", Value: 1}, // Siempre ordenar por _id para asegurar un orden consistente
	}}}
	pipeline = append(pipeline, sortStage)

		limitStage := bson.D{{Key: "$limit", Value: 10}}
	pipeline = append(pipeline, limitStage)
	
	cursor, err := r.categoryRegistrationColl.Aggregate(ctx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for categoryRegistration: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the categoryRegistration: %w", err)
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &categoryRegistration); err != nil {
		return nil, fmt.Errorf("error when decoding categoryRegistration: %w", err)
	}

	return categoryRegistration, nil
}

func (r *Repository) singlesOrDoublesCategoryFilter(pipeline mongo.Pipeline, competitorType models.COMPETITOR_TYPE) mongo.Pipeline {
	if competitorType == models.COMPETITOR_TYPE_SINGLE {
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.M{
			"$or": []bson.M{
				{"users": bson.M{"$size": 1}},
				{"guest_users": bson.M{"$size": 1}},
			},
		}}})
	} else {
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.M{
			"$or": []bson.M{
				{"users": bson.M{"$size": 2}},
				{"guest_users": bson.M{"$size": 2}},
			},
		}}})
	}

	return pipeline
}



func (r *Repository) getParticipantsOfCategoryBasePipeline(categoryOID *primitive.ObjectID, sport models.SPORT) mongo.Pipeline {
	return mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{
			"category_id": categoryOID,
		}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "competitors",
			"localField":   "competitor_id",
			"foreignField": "_id",
			"as":           "competitor",
		}}},
		bson.D{{Key: "$unwind", Value: "$competitor"}},
		bson.D{{Key: "$match", Value: bson.M{
			"competitor.sport": sport,
		}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "competitor_users",
			"localField":   "competitor_id",
			"foreignField": "competitor_id",
			"as":           "competitor_users",
		}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "guest_competitors",
			"localField":   "competitor_id",
			"foreignField": "competitor_id",
			"as":           "guest_competitors",
		}}},
		bson.D{{Key: "$addFields", Value: bson.M{
			"is_guest": bson.M{"$gt": bson.A{bson.M{"$size": "$guest_competitors"}, 0}},
		}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "users",
			"localField":   "competitor_users.user_id",
			"foreignField": "_id",
			"as":           "user",
		}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "guest_users",
			"localField":   "guest_competitors.guest_user_id",
			"foreignField": "_id",
			"as":           "guest_user",
		}}},
	}
}

func (r *Repository) getParticipantsOfCategorySortAndPagination(pipeline mongo.Pipeline, lastOID *primitive.ObjectID, limit int) mongo.Pipeline {
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

func (r *Repository) getParticipantsOfCategoryFinalStages(pipeline mongo.Pipeline) mongo.Pipeline {
	pipeline = append(pipeline, bson.D{{Key: "$group", Value: bson.M{
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
		"points":               bson.M{"$first": "$points"},
		"current_position":     bson.M{"$first": "$current_position"},
		"registered_positions": bson.M{"$first": "$registered_positions"},
	}}},

		// Limpieza de arrays vacíos
		bson.D{{Key: "$project", Value: bson.M{
			"_id":                  "$_id",
			"points":               "$points",
			"current_position":     "$current_position",
			"registered_positions": "$registered_positions",
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
		}}})

	return pipeline
}

func (r *Repository) agetParticipantsOfCategoryNameFilter(pipeline mongo.Pipeline, name string, guest bool) mongo.Pipeline {
    nameParts := strings.Fields(name)
    var firstNameQuery, lastNameQuery string
    
    if len(nameParts) > 0 {
        firstNameQuery = nameParts[0]
    }
    if len(nameParts) > 1 {
        lastNameQuery = strings.Join(nameParts[1:], " ")
    }

    orConditions := bson.A{
        // Coincidencias en usuarios regulares
        bson.M{"$and": bson.A{
            bson.M{"users.first_name": bson.M{"$regex": "^" + firstNameQuery, "$options": "i"}},
            bson.M{"users.last_name": bson.M{"$regex": "^" + lastNameQuery, "$options": "i"}},
        }},
        bson.M{"users.first_name": bson.M{"$regex": "^" + name, "$options": "i"}},
        bson.M{"users.last_name": bson.M{"$regex": "^" + name, "$options": "i"}},
    }

    // Añadir condiciones para usuarios invitados solo si guest es true
    if guest {
        guestConditions := bson.A{
            bson.M{"$and": bson.A{
                bson.M{"guest_users.first_name": bson.M{"$regex": "^" + firstNameQuery, "$options": "i"}},
                bson.M{"guest_users.last_name": bson.M{"$regex": "^" + lastNameQuery, "$options": "i"}},
            }},
            bson.M{"guest_users.first_name": bson.M{"$regex": "^" + name, "$options": "i"}},
            bson.M{"guest_users.last_name": bson.M{"$regex": "^" + name, "$options": "i"}},
        }
        orConditions = append(orConditions, guestConditions...)
    }

    // Aplicar el filtro de búsqueda
    pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.M{
        "$or": orConditions,
    }}})

    return pipeline
}

func (r *Repository) GetCompetitorsOutCategory(ctx context.Context, categoryOID *primitive.ObjectID, competitorOIDs []*primitive.ObjectID) ([]string, error) {
    // Encuentra los documentos que coincidan con los IDs en el slice.
    filter := bson.M{"category_id": categoryOID, "competitor_id": bson.M{"$in": competitorOIDs}}
    cursor, err := r.categoryRegistrationColl.Find(ctx, filter)
    if err != nil {
        return nil, fmt.Errorf("error finding category registrations: %v", err)
    }
    defer cursor.Close(ctx)

    // Extrae los IDs de los documentos que coincidan.
    foundIDsMap := make(map[string]struct{})
    for cursor.Next(ctx) {
        var doc struct {
            CompetitorID primitive.ObjectID `bson:"competitor_id"`
        }
        if err := cursor.Decode(&doc); err != nil {
            return nil, fmt.Errorf("error decoding document: %v", err)
        }
        foundIDsMap[doc.CompetitorID.Hex()] = struct{}{}
    }

    if err := cursor.Err(); err != nil {
        return nil, fmt.Errorf("error iterating cursor: %v", err)
    }

    // Compara los IDs encontrados con el slice original y encuentra los que no coincidan.
    var nonMatchingIDs []string
    for _, id := range competitorOIDs {
        if _, found := foundIDsMap[id.Hex()]; !found {
            nonMatchingIDs = append(nonMatchingIDs, id.Hex())
        }
    }

    return nonMatchingIDs, nil
}

func (r *Repository) GetCategoryRegistrationSortedByPoints(ctx context.Context, categoryOID *primitive.ObjectID) ([]*dao.GetCategoryRegistrationSortedByPointsDAORes, error) {
	// Obtener todos los competidores en la categoría ordenados por puntos en orden descendente
	filter := bson.M{"category_id": categoryOID}

	projection := bson.D{{Key: "points", Value: -1}}

	opts := options.Find().SetSort(projection)

	cursor, err := r.categoryRegistrationColl.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("error retrieving category registrations: %w", err)
	}
	defer cursor.Close(ctx)

	var categoryRegistration []*dao.GetCategoryRegistrationSortedByPointsDAORes

	if err := cursor.All(ctx, &categoryRegistration); err != nil {
		return nil, fmt.Errorf("error decoding category registrations: %w", err)
	}

	return categoryRegistration, nil
}

func (r *Repository) UpdateCategoryRegistrationCurrentPosition(ctx context.Context, categoryOID *primitive.ObjectID, categoryRegistration []*dao.GetCategoryRegistrationSortedByPointsDAORes) error {
	// Actualizar el campo current_position y agregar la nueva posición al historial
	for i, cr := range categoryRegistration {
		newPosition := i + 1 // MongoDB cursor is zero-indexed, positions are 1-indexed

		if cr.CurrentPosition == nil || *cr.CurrentPosition != newPosition {
			update := bson.M{
				"$set": bson.M{
					"current_position": newPosition,
				},
			}

			update["$push"] = bson.M{
				"registered_positions": bson.M{
					"date":     time.Now().UTC(),
					"position": newPosition,
				},
			}

			_, err := r.categoryRegistrationColl.UpdateOne(
				ctx,
				bson.M{"category_id": categoryOID, "competitor_id": cr.CompetitorID},
				update,
			)
			if err != nil {
				return fmt.Errorf("error updating category registration for competitor %s: %w", cr.CompetitorID.Hex(), err)
			}
		}
	}

	return nil
}

func (r *Repository) AddPointsInMultipleCategoryRegistration(ctx context.Context, categoryOID *primitive.ObjectID, competitorOIDs []*primitive.ObjectID, points int) error {
	// Filtro para seleccionar los documentos que coincidan con los IDs de los competidores
	filter := bson.M{"category_id": categoryOID, "competitor_id": bson.M{"$in": competitorOIDs}}

	// Operación de actualización para incrementar el campo total_prize
	update := bson.M{
		"$inc": bson.M{
			"points": points,
		},
	}

	// Ejecutar la actualización en la colección
	_, err := r.categoryRegistrationColl.UpdateMany(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error updating points: %v", err)
	}

	return nil
}
