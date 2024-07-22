package repository

import (
	"context"
	"fmt"
	"time"

	models "github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability"
	availability_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) CreateAvailability(ctx context.Context, userID string) error {
	defaultAvailability := r.generateDefaultAvailability()
	currentDate := time.Now().UTC()

	userOID, err := r.ConvertToObjectID(userID)
	if err != nil {
		return err
	}

	availability := &models.Availability{
		UserID:              *userOID,
		DailyAvailabilities: defaultAvailability,
		CreatedAt:           currentDate,
		UpdatedAt:           currentDate,
	}

	_, err = r.availabilityColl.InsertOne(ctx, &availability)
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

func (r *Repository) generateDefaultAvailability() []*models.DailyAvailability {
	daysOfWeek := []string{"SUNDAY", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY"}

	// Crear franjas horarias de cada hora (00:00 a 23:00) solo una vez
	timeSlots := make([]*models.TimeSlot, 24)
	for hour := 0; hour < 24; hour++ {
		time := fmt.Sprintf("%02d:00", hour)
		timeSlots[hour] = &models.TimeSlot{TimeSlot: time, Status: "not available"}
	}

	// Crear disponibilidad para cada día de la semana utilizando la misma referencia de timeSlots
	dailyAvailability := make([]*models.DailyAvailability, len(daysOfWeek))
	for i, day := range daysOfWeek {
		dailyAvailability[i] = &models.DailyAvailability{Day: day, TimeSlots: timeSlots}
	}

	return dailyAvailability
}

func (r *Repository) UpdateAvailability(ctx context.Context, availabilityID string, availabilityInfoDAO *availability_dao.UpdateDailyAvailabilityDAOReq) error {
	availabilityOID, err := r.ConvertToObjectID(availabilityID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": *availabilityOID, "daily_availabilities.day": availabilityInfoDAO.Day}

	update := bson.M{"daily_availabilities.$.time_slots": availabilityInfoDAO.TimeSlots}

	fmt.Printf("%+v", update)
	result, err := r.availabilityColl.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": update},
	)
	if err != nil {
		return fmt.Errorf("%w: error updating 'availablility': %s", customerrors.ErrUpdated, err.Error())
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no 'availablility' found with id: %s", customerrors.ErrNotFound, availabilityID)
	}

	return nil
}

func (r *Repository) GetAvailabilityInfoByID(ctx context.Context, availabilityID string, day string) (*availability_dao.GetDailyAvailabilityInfoByIDDAORes, error) {
	availabilityOID, err := r.ConvertToObjectID(availabilityID)
	if err != nil {
		return nil, err
	}

	var availability availability_dao.GetAvailabilityInfoByIDDAORes

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

func (r *Repository) GetAvailabilityByUserID(ctx context.Context, userID string) (string, error) {
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
