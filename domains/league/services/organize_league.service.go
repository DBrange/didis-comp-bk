package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/league/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *LeagueService) OrganizeLeague(ctx context.Context, organizerID string, leagueDTO *dto.CreateLeagueDTOReq) error {
	if err := s.leagueQueryer.VerifyOrganizerExists(ctx, organizerID); err != nil {
		return customerrors.HandleErrMsg(err, "league", "error organizer not exits")
	}

	if err := s.leagueQueryer.CreateLeague(ctx, organizerID, leagueDTO); err != nil {
		return customerrors.HandleErrMsg(err, "league", "error when creating lead")
	}

	return nil

}
