package repository

import (
	"context"
	"fmt"
	"sync"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	availability_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
	double_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/double/dao"
	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	role_models "github.com/DBrange/didis-comp-bk/domains/repository/models/role"
	single_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/single/dao"
	team_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/team/dao"
	user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type locationResult struct {
	ID  string
	Err error
}
type roleResult struct {
	Role *role_models.Role
	Err  error
}

// func orDone(ctx context.Context, c <-chan interface{}) <-chan interface{} {
// 	relayStream := make(chan interface{})
// 	go func() {
// 		defer close(relayStream)
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				return
// 			case v, ok := <-c:
// 				if !ok {
// 					return
// 				}
// 				select {
// 				case relayStream <- v:
// 				case <-ctx.Done():
// 					return
// 				}
// 			}
// 		}
// 	}()
// 	return relayStream
// }

// func (r *Repository) CreateUserAndLocation(ctx context.Context, userInfoDAO *user_dao.CreateUserDAOReq, locationInfoDAO *location_dao.CreateLocationDAOReq, organizer bool) error {
// 	// r.InitialiseRole(ctx)
// 	return r.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
// 		wg := &sync.WaitGroup{}

// 		locationCh := make(chan *locationResult, 1)
// 		roleCh := make(chan *roleResult, 1)
// 		wg.Add(2)
// 		go r.createLocationConcurrently(sessCtx, locationInfoDAO, wg, locationCh)
// 		go r.getRoleByNameAndTypeConcurrently(sessCtx, "FREE", "USER", wg, roleCh)

// 		go func() {
// 			wg.Wait()
// 			close(locationCh)
// 			close(roleCh)
// 		}()

// 		var locationID string
// 		var role role_models.Role
// 		var err error

// 		for i := 0; i < 2; i++ {
// 			select {
// 			case lr, ok := <-locationCh:
// 				if !ok {
// 					locationCh = nil // Evitar leer de un canal cerrado
// 					continue
// 				}
// 				if lr.Err != nil {
// 					err = lr.Err
// 					break
// 				}
// 				userInfoDAO.LocationID = &locationID

// 			case rr, ok := <-roleCh:
// 				if !ok {
// 					roleCh = nil // Evitar leer de un canal cerrado
// 					continue
// 				}
// 				if rr.Err != nil {
// 					err = rr.Err
// 					break
// 				}
// 				userInfoDAO.Roles = append(userInfoDAO.Roles, role.ID)

// 			case <-sessCtx.Done():
// 				return sessCtx.Err()
// 			}
// 			if err != nil {
// 				return err
// 			}
// 		}

// 		// Crear usuario
// 		userID, err := r.CreateUser(sessCtx, userInfoDAO)
// 		if err != nil {
// 			return err
// 		}

// 		// Creat organizer
// 		if organizer {
// 			return r.CreateOrganizer(ctx, userID)
// 		} else {
// 			return r.CreateAvailability(sessCtx, &userID, nil)
// 		}
// 	})

// }

func (r *Repository) UpdateProfileAvailability(ctx context.Context, availabilityID string, availabilityInfoDAO *availability_dao.UpdateDailyAvailabilityDAOReq) error {
	return r.UpdateAvailability(ctx, availabilityID, availabilityInfoDAO)
}

func (r *Repository) UpdatePersonalInfo(ctx context.Context, userID string, userInfoDAO *user_dao.UpdateUserDAOReq, locationInfoDAO *location_dao.UpdateLocationDAOReq) error {
	return r.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {

		wg := &sync.WaitGroup{}

		errCh := make(chan error, 2)

		wg.Add(2)
		go r.updateUserConcurrently(sessCtx, userID, userInfoDAO, wg, errCh)
		go r.updateLocationConcurrently(sessCtx, locationInfoDAO.ID, locationInfoDAO, wg, errCh)

		wg.Wait()
		close(errCh)

		for err := range errCh {
			if err != nil {
				return err
			}
		}

		select {
		case <-sessCtx.Done():
			return sessCtx.Err()
		default:
			return nil
		}

	})
}

func (r *Repository) GetPersonalInfoByID(ctx context.Context, userID string) (*user_dao.GetUserByIDDAO, *location_dao.GetLocationByIDDAORes, error) {
	userInfo, err := r.GetUserByID(ctx, userID)
	if err != nil {
		return nil, nil, err
	}

	locationInfo, err := r.GetLocationByID(ctx, *userInfo.LocationID)
	if err != nil {
		return nil, nil, err
	}

	return userInfo, locationInfo, nil
}

func (r *Repository) GetProfileAvailabilityInfoByID(ctx context.Context, availabilityID string, day string) (*availability_dao.GetDailyAvailabilityInfoByIDDAORes, error) {
	return r.GetAvailabilityInfoByID(ctx, availabilityID, day)
}

func (r *Repository) DeleteProfile(ctx context.Context, userID string) error {
	return r.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		projections := bson.M{
			"location_id": 1,
		}

		elemForSetDeletedAt, err := setDeletedAtAndReturnIDs(ctx, r.userColl, userID, "user", projections, &user_dao.UserRelationsToDeleteDAO{})
		if err != nil {
			return err
		}

		locationID := elemForSetDeletedAt.LocationID

		availabilityID, err := r.GetAvailabilityByUserID(ctx, userID)
		if err != nil {
			return err
		}

		err = r.setDeletedAt(ctx, r.locationColl, locationID, "location")
		if err != nil {
			return err
		}

		err = r.setDeletedAt(ctx, r.availabilityColl, availabilityID, "availability")
		if err != nil {
			return err
		}
		return nil
	})

}

func (r *Repository) UpdateProfilePassword(ctx context.Context, userID, newPassword, oldPassword string) error {
	return r.UpdateUserPassword(ctx, userID, newPassword, oldPassword)
}

func (r *Repository) CreateCompetitorType(ctx context.Context, competitorType models.COMPETITOR_TYPE) (*primitive.ObjectID, error) {
	type createTypeCompetitor func(ctx context.Context) (*primitive.ObjectID, error)

	createMap := map[models.COMPETITOR_TYPE]createTypeCompetitor{
		models.COMPETITOR_TYPE_SINGLE: func(ctx context.Context) (*primitive.ObjectID, error) {
			singleDAO := &single_dao.CreateSingleDAOReq{}
			return r.CreateSingle(ctx, singleDAO)
		},
		models.COMPETITOR_TYPE_DOUBLE: func(ctx context.Context) (*primitive.ObjectID, error) {
			doubleDAO := &double_dao.CreateDoubleDAOReq{}
			return r.CreateDouble(ctx, doubleDAO)
		},
		models.COMPETITOR_TYPE_TEAM: func(ctx context.Context) (*primitive.ObjectID, error) {
			teamDAO := &team_dao.CreateTeamDAOReq{}
			return r.CreateTeam(ctx, teamDAO)
		},
	}

	create, ok := createMap[competitorType]
	if !ok {
		return nil, fmt.Errorf("error competitor type no exists: %w", customerrors.ErrNotFound)
	}

	return create(ctx)
}

func (r *Repository) RegisterCompetitor(ctx context.Context, userID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) error {
	return r.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		// Convert to OID
		userOID, err := r.ConvertToObjectID(userID)
		if err != nil {
			return err
		}

		// Create type of competitor
		competiorTypeOID, err := r.CreateCompetitorType(ctx, competitorType)
		if err != nil {
			return nil
		}

		// Create competitor
		competitorID, err := r.CreateCompetitor(ctx, sport, competitorType, competiorTypeOID)
		if err != nil {
			return err
		}

		// Convert to OID
		competitorOID, err := r.ConvertToObjectID(competitorID)
		if err != nil {
			return err
		}

		// Create availability
		if err := r.CreateAvailability(ctx, nil, &competitorID); err != nil {
			return err
		}

		// Create competitor stats
		if err := r.CreateCompetitorStats(ctx, competitorOID); err != nil {
			return err
		}

		// Create competitor_user
		if err := r.CreateCompetitorUser(ctx, userOID, competitorOID); err != nil {
			return err
		}

		return nil
	})
}
