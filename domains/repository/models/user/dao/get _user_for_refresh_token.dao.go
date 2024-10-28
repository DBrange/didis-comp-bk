package dao

type GetUserForRefreshTokenDAO struct {
	ID    string   `bson:"_id"`
	Roles []string `bson:"roles"`
}
