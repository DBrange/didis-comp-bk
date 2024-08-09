package services

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *ProfileService) RegisterCompetitor(ctx context.Context, userIDs []string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) error {
	err := s.profileQueryer.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		// Create type of competitor
		competiorTypeOID, err := s.CreateCompetitorType(sessCtx, competitorType)
		if err != nil {
			return nil
		}

		// Create competitor
		competitorID, err := s.profileQueryer.CreateCompetitor(sessCtx, sport, competitorType, competiorTypeOID)
		if err != nil {
			return err
		}

		switch len(userIDs) {
		case 1:
			err = s.registerCompetitorSingle(sessCtx, userIDs, competitorID)
		case 2:
			err = s.registerCompetitorDouble(sessCtx, userIDs, competitorID)
		default:
			err = fmt.Errorf("unsupported number of users: %d", len(userIDs))
		}

		return err
	})

	if err != nil {
		return customerrors.HandleErrMsg(err, "profile", "error when registering competitor")
	}

	return nil
}

func (s *ProfileService) registerCompetitorSingle(ctx context.Context, userIDs []string, competitorID string) error {
	userID := userIDs[0]
	// Get availability
	availabilitySliceDTO, err := s.profileQueryer.GetAvailabilityDailySlice(ctx, userID, "")
	if err != nil {
		return err
	}

	availabilitySliceOrder := orderAvailability(availabilitySliceDTO)

	// Create availability
	err = s.profileQueryer.CreateAvailabilityForCompetitor(ctx, competitorID, availabilitySliceOrder)
	if err != nil {
		return err
	}

	// Create competitor stats
	if err := s.profileQueryer.CreateCompetitorStats(ctx, competitorID); err != nil {
		return err
	}

	// Create competitor_user
	if err := s.profileQueryer.CreateCompetitorUser(ctx, userID, competitorID); err != nil {
		return err
	}

	return nil
}

func (s *ProfileService) registerCompetitorDouble(ctx context.Context, userIDs []string, competitorID string) error {
	// Availability users
	usersAvailabilitySliceDTO := make([][]*dto.GetDailyAvailabilityByIDDTORes, len(userIDs))

	// Get availability users
	for i, userID := range userIDs {
		// Get availability
		availabilitySliceDTO, err := s.profileQueryer.GetAvailabilityDailySlice(ctx, userID, "")
		if err != nil {
			return err
		}

		usersAvailabilitySliceDTO[i] = availabilitySliceDTO
	}

	// Get availability
	availabilitySliceDTO := intermediateAvailability(usersAvailabilitySliceDTO)

	// Create availability
	err := s.profileQueryer.CreateAvailabilityForCompetitor(ctx, competitorID, availabilitySliceDTO)
	if err != nil {
		return err
	}

	// Create competitor stats
	if err := s.profileQueryer.CreateCompetitorStats(ctx, competitorID); err != nil {
		return err
	}

	for _, userID := range userIDs {
		// Create competitor_user
		if err := s.profileQueryer.CreateCompetitorUser(ctx, userID, competitorID); err != nil {
			return err
		}
	}

	return nil
}

func (r *ProfileService) CreateCompetitorType(ctx context.Context, competitorType models.COMPETITOR_TYPE) (string, error) {
	type createTypeCompetitor func(ctx context.Context) (string, error)

	createMap := map[models.COMPETITOR_TYPE]createTypeCompetitor{
		models.COMPETITOR_TYPE_SINGLE: func(ctx context.Context) (string, error) {
			singleDTO := &dto.CreateSingleDTOReq{}
			return r.profileQueryer.CreateSingle(ctx, singleDTO)
		},
		models.COMPETITOR_TYPE_DOUBLE: func(ctx context.Context) (string, error) {
			doubleDTO := &dto.CreateDoubleDTOReq{}
			return r.profileQueryer.CreateDouble(ctx, doubleDTO)
		},
		// models.COMPETITOR_TYPE_TEAM: func(ctx context.Context) (string, error) {
		// 	teamDTO := &dto.CreateTeamDTOReq{}
		// 	teamDTO.Admins = []string{userID}
		// 	return r.profileQueryer.CreateTeam(ctx, teamDTO)
		// },
	}

	create, ok := createMap[competitorType]
	if !ok {
		err := fmt.Errorf("error competitor type no exists: %w", customerrors.ErrNotFound)
		return "", customerrors.HandleErrMsg(err, "profile", "error when registering competitor")
	}

	return create(ctx)
}

// []string{"SUNDAY", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY"}
func intermediateAvailability(usersAvailabilitySliceDTO [][]*dto.GetDailyAvailabilityByIDDTORes) []*dto.GetDailyAvailabilityByIDDTORes {
	if len(usersAvailabilitySliceDTO) != 2 {
		return nil // o manejar el error apropiadamente
	}

	daysOrder := []models.DAY{
		models.DAY_SUNDAY,
		models.DAY_MONDAY,
		models.DAY_TUESDAY,
		models.DAY_WEDNESDAY,
		models.DAY_THURSDAY,
		models.DAY_FRIDAY,
		models.DAY_SATURDAY,
	}

	// Crear mapas para acceder rápidamente a la disponibilidad por día
	availabilityOneMap := make(map[models.DAY]*dto.GetDailyAvailabilityByIDDTORes)
	availabilityTwoMap := make(map[models.DAY]*dto.GetDailyAvailabilityByIDDTORes)

	for _, daily := range usersAvailabilitySliceDTO[0] {
		availabilityOneMap[daily.Day] = daily
	}
	for _, daily := range usersAvailabilitySliceDTO[1] {
		availabilityTwoMap[daily.Day] = daily
	}

	intermediateAvailability := make([]*dto.GetDailyAvailabilityByIDDTORes, 7)

	for i, day := range daysOrder {
		dailyOne, existsOne := availabilityOneMap[day]
		dailyTwo, existsTwo := availabilityTwoMap[day]

		if !existsOne || !existsTwo {
			// Manejar el caso en que falta información para un día
			intermediateAvailability[i] = &dto.GetDailyAvailabilityByIDDTORes{
				Day:       day,
				TimeSlots: []*dto.GetDailyTimeSlotByIDDTORes{}, // o algún valor predeterminado
			}
			continue
		}

		intermediateAvailability[i] = &dto.GetDailyAvailabilityByIDDTORes{
			Day:       day,
			TimeSlots: intermediateTimeSlots(dailyOne.TimeSlots, dailyTwo.TimeSlots),
		}
	}

	return intermediateAvailability
}

func intermediateTimeSlots(timeSlotsOne, timeSlotsTwo []*dto.GetDailyTimeSlotByIDDTORes) []*dto.GetDailyTimeSlotByIDDTORes {
	intermediateTimeSlots := make([]*dto.GetDailyTimeSlotByIDDTORes, len(timeSlotsOne))
	for i, slotOne := range timeSlotsOne {
		intermediateTimeSlots[i] = &dto.GetDailyTimeSlotByIDDTORes{
			TimeSlot: slotOne.TimeSlot,
			Status:   combineAvailabilityStatuses(slotOne.Status, timeSlotsTwo[i].Status),
		}
	}

	return intermediateTimeSlots
}

func combineAvailabilityStatuses(status1, status2 models.AVAILABILITY_STATUS) models.AVAILABILITY_STATUS {
	// Si alguno de los estados es NOT_AVAILABLE, el resultado es NOT_AVAILABLE
	if status1 == models.AVAILABILITY_STATUS_NOT_AVAILABLE || status2 == models.AVAILABILITY_STATUS_NOT_AVAILABLE {
		return models.AVAILABILITY_STATUS_NOT_AVAILABLE
	}

	// Si ambos estados son AVAILABLE, el resultado es AVAILABLE
	if status1 == models.AVAILABILITY_STATUS_AVAILABLE && status2 == models.AVAILABILITY_STATUS_AVAILABLE {
		return models.AVAILABILITY_STATUS_AVAILABLE
	}

	// En cualquier otro caso (que incluye todas las combinaciones con POSSIBLY_AVAILABLE),
	// el resultado es POSSIBLY_AVAILABLE
	return models.AVAILABILITY_STATUS_POSSIBLY_AVAILABLE
}

func orderAvailability(availabilitySlice []*dto.GetDailyAvailabilityByIDDTORes) []*dto.GetDailyAvailabilityByIDDTORes {
	daysOrder := []models.DAY{
		models.DAY_SUNDAY,
		models.DAY_MONDAY,
		models.DAY_TUESDAY,
		models.DAY_WEDNESDAY,
		models.DAY_THURSDAY,
		models.DAY_FRIDAY,
		models.DAY_SATURDAY,
	}

	// Crear un mapa para acceder rápidamente a la disponibilidad por día
	availabilityMap := make(map[models.DAY]*dto.GetDailyAvailabilityByIDDTORes)
	for _, daily := range availabilitySlice {
		availabilityMap[daily.Day] = daily
	}

	orderedAvailability := make([]*dto.GetDailyAvailabilityByIDDTORes, 7)

	for i, day := range daysOrder {
		if daily, exists := availabilityMap[day]; exists {
			orderedAvailability[i] = daily
		} else {
			// Si no existe la disponibilidad para este día, crear una entrada vacía
			orderedAvailability[i] = &dto.GetDailyAvailabilityByIDDTORes{
				Day:       day,
				TimeSlots: []*dto.GetDailyTimeSlotByIDDTORes{},
			}
		}
	}

	return orderedAvailability
}
