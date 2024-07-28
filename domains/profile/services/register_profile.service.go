package services

import (
	"context"
	"sync"

	"github.com/DBrange/didis-comp-bk/domains/profile/adapters/mappers"
	profile_dto "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	role_dto "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type locationResult struct {
	ID  string
	Err error
}
type roleResult struct {
	Role *role_dto.GetRoleDTOByID
	Err  error
}

func (s *ProfileService) RegisterUser(ctx context.Context, profileInfoDTO *profile_dto.RegisterUserDTOReq) error {
	err := s.profileQueryer.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		wg := &sync.WaitGroup{}

		locationCh := make(chan *locationResult, 1)
		roleCh := make(chan *roleResult, 1)
		var err error

		userDTO, locationDTO := mappers.RegisterUserMapper(profileInfoDTO)

		hashedPassword, err := s.HashPassword(*userDTO.Password)
		if err != nil {
			return err
		}

		userDTO.Password = &hashedPassword

		wg.Add(2)
		go s.createLocationConcurrently(sessCtx, locationDTO, wg, locationCh)
		go s.getRoleByNameAndTypeConcurrently(sessCtx, "FREE", "USER", wg, roleCh)

		go func() {
			wg.Wait()
			close(locationCh)
			close(roleCh)
		}()

		for i := 0; i < 2; i++ {
			select {
			case lr, ok := <-locationCh:
				if !ok {
					locationCh = nil // Evitar leer de un canal cerrado
					continue
				}
				if lr.Err != nil {
					err = lr.Err
					break
				}
				userDTO.LocationID = &lr.ID

			case rr, ok := <-roleCh:
				if !ok {
					roleCh = nil // Evitar leer de un canal cerrado
					continue
				}
				if rr.Err != nil {
					err = rr.Err
					break
				}
				userDTO.Roles = append(userDTO.Roles, rr.Role.ID)

			case <-sessCtx.Done():
				return sessCtx.Err()
			}
			if err != nil {
				return err
			}
		}

		// Create usuario
		userID, err := s.profileQueryer.CreateUser(sessCtx, userDTO)
		if err != nil {
			return err
		}

		// Create organizer
		if profileInfoDTO.Organizer {
			err := s.profileQueryer.CreateOrganizer(ctx, userID)
			if err != nil {
				return err
			}
		} else {
			err := s.profileQueryer.CreateAvailability(sessCtx, &userID, nil)
			if err != nil {
				return err
			}

		}

		//ACTIVAR LUEGO
		// s.profileQueryer.ActivateUserNotification(sessCtx)

		return nil
	})

	if err != nil {
		return customerrors.HandleErrMsg(err, "profile", "error when registering profile")
	}

	return nil
}

func (s *ProfileService) createLocationConcurrently(sessCtx mongo.SessionContext, locationInfoDTO *profile_dto.CreateLocationDTOReq, wg *sync.WaitGroup, locationCh chan<- *locationResult) {
	defer wg.Done()
	locationID, err := s.profileQueryer.CreateLocation(sessCtx, locationInfoDTO)
	locationCh <- &locationResult{ID: locationID, Err: err}
}

func (s *ProfileService) getRoleByNameAndTypeConcurrently(sessCtx mongo.SessionContext, roleName string, roleType string, wg *sync.WaitGroup, roleCh chan<- *roleResult) {
	defer wg.Done()
	role, err := s.profileQueryer.GetRoleByNameAndType(sessCtx, roleName, roleType)
	roleCh <- &roleResult{Role: role, Err: err}

}
