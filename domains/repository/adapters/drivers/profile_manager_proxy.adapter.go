package adapters

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
	organizer_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/organizer/dao"
	role_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/role/dao"
	single_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/single/dao"
	team_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/team/dao"
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

func (a *ProfileManagerProxyAdapter) InitialiseRole(ctx context.Context) error {
	return a.repository.InitialiseRole(ctx)
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

func (a *ProfileManagerProxyAdapter) CreateAvailability(ctx context.Context, userOID, competitorOID, tournamentOID *primitive.ObjectID) error {
	return a.repository.CreateAvailability(ctx, userOID, competitorOID, tournamentOID)
}

func (a *ProfileManagerProxyAdapter) ConvertToObjectID(ID string) (*primitive.ObjectID, error) {
	return a.repository.ConvertToObjectID(ID)
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
	return a.repository.GetUserPasswordByID(ctx, userID)
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

func (a *ProfileManagerProxyAdapter) GetDailyAvailabilityUserID(ctx context.Context, userID, day string) (*availability_dao.GetDailyAvailabilityByIDDAORes, *primitive.ObjectID, error) {
	return a.repository.GetDailyAvailabilityUserID(ctx, userID, day)
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

func (a *ProfileManagerProxyAdapter) GetUserForLogin(ctx context.Context, username string) (*user_dao.GetUserForLoginDAO, error) {
	return a.repository.GetUserForLogin(ctx, username)
}

func (a *ProfileManagerProxyAdapter) GetUserForRefreshToken(ctx context.Context, userID string) (*user_dao.GetUserForRefreshTokenDAO, error) {
	userOID, err := a.ConvertToObjectID(userID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetUserForRefreshToken(ctx, userOID)
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

func (a *ProfileManagerProxyAdapter) CreateSingle(ctx context.Context, singleInfoDAO *single_dao.CreateSingleDAOReq) (string, error) {
	return a.repository.CreateSingle(ctx, singleInfoDAO)
}

func (a *ProfileManagerProxyAdapter) CreateDouble(ctx context.Context, doubleInfoDAO *double_dao.CreateDoubleDAOReq) (string, error) {
	return a.repository.CreateDouble(ctx, doubleInfoDAO)
}

func (a *ProfileManagerProxyAdapter) CreateTeam(ctx context.Context, teamInfoDAO *team_dao.CreateTeamDAOReq) (string, error) {
	return a.repository.CreateTeam(ctx, teamInfoDAO)
}

func (a *ProfileManagerProxyAdapter) CreateFollower(ctx context.Context, followerDAO *follower_dao.CreateFollowerDAOReq) error {
	return a.repository.CreateFollower(ctx, followerDAO)
}

func (a *ProfileManagerProxyAdapter) GetProfileInfoInCategory(ctx context.Context, categoryID, competitorID string) (*category_registration_dao.GetProfileInfoInCategoryDAORes, error) {
	categoryOID, err := a.ConvertToObjectID(categoryID)
	if err != nil {
		return nil, err
	}

	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetProfileInfoInCategory(ctx, categoryOID, competitorOID)

}

func (a *ProfileManagerProxyAdapter) GetAvailabilityDailySlice(ctx context.Context, userID, competitorID string) ([]*availability_dao.GetDailyAvailabilityByIDDAORes, error) {
	var userOID, competitorOID *primitive.ObjectID
	// var err error

	if userID != "" {
		userOIDConv, err := a.ConvertToObjectID(userID)
		if err != nil {
			return nil, err
		}
		userOID = userOIDConv
	}

	if competitorID != "" {
		competitorOIDConv, err := a.ConvertToObjectID(competitorID)
		if err != nil {
			return nil, err
		}
		competitorOID = competitorOIDConv
	}

	return a.repository.GetAvailabilityDailySlice(ctx, userOID, competitorOID)

}

func (a *ProfileManagerProxyAdapter) CreateAvailabilityForCompetitor(ctx context.Context, competitorID string, dailyAvailability []*availability_dao.CreateDailyAvailability) error {
	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	return a.repository.CreateAvailabilityForCompetitor(ctx, competitorOID, dailyAvailability)
}

func (a *ProfileManagerProxyAdapter) GetDailyAvailabilityCompetitorID(ctx context.Context, competitorID string, day string) (*availability_dao.GetDailyAvailabilityByIDDAORes, *primitive.ObjectID, error) {
	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return nil, nil, err
	}

	return a.repository.GetDailyAvailabilityCompetitorID(ctx, competitorOID, day)

}

func (a *ProfileManagerProxyAdapter) GetCompetitorTournamentsInCategory(ctx context.Context, categoryID, competitorID, lastID string, limit int) ([]*category_dao.GetTournamentsFromCategoryDAORes, error) {
	categoryOID, err := a.ConvertToObjectID(categoryID)
	if err != nil {
		return nil, err
	}

	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return nil, err
	}

	var lastOID *primitive.ObjectID
	if lastID != "" {
		lastOID, err = a.ConvertToObjectID(lastID)
		if err != nil {
			return nil, err
		}
	} else {
		lastOID = nil
	}

	return a.repository.GetCompetitorTournamentsInCategory(ctx, categoryOID, competitorOID, lastOID, limit)

}

func (a *ProfileManagerProxyAdapter) VerifyFollowerExistsRelation(ctx context.Context, followerDAO *follower_dao.CreateFollowerDAOReq) error {
	return a.repository.VerifyFollowerExistsRelation(ctx, followerDAO)

}

func (a *ProfileManagerProxyAdapter) GetUserCategories(ctx context.Context, userID string, sport models.SPORT, limit int, lastID string) ([]*competitor_user_dao.GetUserCategoriesCategoryDAO, error) {
	userOID, err := a.ConvertToObjectID(userID)
	if err != nil {
		return nil, err
	}

	var lastOID *primitive.ObjectID
	if lastID != "" {
		lastOID, err = a.ConvertToObjectID(lastID)
		if err != nil {
			return nil, err
		}
	} else {
		lastOID = nil
	}

	return a.repository.GetUserCategories(ctx, userOID, sport, limit, lastOID)

}

func (a *ProfileManagerProxyAdapter) GetNumberFollowers(ctx context.Context, userOID *primitive.ObjectID) (int, error) {
	return a.repository.GetNumberFollowers(ctx, userOID)
}

func (a *ProfileManagerProxyAdapter) GetUserFollowers(ctx context.Context, userID string, name string, limit int, lastCreatedAt *time.Time) (*follower_dao.GetUserFollowersDAORes, error) {
	userOID, err := a.ConvertToObjectID(userID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetUserFollowers(ctx, userOID, name, limit, lastCreatedAt)
}

func (a *ProfileManagerProxyAdapter) GetUserPrimaryData(ctx context.Context, userID string) (*user_dao.GetUserPrimaryDataDAORes, error) {
	userOID, err := a.ConvertToObjectID(userID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetUserPrimaryData(ctx, userOID)
}

func (a *ProfileManagerProxyAdapter) IsFollowing(ctx context.Context, fromOID, userToOID *primitive.ObjectID) (bool, error) {
	return a.repository.IsFollowing(ctx, fromOID, userToOID)
}

func (a *ProfileManagerProxyAdapter) FollowOrUnfollow(ctx context.Context, followerDAO *follower_dao.CreateFollowerDAOReq) error {
	return a.repository.FollowOrUnfollow(ctx, followerDAO)
}

func (a *ProfileManagerProxyAdapter) VerifyUserExists(ctx context.Context, userOID *primitive.ObjectID) error {
	return a.repository.VerifyUserExists(ctx, userOID)
}

func (a *ProfileManagerProxyAdapter) GetRoleString(ctx context.Context, roleID string) (models.ROLE, error) {
	roleOID, err := a.ConvertToObjectID(roleID)
	if err != nil {
		return models.ROLE_COMPETITOR, err
	}

	return a.repository.GetRoleString(ctx, roleOID)
}

func (a *ProfileManagerProxyAdapter) GetOrganizerIDByUserID(ctx context.Context, userID string) (*string, error) {
	userOID, err := a.ConvertToObjectID(userID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetOrganizerIDByUserID(ctx, userOID)
}

func (a *ProfileManagerProxyAdapter) GetUserAllCompetitorSports(ctx context.Context, userID string) ([]models.SPORT, error) {
	userOID, err := a.ConvertToObjectID(userID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetUserAllCompetitorSports(ctx, userOID)
}

func (a *ProfileManagerProxyAdapter) GetOrganizerData(ctx context.Context, userID string) (*organizer_dao.GetOrganizerDataDAORes, error) {
	userOID, err := a.ConvertToObjectID(userID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetOrganizerData(ctx, userOID)
}
