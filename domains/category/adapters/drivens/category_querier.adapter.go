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

type CategoryQuerierAdapter struct {
	adapter ports.ForManagingCategory
}

func NewCategoryQuerierAdapter(adapter ports.ForManagingCategory) *CategoryQuerierAdapter {
	return &CategoryQuerierAdapter{
		adapter: adapter,
	}
}

func (a *CategoryQuerierAdapter) VerifyOrganizerExists(ctx context.Context, organizerID string) error {
	return a.adapter.VerifyOrganizerExists(ctx, organizerID)
}

func (a *CategoryQuerierAdapter) ConvertToObjectID(ID string) (*primitive.ObjectID, error) {
	return a.adapter.ConvertToObjectID(ID)
}

func (a *CategoryQuerierAdapter) CreateCategory(ctx context.Context, organizerID string, categoryDTO *dto.CreateCategoryDTOReq) (string, error) {
	categoryDAO := mappers.CreateCategoryDTOtoDAO(categoryDTO)

	return a.adapter.CreateCategory(ctx, organizerID, categoryDAO)
}

func (a *CategoryQuerierAdapter) AddTournamentInCategory(ctx context.Context, categoryID string, tournamentID string) error {
	return a.adapter.AddTournamentInCategory(ctx, categoryID, tournamentID)
}

func (a *CategoryQuerierAdapter) AddCategoryInTournament(ctx context.Context, tournamentID string, categoryID string) error {
	return a.adapter.AddTournamentInCategory(ctx, tournamentID, categoryID)
}

func (a *CategoryQuerierAdapter) CreateCategoryRegistration(ctx context.Context, categoryRegistrationDTO *dto.CreateCategoryRegistrationDTOReq) error {
	categoryRegistrationDAO, err := mappers.CreateCategoryRegistrationDTOtoDAO(categoryRegistrationDTO, a.ConvertToObjectID)
	if err != nil {
		return nil
	}

	return a.adapter.CreateCategoryRegistration(ctx, categoryRegistrationDAO)
}

func (a *CategoryQuerierAdapter) VerifyCategoryExists(ctx context.Context, categoryID string) error {
	return a.adapter.VerifyCategoryExists(ctx, categoryID)
}

func (a *CategoryQuerierAdapter) VerifyCompetitorExists(ctx context.Context, competitorID string) error {
	return a.adapter.VerifyCompetitorExists(ctx, competitorID)
}

func (a *CategoryQuerierAdapter) VerifyCategoryExistsRelation(ctx context.Context, categoryRegistrationDTO *dto.CreateCategoryRegistrationDTOReq) error {
	categoryRegistrationDAO, err := mappers.CreateCategoryRegistrationDTOtoDAO(categoryRegistrationDTO, a.ConvertToObjectID)
	if err != nil {
		return nil
	}

	return a.adapter.VerifyCategoryExistsRelation(ctx, categoryRegistrationDAO)
}

func (a *CategoryQuerierAdapter) GetCompetitorsOfCategoryByName(ctx context.Context, categoryID string, name string, sport models.SPORT, competitorType models.COMPETITOR_TYPE, limit int, lastID string) ([]*dto.GetCompetitorsOfCategoryDTORes, error) {
	categoryRegistrationDAO, err := a.adapter.GetCompetitorsOfCategoryByName(ctx, categoryID, name, sport, competitorType, limit, lastID)
	if err != nil {
		return nil, err
	}

	categoryRegistrationDTO := mappers.GetUsersOfCategoryByNameDAOtoDTO(categoryRegistrationDAO)

	return categoryRegistrationDTO, nil
}

func (a *CategoryQuerierAdapter) GetCompetitorsFollowed(
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

func (a *CategoryQuerierAdapter) UpdateCategory(ctx context.Context, categoryID string, categoryInfoDTO *dto.UpdateCategoryDTOReq) error {
	categoryInfoDAO, categoryOID, err := mappers.UpdateCategoryDTOtoDAO(categoryInfoDTO, categoryID, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.adapter.UpdateCategory(ctx, categoryOID, categoryInfoDAO)
}

func (a *CategoryQuerierAdapter) GetCategoryInfoByID(ctx context.Context, categoryID string) (*dto.GetCategoryInfoByIDDTORes, error) {
	competitorsDAO, err := a.adapter.GetCategoryInfoByID(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	categoryRegistrationDTO := mappers.GetCategoryInfoByIDDAOtoDTO(competitorsDAO)

	return categoryRegistrationDTO, nil
}

func (a *CategoryQuerierAdapter) IncrementTotalParticipants(ctx context.Context, categoryID string) error {
	categoryOID, err := a.ConvertToObjectID(categoryID)
	if err != nil {
		return err
	}

	return a.adapter.IncrementTotalParticipants(ctx, categoryOID)
}

func (a *CategoryQuerierAdapter) DecrementTotalParticipants(ctx context.Context, categoryID string) error {
	categoryOID, err := a.ConvertToObjectID(categoryID)
	if err != nil {
		return err
	}

	return a.adapter.DecrementTotalParticipants(ctx, categoryOID)
}

func (a *CategoryQuerierAdapter) GetParticipantsOfCategory(ctx context.Context, categoryID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE, limit int, lastID string) ([]*dto.GetCompetitorsOfCategoryDTORes, error) {
	categoryRegistrationDAO, err := a.adapter.GetParticipantsOfCategory(ctx, categoryID, sport, competitorType, limit, lastID)
	if err != nil {
		return nil, err
	}

	categoryRegistrationDTO := mappers.GetUsersOfCategoryByNameDAOtoDTO(categoryRegistrationDAO)

	return categoryRegistrationDTO, nil
}

func (a *CategoryQuerierAdapter) CategoryRegistrationColl() *mongo.Collection {
	return a.adapter.CategoryRegistrationColl()
}

func (a *CategoryQuerierAdapter) PermaDeleteCategoryRegistration(ctx context.Context, mc *mongo.Collection, ID string) error {
	return a.adapter.PermaDeleteCategoryRegistration(ctx, mc, ID)
}

func (a *CategoryQuerierAdapter) AddCategoryInOrganizer(ctx context.Context, organizerID, categoryID string) error {
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

func (a *CategoryQuerierAdapter) GetCategoriesFromOrganizer(ctx context.Context, organizerID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) ([]dto.GetCategoriesFromOrganizerDTORes, error) {
	categoriesDAO, err := a.adapter.GetCategoriesFromOrganizer(ctx, organizerID, sport, competitorType)
	if err != nil {
		return nil, err
	}

	categoryRegistrationDTO := mappers.GetCategoriesFromOrganizerDAOtoDTO(categoriesDAO)

	return categoryRegistrationDTO, nil
}

func (a *CategoryQuerierAdapter) GetTournamentsFromCategory(ctx context.Context, categoryID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE, limit int, lastID string) ([]dto.GetTournamentsFromCategoryDTORes, error) {
	tournamentsDAO, err := a.adapter.GetTournamentsFromCategory(ctx, categoryID, sport, competitorType, limit, lastID)
	if err != nil {
		return nil, err
	}

	tournamentsDTO := mappers.GetTournamentsFromCategoryDAOtoDTO(tournamentsDAO)

	return tournamentsDTO, nil
}

func (a *CategoryQuerierAdapter) UpdateCompetitorPoints(ctx context.Context, categoryID, competitorID string, points int) error {
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

func (a *CategoryQuerierAdapter) CreateGuestUser(ctx context.Context, guestUserDTO *dto.CreateGuestUserDTOReq) (string, error) {
	guestUserDAO := mappers.CreateGuestUserDTOtoDAO(guestUserDTO)

	return a.adapter.CreateGuestUser(ctx, guestUserDAO)
}

func (a *CategoryQuerierAdapter) CreateCompetitor(ctx context.Context, sport models.SPORT, competitorType models.COMPETITOR_TYPE, ID string) (string, error) {
	OID, err := a.ConvertToObjectID(ID)
	if err != nil {
		return "", err
	}
	return a.adapter.CreateCompetitor(ctx, sport, competitorType, OID)
}

func (a *CategoryQuerierAdapter) CreateGuestCompetitor(ctx context.Context, guestCompetitorDTO *dto.CreateGuestCompetitorDTOReq) (string, error) {
	guestCompetitorDAO, err := mappers.CreateGuestCompetitorDTOtoDAO(guestCompetitorDTO, a.ConvertToObjectID)
	if err != nil {
		return "", err
	}

	return a.adapter.CreateGuestCompetitor(ctx, guestCompetitorDAO)
}

func (a *CategoryQuerierAdapter) CreateSingle(ctx context.Context, singleDTO *dto.CreateSingleDTOReq) (string, error) {
	singleDAO := mappers.CreateSingleDTOtoDAO(singleDTO)

	return a.adapter.CreateSingle(ctx, singleDAO)
}

func (a *CategoryQuerierAdapter) CreateDouble(ctx context.Context, doubleDTO *dto.CreateDoubleDTOReq) (string, error) {
	doubleDAO := mappers.CreateDoubleDTOtoDAO(doubleDTO)

	return a.adapter.CreateDouble(ctx, doubleDAO)
}

func (a *CategoryQuerierAdapter) CreateTeam(ctx context.Context, teamDTO *dto.CreateTeamDTOReq) (string, error) {
	teamDAO, err := mappers.CreateTeamDTOtoDAO(teamDTO, a.ConvertToObjectID)
	if err != nil {
		return "", nil
	}

	return a.adapter.CreateTeam(ctx, teamDAO)
}

func (a *CategoryQuerierAdapter) CreateCompetitorStats(ctx context.Context, competitorID string) error {
	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	return a.adapter.CreateCompetitorStats(ctx, competitorOID)
}

func (a *CategoryQuerierAdapter) GetCategoryRegistrationSortedByPoints(ctx context.Context, categoryID string) ([]*dto.GetCategoryRegistrationSortedByPointsDTORes, error) {
	categoryRegistrationSortedDAO, err := a.adapter.GetCategoryRegistrationSortedByPoints(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	categoryRegistrationSortedDTO := mappers.GetCategoryRegistrationSortedByPointsDAOtoDTO(categoryRegistrationSortedDAO)

	return categoryRegistrationSortedDTO, nil
}

func (a *CategoryQuerierAdapter) UpdateCategoryRegistrationCurrentPosition(ctx context.Context, categoryID string, categoryRegistrationDTO []*dto.GetCategoryRegistrationSortedByPointsDTORes) error {
	categoryRegistrationDAO, categoryOID, err := mappers.UpdateCategoryRegistrationCurrentPositionDTOtoDAO(categoryRegistrationDTO, categoryID, a.ConvertToObjectID)
	if err != nil {
		return err
	}

	return a.adapter.UpdateCategoryRegistrationCurrentPosition(ctx, categoryOID, categoryRegistrationDAO)
}
