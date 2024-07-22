//go:generate mockgen -destination=tests/mocks/for_querying_profile.mock.go -package=mocks github.com/DBrange/didis-comp-bk/domains/profile/ports/drivens ForQueryingProfile

package ports

import (
	"context"

	profile_dto "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
)

type ForQueryingProfile interface {
	RegisterUser(ctx context.Context, profileInfoDTO *profile_dto.RegisterUserDTOReq) error
	ModifyProfileAvailability(ctx context.Context, availabilityID string, availabilityInfoDTO *profile_dto.ModifyProfileDailyAvailabilityDTOReq) error
	ModifyPersonalInfo(ctx context.Context, userID string, userInfoDTO *profile_dto.ModifyPersonalInfoDTOReq) error
	GetPersonalInfoByID(ctx context.Context, userID string) (*profile_dto.GetPersonalInfoByIDDTORes, error)
	GetProfileAvailabilityInfoByID(ctx context.Context, userID string, day string) (*profile_dto.GetProfileDailyAvailabilityInfoByIDDTORes, error)
	CloseProfile(ctx context.Context, userID string) error
	ModifyPassword(ctx context.Context, userID, newPassword, oldPassword string) error
}
