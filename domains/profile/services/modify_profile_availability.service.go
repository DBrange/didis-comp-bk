package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/cmd/utils"
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *ProfileService) ModifyProfileAvailability(ctx context.Context, userID, availabilityID string, availabilityInfoDTO *models.UpdateDailyAvailabilityDTOReq) error {
	err := s.profileQuerier.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		// Obtener las disponibilidades de los competidores
		competitorIDs, err := s.profileQuerier.GetCompetitorIDsFromUser(sessCtx, userID)
		if err != nil {
			return customerrors.HandleErrMsg(err, "profile", "error updating profile daily availability")
		}

		// Actualizar la disponibilidad principal
		if err := s.profileQuerier.UpdateAvailability(sessCtx, availabilityID, availabilityInfoDTO); err != nil {
			return customerrors.HandleErrMsg(err, "profile", "error updating profile daily availability")
		}

		// Iterar sobre las disponibilidades de los competidores
		if err := s.UpdateCompetitorAvailability(ctx, competitorIDs, availabilityInfoDTO); err != nil {
			return customerrors.HandleErrMsg(err, "profile", "error updating profile daily availability")
		}

		return nil
	})

	if err != nil {
		return customerrors.HandleErrMsg(err, "profile", "error updating availability")
	}
	return nil
}

func (s *ProfileService) UpdateCompetitorAvailability(ctx context.Context, competitorIDs []string, availabilityInfoDTO *models.UpdateDailyAvailabilityDTOReq) error {
	timeSlot := availabilityInfoDTO.TimeSlots[0].TimeSlot

	for _, cID := range competitorIDs {
		availabiltyDTOs, err := s.profileQuerier.GetUsersAvailability(ctx, cID, availabilityInfoDTO.Day, timeSlot)
		if err != nil {
			return customerrors.HandleErrMsg(err, "profile", "error updating profile daily availability")
		}

		if len(availabiltyDTOs) == 1 {
			if err := s.updateSingleCompetitor(ctx, cID, availabilityInfoDTO); err != nil {
				return customerrors.HandleErrMsg(err, "profile", "error updating profile daily availability")
			}

			return nil
		}

		// if len(availabiltyDTOs) == 2, se realiza esto:
		if err := s.updateDoubleCompetitor(ctx, cID, availabiltyDTOs, availabilityInfoDTO); err != nil {
			return customerrors.HandleErrMsg(err, "profile", "error updating profile daily availability")
		}

	}
	return nil
}

func (s *ProfileService) updateSingleCompetitor(ctx context.Context, competitorID string, availabilityInfoDTO *models.UpdateDailyAvailabilityDTOReq) error {
	if err := s.profileQuerier.UpdateCompetitorAvailability(ctx, competitorID, availabilityInfoDTO); err != nil {
		return customerrors.HandleErrMsg(err, "profile", "error updating profile daily availability")
	}

	return nil
}

func (s *ProfileService) updateDoubleCompetitor(
	ctx context.Context,
	competitorID string,
	availabiltyDTOs []*dto.GetDayTimeSlotDTORes,
	availabilityInfoDTO *models.UpdateDailyAvailabilityDTOReq,
) error {
	for _, aDTO := range availabiltyDTOs {
		statusOne := aDTO.DailyAvailabilities[0].TimeSlots[0].Status
		statusTwo := aDTO.DailyAvailabilities[1].TimeSlots[1].Status

		intermediateStatus := utils.CombineAvailabilityStatuses(statusOne, statusTwo)

		// Crear una copia del availabilityInfoDTO para no modificar el original
		availabilityInfoCopy := availabilityInfoDTODeepCopy(availabilityInfoDTO)

		// Modificar la copia
		availabilityInfoCopy.TimeSlots[0].Status = intermediateStatus

		// Actualizar la disponibilidad con la copia modificada
		if err := s.profileQuerier.UpdateCompetitorAvailability(ctx, competitorID, availabilityInfoCopy); err != nil {
			return customerrors.HandleErrMsg(err, "profile", "error updating profile daily availability")
		}

	}

	return nil
}

func availabilityInfoDTODeepCopy(dto *models.UpdateDailyAvailabilityDTOReq) *models.UpdateDailyAvailabilityDTOReq {
	copy := *dto // Crear una copia superficial

	// Copiar TimeSlots de forma profunda
	copy.TimeSlots = make([]*models.UpdateTimeSlotDTOReq, len(dto.TimeSlots))
	for i, ts := range dto.TimeSlots {
		copy.TimeSlots[i] = &models.UpdateTimeSlotDTOReq{
			TimeSlot: ts.TimeSlot,
			Status:   ts.Status,
		}
	}

	return &copy
}

// package services

// import (
// 	"context"

// 	"github.com/DBrange/didis-comp-bk/cmd/api/models"
// 	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
// )

// func (d *ProfileService) ModifyProfileAvailability(ctx context.Context, availabilityID string, availabilityInfoDTO *models.UpdateDailyAvailabilityDTOReq) error {
// 	if err := s.profileQuerier.UpdateAvailability(ctx, availabilityID, availabilityInfoDTO); err != nil {
// 		return customerrors.HandleErrMsg(err, "profile", "error updating profile daily availability")
// 	}

// 	return nil
// }
