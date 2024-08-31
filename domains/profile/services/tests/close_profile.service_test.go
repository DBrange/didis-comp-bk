package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/DBrange/didis-comp-bk/domains/profile/ports/drivens/tests/mocks"
	"github.com/DBrange/didis-comp-bk/domains/profile/services"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_CloseProfile(t *testing.T) {
	testTable := map[string]struct {
		setup         func(mockForQueryProfile *mocks.MockForQueryingProfile)
		assertionFunc func(subTest *testing.T, err error)
		userID        string
	}{
		"profile successfully closed": {
			setup: func(mockForQueryProfile *mocks.MockForQueryingProfile) {
				mockForQueryProfile.EXPECT().CloseProfile(gomock.Any(), "abc").Return(nil)
			},
			assertionFunc: func(subTest *testing.T, err error) {
				assert.NoError(subTest, err)
			},
			userID: "abc",
		},

		"not found error": {
			setup: func(mockForQueryProfile *mocks.MockForQueryingProfile) {
				mockForQueryProfile.EXPECT().CloseProfile(gomock.Any(), "abc").Return(customerrors.ErrNotFound)
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
			userID: "abc",
		},

		"invalid id error": {
			setup: func(mockForQueryProfile *mocks.MockForQueryingProfile) {
				mockForQueryProfile.EXPECT().CloseProfile(gomock.Any(), "badID").Return(customerrors.ErrInvalidID)
			},
			assertionFunc: func(subTest *testing.T, err error) {
				assert.Error(subTest, err)
				assert.NotNil(subTest, err)

				var appErr customerrors.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, customerrors.ErrCodeInvalidID, appErr.Code)
				} else {
					subTest.Errorf("expected AppError, got '%v'", err)
				}
			},
			userID: "badID",
		},

		"updating error": {
			setup: func(mockForQueryProfile *mocks.MockForQueryingProfile) {
				mockForQueryProfile.EXPECT().CloseProfile(gomock.Any(), "abc").Return(customerrors.ErrUpdated)
			},
			assertionFunc: func(subTest *testing.T, err error) {
				assert.Error(subTest, err)
				assert.NotNil(subTest, err)

				var appErr customerrors.AppError
				if errors.As(err, &appErr) {
					assert.Equal(subTest, customerrors.ErrCodeUpdated, appErr.Code)
				} else {
					subTest.Errorf("expected AppError, got '%v'", err)
				}
			},
			userID: "abc",
		},

		"unexpected error": {
			setup: func(mockForQueryProfile *mocks.MockForQueryingProfile) {
				mockForQueryProfile.EXPECT().CloseProfile(gomock.Any(), "abc").Return(errors.New("unexpected error"))
			},
			assertionFunc: func(subTest *testing.T, err error) {
				assert.Error(subTest, err)
			},
			userID: "abc",
		},
	}

	for testName, test := range testTable {
		t.Run(testName, func(subTest *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockForProfileQuerier := mocks.NewMockForQueryingProfile(ctrl)

			profileService := services.NewProfileService(mockForProfileQuerier)

			test.setup(mockForProfileQuerier)

			err := profileService.CloseProfile(context.Background(), test.userID)

			test.assertionFunc(subTest, err)
		})
	}
}
