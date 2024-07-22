package repository

import (
	"context"
	"fmt"

	availability_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	user_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/user/dao"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateUserAndLocation(ctx context.Context, userInfoDAO *user_dao.CreateUserDAO, locationInfoDAO *location_dao.CreateLocationDAOReq, organizer bool) error {
		// r.InitialiseRole(ctx)
	err := r.withTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		// Crear ubicación
		locationID, err := r.CreateLocation(sessCtx, locationInfoDAO)
		if err != nil {
			return err
		}

		// Obtener rol
		role, err := r.GetRoleByNameAndType(sessCtx, "FREE", "USER")
		if err != nil {
			return err
		}

		// Actualizar DAO del usuario
		userInfoDAO.LocationID = &locationID
		userInfoDAO.Roles = append(userInfoDAO.Roles, role.ID)

		// Crear usuario
		userID, err := r.CreateUser(sessCtx, userInfoDAO)
		if err != nil {
			return err
		}

		// Creat organizer
		if organizer {
			if err := r.CreateOrganizer(ctx, userID); err != nil {
				return err
			}
		}

		// Crear disponibilidad
		if !organizer {
			if err := r.CreateAvailability(sessCtx, userID); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateProfileAvailability(ctx context.Context, availabilityID string, availabilityInfoDAO *availability_dao.UpdateDailyAvailabilityDAOReq) error {
	return r.UpdateAvailability(ctx, availabilityID, availabilityInfoDAO)
}

func (r *Repository) UpdatePersonalInfo(ctx context.Context, userID string, userInfoDAO *user_dao.UpdateUserDAOReq, locationInfoDAO *location_dao.UpdateLocationDAOReq) error {
	err := r.withTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		if err := r.UpdateUser(sessCtx, userID, userInfoDAO); err != nil {
			return fmt.Errorf("failed to update user: %w", err)
		}

		if locationInfoDAO != nil {
			if err := r.UpdateLocation(sessCtx, locationInfoDAO.ID, locationInfoDAO); err != nil {
				return fmt.Errorf("failed to update location: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
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
	err := r.withTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		projections := bson.M{
			"location_id": 1,
		}

		elemForSetDeletedAt, err := setDeletedAtAndReturnIDs(ctx, r.userColl, userID, "user", projections, &user_dao.UserRelationsToDeleteDAO{})
		if err != nil {
			return err
		}

		locationID := elemForSetDeletedAt.LocationID

		err = r.setDeletedAt(ctx, r.locationColl, locationID, "location")
		if err != nil {
			return err
		}

		availabilityID, err := r.GetAvailabilityByUserID(ctx, userID)
		if err != nil {
			return err
		}

		err = r.setDeletedAt(ctx, r.availabilityColl, availabilityID, "availability")
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateProfilePassword(ctx context.Context, userID, newPassword, oldPassword string) error {
	return r.UpdateUserPassword(ctx, userID, newPassword, oldPassword)
}
