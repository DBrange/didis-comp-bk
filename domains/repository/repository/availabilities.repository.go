package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	availability_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) AvailabilityColl() *mongo.Collection {
	return r.availabilityColl
}

func (r *Repository) CreateAvailability(ctx context.Context, userOID, competitorOID, tournamentOID *primitive.ObjectID) error {
	defaultAvailability := r.generateDefaultAvailability()
	currentDate := time.Now().UTC()

	availability := &availability_dao.CreateAvailability{
		DailyAvailabilities: defaultAvailability,
		CreatedAt:           currentDate,
		UpdatedAt:           currentDate,
	}

	if userOID != nil {
		availability.UserID = userOID
	} else if competitorOID != nil {
		availability.CompetitorID = competitorOID
	} else {
		availability.TournamentID = tournamentOID
	}

	_, err := r.availabilityColl.InsertOne(ctx, &availability)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return fmt.Errorf("%w: error 'availability' scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return fmt.Errorf("error when inserting 'availability': %w", err)
	}

	return nil
}

func (r *Repository) CreateAvailabilityForCompetitor(ctx context.Context, competitorOID *primitive.ObjectID, dailyAvailability []*availability_dao.CreateDailyAvailability) error {
	currentDate := time.Now().UTC()

	availability := &availability_dao.CreateAvailability{
		DailyAvailabilities: dailyAvailability,
		CreatedAt:           currentDate,
		UpdatedAt:           currentDate,
	}

	availability.CompetitorID = competitorOID

	_, err := r.availabilityColl.InsertOne(ctx, &availability)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return fmt.Errorf("%w: error 'availability' scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return fmt.Errorf("error when inserting 'availability': %w", err)
	}

	return nil
}

func (r *Repository) generateDefaultAvailability() []*availability_dao.CreateDailyAvailability {
	daysOfWeek := []models.DAY{"SUNDAY", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY"}

	// Crear franjas horarias de cada hora (00:00 a 23:00) solo una vez
	timeSlots := make([]*availability_dao.CreateTimeSlot, 24)
	for hour := 0; hour < 24; hour++ {
		time := fmt.Sprintf("%02d:00", hour)
		timeSlots[hour] = &availability_dao.CreateTimeSlot{TimeSlot: time, Status: models.AVAILABILITY_STATUS_NOT_AVAILABLE}
	}

	// Crear disponibilidad para cada día de la semana utilizando la misma referencia de timeSlots
	dailyAvailability := make([]*availability_dao.CreateDailyAvailability, len(daysOfWeek))
	for i, day := range daysOfWeek {
		dailyAvailability[i] = &availability_dao.CreateDailyAvailability{Day: day, TimeSlots: timeSlots}
	}

	return dailyAvailability
}

func (r *Repository) GetAvailabilityDailySlice(ctx context.Context, userOID, competitorOID *primitive.ObjectID) ([]*availability_dao.GetDailyAvailabilityByIDDAORes, error) {
	var filter bson.M

	if userOID != nil {
		filter = bson.M{"user_id": userOID}
	} else {
		filter = bson.M{"competitor_id": competitorOID}

	}

	pipeline := mongo.Pipeline{
		// Match the document with the given availability ID
		{{
			Key: "$match", Value: filter,
		}},
		// Unwind the daily availabilities array
		{{
			Key: "$unwind", Value: "$daily_availabilities",
		}},
		// Unwind the time slots array
		{{
			Key: "$unwind", Value: "$daily_availabilities.time_slots",
		}},
		// Group by the daily availability to collect all time slots
		{{
			Key: "$group", Value: bson.M{
				"_id": "$daily_availabilities.day",
				"time_slots": bson.M{
					"$push": bson.M{
						"time_slot": "$daily_availabilities.time_slots.time_slot",
						"status":    "$daily_availabilities.time_slots.status",
					},
				},
			},
		}},
		// Project the structure to match CreateDailyAvailability
		{{
			Key: "$project", Value: bson.M{
				"day":        "$_id",
				"time_slots": "$time_slots",
			},
		}},
	}

	var dailyAvailabilities []*availability_dao.GetDailyAvailabilityByIDDAORes

	cursor, err := r.availabilityColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error executing aggregation: %w", err)
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &dailyAvailabilities); err != nil {
		return nil, fmt.Errorf("error decoding result: %w", err)
	}

	return dailyAvailabilities, nil
}

func (r *Repository) UpdateAvailability(ctx context.Context, availabilityID string, availabilityInfoDAO *availability_dao.UpdateDailyAvailabilityDAOReq) error {
	availabilityOID, err := r.ConvertToObjectID(availabilityID)
	if err != nil {
		return err
	}

	// Filtra por el ID de disponibilidad y el día específico
	filter := bson.M{
		"_id":                      *availabilityOID,
		"daily_availabilities.day": availabilityInfoDAO.Day,
	}

	// Usar el operador $[<identifier>] para actualizar solo el time_slot correspondiente
	update := bson.M{
		"$set": bson.M{"daily_availabilities.$.time_slots.$[elem]": availabilityInfoDAO.TimeSlots[0]}, // Actualiza el time_slot con la información nueva
	}

	// Definir el filtro de array (usando el campo TimeSlot como identificador)
	arrayFilters := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{
			bson.M{"elem.time_slot": availabilityInfoDAO.TimeSlots[0].TimeSlot}, // Filtra por el campo TimeSlot
		},
	})

	result, err := r.availabilityColl.UpdateOne(
		ctx,
		filter,
		update,
		arrayFilters, // Aplica el filtro de array para afectar solo el time_slot correcto
	)
	if err != nil {
		return fmt.Errorf("%w: error updating 'availability': %s", customerrors.ErrUpdated, err.Error())
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no 'availability' found with id: %s", customerrors.ErrNotFound, availabilityID)
	}

	return nil
}

func (r *Repository) GetDailyAvailabilityByID(ctx context.Context, availabilityID string, day string) (*availability_dao.GetDailyAvailabilityByIDDAORes, error) {
	availabilityOID, err := r.ConvertToObjectID(availabilityID)
	if err != nil {
		return nil, err
	}

	var availability availability_dao.GetAvailabilityByIDDAORes

	projection := bson.M{
		"daily_availabilities.$": 1,
	}

	filter := bson.M{"_id": *availabilityOID, "daily_availabilities.day": day}

	opts := options.FindOne().SetProjection(projection)

	err = r.availabilityColl.FindOne(ctx, filter, opts).Decode(&availability)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for the 'availability': %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the availability: %w", err)
	}

	return availability.DailyAvailabilities[0], nil
}

func (r *Repository) GetDailyAvailabilityUserID(ctx context.Context, userID, day string) (*availability_dao.GetDailyAvailabilityByIDDAORes, *primitive.ObjectID, error) {
	userOID, err := r.ConvertToObjectID(userID)
	if err != nil {
		return nil, nil, err
	}

	var availability availability_dao.GetAvailabilityByIDDAORes

	projection := bson.M{
		"_id":                    1,
		"daily_availabilities.$": 1,
	}

	filter := bson.M{"user_id": *userOID, "daily_availabilities.day": day}

	opts := options.FindOne().SetProjection(projection)

	if err := r.availabilityColl.FindOne(ctx, filter, opts).Decode(&availability); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil, fmt.Errorf("%w: error when searching for the 'availability': %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, nil, fmt.Errorf("error when searching for the 'availability': %w", err)
	}

	return availability.DailyAvailabilities[0], availability.ID, nil
}

func (r *Repository) GetDailyAvailabilityCompetitorID(ctx context.Context, competitorOID *primitive.ObjectID, day string) (*availability_dao.GetDailyAvailabilityByIDDAORes, *primitive.ObjectID, error) {
	var availability availability_dao.GetAvailabilityByIDDAORes

	projection := bson.M{
		"_id":                    1,
		"daily_availabilities.$": 1,
	}

	filter := bson.M{"competitor_id": competitorOID, "daily_availabilities.day": day}

	opts := options.FindOne().SetProjection(projection)

	if err := r.availabilityColl.FindOne(ctx, filter, opts).Decode(&availability); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil, fmt.Errorf("%w: error when searching for the 'availability': %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, nil, fmt.Errorf("error when searching for the 'availability': %w", err)
	}

	return availability.DailyAvailabilities[0], availability.ID, nil
}

func (r *Repository) GetDailyAvailabilityTournamentID(ctx context.Context, tournamentOID *primitive.ObjectID, day string) (*availability_dao.GetDailyAvailabilityByIDDAORes, *primitive.ObjectID, error) {
	var availability availability_dao.GetAvailabilityByIDDAORes

	projection := bson.M{
		"_id":                    1,
		"daily_availabilities.$": 1,
	}

	filter := bson.M{"tournament_id": tournamentOID, "daily_availabilities.day": day}

	opts := options.FindOne().SetProjection(projection)

	if err := r.availabilityColl.FindOne(ctx, filter, opts).Decode(&availability); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil, fmt.Errorf("%w: error when searching for the 'availability': %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, nil, fmt.Errorf("error when searching for the 'availability': %w", err)
	}

	return availability.DailyAvailabilities[0], availability.ID, nil
}

func (r *Repository) GetAvailabilityIDByUserID(ctx context.Context, userID string) (string, error) {
	userOID, err := r.ConvertToObjectID(userID)
	if err != nil {
		return "", err
	}

	filter := bson.M{"user_id": *userOID}

	var result struct {
		ID primitive.ObjectID `bson:"_id"`
	}

	projection := bson.M{"_id": 1}

	opts := options.FindOne().SetProjection(projection)

	if err := r.availabilityColl.FindOne(ctx, filter, opts).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return "", fmt.Errorf("%w: error when searching for the 'availability': %s", customerrors.ErrNotFound, err.Error())
		}
		return "", fmt.Errorf("error when searching for the 'availability': %w", err)
	}

	return result.ID.Hex(), nil
}

func (r *Repository) GetAvailabilityByTournamentID(ctx context.Context, tournamentOID *primitive.ObjectID) ([]*availability_dao.GetDailyAvailabilityByIDDAORes, error) {
	filter := bson.M{"tournament_id": tournamentOID}

	var result availability_dao.GetAvailabilityByIDDAORes

	projection := bson.M{"daily_availabilities": 1, "_id": 0}

	opts := options.FindOne().SetProjection(projection)

	err := r.availabilityColl.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for the 'availability': %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the availability: %w", err)
	}

	return result.DailyAvailabilities, nil
}
