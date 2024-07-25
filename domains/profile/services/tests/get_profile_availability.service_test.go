package services_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/profile/ports/drivens/tests/mocks"
	"github.com/DBrange/didis-comp-bk/domains/profile/services"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestProfileService_GetProfileAvailability(t *testing.T) {

	// Define los días de la semana
	daysOfWeek := []string{"SUNDAY", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY"}

	// Define las 24 horas del día en formato de 24 horas
	timeSlots := make([]*dto.GetProfileDailyTimeSlotInfoByIDDTORes, 24)
	for i := 0; i < 24; i++ {
		timeSlot := &dto.GetProfileDailyTimeSlotInfoByIDDTORes{
			TimeSlot: fmt.Sprintf("%02d:00", i),
			Status:   models.AVAILABILITY_STATUS("AVAILABLE"), // Cambia esto según tus necesidades
		}
		timeSlots[i] = timeSlot
	}

	// Crea la disponibilidad diaria para cada día de la semana
	profileAvailability := make([]*dto.GetProfileDailyAvailabilityInfoByIDDTORes, 7)
	for i, day := range daysOfWeek {
		profileAvailability[i] = &dto.GetProfileDailyAvailabilityInfoByIDDTORes{
			Day:       day,
			TimeSlots: timeSlots,
		}
	}
	// Este es el que vamos a usar
	var profileAvailabilityDefault *dto.GetProfileDailyAvailabilityInfoByIDDTORes
	// Imprime el resultado para verificar
	for _, availability := range profileAvailability {
		fmt.Printf("Day: %s\n", availability.Day)
		profileAvailabilityDefault = availability
		for _, slot := range availability.TimeSlots {
			fmt.Printf("TimeSlot: %s, Status: %s\n", slot.TimeSlot, slot.Status)
		}
		fmt.Println()
	}

	testTable := map[string]struct {
		setup         func(mockForQueryProfile *mocks.MockForQueryingProfile)
		assertionFunc func(subTest *testing.T, dailyAvailabilityInfo *dto.GetProfileDailyAvailabilityInfoByIDDTORes, err error)
		userID        string
		day           string
	}{

		"personal information successfully obteined": {
			setup: func(mockForQueryProfile *mocks.MockForQueryingProfile) {
				mockForQueryProfile.EXPECT().GetProfileAvailabilityInfoByID(gomock.Any(), "defaultID", "THUESDAY").Return(profileAvailabilityDefault, nil)
			},
			assertionFunc: func(subTest *testing.T, dailyAvailabilityInfo *dto.GetProfileDailyAvailabilityInfoByIDDTORes, err error) {
				assert.NoError(subTest, err)
				assert.NotNil(subTest, dailyAvailabilityInfo)
				assert.Equal(subTest, profileAvailabilityDefault.Day, dailyAvailabilityInfo.Day)
			},
			userID: "defaultID",
			day:    "THUESDAY",
		},

		"incorrect day": {
			setup: func(mockForQueryProfile *mocks.MockForQueryingProfile) {
				mockForQueryProfile.EXPECT().GetProfileAvailabilityInfoByID(gomock.Any(), "defaultID", "asdasd").Return(nil, customerrors.ErrNotFound)
			},
			assertionFunc: func(subTest *testing.T, dailyAvailabilityInfo *dto.GetProfileDailyAvailabilityInfoByIDDTORes, err error) {
				assert.Error(subTest, err)
				assert.NotNil(subTest, err)
				assert.Nil(subTest, dailyAvailabilityInfo)

				var appErr customerrors.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, customerrors.ErrCodeNotFound, appErr.Code)
				} else {
					subTest.Errorf("expected AppError, got '%v'", err)
				}
			},
			userID: "defaultID",
			day:    "asdasd",
		},

		"not found error": {
			setup: func(mockForQueryProfile *mocks.MockForQueryingProfile) {
				mockForQueryProfile.EXPECT().GetProfileAvailabilityInfoByID(gomock.Any(), "defaultID", "THUESDAY").Return(nil, customerrors.ErrNotFound)
			},
			assertionFunc: func(subTest *testing.T, dailyAvailabilityInfo *dto.GetProfileDailyAvailabilityInfoByIDDTORes, err error) {
				assert.Error(subTest, err)
				assert.NotNil(subTest, err)
				assert.Nil(subTest, dailyAvailabilityInfo)

				var appErr customerrors.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, customerrors.ErrCodeNotFound, appErr.Code)
				} else {
					subTest.Errorf("expected AppError, got '%v'", err)
				}
			},
			userID: "defaultID",
			day:    "THUESDAY",
		},

		"invalid id error": {
			setup: func(mockForQueryProfile *mocks.MockForQueryingProfile) {
				mockForQueryProfile.EXPECT().GetProfileAvailabilityInfoByID(gomock.Any(), "anything", "THUESDAY").Return(nil, customerrors.ErrInvalidID)
			},
			assertionFunc: func(subTest *testing.T, dailyAvailabilityInfo *dto.GetProfileDailyAvailabilityInfoByIDDTORes, err error) {
				assert.Error(subTest, err)
				assert.NotNil(subTest, err)
				assert.Nil(subTest, dailyAvailabilityInfo)

				var appErr customerrors.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, customerrors.ErrCodeInvalidID, appErr.Code)
				} else {
					subTest.Errorf("expected AppError, got '%v'", err)
				}
			},
			userID: "anything",
			day:    "THUESDAY",
		},

		"general error": {
			setup: func(mockForQueryProfile *mocks.MockForQueryingProfile) {
				mockForQueryProfile.EXPECT().GetProfileAvailabilityInfoByID(gomock.Any(), "defaultID", "THUESDAY").Return(nil, errors.New("general error"))
			},
			assertionFunc: func(subTest *testing.T, dailyAvailabilityInfo *dto.GetProfileDailyAvailabilityInfoByIDDTORes, err error) {
				assert.Error(subTest, err)
			},
			userID: "defaultID",
			day:    "THUESDAY",
		},
	}

	for testName, test := range testTable {
		t.Run(testName, func(subTest *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockForProfileQueryer := mocks.NewMockForQueryingProfile(ctrl)

			profileService := services.NewProfileService(mockForProfileQueryer)

			test.setup(mockForProfileQueryer)

			dailyAvailabilityInfo, err := profileService.GetProfileAvailabilityInfoByID(context.Background(), test.userID, test.day)

			test.assertionFunc(subTest, dailyAvailabilityInfo, err)
		})
	}
}
