package adapters

import (
	"context"
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
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

func (a *ProfileProxyAdapter) ModifyProfileAvailability(ctx context.Context, availabilityID string, availabilityDTO *models.UpdateDailyAvailabilityDTOReq) error {
	return a.profileService.ModifyProfileAvailability(ctx, availabilityID, availabilityDTO)
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
func (a *ProfileProxyAdapter) GetProfileDailyAvailabilityByID(ctx context.Context, userID string, day string) (*models.GetDailyAvailabilityByIDDTORes, string, error) {
	return a.profileService.GetProfileDailyAvailabilityByID(ctx, userID, day)
}

func (a *ProfileProxyAdapter) CloseProfile(ctx context.Context, userID string) error {
	return a.profileService.CloseProfile(ctx, userID)
}

func (a *ProfileProxyAdapter) ModifyPassword(ctx context.Context, userID, newPassword, oldPassword string) error {
	return a.profileService.ModifyPassword(ctx, userID, newPassword, oldPassword)
}

func (a *ProfileProxyAdapter) RegisterCompetitor(ctx context.Context, userIDs []string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) error {
	return a.profileService.RegisterCompetitor(ctx, userIDs, sport, competitorType)
}

func (a *ProfileProxyAdapter) Login(ctx context.Context, loginDTO *profile_dto.LoginDTOReq) (*profile_dto.GetUserForLoginDTO, string, string, error) {
	return a.profileService.Login(ctx, loginDTO)
}

func (a *ProfileProxyAdapter) RefreshToken(ctx context.Context, refreshToken string) (string, string, error) {
	return a.profileService.RefreshToken(ctx, refreshToken)
}

func (a *ProfileProxyAdapter) FollowProfile(ctx context.Context, fromUserID, toUserID string) error {
	return a.profileService.FollowProfile(ctx, fromUserID, toUserID)
}

func (a *ProfileProxyAdapter) GetProfileInfoInCategory(ctx context.Context, categoryID, competitorID string) (*profile_dto.GetProfileInfoInCategoryDTORes, error) {
	return a.profileService.GetProfileInfoInCategory(ctx, categoryID, competitorID)
}

func (a *ProfileProxyAdapter) GetProfileAvailabilityInCategory(ctx context.Context, competitorID, day string) (*models.GetDailyAvailabilityByIDDTORes, string, error) {
	return a.profileService.GetProfileAvailabilityInCategory(ctx, competitorID, day)
}

func (a *ProfileProxyAdapter) GetProfileTournamentsInCategory(ctx context.Context, categoryID, competitorID, lastID string, limit int) ([]*profile_dto.GetTournamentsFromCategoryDTORes, error) {
	return a.profileService.GetProfileTournamentsInCategory(ctx, categoryID, competitorID, lastID, limit)
}

func (a *ProfileProxyAdapter) GetProfileCategories(ctx context.Context, userID string, sport models.SPORT, limit int, lastID string) ([]*profile_dto.GetUserCategoriesCategoryDTO, error) {
	return a.profileService.GetProfileCategories(ctx, userID, sport, limit, lastID)
}

func (a *ProfileProxyAdapter) GetNumberFollowers(ctx context.Context, userID string) (int, error) {
	return a.profileService.GetNumberFollowers(ctx, userID)
}

func (a *ProfileProxyAdapter) GetUserFollowers(ctx context.Context, userID string, name string, limit int, lastCreatedAt *time.Time) (*profile_dto.GetUserFollowersDTORes, error) {
	return a.profileService.GetUserFollowers(ctx, userID, name, limit, lastCreatedAt)
}

func (a *ProfileProxyAdapter) GetUserPrimaryInfo(ctx context.Context, fromID, userToID string) (*profile_dto.GetUserPrimatyInfoDTORes, error) {
	return a.profileService.GetUserPrimaryInfo(ctx, fromID, userToID)
}

func (a *ProfileProxyAdapter) GetOrganizerData(ctx context.Context, userID string) (*profile_dto.GetOrganizerDataDTORes, error) {
	return a.profileService.GetOrganizerData(ctx, userID)
}
