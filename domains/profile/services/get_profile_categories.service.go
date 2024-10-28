package services
import (
	"context"
	"github.com/DBrange/didis-comp-bk/cmd/api/models"

	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *ProfileService) GetProfileCategories(ctx context.Context, userID string,sport models.SPORT, limit int ,lastID string) ([]*dto.GetUserCategoriesCategoryDTO, error){
	categoriesDTO, err := s.profileQuerier.GetUserCategories(ctx, userID, sport, limit, lastID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "profile", "error getting user categories")
	}

	return categoriesDTO, nil
	
}