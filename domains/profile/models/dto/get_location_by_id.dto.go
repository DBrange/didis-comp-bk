package dto

type GetLocationByIDDTORes struct {
	ID      string  `json:"_id"`
	State   *string `json:"state"`
	Country *string `json:"country"`
	City    *string `json:"city"`
	Lat     *string `json:"lat"`
	Long    *string `json:"long"`
}