package dao

import "time"

type GetUserFollowersDAORes struct {
	LastCreatedAt *time.Time                         `bson:"last_created_at"`
	Followers     []*GetUserCompetitorFollowedDAORes `bson:"followers"`
	Total int `bson:"total"`
}
