package dto

type EndTournamentDTORes struct {
	ChampionCompetitorID           string `bson:"champion_competitor_id"`
	DoubleElimChampionCompetitorID string `bson:"double_elimination_champion_competitor_id"`
	DoubleElimID                   string `bson:"double_elimination_id"`
}
