package models

type OrganizeTournamentOptions struct {
	// DoubleElimination bool `json:"double_elimination"`
	QuantityPots      int  `json:"quantity_pots"`
	QuantityGroups    int  `json:"quantity_groups"`
	// Classify          int  `json:"classify"`
	// BestThird         int  `json:"best_third"`
}
