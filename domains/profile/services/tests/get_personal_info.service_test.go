package services_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"

	"github.com/DBrange/didis-comp-bk/domains/profile/ports/drivens/tests/mocks"
	"github.com/DBrange/didis-comp-bk/domains/profile/services"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestProfileSevice_GetPersonalInfo(t *testing.T) {
	username := "defaultUsername"
	password := "defaultPassword"
	phone := "123456789"
	image := "defaultImageURL"
	state := "defaultState"
	country := "defaultCountry"
	city := "defaultCity"
	lat := "0.0"
	long := "0.0"
	now := time.Now()

	personalInfoDTODefault := &dto.GetPersonalInfoByIDDTORes{
		ID:        "defaultID",
		FirstName: "John",
		LastName:  "Doe",
		Username:  &username,
		Birthdate: &now,
		Password:  &password,
		Email:     "john.doe@example.com",
		Phone:     &phone,
		Image:     &image,
		Genre:     models.GENRE("M"),
		Location: &dto.GetPersonalInfoLocationByIDRes{
			ID:      "defaultLocationID",
			State:   &state,
			Country: &country,
			City:    &city,
			Lat:     &lat,
			Long:    &long,
		},
	}
	
	testTable := map[string]struct {
		setup         func(mockForQueryProfile *mocks.MockForQueryingProfile)
		assertionFunc func(subTest *testing.T, personalInfo *dto.GetPersonalInfoByIDDTORes, err error)
		userID        string
	}{
		"personal information successfully obteined": {
			setup: func(mockForQueryProfile *mocks.MockForQueryingProfile) {
				mockForQueryProfile.EXPECT().GetPersonalInfoByID(gomock.Any(), "defaultID").Return(personalInfoDTODefault, nil)
			},
			assertionFunc: func(subTest *testing.T, personalInfo *dto.GetPersonalInfoByIDDTORes, err error) {
				assert.NoError(subTest, err)
				assert.NotNil(subTest, personalInfo)
				assert.Equal(subTest, personalInfoDTODefault.ID, personalInfo.ID)
			},
			userID: "defaultID",
		},

		"not found error": {
			setup: func(mockForQueryProfile *mocks.MockForQueryingProfile) {
				mockForQueryProfile.EXPECT().GetPersonalInfoByID(gomock.Any(), "").Return(nil, customerrors.ErrNotFound)
			},
			assertionFunc: func(subTest *testing.T, personalInfo *dto.GetPersonalInfoByIDDTORes, err error) {
				assert.Error(subTest, err)
				assert.NotNil(subTest, err)
				assert.Nil(subTest, personalInfo)

				var appErr customerrors.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, customerrors.ErrCodeNotFound, appErr.Code)
				} else {
					subTest.Errorf("expected AppError, got '%v'", err)
				}
			},
			userID: "",
		},

		"invalid id error": {
			setup: func(mockForQueryProfile *mocks.MockForQueryingProfile) {
				mockForQueryProfile.EXPECT().GetPersonalInfoByID(gomock.Any(), "badID").Return(nil, customerrors.ErrInvalidID)
			},
			assertionFunc: func(subTest *testing.T, personalInfo *dto.GetPersonalInfoByIDDTORes, err error) {
				assert.Error(subTest, err)
				assert.NotNil(subTest, err)
				assert.Nil(subTest, personalInfo)

				var appErr customerrors.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, customerrors.ErrCodeInvalidID, appErr.Code)
				} else {
					subTest.Errorf("expected AppError, got '%v'", err)
				}
			},
			userID: "badID",
		},

		"general error": {
			setup: func(mockForQueryProfile *mocks.MockForQueryingProfile) {
				mockForQueryProfile.EXPECT().GetPersonalInfoByID(gomock.Any(), "defaultID").Return(nil,errors.New("general error"))
			},
			assertionFunc: func(subTest *testing.T, personalInfo *dto.GetPersonalInfoByIDDTORes, err error) {
				assert.Error(subTest, err)
			},
			userID: "defaultID",
		},
	}

	for testName, test := range testTable {
		t.Run(testName, func(subTest *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockForProfileQueryer := mocks.NewMockForQueryingProfile(ctrl)

			profileService := services.NewProfileService(mockForProfileQueryer)

			test.setup(mockForProfileQueryer)

			personalInfo, err := profileService.GetPersonalInfoByID(context.Background(), test.userID)

			test.assertionFunc(subTest, personalInfo, err)
		})
	}
}
