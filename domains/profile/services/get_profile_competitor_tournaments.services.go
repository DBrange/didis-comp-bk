package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *ProfileService) GetProfileCompetitorTournaments(
	ctx context.Context,
	competitorID,categoryID string,
	sport models.SPORT,
	limit int,
	lastID string,
) (*dto.GetProfileUserTournamentsDTORes, error) {
	tournametsDTO, err := s.profileQuerier.GetProfileCompetitorTournaments(ctx, competitorID,categoryID, sport, limit, lastID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "profile", "error getting profile tournaments")
	}

	return tournametsDTO, nil
}