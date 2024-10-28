package dto

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type GetRoundGroupsDTORes struct {
	ID              string                 `json:"id"`
	Round           models.ROUND           `json:"round"`
	TotalPrize      float64                `json:"total_prize"`
	Points          int                    `json:"points"`
	TotalClassified int                    `json:"total_classified"`
	BestThird       int                    `json:"best_third"`
	Groups          []*GetRoundGroupDTORes `json:"groups"`
}

type GetRoundGroupDTORes struct {
	ID          string                                    `json:"id"`
	Position    int                                       `json:"position"`
	Matches     []*GetRoundWithMatchesMatchDTORes         `json:"matches"`
	Competitors []*GetRoundGroupCompetitorWithStatsDTORes `json:"competitors"`
}

type GetRoundGroupCompetitorWithStatsDTORes struct {
	ID              string                               `json:"id"`
	CurrentPosition *int                                 `json:"current_position"`
	Position        int                                  `json:"position"`
	Stats           TournamentGroupCompetitorStatsDTOReq `json:"stats"`
	Users           []*GetRoundWithMatchesUserDTORes     `json:"users"`
	GuestUsers      []*GetRoundWithMatchesUserDTORes     `json:"guest_users"`
}

type TournamentGroupCompetitorStatsDTOReq struct {
	MatchesPlayed   int   `json:"matches_played"`
	MatchesLost     int   `json:"matches_lost"`
	MatchesWon      int   `json:"matches_won"`
	SetsWon         int   `json:"sets_won"`
	SetsLost        int   `json:"sets_lost"`
	GamesWon        int   `json:"games_won"`
	GamesLost       int   `json:"games_lost"`
	LastFiveMatches []int `json:"last_five_matches"`
}
