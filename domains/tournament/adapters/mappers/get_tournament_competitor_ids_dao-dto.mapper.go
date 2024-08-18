package mappers

import "go.mongodb.org/mongo-driver/bson/primitive"

func GetTournamentCompetitorIDsDAOtoDTO(competitorOIDs []*primitive.ObjectID) []string {
	competitorsIDs := make([]string, len(competitorOIDs))

	for i, competitorOID := range competitorOIDs {
		competitorsIDs[i] = competitorOID.Hex()
	}

	return competitorsIDs
}
