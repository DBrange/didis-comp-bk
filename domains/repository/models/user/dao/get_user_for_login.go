package dao

type GetUserForLoginDAO struct {
	ID        string   `bson:"_id"`
	Password  string   `bson:"password"`
	FirstName string   `bson:"first_name"`
	LastName  string   `bson:"last_name"`
	Username  string   `bson:"username"`
	Image     string   `bson:"image"`
	Roles     []string `bson:"roles"`
}
