package services_test

import (
	"context"
	"errors"
	"testing"

	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/user/ports/drivens/tests/mocks"
	"github.com/DBrange/didis-comp-bk/domains/user/services"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_CreateUser(t *testing.T) {
	var userDTO = &user_dto.CreateUserDTOReq{
		FirstName: "Didier",
		LastName:  "Brange",
	}

	testTable := map[string]struct {
		setup         func(mockForQueryUser *mocks.MockForQueryingUser)
		assertionFunc func(subTest *testing.T, err error)
		userDTO       *user_dto.CreateUserDTOReq
	}{

		"successful creation": {
			setup: func(mockForQueryUser *mocks.MockForQueryingUser) {
				mockForQueryUser.EXPECT().CreateUser(gomock.Any(), userDTO).Return(nil)
			},
			assertionFunc: func(subTest *testing.T, err error) {
				assert.NoError(subTest, err)
			},
			userDTO: &user_dto.CreateUserDTOReq{
				FirstName: "Didier",
				LastName:  "Brange",
			},
		},

		"duplicate key error": {
			setup: func(mockForQueryUser *mocks.MockForQueryingUser) {
				mockForQueryUser.EXPECT().CreateUser(gomock.Any(), userDTO).Return(customerrors.ErrDuplicateKey)
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
			userDTO: &user_dto.CreateUserDTOReq{
				FirstName: "Didier",
				LastName:  "Brange",
			},
		},

		"general error": {
			setup: func(mockForQueryUser *mocks.MockForQueryingUser) {
				mockForQueryUser.EXPECT().CreateUser(gomock.Any(), userDTO).Return(errors.New("general error"))
			},
			assertionFunc: func(subTest *testing.T, err error) {
				assert.Error(subTest, err)
			},
			userDTO: &user_dto.CreateUserDTOReq{
				FirstName: "Didier",
				LastName:  "Brange",
			},
		},
	}

	for testName, test := range testTable {
		t.Run(testName, func(subTest *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockForUserQueryer := mocks.NewMockForQueryingUser(ctrl)

			userService := services.NewUserService(mockForUserQueryer)

			test.setup(mockForUserQueryer)

			err := userService.CreateUser(context.Background(), test.userDTO)

			test.assertionFunc(subTest, err)
		})
	}
}
