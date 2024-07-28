package services

import (
	"context"
	"fmt"

	models "github.com/DBrange/didis-comp-bk/cmd/api/models/options/tournament"
	"github.com/DBrange/didis-comp-bk/domains/tournament/adapters/mappers"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *TournamentService) OrganizeTournament(ctx context.Context, organizeTournamentDTO *dto.OrganizeTournamentDTOReq, options *models.OrganizeTournamentOptions) error {
	err := s.tournamentQueryer.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		// Mapping info
		tournamentDTO, locationDTO, leagueID, organizerID := mappers.OrganizeTournamentMapper(organizeTournamentDTO)

		// Verity if organizer exists
		if err := s.tournamentQueryer.VerifyOrganizerExists(sessCtx, organizerID); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error organizer not exists")
		}

		// Verity if league exists
		if leagueID != nil {
			if err := s.tournamentQueryer.VerifyLeagueExists(sessCtx, *leagueID); err != nil {
				return customerrors.HandleErrMsg(err, "tournament", "error league not exists")
			}
		}

		// Create location for tournament
		locationID, err := s.tournamentQueryer.CreateLocation(sessCtx, locationDTO)
		if err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when creating tournament location")
		}

		// Create tournament
		tournamentID, err := s.tournamentQueryer.CreateTournament(
			sessCtx,
			tournamentDTO,
			locationID,
			options,
			leagueID,
			organizerID,
		)
		if err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when creating tournament")
		}

		// If a league ID is available, add this tournament to the league
		if leagueID != nil {
			if err := s.tournamentQueryer.AddTournamentInLeague(sessCtx, *leagueID, tournamentID); err != nil {
				return customerrors.HandleErrMsg(err, "tournament", "error when adding tournament to league")
			}
		}

		// Optional actions
		if options.DoubleElimination || options.QuantityGroups > 0 || options.QuantityPots > 0 {
			if err = s.optionalActions(sessCtx, options, tournamentID); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when creating tournament location")
	}

	return nil
}

func (s *TournamentService) optionalActions(
	sessCtx mongo.SessionContext,
	options *models.OrganizeTournamentOptions,
	tournamentID string,
) error {
	var tournamentOptions dto.UpdateTournamentOptionsDTOReq

	// update tournamentOptions
	err := s.updateOptions(sessCtx, options, &tournamentOptions, tournamentID)
	if err != nil {
		return nil
	}

	//Update tournament
	if err = s.tournamentQueryer.UpdateTournamentOptions(sessCtx, tournamentID, &tournamentOptions, true); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when updating tournament")
	}

	return nil
}

func (s *TournamentService) updateOptions(
	sessCtx mongo.SessionContext,
	options *models.OrganizeTournamentOptions,
	tournamentOptions *dto.UpdateTournamentOptionsDTOReq,
	tournamentID string,
) error {
	// Create Pots
	if options.QuantityPots > 0 {
		tournamentOptions.Pots = &[]string{}
		err := s.createOptionals(sessCtx, tournamentID, options.QuantityPots, "pot", tournamentOptions.Pots, s.tournamentQueryer.CreatePot)
		if err != nil {
			return err
		}
	}

	//Create groups
	if options.QuantityGroups > 0 {
		tournamentOptions.Groups = &[]string{}
		err := s.createOptionals(sessCtx, tournamentID, options.QuantityGroups, "group", tournamentOptions.Groups, s.tournamentQueryer.CreateTournamentGroup)
		if err != nil {
			return err
		}
	}

	//Crete double elimination
	if options.DoubleElimination {
		doubleEliminationID, err := s.tournamentQueryer.CreateDoubleElimination(sessCtx)
		if err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when creating double elimination")
		}

		tournamentOptions.DoubleEliminationID = &doubleEliminationID
	}

	return nil
}

type FnCreate func(ctx context.Context, tournamentID string) (string, error)

func (s *TournamentService) createOptionals(
	sessCtx mongo.SessionContext,
	tournamentID string,
	quantity int,
	name string,
	dest *[]string,
	fnCreate FnCreate,
) error {
	IDs := make([]string, quantity)
	for i := 0; i < quantity; i++ {
		vID, err := fnCreate(sessCtx, tournamentID)
		if err != nil {
			errMsg := fmt.Sprintf("error when creating %sID", name)
			return customerrors.HandleErrMsg(err, "tournament", errMsg)
		}

		IDs[i] = vID
	}

	*dest = IDs

	return nil
}
