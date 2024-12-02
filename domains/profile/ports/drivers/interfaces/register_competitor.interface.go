package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type RegisterCompetitor interface {
	RegisterCompetitor(ctx context.Context, userIDs []string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) error
}