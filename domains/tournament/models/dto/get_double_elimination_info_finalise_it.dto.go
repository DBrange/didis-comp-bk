package dto

type GetDoubleElimInfoToFinaliseItDTORes struct {
	TotalPrize float64 `bson:"total_prize"`
	Points     int     `bson:"points"`
}
