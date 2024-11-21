package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	 "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
)

type GetProfileCategories interface {
	GetProfileCategories(ctx context.Context, userID string, sport models.SPORT, limit int ,lastID string) (*dto.GetUserCategoriesDTO, error)
}
