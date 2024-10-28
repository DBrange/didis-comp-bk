package dao

import "go.mongodb.org/mongo-driver/bson/primitive"
type GetCompetitorsInTournamentDAORes struct {
	CompetitorID    *primitive.ObjectID                     `bson:"_id"`
	CurrentPosition *int                                    `bson:"current_position"`
	Users           []*GetCompetitorsInTournamentUserDAORes `bson:"users"`
	GuestUsers      []*GetCompetitorsInTournamentUserDAORes `bson:"guest_users"`
}

type GetCompetitorsInTournamentUserDAORes struct {
	ID        primitive.ObjectID `bson:"_id"`
	FirstName string              `bson:"first_name"`
	LastName  string              `bson:"last_name"`
	Image     string              `bson:"image"`
	Username     *string              `bson:"username"`
}
