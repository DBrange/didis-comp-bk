package adapters

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/category/services"
)

type CategoryProxyAdapter struct {
	categoryService *services.CategoryService
}

func NewCategoryProxyAdapter(categoryService *services.CategoryService) *CategoryProxyAdapter {
	return &CategoryProxyAdapter{
		categoryService: categoryService,
	}
}

func (a *CategoryProxyAdapter) OrganizeCategory(ctx context.Context, organizerID string, categoryInfoDTO *dto.CreateCategoryDTOReq) error {
	return a.categoryService.OrganizeCategory(ctx, organizerID, categoryInfoDTO)
}

func (a *CategoryProxyAdapter) AddTournamentInCategory(ctx context.Context, categoryID string, tournamentID string) error {
	return a.categoryService.AddTournamentInCategory(ctx, categoryID, tournamentID)
}

func (a *CategoryProxyAdapter) AddCompetitorInCategory(ctx context.Context, categoryID, competitorID string) error {
	return a.categoryService.AddCompetitorInCategory(ctx, categoryID, competitorID)
}

func (a *CategoryProxyAdapter) SearchCompetitorInCategory(ctx context.Context, categoryID string, name string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) ([]*dto.GetCompetitorsOfCategoryCompetitorDTORes, error) {
	return a.categoryService.SearchCompetitorInCategory(ctx, categoryID, name, sport, competitorType)
}

func (a *CategoryProxyAdapter) SearchCompetitorForCategory(ctx context.Context, userID string, name string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) ([]*dto.GetCompetitorFollowedDTORes, error) {
	return a.categoryService.SearchCompetitorForCategory(ctx, userID, name, sport, competitorType)
}

func (a *CategoryProxyAdapter) ModifyCategoryInfo(ctx context.Context, categoryID string, categoryInfoDTO *dto.UpdateCategoryDTOReq) error {
	return a.categoryService.ModifyCategoryInfo(ctx, categoryID, categoryInfoDTO)
}

func (a *CategoryProxyAdapter) GetCategoryInfo(ctx context.Context, categoryID string) (*dto.GetCategoryInfoByIDDTORes, error) {
	return a.categoryService.GetCategoryInfo(ctx, categoryID)
}

func (a *CategoryProxyAdapter) GetParticipantsOfCategory(ctx context.Context, categoryID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE, limit int, lastID string) (*dto.GetCompetitorsOfCategoryDTORes, error) {
	return a.categoryService.GetParticipantsOfCategory(ctx, categoryID, sport, competitorType, limit, lastID)
}

func (a *CategoryProxyAdapter) RemoveCompetitorFromCategory(ctx context.Context, categoryID, competitorID string) error {
	return a.categoryService.RemoveCompetitorFromCategory(ctx, categoryID, competitorID)
}

func (a *CategoryProxyAdapter) ListCategories(ctx context.Context, organizerID string, sport models.SPORT, competitorType *models.COMPETITOR_TYPE) ([]dto.GetCategoriesFromOrganizerDTORes, error) {
	return a.categoryService.ListCategories(ctx, organizerID, sport, competitorType)
}

func (a *CategoryProxyAdapter) GetTournamentsFromCategory(ctx context.Context, categoryID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE, limit int, lastID string) (*dto.GetTournamentsFromCategoryDTORes, error) {
	return a.categoryService.GetTournamentsFromCategory(ctx, categoryID, sport, competitorType, limit, lastID)
}

func (a *CategoryProxyAdapter) ModifyCompetitorPoints(ctx context.Context, categoryID, competitorID string, points int) error {
	return a.categoryService.ModifyCompetitorPoints(ctx, categoryID, competitorID, points)
}

func (a *CategoryProxyAdapter) AddGuestUserInCategory(ctx context.Context, categoryID string, guestUsersDTO []*dto.CreateGuestUserDTOReq, sport models.SPORT, competitorType models.COMPETITOR_TYPE) error {
	return a.categoryService.AddGuestUserInCategory(ctx, categoryID, guestUsersDTO, sport, competitorType)
}

func (a *CategoryProxyAdapter) GetTournamentsByNameFromCategory(ctx context.Context, categoryID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE, tournamentName string) ([]*dto.GetTournamentsFromCategoryTournamentDTORes, error) {
	return a.categoryService.GetTournamentsByNameFromCategory(ctx, categoryID, sport, competitorType, tournamentName)
}
