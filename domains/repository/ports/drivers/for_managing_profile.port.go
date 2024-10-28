package ports

import (
	"context"
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	availability_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
	category_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/category/dao"
	double_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/double/dao"
	category_registration_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/category_registration/dao"
	competitor_user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/competitor_user/dao"
	follower_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/follower/dao"
	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	role_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/role/dao"
	organizer_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/organizer/dao"
	single_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/single/dao"
	team_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/team/dao"
	user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ForManagingProfile interface {
	InitialiseRole(ctx context.Context) error
	UpdateAvailability(ctx context.Context, availabilityID string, availabilityInfoDAO *availability_dao.UpdateDailyAvailabilityDAOReq) error
	UpdateUserPassword(ctx context.Context, userID, newPassword string) error
	GetUserPasswordByID(ctx context.Context, userID string) (string, error)
	WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error
	CreateUser(ctx context.Context, userDAO *user_dao.CreateUserDAOReq) (string, error)
	CreateLocation(ctx context.Context, locationInfoDAO *location_dao.CreateLocationDAOReq) (string, error)
	GetRoleByNameAndType(ctx context.Context, roleName, roleType string) (*role_dao.GetRoleDAOByID, error)
	CreateOrganizer(ctx context.Context, userID string) error
	CreateAvailability(ctx context.Context, userOID, competitorOID, tournamentOID *primitive.ObjectID) error
	UpdateUser(ctx context.Context, userID string, userDAO *user_dao.UpdateUserDAOReq) error
	UpdateLocation(ctx context.Context, locationID string, locationDAO *location_dao.UpdateLocationDAOReq) error
	GetUserByID(ctx context.Context, userID string) (*user_dao.GetUserByIDDAORes, error)
	GetLocationByID(ctx context.Context, locationID string) (*location_dao.GetLocationByIDDAORes, error)
	GetDailyAvailabilityByID(ctx context.Context, availabilityID string, day string) (*availability_dao.GetDailyAvailabilityByIDDAORes, error)
	ConvertToObjectID(ID string) (*primitive.ObjectID, error)
	CreateCompetitor(ctx context.Context, sport models.SPORT, competitorType models.COMPETITOR_TYPE, OID *primitive.ObjectID) (string, error)
	CreateCompetitorStats(ctx context.Context, competitorOID *primitive.ObjectID) error
	CreateCompetitorUser(ctx context.Context, userOID *primitive.ObjectID, competitorOID *primitive.ObjectID) error
	DeleteUser(ctx context.Context, userID string) (*user_dao.UserRelationsToDeleteDAOReq, error)
	GetDailyAvailabilityUserID(ctx context.Context, userID, day string) (*availability_dao.GetDailyAvailabilityByIDDAORes, *primitive.ObjectID, error)
	SetDeletedAt(ctx context.Context, mc *mongo.Collection, ID string, name string) error
	AvailabilityColl() *mongo.Collection
	LocationColl() *mongo.Collection
	GetUserForLogin(ctx context.Context, username string) (*user_dao.GetUserForLoginDAO, error)
	GetUserForRefreshToken(ctx context.Context, userID string) (*user_dao.GetUserForRefreshTokenDAO, error)
	GetUserRoles(ctx context.Context, userID string) ([]string, error)
	ActivateUserNotification(ctx context.Context)
	GetAvailabilityIDByUserID(ctx context.Context, userID string) (string, error)
	CreateSingle(ctx context.Context, singleInfoDAO *single_dao.CreateSingleDAOReq) (string, error)
	CreateDouble(ctx context.Context, doubleInfoDAO *double_dao.CreateDoubleDAOReq) (string, error)
	CreateTeam(ctx context.Context, teamInfoDAO *team_dao.CreateTeamDAOReq) (string, error)
	CreateFollower(ctx context.Context, followerDAO *follower_dao.CreateFollowerDAOReq) error
	GetProfileInfoInCategory(ctx context.Context, categoryID, competitorID string) (*category_registration_dao.GetProfileInfoInCategoryDAORes, error)
	GetAvailabilityDailySlice(ctx context.Context, userID, competitorID string) ([]*availability_dao.GetDailyAvailabilityByIDDAORes, error)
	CreateAvailabilityForCompetitor(ctx context.Context, competitorID string, dailyAvailability []*availability_dao.CreateDailyAvailability) error
	GetDailyAvailabilityCompetitorID(ctx context.Context, competitorID string, day string) (*availability_dao.GetDailyAvailabilityByIDDAORes, *primitive.ObjectID, error)
	GetCompetitorTournamentsInCategory(ctx context.Context, categoryID, competitorID, lastID string, limit int) ([]*category_dao.GetTournamentsFromCategoryDAORes, error)
	VerifyFollowerExistsRelation(ctx context.Context, followerDAO *follower_dao.CreateFollowerDAOReq) error
	GetUserCategories(ctx context.Context, userID string, sport models.SPORT, limit int, lastID string) ([]*competitor_user_dao.GetUserCategoriesCategoryDAO, error)
	GetNumberFollowers(ctx context.Context, userOID *primitive.ObjectID) (int, error)
	GetUserFollowers(ctx context.Context, userID string, name string, limit int, lastCreatedAt *time.Time) (*follower_dao.GetUserFollowersDAORes, error)
	GetUserPrimaryData(ctx context.Context, userID string) (*user_dao.GetUserPrimaryDataDAORes, error)
	IsFollowing(ctx context.Context, fromOID, userToOID *primitive.ObjectID) (bool, error)
	FollowOrUnfollow(ctx context.Context, followerDAO *follower_dao.CreateFollowerDAOReq) error
	VerifyUserExists(ctx context.Context, userOID *primitive.ObjectID) error
	GetRoleString(ctx context.Context, roleID string) (models.ROLE, error)
	GetOrganizerIDByUserID(ctx context.Context, userID string) (*string, error)
	GetUserAllCompetitorSports(ctx context.Context, userID string) ([]models.SPORT, error)
	GetOrganizerData(ctx context.Context, userID string) (*organizer_dao.GetOrganizerDataDAORes, error)
}
