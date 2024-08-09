package dto

type CreateDoubleEliminationDTOReq struct {
	Matches []string `bson:"matches"`
	Rounds  []string `bson:"rounds"`
}
