package mappers

import "go.mongodb.org/mongo-driver/bson/primitive"

func FindCompetitorsOutCategoryDAOtoDTO(competitorOIDs []*primitive.ObjectID, categoryOID *primitive.ObjectID) ([]string, string) {
	competitorIDs := make([]string, len(competitorOIDs))

	for i, competitorOID := range competitorOIDs {
		competitorIDs[i] = competitorOID.Hex()
	}

	return competitorIDs, categoryOID.Hex()
}
