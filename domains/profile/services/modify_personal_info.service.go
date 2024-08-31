package services

import (
	"context"
	"sync"

	"github.com/DBrange/didis-comp-bk/domains/profile/adapters/mappers"
	profile_dto "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *ProfileService) ModifyPersonalInfo(ctx context.Context, userID string, userInfoDTO *profile_dto.ModifyPersonalInfoDTOReq) error {
	err := s.profileQuerier.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		userDTO, locationDTO := mappers.ModifyPersonalInfoMapper(userInfoDTO)
		wg := &sync.WaitGroup{}

		errCh := make(chan error, 2)

		wg.Add(2)
		go s.updateUserConcurrently(sessCtx, userID, userDTO, wg, errCh)
		go s.updateLocationConcurrently(sessCtx, locationDTO.ID, locationDTO, wg, errCh)

		wg.Wait()
		close(errCh)

		for err := range errCh {
			if err != nil {
				return err
			}
		}

		// select {
		// case <-sessCtx.Done():
		// 	return sessCtx.Err()
		// default:
		// 	return nil
		// }

		return nil

	})

	if err != nil {
		return customerrors.HandleErrMsg(err, "profile", "error updating profile")
	}

	return nil
}

func (s *ProfileService) updateUserConcurrently(sessCtx mongo.SessionContext, userID string, userInfoDAO *profile_dto.UpdateUserDTOReq, wg *sync.WaitGroup, errCh chan<- error) {
	defer wg.Done()
	if err := s.profileQuerier.UpdateUser(sessCtx, userID, userInfoDAO); err != nil {
		errCh <- err
	}
}

func (s *ProfileService) updateLocationConcurrently(sessCtx mongo.SessionContext, locationID string, locationInfoDAO *profile_dto.UpdateLocationDTOReq, wg *sync.WaitGroup, errCh chan<- error) {
	defer wg.Done()
	if err := s.profileQuerier.UpdateLocation(sessCtx, locationID, locationInfoDAO); err != nil {
		errCh <- err
	}
}
