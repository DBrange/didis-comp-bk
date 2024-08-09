package services

import (
	"context"
	"sync"

	"github.com/DBrange/didis-comp-bk/domains/profile/adapters/mappers"
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

type locationResult struct {
	ID  string
	Err error
}
type roleResult struct {
	Role *dto.GetRoleDTOByID
	Err  error
}

func (s *ProfileService) RegisterUser(ctx context.Context, profileInfoDTO *dto.RegisterUserDTOReq) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// err := s.profileQueryer.WithTransaction(ctx, func(ctx mongo.SessionContext) error {
// s.profileQueryer.InitialiseRole(ctx)
	wg := &sync.WaitGroup{}

	locationCh := make(chan *locationResult, 1)
	roleCh := make(chan *roleResult, 1)
	var err error

	userDTO, locationDTO := mappers.RegisterUserMapper(profileInfoDTO)

	hashedPassword, err := s.HashPassword(*userDTO.Password)
	if err != nil {
		return customerrors.HandleErrMsg(err, "profile", "error when hashing password")
	}

	userDTO.Password = &hashedPassword

	wg.Add(2)
	go s.createLocationConcurrently(ctx, locationDTO, wg, locationCh)
	go s.getRoleByNameAndTypeConcurrently(ctx, "FREE", "USER", wg, roleCh)

	// Esperar a que todas las goroutines terminen y cerrar los canales
	go func() {
		wg.Wait()
		close(locationCh)
		close(roleCh)
	}()

	// Procesar los resultados
	for i := 0; i < 2; i++ {
		select {
		case lr := <-locationCh:
			if lr.Err != nil {
				cancel() // Cancelar el contexto en caso de error
				return lr.Err
			}
			userDTO.LocationID = &lr.ID

		case rr := <-roleCh:
			if rr.Err != nil {
				cancel() // Cancelar el contexto en caso de error
				return rr.Err
			}
			userDTO.Roles = append(userDTO.Roles, rr.Role.ID)

		case <-ctx.Done():
			return ctx.Err()
		}
	}

	// Crear el usuario
	userID, err := s.profileQueryer.CreateUser(ctx, userDTO)
	if err != nil {
		return customerrors.HandleErrMsg(err, "profile", "error when creating user")
	}

	// Crear el organizador o disponibilidad según sea necesario
	if profileInfoDTO.Organizer {
		err := s.profileQueryer.CreateOrganizer(ctx, userID)
		if err != nil {
			return customerrors.HandleErrMsg(err, "profile", "error when creating organizer")
		}
	} else {
		err := s.profileQueryer.CreateAvailability(ctx, &userID, nil)
		if err != nil {
			return customerrors.HandleErrMsg(err, "profile", "error when creating availability")
		}
	}

	return nil
	// })
	// if err != nil {
	// 	// return customerrors.HandleErrMsg(err, "profile", "error when registering profile")
	// 	return err
	// }

	//ACTIVAR LUEGO
	// s.profileQueryer.ActivateUserNotification(ctx)

}

func (s *ProfileService) createLocationConcurrently(ctx context.Context, locationInfoDTO *dto.CreateLocationDTOReq, wg *sync.WaitGroup, locationCh chan<- *locationResult) {
	defer wg.Done()
	locationID, err := s.profileQueryer.CreateLocation(ctx, locationInfoDTO)
	locationCh <- &locationResult{ID: locationID, Err: err}
}

func (s *ProfileService) getRoleByNameAndTypeConcurrently(ctx context.Context, roleName string, roleType string, wg *sync.WaitGroup, roleCh chan<- *roleResult) {
	defer wg.Done()
	role, err := s.profileQueryer.GetRoleByNameAndType(ctx, roleName, roleType)
	roleCh <- &roleResult{Role: role, Err: err}
}
