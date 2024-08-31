package services

import (
	"context"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (d *CategoryService) AddTournamentInCategory(ctx context.Context, categoryID string, tournamentID string) error {
	if err := d.categoryQuerier.AddTournamentInCategory(ctx, categoryID, tournamentID); err != nil {
		return customerrors.HandleErrMsg(err, "category", "error when add tournament in category")
	}

	return nil
}
