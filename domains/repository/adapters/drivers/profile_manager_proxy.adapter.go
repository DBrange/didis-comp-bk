package adapters

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	availability_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	role_models "github.com/DBrange/didis-comp-bk/domains/repository/models/role"
	user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
	"github.com/DBrange/didis-comp-bk/domains/repository/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileManagerProxyAdapter struct {
	repository *repository.Repository
}

func NewProfileManagerProxyAdapter(repository *repository.Repository) *ProfileManagerProxyAdapter {
	return &ProfileManagerProxyAdapter{
		repository: repository,
	}
}

// func (a *ProfileManagerProxyAdapter) CreateUserAndLocation(ctx context.Context, userInfoDAO *user_dao.CreateUserDAO, locationInfoDAO *location_dao.CreateLocationDAOReq, organizer bool) error {
// 	return a.repository.CreateUserAndLocation(ctx, userInfoDAO, locationInfoDAO, organizer)
// }

func (a *ProfileManagerProxyAdapter) WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error {
	return a.repository.WithTransaction(ctx, fn)
}

func (a *ProfileManagerProxyAdapter) CreateUser(ctx context.Context, userDAO *user_dao.CreateUserDAOReq) (string, error) {
	return a.repository.CreateUser(ctx, userDAO)
}

func (a *ProfileManagerProxyAdapter) CreateLocation(ctx context.Context, locationInfoDAO *location_dao.CreateLocationDAOReq) (string, error) {
	return a.repository.CreateLocation(ctx, locationInfoDAO)
}

func (a *ProfileManagerProxyAdapter) GetRoleByNameAndType(ctx context.Context, roleName, roleType string) (*role_models.Role, error) {
	return a.repository.GetRoleByNameAndType(ctx, roleName, roleType)
}

func (a *ProfileManagerProxyAdapter) CreateOrganizer(ctx context.Context, userID string) error {
	return a.repository.CreateOrganizer(ctx, userID)
}

func (a *ProfileManagerProxyAdapter) CreateAvailability(ctx context.Context, userID, competitorID *string) error {
	return a.repository.CreateAvailability(ctx, userID, competitorID)
}

func (a *ProfileManagerProxyAdapter) UpdateProfileAvailability(ctx context.Context, availabilityID string, availabilityInfoDAO *availability_dao.UpdateDailyAvailabilityDAOReq) error {
	return a.repository.UpdateProfileAvailability(ctx, availabilityID, availabilityInfoDAO)
}

func (a *ProfileManagerProxyAdapter) UpdatePersonalInfo(ctx context.Context, userID string, userInfoDAO *user_dao.UpdateUserDAOReq, locationInfoDAO *location_dao.UpdateLocationDAOReq) error {
	return a.repository.UpdatePersonalInfo(ctx, userID, userInfoDAO, locationInfoDAO)
}

func (a *ProfileManagerProxyAdapter) GetPersonalInfoByID(ctx context.Context, userID string) (*user_dao.GetUserByIDDAO, *location_dao.GetLocationByIDDAORes, error) {
	return a.repository.GetPersonalInfoByID(ctx, userID)
}

func (a *ProfileManagerProxyAdapter) GetProfileAvailabilityInfoByID(ctx context.Context, availabilityID string, day string) (*availability_dao.GetDailyAvailabilityInfoByIDDAORes, error) {
	return a.repository.GetProfileAvailabilityInfoByID(ctx, availabilityID, day)
}

func (a *ProfileManagerProxyAdapter) DeleteProfile(ctx context.Context, userID string) error {
	return a.repository.DeleteProfile(ctx, userID)
}

func (a *ProfileManagerProxyAdapter) UpdateProfilePassword(ctx context.Context, userID, newPassword, oldPassword string) error {
	return a.repository.UpdateProfilePassword(ctx, userID, newPassword, oldPassword)
}

func (a *ProfileManagerProxyAdapter) RegisterCompetitor(ctx context.Context, userID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) error {
	return a.repository.RegisterCompetitor(ctx, userID, sport, competitorType)
}

// func (a *ProfileManagerProxyAdapter) UpdateAvailability(ctx context.Context, availabilityID string, availabilityInfoDAO *availability_dao.UpdateDailyAvailabilityDAOReq) error {
// 	return a.repository.UpdateAvailability(ctx, userID, sport, competitorType)
// }

// func (a *UserManagerProxyAdapter) CreateUser(ctx context.Context, user *profile_dao.CreateUserDAO) (string, error) {
// 	return a.repository.CreateUser(ctx, user)
// }

// func (a *UserManagerProxyAdapter) GetUserByID(ctx context.Context, id string) (*profile_dao.GetUserByIDDAO, error) {
// 	return a.repository.GetUserByID(ctx, id)
// }

// func (a *UserManagerProxyAdapter) UpdateUser(ctx context.Context, userID string, newUserInfo *profile_dao.UpdateUserDAOReq) error {
// 	return a.repository.UpdateUser(ctx, userID, newUserInfo)
// }

// func (a *UserManagerProxyAdapter) DeleteUser(ctx context.Context, userID string) (*profile_dao.UserRelationsToDeleteDAO, error) {
// 	return a.repository.DeleteUser(ctx, userID)
// }
