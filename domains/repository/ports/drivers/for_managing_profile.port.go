package ports

import (
	"context"

	avaliability_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
)

type ForManagingProfile interface {
	CreateUserAndLocation(ctx context.Context, userDAO *user_dao.CreateUserDAO, locationDAO *location_dao.CreateLocationDAOReq, organizer bool) error
	UpdateProfileAvailability(ctx context.Context, availabilityID string, availabilityInfoDAO *avaliability_dao.UpdateDailyAvailabilityDAOReq) error
	UpdatePersonalInfo(ctx context.Context, userID string, userInfoDAO *user_dao.UpdateUserDAOReq, locationInfoDAO *location_dao.UpdateLocationDAOReq) error
	GetPersonalInfoByID(ctx context.Context, userID string) (*user_dao.GetUserByIDDAO, *location_dao.GetLocationByIDDAORes, error)
	GetProfileAvailabilityInfoByID(ctx context.Context, availabilityID string, day string) (*avaliability_dao.GetDailyAvailabilityInfoByIDDAORes, error)
	DeleteProfile(ctx context.Context, userID string) error
	UpdateProfilePassword(ctx context.Context, userID, newPassword, oldPassword string) error
	// CreateUser(ctx context.Context, userDAO *profile_dao.CreateUserDAO) (string, error)
	// GetUserByID(ctx context.Context, userID string) (*profile_dao.GetUserByIDDAO, error)
	// UpdateUser(ctx context.Context, userID string, newUserInfo *profile_dao.UpdateUserDAOReq) error
	// DeleteUser(ctx context.Context, userID string) (*profile_dao.UserRelationsToDeleteDAO, error)
}
