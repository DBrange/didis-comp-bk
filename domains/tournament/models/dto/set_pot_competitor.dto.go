package dto

type SetPotCompetitorDTOReq struct {
	PotID       string   `json:"pot_id"`
	Competitors []string `json:"competitors"`
}
