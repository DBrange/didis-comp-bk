package adapters

import (
	"context"

	profile_dto "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/profile/services"
)

type ProfileProxyAdapter struct {
	profileService *services.ProfileService
}

func NewProfileProxyAdapter(profileService *services.ProfileService) *ProfileProxyAdapter {
	return &ProfileProxyAdapter{
		profileService: profileService,
	}
}

func (a *ProfileProxyAdapter) RegisterUser(ctx context.Context, profileInfoDTO *profile_dto.RegisterUserDTOReq) error {
	return a.profileService.RegisterUser(ctx, profileInfoDTO)
}

func (a *ProfileProxyAdapter) ModifyProfileAvailability(ctx context.Context, availabilityID string, availabilityInfoDTO *profile_dto.ModifyProfileDailyAvailabilityDTOReq) error {
	return a.profileService.ModifyProfileAvailability(ctx, availabilityID, availabilityInfoDTO)
}

func (a *ProfileProxyAdapter) ModifyPersonalInfo(ctx context.Context, userID string, userInfoDTO *profile_dto.ModifyPersonalInfoDTOReq) error {
	return a.profileService.ModifyPersonalInfo(ctx, userID, userInfoDTO)
}

func (a *ProfileProxyAdapter) GetProfileByID(ctx context.Context, userID string) (*profile_dto.GetPersonalInfoByIDDTORes, error) {
	return a.profileService.GetPersonalInfoByID(ctx, userID)
}

func (a *ProfileProxyAdapter) GetPersonalInfoByID(ctx context.Context, userID string) (*profile_dto.GetPersonalInfoByIDDTORes, error) {
	return a.profileService.GetPersonalInfoByID(ctx, userID)
}

func (a *ProfileProxyAdapter) GetProfileAvailabilityInfoByID(ctx context.Context, userID string, day string) (*profile_dto.GetProfileDailyAvailabilityInfoByIDDTORes, error) {
	return a.profileService.GetProfileAvailabilityInfoByID(ctx, userID, day)
}

func (a *ProfileProxyAdapter) CloseProfile(ctx context.Context, userID string) error {
	return a.profileService.CloseProfile(ctx, userID)
}

func (a *ProfileProxyAdapter) ModifyPassword(ctx context.Context, userID, newPassword,oldPassword string) error {
	return a.profileService.ModifyPassword(ctx, userID, newPassword, oldPassword)
}