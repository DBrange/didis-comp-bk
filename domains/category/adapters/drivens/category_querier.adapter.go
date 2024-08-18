package adapters

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/category/adapters/mappers"
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	ports "github.com/DBrange/didis-comp-bk/domains/repository/ports/drivers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryQueryerAdapter struct {
	adapter ports.ForManagingCategory
}

func NewCategoryQueryerAdapter(adapter ports.ForManagingCategory) *CategoryQueryerAdapter {
	return &CategoryQueryerAdapter{
		adapter: adapter,
	}
}

func (a *CategoryQueryerAdapter) VerifyOrganizerExists(ctx context.Context, organizerID string) error {
	return a.adapter.VerifyOrganizerExists(ctx, organizerID)
}

func (a *CategoryQueryerAdapter) ConvertToObjectID(ID string) (*primitive.ObjectID, error) {
	return a.adapter.ConvertToObjectID(ID)
}

func (a *CategoryQueryerAdapter) CreateCategory(ctx context.Context, organizerID string, categoryDTO *dto.CreateCategoryDTOReq) (string, error) {
	categoryDAO := mappers.CreateCategoryDTOtoDAO(categoryDTO)

	return a.adapter.CreateCategory(ctx, organizerID, categoryDAO)
}

func (a *CategoryQueryerAdapter) AddTournamentInCategory(ctx context.Context, categoryID string, tournamentID string) error {
	return a.adapter.AddTournamentInCategory(ctx, categoryID, tournamentID)
}

func (a *CategoryQueryerAdapter) AddCategoryInTournament(ctx context.Context, tournamentID string, categoryID string) error {
	return a.adapter.AddTournamentInCategory(ctx, tournamentID, categoryID)
}

func (a *CategoryQueryerAdapter) CreateCategoryRegistration(ctx context.Context, categoryRegistrationDTO *dto.CreateCategoryRegistrationDTOReq) error {
	categoryRegistrationDAO, err := mappers.CreateCategoryRegistrationDTOtoDAO(categoryRegistrationDTO, a.ConvertToObjectID)
	if err != nil {
		return nil
	}

	return a.adapter.CreateCategoryRegistration(ctx, categoryRegistrationDAO)
}

func (a *CategoryQueryerAdapter) VerifyCategoryExists(ctx context.Context, categoryID string) error {
	return a.adapter.VerifyCategoryExists(ctx, categoryID)
}

func (a *CategoryQueryerAdapter) VerifyCompetitorExists(ctx context.Context, competitorID string) error {
	return a.adapter.VerifyCompetitorExists(ctx, competitorID)
}

func (a *CategoryQueryerAdapter) VerifyCategoryExistsRelation(ctx context.Context, categoryRegistrationDTO *dto.CreateCategoryRegistrationDTOReq) error {
	categoryRegistrationDAO, err := mappers.CreateCategoryRegistrationDTOtoDAO(categoryRegistrationDTO, a.ConvertToObjectID)
	if err != nil {
		return nil
	}

	return a.adapter.VerifyCategoryExistsRelation(ctx, categoryRegistrationDAO)
}

func (a *CategoryQueryerAdapter) GetCompetitorsOfCategoryByName(ctx context.Context, categoryID string, name string, sport models.SPORT, competitorType models.COMPETITOR_TYPE, limit int, lastID string) ([]*dto.GetCompetitorsOfCategoryDTORes, error) {
	categoryRegistrationDAO, err := a.adapter.GetCompetitorsOfCategoryByName(ctx, categoryID, name, sport, competitorType, limit, lastID)
	if err != nil {
		return nil, err
	}

	categoryRegistrationDTO := mappers.GetUsersOfCategoryByNameDAOtoDTO(categoryRegistrationDAO)

	return categoryRegistrationDTO, nil
}

func (a *CategoryQueryerAdapter) GetCompetitorsFollowed(
	ctx context.Context,
	userID string,
	name string,
	sport models.SPORT,
	competitorType models.COMPETITOR_TYPE,
) ([]*dto.GetCompetitorFollowedDTORes, error) {
	competitorsDAO, err := a.adapter.GetCompetitorsFollowed(ctx, userID, name, sport, competitorType)
	if err != nil {
		return nil, err
	}

	categoryRegistrationDTO := mappers.GetCompetitorsFollowedDAOtoDTO(competitorsDAO)

	return categoryRegistrationDTO, nil
}

func (a *CategoryQueryerAdapter) UpdateCategory(ctx context.Context, categoryID string, categoryInfoDTO *dto.UpdateCategoryDTOReq) error {
	categoryInfoDAO, categoryOID, err := mappers.UpdateCategoryDTOtoDAO(categoryInfoDTO, categoryID, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.adapter.UpdateCategory(ctx, categoryOID, categoryInfoDAO)
}

func (a *CategoryQueryerAdapter) GetCategoryInfoByID(ctx context.Context, categoryID string) (*dto.GetCategoryInfoByIDDTORes, error) {
	competitorsDAO, err := a.adapter.GetCategoryInfoByID(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	categoryRegistrationDTO := mappers.GetCategoryInfoByIDDAOtoDTO(competitorsDAO)

	return categoryRegistrationDTO, nil
}

func (a *CategoryQueryerAdapter) IncrementTotalParticipants(ctx context.Context, categoryID string) error {
	categoryOID, err := a.ConvertToObjectID(categoryID)
	if err != nil {
		return err
	}

	return a.adapter.IncrementTotalParticipants(ctx, categoryOID)
}

func (a *CategoryQueryerAdapter) DecrementTotalParticipants(ctx context.Context, categoryID string) error {
	categoryOID, err := a.ConvertToObjectID(categoryID)
	if err != nil {
		return err
	}

	return a.adapter.DecrementTotalParticipants(ctx, categoryOID)
}

func (a *CategoryQueryerAdapter) GetParticipantsOfCategory(ctx context.Context, categoryID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE, limit int, lastID string) ([]*dto.GetCompetitorsOfCategoryDTORes, error) {
	categoryRegistrationDAO, err := a.adapter.GetParticipantsOfCategory(ctx, categoryID, sport, competitorType, limit, lastID)
	if err != nil {
		return nil, err
	}

	categoryRegistrationDTO := mappers.GetUsersOfCategoryByNameDAOtoDTO(categoryRegistrationDAO)

	return categoryRegistrationDTO, nil
}

func (a *CategoryQueryerAdapter) CategoryRegistrationColl() *mongo.Collection {
	return a.adapter.CategoryRegistrationColl()
}

func (a *CategoryQueryerAdapter) PermaDeleteCategoryRegistration(ctx context.Context, mc *mongo.Collection, ID string) error {
	return a.adapter.PermaDeleteCategoryRegistration(ctx, mc, ID)
}

func (a *CategoryQueryerAdapter) AddCategoryInOrganizer(ctx context.Context, organizerID, categoryID string) error {
	organizerOID, err := a.ConvertToObjectID(organizerID)
	if err != nil {
		return err
	}

	categoryOID, err := a.ConvertToObjectID(categoryID)
	if err != nil {
		return err
	}

	return a.adapter.AddCategoryInOrganizer(ctx, organizerOID, categoryOID)
}

func (a *CategoryQueryerAdapter) GetCategoriesFromOrganizer(ctx context.Context, organizerID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) ([]dto.GetCategoriesFromOrganizerDTORes, error) {
	categoriesDAO, err := a.adapter.GetCategoriesFromOrganizer(ctx, organizerID, sport, competitorType)
	if err != nil {
		return nil, err
	}

	categoryRegistrationDTO := mappers.GetCategoriesFromOrganizerDAOtoDTO(categoriesDAO)

	return categoryRegistrationDTO, nil
}

func (a *CategoryQueryerAdapter) GetTournamentsFromCategory(ctx context.Context, categoryID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE, limit int, lastID string) ([]dto.GetTournamentsFromCategoryDTORes, error) {
	tournamentsDAO, err := a.adapter.GetTournamentsFromCategory(ctx, categoryID, sport, competitorType, limit, lastID)
	if err != nil {
		return nil, err
	}

	tournamentsDTO := mappers.GetTournamentsFromCategoryDAOtoDTO(tournamentsDAO)

	return tournamentsDTO, nil
}

func (a *CategoryQueryerAdapter) UpdateCompetitorPoints(ctx context.Context, categoryID, competitorID string, points int) error {
	categoryOID, err := a.ConvertToObjectID(categoryID)
	if err != nil {
		return err
	}

	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	return a.adapter.UpdateCompetitorPoints(ctx, categoryOID, competitorOID, points)
}

func (a *CategoryQueryerAdapter) CreateGuestUser(ctx context.Context, guestUserDTO *dto.CreateGuestUserDTOReq) (string, error) {
	guestUserDAO := mappers.CreateGuestUserDTOtoDAO(guestUserDTO)

	return a.adapter.CreateGuestUser(ctx, guestUserDAO)
}

func (a *CategoryQueryerAdapter) CreateCompetitor(ctx context.Context, sport models.SPORT, competitorType models.COMPETITOR_TYPE, ID string) (string, error) {
	OID, err := a.ConvertToObjectID(ID)
	if err != nil {
		return "", err
	}
	return a.adapter.CreateCompetitor(ctx, sport, competitorType, OID)
}

func (a *CategoryQueryerAdapter) CreateGuestCompetitor(ctx context.Context, guestCompetitorDTO *dto.CreateGuestCompetitorDTOReq) (string, error) {
	guestCompetitorDAO, err := mappers.CreateGuestCompetitorDTOtoDAO(guestCompetitorDTO, a.ConvertToObjectID)
	if err != nil {
		return "", err
	}

	return a.adapter.CreateGuestCompetitor(ctx, guestCompetitorDAO)
}

func (a *CategoryQueryerAdapter) CreateSingle(ctx context.Context, singleDTO *dto.CreateSingleDTOReq) (string, error) {
	singleDAO := mappers.CreateSingleDTOtoDAO(singleDTO)

	return a.adapter.CreateSingle(ctx, singleDAO)
}

func (a *CategoryQueryerAdapter) CreateDouble(ctx context.Context, doubleDTO *dto.CreateDoubleDTOReq) (string, error) {
	doubleDAO := mappers.CreateDoubleDTOtoDAO(doubleDTO)

	return a.adapter.CreateDouble(ctx, doubleDAO)
}

func (a *CategoryQueryerAdapter) CreateTeam(ctx context.Context, teamDTO *dto.CreateTeamDTOReq) (string, error) {
	teamDAO, err := mappers.CreateTeamDTOtoDAO(teamDTO, a.ConvertToObjectID)
	if err != nil {
		return "", nil
	}

	return a.adapter.CreateTeam(ctx, teamDAO)
}

func (a *CategoryQueryerAdapter) CreateCompetitorStats(ctx context.Context, competitorID string) error {
	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	return a.adapter.CreateCompetitorStats(ctx, competitorOID)
}

func (a *CategoryQueryerAdapter) GetCategoryRegistrationSortedByPoints(ctx context.Context, categoryID string) ([]*dto.GetCategoryRegistrationSortedByPointsDTORes, error) {
	categoryRegistrationSortedDAO, err := a.adapter.GetCategoryRegistrationSortedByPoints(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	categoryRegistrationSortedDTO := mappers.GetCategoryRegistrationSortedByPointsDAOtoDTO(categoryRegistrationSortedDAO)

	return categoryRegistrationSortedDTO, nil
}

func (a *CategoryQueryerAdapter) UpdateCategoryRegistrationCurrentPosition(ctx context.Context, categoryID string, categoryRegistrationDTO []*dto.GetCategoryRegistrationSortedByPointsDTORes) error {
	categoryRegistrationDAO, categoryOID, err := mappers.UpdateCategoryRegistrationCurrentPositionDTOtoDAO(categoryRegistrationDTO, categoryID, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.adapter.UpdateCategoryRegistrationCurrentPosition(ctx, categoryOID, categoryRegistrationDAO)
}