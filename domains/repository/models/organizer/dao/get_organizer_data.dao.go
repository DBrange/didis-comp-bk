package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetOrganizerDataDAORes struct {
	ID                     *primitive.ObjectID `bson:"_id"`
	AverageScore           int                 `bson:"average_score"`
	AverageTournamentScore int                 `bson:"average_tournament_score"`
	TotalCategories        int                 `bson:"total_categories"`
	TotalTournaments       int                 `bson:"total_tournaments"`
}
