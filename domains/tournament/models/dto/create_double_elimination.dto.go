package dto

type CreateDoubleEliminationDTOReq struct {
	Matches []string `bson:"matches"`
	Rounds  []string `bson:"rounds"`
	TotalPrize float64               `bson:"total_prize"`
	Points     int                   `bson:"points"`
}