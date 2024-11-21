package services

import (
	"context"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *CategoryService) ModifyCompetitorPoints(ctx context.Context, categoryID, competitorID string, points int) error {
	if err := s.categoryQuerier.UpdateCompetitorPoints(ctx, categoryID, competitorID, points); err != nil {
		return customerrors.HandleErrMsg(err, "category", "error when modifing competitor points")
	}

	// Si hay cambios en el ranking de algun competidor, agregarlos al slice de registered_positions (numero y hora)
	if err := s.UpdateCategoryRanking(ctx, categoryID); err != nil {
		return err
	}
	return nil
}

func (s *CategoryService) UpdateCategoryRanking(ctx context.Context, categoryID string) error {
	rankingSorted, err := s.categoryQuerier.GetCategoryRegistrationSortedByPoints(ctx, categoryID)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when getting categoryRegistration serted")
	}

	if err := s.categoryQuerier.UpdateCategoryRegistrationCurrentPosition(ctx, categoryID, rankingSorted); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when updating ranking after the end of tournament")
	}

	return nil
}
