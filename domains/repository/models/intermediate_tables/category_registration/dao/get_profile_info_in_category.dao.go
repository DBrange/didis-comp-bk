package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetProfileInfoInCategoryDAORes struct {
	ID              *primitive.ObjectID                    `bson:"_id"`
	Points          int                                    `bson:"points"`
	CurrentPosition int                                    `bson:"current_position"`
	Users           []*GetProfileInfoInCategoryUsersDAORes `bson:"users"`
	GuestUsers      []*GetProfileInfoInCategoryUsersDAORes `bson:"guest_users"`
	CompetitorStats *GetProfileInfoInCategoryStatsDAORes   `bson:"competitor_stats"`
}

type GetProfileInfoInCategoryUsersDAORes struct {
	ID        *primitive.ObjectID `bson:"_id"`
	FirstName string              `bson:"first_name"`
	LastName  string              `bson:"last_name"`
	Image     string              `bson:"image"`
}

type GetProfileInfoInCategoryStatsDAORes struct {
	ID             *primitive.ObjectID   `bson:"_id"`
	TotalWins      int                   `bson:"total_wins"`
	TotalLosses    int                   `bson:"total_losses"`
	MoneyEarned    float64               `bson:"money_earned"`
	TournamentsWon []*primitive.ObjectID `bson:"tournaments_won"`
}
