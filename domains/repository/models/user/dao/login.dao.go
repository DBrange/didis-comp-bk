package dao

type LoginDAOReq struct {
	Username string `bson:"username"`
	Password string  `bson:"password"`
}
