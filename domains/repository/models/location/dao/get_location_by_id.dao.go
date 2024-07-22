package dao

type GetLocationByIDDAORes struct {
	ID      string  `bson:"_id"`
	State   *string `bson:"state"`
	Country *string `bson:"country"`
	City    *string `bson:"city"`
	Lat     *string `bson:"lat"`
	Long    *string `bson:"long"`
}
