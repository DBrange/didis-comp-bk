//go:generate mockgen -destination=tests/mocks/for_querying_profile.mock.go -package=mocks github.com/DBrange/didis-comp-bk/domains/profile/ports/drivens ForQueryingProfile

package ports

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	profile_dto "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ForQueryingProfile interface {
	InitialiseRole(ctx context.Context) error
	WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error
	CreateUser(ctx context.Context, userDAO *profile_dto.CreateUserDTOReq) (string, error)
	CreateLocation(ctx context.Context, locationInfoDAO *profile_dto.CreateLocationDTOReq) (string, error)
	GetRoleByNameAndType(ctx context.Context, roleName, roleType string) (*profile_dto.GetRoleDTOByID, error)
	CreateOrganizer(ctx context.Context, userID string) error
	CreateAvailability(ctx context.Context, userID, competitorID *string) error
	UpdateAvailability(ctx context.Context, availabilityID string, availabilityInfoDAO *profile_dto.UpdateDailyAvailabilityDTOReq) error
	UpdateUser(ctx context.Context, userID string, userInfoDAO *profile_dto.UpdateUserDTOReq) error
	UpdateLocation(ctx context.Context, locationID string, locationDAO *profile_dto.UpdateLocationDTOReq) error
	GetUserByID(ctx context.Context, userID string) (*profile_dto.GetUserByIDDTORes, error)
	GetLocationByID(ctx context.Context, locationID string) (*profile_dto.GetLocationByIDDTORes, error)
	GetDailyAvailabilityByID(ctx context.Context, availabilityID string, day string) (*profile_dto.GetDailyAvailabilityByIDDTORes, error)
	ConvertToObjectID(ID string) (*primitive.ObjectID, error)
	CreateCompetitor(ctx context.Context, sport models.SPORT, competitorType models.COMPETITOR_TYPE, ID string) (string, error)
	CreateCompetitorStats(ctx context.Context, competitorOID string) error
	CreateCompetitorUser(ctx context.Context, userID, competitorID string) error
	DeleteUser(ctx context.Context, userID string) (*user_dao.UserRelationsToDeleteDAOReq, error)
	GetDailyAvailabilityUserID(ctx context.Context, userID, day string) (*profile_dto.GetDailyAvailabilityByIDDTORes, error)
	SetDeletedAt(ctx context.Context, mc *mongo.Collection, ID string, name string) error
	AvailabilityColl() *mongo.Collection
	LocationColl() *mongo.Collection
	GetUserPasswordForLogin(ctx context.Context, username string) (string, string, error)
	GetUserRoles(ctx context.Context, userID string) ([]string, error)
	ActivateUserNotification(ctx context.Context)
	GetAvailabilityIDByUserID(ctx context.Context, userID string) (string, error)
	UpdateUserPassword(ctx context.Context, userID, newPassword string) error
	GetUserPasswordByID(ctx context.Context, userID string) (string, error)
	CreateSingle(ctx context.Context, singleInfoDTO *profile_dto.CreateSingleDTOReq) (string, error)
	CreateDouble(ctx context.Context, doubleInfoDTO *profile_dto.CreateDoubleDTOReq) (string, error)
	CreateTeam(ctx context.Context, teamInfoDTO *profile_dto.CreateTeamDTOReq) (string, error)
	CreateFollower(ctx context.Context, followerDTO *profile_dto.CreateFollowerDTOReq) error
	GetProfileInfoInCategory(ctx context.Context, categoryID, competitorID string) (*profile_dto.GetProfileInfoInCategoryDTORes, error)
	CreateAvailabilityForCompetitor(ctx context.Context, competitorID string, dailyAvailability []*profile_dto.GetDailyAvailabilityByIDDTORes) error
	GetAvailabilityDailySlice(ctx context.Context, userID, competitorID string) ([]*profile_dto.GetDailyAvailabilityByIDDTORes, error)
	GetDailyAvailabilityCompetitorID(ctx context.Context, competitorID string, day string) (*profile_dto.GetDailyAvailabilityByIDDTORes, error)
	GetCompetitorTournamentsInCategory(ctx context.Context, categoryID, competitorID, lastID string, limit int) ([]*profile_dto.GetTournamentsFromCategoryDTORes, error)
	VerifyFollowerExistsRelation(ctx context.Context, followerDAO *profile_dto.CreateFollowerDTOReq) error
}
