package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (d *TournamentService) OrganizeTournament(ctx context.Context, tournamentInfoDTO *dto.OrganizeTournamentDTOReq) error {
	if err := d.tournamentQueryer.OrganizeTournament(ctx, tournamentInfoDTO); err != nil {
		tournamentErrorHandlers := customerrors.CreateErrorHandlers("tournament")
		errMsgTemplate := "error when registering tournament"
		return customerrors.HandleError(err, tournamentErrorHandlers, errMsgTemplate)
	}

	return nil
}
