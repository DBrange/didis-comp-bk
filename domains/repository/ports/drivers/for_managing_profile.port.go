package ports

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	availability_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/role/dao"
	user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ForManagingProfile interface {
	UpdateAvailability(ctx context.Context, availabilityID string, availabilityInfoDAO *availability_dao.UpdateDailyAvailabilityDAOReq) error
	UpdateUserPassword(ctx context.Context, userID, newPassword string) error
	GetUserPasswordByID(ctx context.Context, userID string) (string, error)
	WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error
	CreateUser(ctx context.Context, userDAO *user_dao.CreateUserDAOReq) (string, error)
	CreateLocation(ctx context.Context, locationInfoDAO *location_dao.CreateLocationDAOReq) (string, error)
	GetRoleByNameAndType(ctx context.Context, roleName, roleType string) (*dao.GetRoleDAOByID, error)
	CreateOrganizer(ctx context.Context, userID string) error
	CreateAvailability(ctx context.Context, userID, competitorID *string) error
	UpdateUser(ctx context.Context, userID string, userDAO *user_dao.UpdateUserDAOReq) error
	UpdateLocation(ctx context.Context, locationID string, locationDAO *location_dao.UpdateLocationDAOReq) error
	GetUserByID(ctx context.Context, userID string) (*user_dao.GetUserByIDDAORes, error)
	GetLocationByID(ctx context.Context, locationID string) (*location_dao.GetLocationByIDDAORes, error)
	GetDailyAvailabilityByID(ctx context.Context, availabilityID string, day string) (*availability_dao.GetDailyAvailabilityByIDDAORes, error)
	ConvertToObjectID(ID string) (*primitive.ObjectID, error)
	CreateCompetitorType(ctx context.Context, competitorType models.COMPETITOR_TYPE) (*primitive.ObjectID, error)
	CreateCompetitor(ctx context.Context, sport models.SPORT, competitorType models.COMPETITOR_TYPE, OID *primitive.ObjectID) (string, error)
	CreateCompetitorStats(ctx context.Context, competitorOID *primitive.ObjectID) error
	CreateCompetitorUser(ctx context.Context, userOID *primitive.ObjectID, competitorOID *primitive.ObjectID) error
	DeleteUser(ctx context.Context, userID string) (*user_dao.UserRelationsToDeleteDAOReq, error)
	GetDailyAvailabilityByUserID(ctx context.Context, userID, day string) (*availability_dao.GetDailyAvailabilityByIDDAORes, error)
	SetDeletedAt(ctx context.Context, mc *mongo.Collection, ID string, name string) error
	AvailabilityColl() *mongo.Collection
	LocationColl() *mongo.Collection
	GetUserPasswordForLogin(ctx context.Context, username string) (string, string, error)
	GetUserRoles(ctx context.Context, userID string) ([]string, error)
	ActivateUserNotification(ctx context.Context)
	GetAvailabilityIDByUserID(ctx context.Context, userID string) (string, error)
}
