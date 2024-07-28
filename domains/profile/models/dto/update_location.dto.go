package dto

type UpdateLocationDTOReq struct {
	ID      string  `json:"id"`
	State   *string `json:"state,omitempty"`
	Country *string `json:"country,omitempty"`
	City    *string `json:"city,omitempty"`
	Lat     *string `json:"lat,omitempty"`
	Long    *string `json:"long,omitempty"`
}
