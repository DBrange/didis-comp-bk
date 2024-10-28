package dto

import "time"

type GetUserFollowersDTORes struct {
	LastCreatedAt *time.Time                    `json:"last_created_at"`
	Followers     []*GetUserFollowersUserDTORes `json:"followers"`
	Total int `json:"total"`
}

type GetUserFollowersUserDTORes struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Image     string `json:"image"`
	Username  string `json:"username"`
}
