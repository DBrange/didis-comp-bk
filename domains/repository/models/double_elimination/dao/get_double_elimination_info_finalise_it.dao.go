package dao

type GetDoubleElimInfoToFinaliseItDAORes struct {
	TotalPrize float64 `bson:"total_prize"`
	Points     int     `bson:"points"`
}
