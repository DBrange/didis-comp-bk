package adapters

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/profile/adapters/mappers"
	profile_dto "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
	ports "github.com/DBrange/didis-comp-bk/domains/repository/ports/drivers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileQuerierAdapter struct {
	adapter ports.ForManagingProfile
}

func NewProfileQuerierAdapter(adapter ports.ForManagingProfile) *ProfileQuerierAdapter {
	return &ProfileQuerierAdapter{
		adapter: adapter,
	}
}

func (a *ProfileQuerierAdapter) InitialiseRole(ctx context.Context) error {
	return a.adapter.InitialiseRole(ctx)
}

func (a *ProfileQuerierAdapter) WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error {
	return a.adapter.WithTransaction(ctx, fn)
}

func (a *ProfileQuerierAdapter) CreateUser(ctx context.Context, userDTO *profile_dto.CreateUserDTOReq) (string, error) {
	userDAO, err := mappers.CreateUserDTOtoDAO(userDTO, a.ConvertToObjectID)
	if err != nil {
		return "", err
	}

	return a.adapter.CreateUser(ctx, userDAO)
}

func (a *ProfileQuerierAdapter) CreateLocation(ctx context.Context, locationDTO *profile_dto.CreateLocationDTOReq) (string, error) {
	locationDAO := mappers.CreateLocationDTOtoDAO(locationDTO)

	return a.adapter.CreateLocation(ctx, locationDAO)
}

func (a *ProfileQuerierAdapter) GetRoleByNameAndType(ctx context.Context, roleName, roleType string) (*profile_dto.GetRoleDTOByID, error) {
	roleDAO, err := a.adapter.GetRoleByNameAndType(ctx, roleName, roleType)
	if err != nil {
		return nil, err
	}

	roleDTO := mappers.CreateRoleDAOtoDTO(roleDAO)

	return roleDTO, nil
}

func (a *ProfileQuerierAdapter) CreateOrganizer(ctx context.Context, userID string) error {
	return a.adapter.CreateOrganizer(ctx, userID)
}

func (a *ProfileQuerierAdapter) CreateAvailability(ctx context.Context, userID, competitorID, tournamentID *string) error {
	userOID, competitorOID, tournamentOID, err := mappers.CreateAvailabilityDTOtODAO(userID, competitorID, tournamentID, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.adapter.CreateAvailability(ctx, userOID, competitorOID, tournamentOID)
}

func (a *ProfileQuerierAdapter) UpdateAvailability(ctx context.Context, availabilityID string, availabilityDTO *profile_dto.UpdateDailyAvailabilityDTOReq) error {
	availabilityDAO := mappers.UpdateDailyAvailabilityDTOtoDAO(availabilityDTO)

	return a.adapter.UpdateAvailability(ctx, availabilityID, availabilityDAO)
}

func (a *ProfileQuerierAdapter) UpdateUser(ctx context.Context, userID string, userDTO *profile_dto.UpdateUserDTOReq) error {
	userDAO := mappers.UpdateUserDTOtoDAO(userDTO)

	return a.adapter.UpdateUser(ctx, userID, userDAO)
}

func (a *ProfileQuerierAdapter) UpdateLocation(ctx context.Context, locationID string, locationDTO *profile_dto.UpdateLocationDTOReq) error {
	locationDAO := mappers.UpdateLocationDTOtoDAO(locationDTO)

	return a.adapter.UpdateLocation(ctx, locationID, locationDAO)
}

func (a *ProfileQuerierAdapter) GetUserByID(ctx context.Context, userID string) (*profile_dto.GetUserByIDDTORes, error) {
	userDAO, err := a.adapter.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	userDTO := mappers.GetUserByIDDAOtoDTO(userDAO)

	return userDTO, err
}

func (a *ProfileQuerierAdapter) GetLocationByID(ctx context.Context, locationID string) (*profile_dto.GetLocationByIDDTORes, error) {
	locationDAO, err := a.adapter.GetLocationByID(ctx, locationID)
	if err != nil {
		return nil, err
	}

	locationDTO := mappers.GetLocationByIDDAOtoDTO(locationDAO)

	return locationDTO, nil
}

func (a *ProfileQuerierAdapter) GetDailyAvailabilityByID(ctx context.Context, availabilityID string, day string) (*profile_dto.GetDailyAvailabilityByIDDTORes, error) {
	availabilityDAO, err := a.adapter.GetDailyAvailabilityByID(ctx, availabilityID, day)
	if err != nil {
		return nil, err
	}

	availabilityDTO := mappers.GetDailyAvailabilityByIDDAOtoDTO(availabilityDAO)

	return availabilityDTO, nil
}

func (a *ProfileQuerierAdapter) ConvertToObjectID(ID string) (*primitive.ObjectID, error) {
	return a.adapter.ConvertToObjectID(ID)
}

func (a *ProfileQuerierAdapter) CreateCompetitor(ctx context.Context, sport models.SPORT, competitorType models.COMPETITOR_TYPE, ID string) (string, error) {
	OID, err := a.ConvertToObjectID(ID)
	if err != nil {
		return "", nil
	}

	return a.adapter.CreateCompetitor(ctx, sport, competitorType, OID)
}

func (a *ProfileQuerierAdapter) CreateCompetitorStats(ctx context.Context, competitorID string) error {
	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	return a.adapter.CreateCompetitorStats(ctx, competitorOID)
}

func (a *ProfileQuerierAdapter) CreateCompetitorUser(ctx context.Context, userID, competitorID string) error {
	userOID, err := a.ConvertToObjectID(userID)
	if err != nil {
		return err
	}
	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	return a.adapter.CreateCompetitorUser(ctx, userOID, competitorOID)
}

func (a *ProfileQuerierAdapter) DeleteUser(ctx context.Context, userID string) (*user_dao.UserRelationsToDeleteDAOReq, error) {
	return a.adapter.DeleteUser(ctx, userID)
}

func (a *ProfileQuerierAdapter) GetDailyAvailabilityUserID(ctx context.Context, userID, day string) (*profile_dto.GetDailyAvailabilityByIDDTORes, error) {
	availabilityDAO, err := a.adapter.GetDailyAvailabilityUserID(ctx, userID, day)
	if err != nil {
		return nil, err
	}

	availabilityDTO := mappers.GetDailyAvailabilityByIDDAOtoDTO(availabilityDAO)

	return availabilityDTO, nil
}

func (a *ProfileQuerierAdapter) SetDeletedAt(ctx context.Context, mc *mongo.Collection, ID string, name string) error {
	return a.adapter.SetDeletedAt(ctx, mc, ID, name)
}

func (a *ProfileQuerierAdapter) AvailabilityColl() *mongo.Collection {
	return a.adapter.AvailabilityColl()
}

func (a *ProfileQuerierAdapter) LocationColl() *mongo.Collection {
	return a.adapter.LocationColl()
}

func (a *ProfileQuerierAdapter) UpdateUserPassword(ctx context.Context, userID, newPassword string) error {
	return a.adapter.UpdateUserPassword(ctx, userID, newPassword)
}

func (a *ProfileQuerierAdapter) GetUserPasswordByID(ctx context.Context, userID string) (string, error) {
	return a.adapter.GetUserPasswordByID(ctx, userID)
}

func (a *ProfileQuerierAdapter) GetUserPasswordForLogin(ctx context.Context, username string) (string, string, error) {
	return a.adapter.GetUserPasswordForLogin(ctx, username)
}

func (a *ProfileQuerierAdapter) GetUserRoles(ctx context.Context, userID string) ([]string, error) {
	return a.adapter.GetUserRoles(ctx, userID)
}

func (a *ProfileQuerierAdapter) ActivateUserNotification(ctx context.Context) {
	a.adapter.ActivateUserNotification(ctx)
}

func (a *ProfileQuerierAdapter) GetAvailabilityIDByUserID(ctx context.Context, userID string) (string, error) {
	return a.adapter.GetAvailabilityIDByUserID(ctx, userID)
}

func (a *ProfileQuerierAdapter) CreateSingle(ctx context.Context, singleDTO *profile_dto.CreateSingleDTOReq) (string, error) {
	singleDAO := mappers.CreateSingleDTOtoDAO(singleDTO)

	return a.adapter.CreateSingle(ctx, singleDAO)
}

func (a *ProfileQuerierAdapter) CreateDouble(ctx context.Context, doubleDTO *profile_dto.CreateDoubleDTOReq) (string, error) {
	doubleDAO := mappers.CreateDoubleDTOtoDAO(doubleDTO)

	return a.adapter.CreateDouble(ctx, doubleDAO)
}

func (a *ProfileQuerierAdapter) CreateTeam(ctx context.Context, teamDTO *profile_dto.CreateTeamDTOReq) (string, error) {
	teamDAO, err := mappers.CreateTeamDTOtoDAO(teamDTO, a.ConvertToObjectID)
	if err != nil {
		return "", nil
	}

	return a.adapter.CreateTeam(ctx, teamDAO)
}

func (a *ProfileQuerierAdapter) CreateFollower(ctx context.Context, followerDTO *profile_dto.CreateFollowerDTOReq) error {
	teamDAO, err := mappers.CreateFollowerDTOtoDAO(followerDTO, a.ConvertToObjectID)
	if err != nil {
		return nil
	}

	return a.adapter.CreateFollower(ctx, teamDAO)
}

func (a *ProfileQuerierAdapter) GetProfileInfoInCategory(ctx context.Context, categoryID, competitorID string) (*profile_dto.GetProfileInfoInCategoryDTORes, error) {
	profileDAO, err := a.adapter.GetProfileInfoInCategory(ctx, categoryID, competitorID)
	if err != nil {
		return nil, err
	}

	profileDTO := mappers.GetProfileInfoInCategoryDAOtoDTO(profileDAO)

	return profileDTO, nil
}

func (a *ProfileQuerierAdapter) GetAvailabilityDailySlice(ctx context.Context, userID, competitorID string) ([]*profile_dto.GetDailyAvailabilityByIDDTORes, error) {
	dailyAvailabilityDAO, err := a.adapter.GetAvailabilityDailySlice(ctx, userID, competitorID)
	if err != nil {
		return nil, err
	}

	dailyAvailabilitiesDTO := mappers.GetAvailabilityDailySliceDTOtoDAO(dailyAvailabilityDAO)

	return dailyAvailabilitiesDTO, nil
}

func (a *ProfileQuerierAdapter) CreateAvailabilityForCompetitor(ctx context.Context, competitorID string, dailyAvailabilityDTO []*profile_dto.GetDailyAvailabilityByIDDTORes) error {
	dailyAvailabilityDAO := mappers.CreateAvailabilityDailySliceDTOtoDAO(dailyAvailabilityDTO)

	return a.adapter.CreateAvailabilityForCompetitor(ctx, competitorID, dailyAvailabilityDAO)
}

func (a *ProfileQuerierAdapter) GetDailyAvailabilityCompetitorID(ctx context.Context, competitorID string, day string) (*profile_dto.GetDailyAvailabilityByIDDTORes, error) {
	availabilityDAO, err := a.adapter.GetDailyAvailabilityCompetitorID(ctx, competitorID, day)
	if err != nil {
		return nil, err
	}

	availabilityDTO := mappers.GetDailyAvailabilityCompetitorIDDAOtoDTO(availabilityDAO)

	return availabilityDTO, nil
}

func (a *ProfileQuerierAdapter) GetCompetitorTournamentsInCategory(ctx context.Context, categoryID, competitorID, lastID string, limit int) ([]*profile_dto.GetTournamentsFromCategoryDTORes, error) {
	tournamentsDAO, err := a.adapter.GetCompetitorTournamentsInCategory(ctx, categoryID, competitorID, lastID, limit)
	if err != nil {
		return nil, err
	}

	tournamentsDTO := mappers.GetCompetitorTournamentsInCategoryDAOtoDTO(tournamentsDAO)

	return tournamentsDTO, nil
}

func (a *ProfileQuerierAdapter) VerifyFollowerExistsRelation(ctx context.Context, followerDTO *profile_dto.CreateFollowerDTOReq) error {
	followerDAO, err := mappers.CreateFollowerDTOtoDAO(followerDTO, a.ConvertToObjectID)
	if err != nil {
		return nil
	}

	return a.adapter.VerifyFollowerExistsRelation(ctx, followerDAO)
}
