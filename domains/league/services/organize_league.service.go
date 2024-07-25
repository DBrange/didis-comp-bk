package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/league/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (d *LeagueService) OrganizeLeague(ctx context.Context, organizerID string,leagueInfoDTO *dto.OrganizeLeagueDTOReq) error {
	if err := d.leagueQueryer.OrganizeLeague(ctx,organizerID , leagueInfoDTO); err != nil {
		leagueErrorHandlers := customerrors.CreateErrorHandlers("league")
		errMsgTemplate := "error when registering league"
		return customerrors.HandleError(err, leagueErrorHandlers, errMsgTemplate)
	}

	return nil
}
