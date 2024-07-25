package services_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	profile_dto "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/profile/ports/drivens/tests/mocks"
	"github.com/DBrange/didis-comp-bk/domains/profile/services"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_RegisterUser(t *testing.T) {
	var username = "johndoe"
	var password = "Password@123"
	var phone = "1234567890"
	var image = "https://example.com/image.jpg"
	var state = "California"
	var country = "USA"
	var city = "Los Angeles"
	var lat = "34.0522"
	var long = "-118.2437"
	var birthdate = time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)

	var profileDTO = &profile_dto.RegisterUserDTOReq{
		FirstName: "John",
		LastName:  "Doe",
		Username:  &username,
		Birthdate: &birthdate,
		Password:  &password,
		Email:     "johndoe@example.com",
		Phone:     &phone,
		Image:     &image,
		Genre:     models.GENRE_MALE,
		Location: &profile_dto.CreateLocationDTOReq{
			State:   &state,
			Country: &country,
			City:    &city,
			Lat:     &lat,
			Long:    &long,
		},
	}

	testTable := map[string]struct {
		setup         func(mockForQueryProfile *mocks.MockForQueryingProfile)
		assertionFunc func(subTest *testing.T, err error)
		profileDTO    *profile_dto.RegisterUserDTOReq
	}{

		"successful creation": {
			setup: func(mockForQueryProfile *mocks.MockForQueryingProfile) {
				mockForQueryProfile.EXPECT().RegisterUser(gomock.Any(), profileDTO).Return(nil)
			},
			assertionFunc: func(subTest *testing.T, err error) {
				assert.NoError(subTest, err)
			},
			profileDTO: profileDTO,
		},

		"duplicate key error": {
			setup: func(mockForQueryProfile *mocks.MockForQueryingProfile) {
				mockForQueryProfile.EXPECT().RegisterUser(gomock.Any(), profileDTO).Return(customerrors.ErrDuplicateKey)
			},
			assertionFunc: func(subTest *testing.T, err error) {
				assert.Error(subTest, err)
				assert.NotNil(subTest, err)

				var appErr customerrors.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, customerrors.ErrCodeDuplicateKey, appErr.Code)
				} else {
					subTest.Errorf("expected AppError, got '%v'", err)
				}
			},
			profileDTO: profileDTO,
		},

		"schema type error": {
			setup: func(mockForQueryProfile *mocks.MockForQueryingProfile) {
				mockForQueryProfile.EXPECT().RegisterUser(gomock.Any(), profileDTO).Return(customerrors.ErrSchemaViolation)
			},
			assertionFunc: func(subTest *testing.T, err error) {
				assert.Error(subTest, err)
				assert.NotNil(subTest, err)

				var appErr customerrors.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, customerrors.ErrCodeSchemaViolation, appErr.Code)
				} else {
					subTest.Errorf("expected AppError, got '%v'", err)
				}
			},
			profileDTO: profileDTO,
		},

		"not found error": {
			setup: func(mockForQueryProfile *mocks.MockForQueryingProfile) {
				mockForQueryProfile.EXPECT().RegisterUser(gomock.Any(), profileDTO).Return(customerrors.ErrNotFound)
			},
			assertionFunc: func(subTest *testing.T, err error) {
				assert.Error(subTest, err)
				assert.NotNil(subTest, err)

				var appErr customerrors.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, customerrors.ErrCodeNotFound, appErr.Code)
				} else {
					subTest.Errorf("expected AppError, got '%v'", err)
				}
			},
			profileDTO: profileDTO,
		},

		"unexpected error": {
			setup: func(mockForQueryProfile *mocks.MockForQueryingProfile) {
				mockForQueryProfile.EXPECT().RegisterUser(gomock.Any(), profileDTO).Return(errors.New("unexpected error"))
			},
			assertionFunc: func(subTest *testing.T, err error) {
				assert.Error(subTest, err)
			},
			profileDTO: profileDTO,
		},
	}

	for testName, test := range testTable {
		t.Run(testName, func(subTest *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockForProfileQueryer := mocks.NewMockForQueryingProfile(ctrl)

			profileService := services.NewProfileService(mockForProfileQueryer)

			test.setup(mockForProfileQueryer)

			err := profileService.RegisterUser(context.Background(), test.profileDTO)

			test.assertionFunc(subTest, err)
		})
	}
}
