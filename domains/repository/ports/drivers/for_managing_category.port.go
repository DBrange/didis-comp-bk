package ports

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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ForManagingCategory interface {
	VerifyOrganizerExists(ctx context.Context, organizerID string) error
	CreateCategory(ctx context.Context, organizerID string, categoryInfoDAO *category_dao.CreateCategoryDAOReq) (string, error)
	ConvertToObjectID(ID string) (*primitive.ObjectID, error)
	AddTournamentInCategory(ctx context.Context, categoryID string, tournamentID string) error
	AddCategoryInTournament(ctx context.Context, tournamentID string, categoryID string) error
	CreateCategoryRegistration(ctx context.Context, categoryRegistrationInfoDAO *category_registration_dao.CreateCategoryRegistrationDAOReq) error
	VerifyCategoryExists(ctx context.Context, categoryID string) error
	VerifyCompetitorExists(ctx context.Context, competitorID string) error
	VerifyCategoryExistsRelation(ctx context.Context, categoryRegistrationDAO *category_registration_dao.CreateCategoryRegistrationDAOReq) error
	GetCompetitorsOfCategoryByName(ctx context.Context, categoryOID string, name string, sport models.SPORT, competitorType models.COMPETITOR_TYPE, limit int, lastOID string) ([]*category_registration_dao.GetCompetitorsOfCategoryDAORes, error)
	GetCompetitorsFollowed(ctx context.Context, userID string, name string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) ([]*follower_dao.GetCompetitorFollowedDAORes, error)
	UpdateCategory(ctx context.Context, categoryID *primitive.ObjectID, categoryInfoDAO *category_dao.UpdateCategoryDAOReq) error
	GetCategoryInfoByID(ctx context.Context, categoryOID string) (*category_dao.GetCategoryInfoByIDDAORes, error)
	IncrementTotalParticipants(ctx context.Context, categoryOID *primitive.ObjectID) error
	DecrementTotalParticipants(ctx context.Context, categoryOID *primitive.ObjectID) error
	GetParticipantsOfCategory(ctx context.Context, categoryID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE, limit int, lastID string) ([]*category_registration_dao.GetCompetitorsOfCategoryDAORes, error)
	CategoryRegistrationColl() *mongo.Collection
	PermaDeleteCategoryRegistration(ctx context.Context, mc *mongo.Collection, ID string) error
	AddCategoryInOrganizer(ctx context.Context, organizerOID, categoryOID *primitive.ObjectID) error
	GetCategoriesFromOrganizer(ctx context.Context, organizerOID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) ([]organizer_dao.GetCategoriesFromOrganizerDAORes, error)
	GetTournamentsFromCategory(ctx context.Context, categoryOID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE, limit int, lastID string) ([]category_dao.GetTournamentsFromCategoryDAORes, error)
	UpdateCompetitorPoints(ctx context.Context, categoryOID, competitorOID *primitive.ObjectID, points int) error
	CreateGuestUser(ctx context.Context, guestUserInfoDAO *guest_user_dao.CreateGuestUserDAOReq) (string, error)
	CreateCompetitor(ctx context.Context, sport models.SPORT, competitorType models.COMPETITOR_TYPE, OID *primitive.ObjectID) (string, error)
	CreateGuestCompetitor(ctx context.Context, guestCompetitorInfoDAO *guest_competitor_dao.CreateGuestCompetitorDAOReq) (string, error)
	CreateSingle(ctx context.Context, singleInfoDAO *single_dao.CreateSingleDAOReq) (string, error)
	CreateDouble(ctx context.Context, doubleInfoDAO *double_dao.CreateDoubleDAOReq) (string, error)
	CreateTeam(ctx context.Context, teamInfoDAO *team_dao.CreateTeamDAOReq) (string, error)
	CreateCompetitorStats(ctx context.Context, competitorOID *primitive.ObjectID) error
	GetCategoryRegistrationSortedByPoints(ctx context.Context, categoryID string) ([]*category_registration_dao.GetCategoryRegistrationSortedByPointsDAORes, error)
	UpdateCategoryRegistrationCurrentPosition(ctx context.Context, categoryOID *primitive.ObjectID, categoryRegistration []*category_registration_dao.GetCategoryRegistrationSortedByPointsDAORes) error
	
}
