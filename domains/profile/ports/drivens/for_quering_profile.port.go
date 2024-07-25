//go:generate mockgen -destination=tests/mocks/for_querying_profile.mock.go -package=mocks github.com/DBrange/didis-comp-bk/domains/profile/ports/drivens ForQueryingProfile

package ports

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	profile_dto "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	models_role "github.com/DBrange/didis-comp-bk/domains/repository/models/role"
	"go.mongodb.org/mongo-driver/mongo"
)

type ForQueryingProfile interface {
	// RegisterUser(ctx context.Context, profileInfoDTO *profile_dto.RegisterUserDTOReq) error

	WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error
	CreateUser(ctx context.Context, userDAO *profile_dto.CreateUserDTOReq) (string, error)
	CreateLocation(ctx context.Context, locationInfoDAO *profile_dto.CreateLocationDTOReq) (string, error)
	GetRoleByNameAndType(ctx context.Context, roleName, roleType string) (*models_role.Role, error)
	CreateOrganizer(ctx context.Context, userID string) error
	CreateAvailability(ctx context.Context, userID, competitorID *string) error

	ModifyProfileAvailability(ctx context.Context, availabilityID string, availabilityInfoDTO *profile_dto.ModifyProfileDailyAvailabilityDTOReq) error
	ModifyPersonalInfo(ctx context.Context, userID string, userInfoDTO *profile_dto.ModifyPersonalInfoDTOReq) error
	GetPersonalInfoByID(ctx context.Context, userID string) (*profile_dto.GetPersonalInfoByIDDTORes, error)
	GetProfileAvailabilityInfoByID(ctx context.Context, userID string, day string) (*profile_dto.GetProfileDailyAvailabilityInfoByIDDTORes, error)
	CloseProfile(ctx context.Context, userID string) error
	ModifyPassword(ctx context.Context, userID, newPassword, oldPassword string) error
	RegisterCompetitor(ctx context.Context, userID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) error
}
