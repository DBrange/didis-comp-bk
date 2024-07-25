package adapters

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/profile/adapters/mappers"
	profile_dto "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	models_role "github.com/DBrange/didis-comp-bk/domains/repository/models/role"
	ports "github.com/DBrange/didis-comp-bk/domains/repository/ports/drivers"
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

// func (a *ProfileQueryerAdapter) RegisterUser(ctx context.Context, profileInfoDTO *profile_dto.RegisterUserDTOReq) error {
// 	userInfoDAO, locationInfoDAO := mappers.RegisterUserMapper(profileInfoDTO)

// 	return a.adapter.CreateUserAndLocation(ctx, userInfoDAO, locationInfoDAO, profileInfoDTO.Organizer)
// }

func (a *ProfileQueryerAdapter) WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error {
	return a.adapter.WithTransaction(ctx, fn)
}

func (a *ProfileQueryerAdapter) CreateUser(ctx context.Context, userDTO *profile_dto.CreateUserDTOReq) (string, error) {
	userDAO := mappers.CreateUserDTOtoDAO(userDTO)

	return a.adapter.CreateUser(ctx, userDAO)
}

func (a *ProfileQueryerAdapter) CreateLocation(ctx context.Context, locationDTO *profile_dto.CreateLocationDTOReq) (string, error) {
	locationDAO := mappers.CreateLocationDTOtoDAO(locationDTO)

	return a.adapter.CreateLocation(ctx, locationDAO)
}

func (a *ProfileQueryerAdapter) GetRoleByNameAndType(ctx context.Context, roleName, roleType string) (*models_role.Role, error) {
	return a.adapter.GetRoleByNameAndType(ctx, roleName, roleType)
}

func (a *ProfileQueryerAdapter) CreateOrganizer(ctx context.Context, userID string) error {
	return a.adapter.CreateOrganizer(ctx, userID)
}

func (a *ProfileQueryerAdapter) CreateAvailability(ctx context.Context, userID, competitorID *string) error {
	return a.adapter.CreateAvailability(ctx, userID, competitorID)
}

func (a *ProfileQueryerAdapter) ModifyProfileAvailability(ctx context.Context, availabilityID string, availabilityInfoDTO *profile_dto.ModifyProfileDailyAvailabilityDTOReq) error {
	availabilityInfoDAO := mappers.ModifyProfileDailyAvailabilityMapper(availabilityInfoDTO)

	return a.adapter.UpdateProfileAvailability(ctx, availabilityID, availabilityInfoDAO)
}

func (a *ProfileQueryerAdapter) ModifyPersonalInfo(ctx context.Context, userID string, userInfoDTO *profile_dto.ModifyPersonalInfoDTOReq) error {
	userInfoDAO, locationInfoDAO := mappers.ModifyPersonalInfoMapper(userInfoDTO)

	return a.adapter.UpdatePersonalInfo(ctx, userID, userInfoDAO, locationInfoDAO)
}

func (a *ProfileQueryerAdapter) GetPersonalInfoByID(ctx context.Context, userID string) (*profile_dto.GetPersonalInfoByIDDTORes, error) {
	userInfo, locationInfo, err := a.adapter.GetPersonalInfoByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	personalInfo := mappers.GetPersonalInfoByIDMapper(userInfo, locationInfo)

	return personalInfo, nil
}

func (a *ProfileQueryerAdapter) GetProfileAvailabilityInfoByID(ctx context.Context, userID string, day string) (*profile_dto.GetProfileDailyAvailabilityInfoByIDDTORes, error) {
	availabilityInfo, err := a.adapter.GetProfileAvailabilityInfoByID(ctx, userID, day)
	if err != nil {
		return nil, err
	}

	availabilityInfoDTO := mappers.GetProfileAvailabilityInfoByIDMapper(availabilityInfo)

	return availabilityInfoDTO, nil
}

func (a *ProfileQueryerAdapter) CloseProfile(ctx context.Context, userID string) error {
	return a.adapter.DeleteProfile(ctx, userID)
}

func (a *ProfileQueryerAdapter) ModifyPassword(ctx context.Context, userID, newPassword, oldPassword string) error {
	return a.adapter.UpdateProfilePassword(ctx, userID, newPassword, oldPassword)
}

func (a *ProfileQueryerAdapter) RegisterCompetitor(ctx context.Context, userID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) error {
	return a.adapter.RegisterCompetitor(ctx, userID, sport, competitorType)
}
