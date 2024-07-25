package ports

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	avaliability_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	models_role "github.com/DBrange/didis-comp-bk/domains/repository/models/role"
	user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
	"go.mongodb.org/mongo-driver/mongo"
)

type ForManagingProfile interface {
	// CreateUserAndLocation(ctx context.Context, userDAO *user_dao.CreateUserDAO, locationDAO *location_dao.CreateLocationDAOReq, organizer bool) error
	UpdateProfileAvailability(ctx context.Context, availabilityID string, availabilityInfoDAO *avaliability_dao.UpdateDailyAvailabilityDAOReq) error
	UpdatePersonalInfo(ctx context.Context, userID string, userInfoDAO *user_dao.UpdateUserDAOReq, locationInfoDAO *location_dao.UpdateLocationDAOReq) error
	GetPersonalInfoByID(ctx context.Context, userID string) (*user_dao.GetUserByIDDAO, *location_dao.GetLocationByIDDAORes, error)
	GetProfileAvailabilityInfoByID(ctx context.Context, availabilityID string, day string) (*avaliability_dao.GetDailyAvailabilityInfoByIDDAORes, error)
	DeleteProfile(ctx context.Context, userID string) error
	UpdateProfilePassword(ctx context.Context, userID, newPassword, oldPassword string) error
	RegisterCompetitor(ctx context.Context, userID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) error

	CreateUser(ctx context.Context, userDAO *user_dao.CreateUserDAOReq) (string, error)
	// GetUserByID(ctx context.Context, userID string) (*user_dao.GetUserByIDDAO, error)
	// UpdateUser(ctx context.Context, userID string, newUserInfo *user_dao.UpdateUserDAOReq) error
	// DeleteUser(ctx context.Context, userID string) (*user_dao.UserRelationsToDeleteDAO, error)

	WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error
	CreateLocation(ctx context.Context, locationInfoDAO *location_dao.CreateLocationDAOReq) (string, error)
	GetRoleByNameAndType(ctx context.Context, roleName, roleType string) (*models_role.Role, error)
	CreateOrganizer(ctx context.Context, userID string) error
	CreateAvailability(ctx context.Context, userID, competitorID *string) error 
}
