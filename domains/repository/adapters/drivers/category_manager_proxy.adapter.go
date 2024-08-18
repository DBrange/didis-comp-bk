package adapters

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	category_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/category/dao"
	double_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/double/dao"
	guest_user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/guest_user/dao"
	category_registration_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/category_registration/dao"
	follower_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/follower/dao"
	guest_competitor_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/guest_competitor/dao"
	organizer_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/organizer/dao"
	single_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/single/dao"
	team_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/team/dao"
	"github.com/DBrange/didis-comp-bk/domains/repository/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryManagerProxyAdapter struct {
	repository *repository.Repository
}

func NewCategoryManagerProxyAdapter(repository *repository.Repository) *CategoryManagerProxyAdapter {
	return &CategoryManagerProxyAdapter{
		repository: repository,
	}
}

func (a *CategoryManagerProxyAdapter) VerifyOrganizerExists(ctx context.Context, organizerID string) error {
	return a.repository.VerifyOrganizerExists(ctx, organizerID)
}

func (a *CategoryManagerProxyAdapter) CreateCategory(ctx context.Context, organizerID string, categoryInfoDAO *category_dao.CreateCategoryDAOReq) (string, error) {
	return a.repository.CreateCategory(ctx, organizerID, categoryInfoDAO)
}

func (a *CategoryManagerProxyAdapter) ConvertToObjectID(ID string) (*primitive.ObjectID, error) {
	return a.repository.ConvertToObjectID(ID)
}

func (a *CategoryManagerProxyAdapter) AddTournamentInCategory(ctx context.Context, categoryID string, tournamentID string) error {
	return a.repository.AddTournamentInCategory(ctx, categoryID, tournamentID)
}

func (a *CategoryManagerProxyAdapter) AddCategoryInTournament(ctx context.Context, tournamentID string, categoryID string) error {
	return a.repository.AddCategoryInTournament(ctx, tournamentID, categoryID)
}

func (a *CategoryManagerProxyAdapter) CreateCategoryRegistration(ctx context.Context, categoryRegistrationInfoDAO *category_registration_dao.CreateCategoryRegistrationDAOReq) error {
	return a.repository.CreateCategoryRegistration(ctx, categoryRegistrationInfoDAO)
}

func (a *CategoryManagerProxyAdapter) VerifyCategoryExists(ctx context.Context, categoryID string) error {
	return a.repository.VerifyCategoryExists(ctx, categoryID)
}

func (a *CategoryManagerProxyAdapter) VerifyCompetitorExists(ctx context.Context, competitorID string) error {
	competitorOID, err := a.ConvertToObjectID(competitorID)
	if err != nil {
		return err
	}

	return a.repository.VerifyCompetitorExists(ctx, competitorOID)
}

func (a *CategoryManagerProxyAdapter) VerifyCategoryExistsRelation(ctx context.Context, categoryRegistrationDAO *category_registration_dao.CreateCategoryRegistrationDAOReq) error {
	return a.repository.VerifyCategoryExistsRelation(ctx, categoryRegistrationDAO)
}

func (a *CategoryManagerProxyAdapter) GetCompetitorsOfCategoryByName(ctx context.Context, categoryID string, name string, sport models.SPORT, competitorType models.COMPETITOR_TYPE, limit int, lastID string) ([]*category_registration_dao.GetCompetitorsOfCategoryDAORes, error) {
	categoryOID, err := a.ConvertToObjectID(categoryID)
	if err != nil {
		return nil, err
	}

	var lastOID *primitive.ObjectID
	if lastID != "" {
		lastOID, err = a.ConvertToObjectID(lastID)
		if err != nil {
			return nil, err
		}
	} else {
		lastOID = nil
	}
	return a.repository.GetCompetitorsOfCategoryByName(ctx, categoryOID, sport, competitorType, limit, lastOID, name)
}

func (a *CategoryManagerProxyAdapter) GetCompetitorsFollowed(
	ctx context.Context,
	userID string,
	name string,
	sport models.SPORT,
	competitorType models.COMPETITOR_TYPE,
) ([]*follower_dao.GetCompetitorFollowedDAORes, error) {
	userOID, err := a.ConvertToObjectID(userID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetCompetitorsFollowed(ctx, userOID, name, sport, competitorType)
}

func (a *CategoryManagerProxyAdapter) UpdateCategory(ctx context.Context, categoryOID *primitive.ObjectID, categoryInfoDAO *category_dao.UpdateCategoryDAOReq) error {
	return a.repository.UpdateCategory(ctx, categoryOID, categoryInfoDAO)
}

func (a *CategoryManagerProxyAdapter) GetCategoryInfoByID(ctx context.Context, categoryID string) (*category_dao.GetCategoryInfoByIDDAORes, error) {
	categoryOID, err := a.ConvertToObjectID(categoryID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetCategoryInfoByID(ctx, categoryOID)
}

func (a *CategoryManagerProxyAdapter) IncrementTotalParticipants(ctx context.Context, categoryOID *primitive.ObjectID) error {
	return a.repository.IncrementTotalParticipants(ctx, categoryOID)
}

func (a *CategoryManagerProxyAdapter) DecrementTotalParticipants(ctx context.Context, categoryOID *primitive.ObjectID) error {
	return a.repository.DecrementTotalParticipants(ctx, categoryOID)
}

func (a *CategoryManagerProxyAdapter) GetParticipantsOfCategory(ctx context.Context, categoryID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE, limit int, lastID string) ([]*category_registration_dao.GetCompetitorsOfCategoryDAORes, error) {
	categoryOID, err := a.ConvertToObjectID(categoryID)
	if err != nil {
		return nil, err
	}

	var lastOID *primitive.ObjectID
	if lastID != "" {
		lastOID, err = a.ConvertToObjectID(lastID)
		if err != nil {
			return nil, err
		}
	} else {
		lastOID = nil
	}

	return a.repository.GetParticipantsOfCategory(ctx, categoryOID, sport, competitorType, limit, lastOID)
}

func (a *CategoryManagerProxyAdapter) CategoryRegistrationColl() *mongo.Collection {
	return a.repository.CategoryRegistrationColl()
}

func (a *CategoryManagerProxyAdapter) PermaDeleteCategoryRegistration(ctx context.Context, mc *mongo.Collection, categoryRegistrationID string) error {
	return a.repository.PermaDeleteCategoryRegistration(ctx, mc, categoryRegistrationID)
}

func (a *CategoryManagerProxyAdapter) AddCategoryInOrganizer(ctx context.Context, organizerOID, categoryOID *primitive.ObjectID) error {
	return a.repository.AddCategoryInOrganizer(ctx, organizerOID, categoryOID)
}

func (a *CategoryManagerProxyAdapter) GetCategoriesFromOrganizer(ctx context.Context, organizerID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) ([]organizer_dao.GetCategoriesFromOrganizerDAORes, error) {
	organizerOID, err := a.ConvertToObjectID(organizerID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetCategoriesFromOrganizer(ctx, organizerOID, sport, competitorType)
}

func (a *CategoryManagerProxyAdapter) GetTournamentsFromCategory(ctx context.Context, categoryID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE, limit int, lastID string) ([]category_dao.GetTournamentsFromCategoryDAORes, error) {
	categoryOID, err := a.ConvertToObjectID(categoryID)
	if err != nil {
		return nil, err
	}

	var lastOID *primitive.ObjectID
	if lastID != "" {
		lastOID, err = a.ConvertToObjectID(lastID)
		if err != nil {
			return nil, err
		}
	} else {
		lastOID = nil
	}

	return a.repository.GetTournamentsFromCategory(ctx, categoryOID, sport, competitorType, limit, lastOID)

}

func (a *CategoryManagerProxyAdapter) UpdateCompetitorPoints(ctx context.Context, categoryOID, competitorOID *primitive.ObjectID, points int) error {
	return a.repository.UpdateCompetitorPoints(ctx, categoryOID, competitorOID, points)
}

func (a *CategoryManagerProxyAdapter) CreateGuestUser(ctx context.Context, guestUserInfoDAO *guest_user_dao.CreateGuestUserDAOReq) (string, error) {
	return a.repository.CreateGuestUser(ctx, guestUserInfoDAO)
}

func (a *CategoryManagerProxyAdapter) CreateCompetitor(ctx context.Context, sport models.SPORT, competitorType models.COMPETITOR_TYPE, OID *primitive.ObjectID) (string, error) {
	return a.repository.CreateCompetitor(ctx, sport, competitorType, OID)
}

func (a *CategoryManagerProxyAdapter) CreateGuestCompetitor(ctx context.Context, guestCompetitorInfoDAO *guest_competitor_dao.CreateGuestCompetitorDAOReq) (string, error) {
	return a.repository.CreateGuestCompetitor(ctx, guestCompetitorInfoDAO)
}

func (a *CategoryManagerProxyAdapter) CreateSingle(ctx context.Context, singleInfoDAO *single_dao.CreateSingleDAOReq) (string, error) {
	return a.repository.CreateSingle(ctx, singleInfoDAO)
}

func (a *CategoryManagerProxyAdapter) CreateDouble(ctx context.Context, doubleInfoDAO *double_dao.CreateDoubleDAOReq) (string, error) {
	return a.repository.CreateDouble(ctx, doubleInfoDAO)
}

func (a *CategoryManagerProxyAdapter) CreateTeam(ctx context.Context, teamInfoDAO *team_dao.CreateTeamDAOReq) (string, error) {
	return a.repository.CreateTeam(ctx, teamInfoDAO)
}

func (a *CategoryManagerProxyAdapter) CreateCompetitorStats(ctx context.Context, competitorOID *primitive.ObjectID) error {
	return a.repository.CreateCompetitorStats(ctx, competitorOID)

}

func (a *CategoryManagerProxyAdapter) GetCategoryRegistrationSortedByPoints(ctx context.Context, categoryID string) ([]*category_registration_dao.GetCategoryRegistrationSortedByPointsDAORes, error) {
	categoryOID, err := a.ConvertToObjectID(categoryID)
	if err != nil {
		return nil, err
	}

	return a.repository.GetCategoryRegistrationSortedByPoints(ctx, categoryOID)
}

func (a *CategoryManagerProxyAdapter) UpdateCategoryRegistrationCurrentPosition(ctx context.Context, categoryOID *primitive.ObjectID, categoryRegistration []*category_registration_dao.GetCategoryRegistrationSortedByPointsDAORes) error {
	return a.repository.UpdateCategoryRegistrationCurrentPosition(ctx, categoryOID, categoryRegistration)
}

func (a *CategoryManagerProxyAdapter) VerifyTournamentGroupInTournament(ctx context.Context, tournamentOID *primitive.ObjectID, groupOIDs []*primitive.ObjectID) error {
	return a.repository.VerifyTournamentGroupInTournament(ctx, tournamentOID, groupOIDs)
}

func (a *CategoryManagerProxyAdapter) DeletePotByPosition(ctx context.Context, position int, tournamentOID *primitive.ObjectID) error {
	return a.repository.DeletePotByPosition(ctx, position, tournamentOID)
}
