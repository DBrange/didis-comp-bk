package ports

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ForQueryingCategory interface {
	VerifyOrganizerExists(ctx context.Context, organizerID string) error
	CreateCategory(ctx context.Context, organizerID string, categoryDtO *dto.CreateCategoryDTOReq) (string, error)
	ConvertToObjectID(ID string) (*primitive.ObjectID, error)
	AddTournamentInCategory(ctx context.Context, categoryID string, tournamentID string) error
	AddCategoryInTournament(ctx context.Context, tournamentID string, categoryID string) error
	CreateCategoryRegistration(ctx context.Context, categoryRegistrationInfoDAO *dto.CreateCategoryRegistrationDTOReq) error
	VerifyCategoryExists(ctx context.Context, categoryID string) error
	VerifyCompetitorExists(ctx context.Context, competitorID string) error
	VerifyCategoryExistsRelation(ctx context.Context, categoryRegistrationDTO *dto.CreateCategoryRegistrationDTOReq) error
	GetCompetitorsOfCategoryByName(ctx context.Context, categoryID string, name string, sport models.SPORT, competitorType models.COMPETITOR_TYPE, limit int, lastID string) ([]*dto.GetCompetitorsOfCategoryDTORes, error)
	GetCompetitorsFollowed(
		ctx context.Context,
		userID string,
		name string,
		sport models.SPORT,
		competitorType models.COMPETITOR_TYPE,
	) ([]*dto.GetCompetitorFollowedDTORes, error)
	UpdateCategory(ctx context.Context, categoryID string, categoryInfoDTO *dto.UpdateCategoryDTOReq) error
	GetCategoryInfoByID(ctx context.Context, categoryOID string) (*dto.GetCategoryInfoByIDDTORes, error)
	IncrementTotalParticipants(ctx context.Context, categoryID string) error
	DecrementTotalParticipants(ctx context.Context, categoryID string) error
	GetParticipantsOfCategory(ctx context.Context, categoryID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE, limit int, lastID string) ([]*dto.GetCompetitorsOfCategoryDTORes, error)
	CategoryRegistrationColl() *mongo.Collection
	PermaDeleteCategoryRegistration(ctx context.Context, mc *mongo.Collection, categoryRegistrationID string) error
	AddCategoryInOrganizer(ctx context.Context, organizerID, categoryID string) error
	GetCategoriesFromOrganizer(ctx context.Context, organizerID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) ([]dto.GetCategoriesFromOrganizerDTORes, error)
	GetTournamentsFromCategory(ctx context.Context, categoryID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE, limit int, lastID string) ([]dto.GetTournamentsFromCategoryDTORes, error)
	UpdateCompetitorPoints(ctx context.Context, categoryID, competitorID string, points int) error
	CreateGuestUser(ctx context.Context, guestUserInfoDTO *dto.CreateGuestUserDTOReq) (string, error)
	CreateCompetitor(ctx context.Context, sport models.SPORT, competitorType models.COMPETITOR_TYPE, ID string) (string, error)
	CreateGuestCompetitor(ctx context.Context, guestCompetitorInfoDTO *dto.CreateGuestCompetitorDTOReq) (string, error)
	CreateSingle(ctx context.Context, singleInfoDTO *dto.CreateSingleDTOReq) (string, error)
	CreateDouble(ctx context.Context, doubleInfoDTO *dto.CreateDoubleDTOReq) (string, error)
	CreateTeam(ctx context.Context, teamInfoDTO *dto.CreateTeamDTOReq) (string, error)
	CreateCompetitorStats(ctx context.Context, competitorID string) error
	GetCategoryRegistrationSortedByPoints(ctx context.Context, categoryID string) ([]*dto.GetCategoryRegistrationSortedByPointsDTORes, error)
	UpdateCategoryRegistrationCurrentPosition(ctx context.Context, categoryID string, categoryRegistrationDTO []*dto.GetCategoryRegistrationSortedByPointsDTORes) error
}
