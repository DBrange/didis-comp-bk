package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) GetTournamentsInOrganizer(
	ctx context.Context,
	organizerID string,
	sport models.SPORT,
	limit int,
	lastID string,
) (*dto.GetUserTournamentsDTORes, error) {
	tournamentsDTO, err := s.tournamentQuerier.GetTournamentsInOrganizer(ctx, organizerID,  sport, limit, lastID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "tournament", "error when getting user tournamets")
	}

	return tournamentsDTO, nil
}
