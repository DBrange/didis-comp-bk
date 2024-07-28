package adapters

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	availability_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	role_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/role/dao"
	user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
	"github.com/DBrange/didis-comp-bk/domains/repository/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (a *ProfileManagerProxyAdapter) WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error {
	return a.repository.WithTransaction(ctx, fn)
}

func (a *ProfileManagerProxyAdapter) CreateUser(ctx context.Context, userDAO *user_dao.CreateUserDAOReq) (string, error) {
	return a.repository.CreateUser(ctx, userDAO)
}

func (a *ProfileManagerProxyAdapter) CreateLocation(ctx context.Context, locationInfoDAO *location_dao.CreateLocationDAOReq) (string, error) {
	return a.repository.CreateLocation(ctx, locationInfoDAO)
}

func (a *ProfileManagerProxyAdapter) GetRoleByNameAndType(ctx context.Context, roleName, roleType string) (*role_dao.GetRoleDAOByID, error) {
	return a.repository.GetRoleByNameAndType(ctx, roleName, roleType)
}

func (a *ProfileManagerProxyAdapter) CreateOrganizer(ctx context.Context, userID string) error {
	return a.repository.CreateOrganizer(ctx, userID)
}

func (a *ProfileManagerProxyAdapter) CreateAvailability(ctx context.Context, userID, competitorID *string) error {
	return a.repository.CreateAvailability(ctx, userID, competitorID)
}

func (a *ProfileManagerProxyAdapter) ConvertToObjectID(ID string) (*primitive.ObjectID, error) {
	return a.repository.ConvertToObjectID(ID)
}

func (a *ProfileManagerProxyAdapter) CreateCompetitorType(ctx context.Context, competitorType models.COMPETITOR_TYPE) (*primitive.ObjectID, error) {
	return a.repository.CreateCompetitorType(ctx, competitorType)
}

func (a *ProfileManagerProxyAdapter) CreateCompetitorUser(ctx context.Context, userOID *primitive.ObjectID, competitorOID *primitive.ObjectID) error {
	return a.repository.CreateCompetitorUser(ctx, userOID, competitorOID)
}

func (a *ProfileManagerProxyAdapter) CreateCompetitor(ctx context.Context, sport models.SPORT, competitorType models.COMPETITOR_TYPE, OID *primitive.ObjectID) (string, error) {
	return a.repository.CreateCompetitor(ctx, sport, competitorType, OID)
}

func (a *ProfileManagerProxyAdapter) CreateCompetitorStats(ctx context.Context, competitorOID *primitive.ObjectID) error {
	return a.repository.CreateCompetitorStats(ctx, competitorOID)
}

func (a *ProfileManagerProxyAdapter) UpdateUserPassword(ctx context.Context, userID, newPassword string) error {
	return a.repository.UpdateUserPassword(ctx, userID, newPassword)
}

func (a *ProfileManagerProxyAdapter) GetUserPasswordByID(ctx context.Context, userID string) (string, error) {
	return a.repository.GetUserPasswordByID(ctx , userID ) 
}

func (a *ProfileManagerProxyAdapter) UpdateAvailability(ctx context.Context, availabilityID string, availabilityInfoDAO *availability_dao.UpdateDailyAvailabilityDAOReq) error {
	return a.repository.UpdateAvailability(ctx, availabilityID, availabilityInfoDAO)
}

func (a *ProfileManagerProxyAdapter) UpdateUser(ctx context.Context, userID string, userDAO *user_dao.UpdateUserDAOReq) error {
	return a.repository.UpdateUser(ctx, userID, userDAO)
}

func (a *ProfileManagerProxyAdapter) UpdateLocation(ctx context.Context, locationID string, locationDAO *location_dao.UpdateLocationDAOReq) error {
	return a.repository.UpdateLocation(ctx, locationID, locationDAO)
}

func (a *ProfileManagerProxyAdapter) GetUserByID(ctx context.Context, userID string) (*user_dao.GetUserByIDDAORes, error) {
	return a.repository.GetUserByID(ctx, userID)
}

func (a *ProfileManagerProxyAdapter) GetLocationByID(ctx context.Context, locationID string) (*location_dao.GetLocationByIDDAORes, error) {
	return a.repository.GetLocationByID(ctx, locationID)
}

func (a *ProfileManagerProxyAdapter) GetDailyAvailabilityByID(ctx context.Context, availabilityID string, day string) (*availability_dao.GetDailyAvailabilityByIDDAORes, error) {
	return a.repository.GetDailyAvailabilityByID(ctx, availabilityID, day)
}

func (a *ProfileManagerProxyAdapter) DeleteUser(ctx context.Context, userID string) (*user_dao.UserRelationsToDeleteDAOReq, error) {
	return a.repository.DeleteUser(ctx, userID)
}

func (a *ProfileManagerProxyAdapter) GetDailyAvailabilityByUserID(ctx context.Context, userID, day string) (*availability_dao.GetDailyAvailabilityByIDDAORes, error) {
	return a.repository.GetDailyAvailabilityByUserID(ctx, userID, day)
}

func (a *ProfileManagerProxyAdapter) SetDeletedAt(ctx context.Context, mc *mongo.Collection, ID string, name string) error {
	return a.repository.SetDeletedAt(ctx, mc, ID, name)
}

func (a *ProfileManagerProxyAdapter) AvailabilityColl() *mongo.Collection {
	return a.repository.AvailabilityColl()
}

func (a *ProfileManagerProxyAdapter) LocationColl() *mongo.Collection {
	return a.repository.LocationColl()
}

func (a *ProfileManagerProxyAdapter) GetUserPasswordForLogin(ctx context.Context, username string) (string, string, error) {
	return a.repository.GetUserPasswordForLogin(ctx, username)
}

func (a *ProfileManagerProxyAdapter) GetUserRoles(ctx context.Context, userID string) ([]string, error) {
	return a.repository.GetUserRoles(ctx, userID)
}

func (a *ProfileManagerProxyAdapter) ActivateUserNotification(ctx context.Context) {
	a.repository.ActivateUserNotification(ctx)
}

func (a *ProfileManagerProxyAdapter) GetAvailabilityIDByUserID(ctx context.Context, userID string) (string, error) {
	return a.repository.GetAvailabilityIDByUserID(ctx, userID)
}
