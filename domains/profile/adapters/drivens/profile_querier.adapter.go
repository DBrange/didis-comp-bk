package adapters

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/profile/adapters/mappers"
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	profile_dto "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
	ports "github.com/DBrange/didis-comp-bk/domains/repository/ports/drivers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileQueryerAdapter struct {
	adapter ports.ForManagingProfile
}

func NewProfileQueryerAdapter(adapter ports.ForManagingProfile) *ProfileQueryerAdapter {
	return &ProfileQueryerAdapter{
		adapter: adapter,
	}
}

func (a *ProfileQueryerAdapter) WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error {
	return a.adapter.WithTransaction(ctx, fn)
}

func (a *ProfileQueryerAdapter) CreateUser(ctx context.Context, userDTO *profile_dto.CreateUserDTOReq) (string, error) {
	userDAO := mappers.CreateUserDTOtoDAO(userDTO)
	locationIOD, err := a.ConvertToObjectID(*userDTO.LocationID)
	if err != nil {
		return "", nil
	}

	userDAO.LocationID = locationIOD

	return a.adapter.CreateUser(ctx, userDAO)
}

func (a *ProfileQueryerAdapter) CreateLocation(ctx context.Context, locationDTO *profile_dto.CreateLocationDTOReq) (string, error) {
	locationDAO := mappers.CreateLocationDTOtoDAO(locationDTO)

	return a.adapter.CreateLocation(ctx, locationDAO)
}

func (a *ProfileQueryerAdapter) GetRoleByNameAndType(ctx context.Context, roleName, roleType string) (*dto.GetRoleDTOByID, error) {
	roleDAO, err := a.adapter.GetRoleByNameAndType(ctx, roleName, roleType)
	if err != nil {
		return nil, err
	}

	roleDTO := mappers.CreateRoleDAOtoDTO(roleDAO)

	return roleDTO, nil
}

func (a *ProfileQueryerAdapter) CreateOrganizer(ctx context.Context, userID string) error {
	return a.adapter.CreateOrganizer(ctx, userID)
}

func (a *ProfileQueryerAdapter) CreateAvailability(ctx context.Context, userID, competitorID *string) error {
	return a.adapter.CreateAvailability(ctx, userID, competitorID)
}

func (a *ProfileQueryerAdapter) UpdateAvailability(ctx context.Context, availabilityID string, availabilityDTO *profile_dto.UpdateDailyAvailabilityDTOReq) error {
	availabilityDAO := mappers.UpdateDailyAvailabilityDTOtoDAO(availabilityDTO)

	return a.adapter.UpdateAvailability(ctx, availabilityID, availabilityDAO)
}

func (a *ProfileQueryerAdapter) UpdateUser(ctx context.Context, userID string, userDTO *profile_dto.UpdateUserDTOReq) error {
	userDAO := mappers.UpdateUserDTOtoDAO(userDTO)

	return a.adapter.UpdateUser(ctx, userID, userDAO)
}

func (a *ProfileQueryerAdapter) UpdateLocation(ctx context.Context, locationID string, locationDTO *profile_dto.UpdateLocationDTOReq) error {
	locationDAO := mappers.UpdateLocationDTOtoDAO(locationDTO)

	return a.adapter.UpdateLocation(ctx, locationID, locationDAO)
}

func (a *ProfileQueryerAdapter) GetUserByID(ctx context.Context, userID string) (*profile_dto.GetUserByIDDTORes, error) {
	userDAO, err := a.adapter.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	userDTO := mappers.GetUserByIDDAOtoDTO(userDAO)

	return userDTO, err
}

func (a *ProfileQueryerAdapter) GetLocationByID(ctx context.Context, locationID string) (*profile_dto.GetLocationByIDDTORes, error) {
	locationDAO, err := a.adapter.GetLocationByID(ctx, locationID)
	if err != nil {
		return nil, err
	}

	locationDTO := mappers.GetLocationByIDDAOtoDTO(locationDAO)

	return locationDTO, nil
}

func (a *ProfileQueryerAdapter) GetDailyAvailabilityByID(ctx context.Context, availabilityID string, day string) (*profile_dto.GetDailyAvailabilityByIDDTORes, error) {
	availabilityDAO, err := a.adapter.GetDailyAvailabilityByID(ctx, availabilityID, day)
	if err != nil {
		return nil, err
	}

	availabilityDTO := mappers.GetDailyAvailabilityByIDDAOtoDTO(availabilityDAO)

	return availabilityDTO, nil
}

func (a *ProfileQueryerAdapter) ConvertToObjectID(ID string) (*primitive.ObjectID, error) {
	return a.adapter.ConvertToObjectID(ID)
}

func (a *ProfileQueryerAdapter) CreateCompetitorType(ctx context.Context, competitorType models.COMPETITOR_TYPE) (*primitive.ObjectID, error) {
	return a.adapter.CreateCompetitorType(ctx, competitorType)
}

func (a *ProfileQueryerAdapter) CreateCompetitor(ctx context.Context, sport models.SPORT, competitorType models.COMPETITOR_TYPE, OID *primitive.ObjectID) (string, error) {
	return a.adapter.CreateCompetitor(ctx, sport, competitorType, OID)
}

func (a *ProfileQueryerAdapter) CreateCompetitorStats(ctx context.Context, competitorOID *primitive.ObjectID) error {
	return a.adapter.CreateCompetitorStats(ctx, competitorOID)
}

func (a *ProfileQueryerAdapter) CreateCompetitorUser(ctx context.Context, userOID *primitive.ObjectID, competitorOID *primitive.ObjectID) error {
	return a.adapter.CreateCompetitorUser(ctx, userOID, competitorOID)
}

func (a *ProfileQueryerAdapter) DeleteUser(ctx context.Context, userID string) (*user_dao.UserRelationsToDeleteDAOReq, error) {
	return a.adapter.DeleteUser(ctx, userID)
}

func (a *ProfileQueryerAdapter) GetDailyAvailabilityByUserID(ctx context.Context, userID, day string) (*profile_dto.GetDailyAvailabilityByIDDTORes, error) {
	availabilityDAO, err := a.adapter.GetDailyAvailabilityByUserID(ctx, userID, day)
	if err != nil {
		return nil, err
	}

	availabilityDTO := mappers.GetDailyAvailabilityByIDDAOtoDTO(availabilityDAO)

	return availabilityDTO, nil
}

func (a *ProfileQueryerAdapter) SetDeletedAt(ctx context.Context, mc *mongo.Collection, ID string, name string) error {
	return a.adapter.SetDeletedAt(ctx, mc, ID, name)
}

func (a *ProfileQueryerAdapter) AvailabilityColl() *mongo.Collection {
	return a.adapter.AvailabilityColl()
}

func (a *ProfileQueryerAdapter) LocationColl() *mongo.Collection {
	return a.adapter.LocationColl()
}

func (a *ProfileQueryerAdapter) UpdateUserPassword(ctx context.Context, userID, newPassword string) error  {
	return a.adapter.UpdateUserPassword(ctx , userID, newPassword )  
}

func (a *ProfileQueryerAdapter) GetUserPasswordByID(ctx context.Context, userID string) (string, error)  {
	return a.adapter.GetUserPasswordByID(ctx, userID ) 
}

func (a *ProfileQueryerAdapter) GetUserPasswordForLogin(ctx context.Context, username string) (string, string, error) {
	return a.adapter.GetUserPasswordForLogin(ctx, username)
}

func (a *ProfileQueryerAdapter) GetUserRoles(ctx context.Context, userID string) ([]string, error) {
	return a.adapter.GetUserRoles(ctx, userID)
}

func (a *ProfileQueryerAdapter) ActivateUserNotification(ctx context.Context) {
	a.adapter.ActivateUserNotification(ctx)
}

func (a *ProfileQueryerAdapter) GetAvailabilityIDByUserID(ctx context.Context, userID string) (string, error) {
	return a.adapter.GetAvailabilityIDByUserID(ctx, userID)
}


