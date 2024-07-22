package services

import (
	"context"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (d *LeagueService) OrganizeLeague(ctx context.Context, leagueInfoDTO any) error {
	if err := d.leagueQueryer.OrganizeLeague(ctx, leagueInfoDTO); err != nil {
		leagueErrorHandlers := customerrors.CreateErrorHandlers("league")
		errMsgTemplate := "error when registering league"
		return customerrors.HandleError(err, leagueErrorHandlers, errMsgTemplate)
	}

	return nil
}
