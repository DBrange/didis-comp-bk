package dto

type CreateLocationDTO struct {
	State   *string `json:"state" validate:"omitempty,min=2"`
	Country *string `json:"country" validate:"omitempty,min=2"`
	City    *string `json:"city" validate:"omitempty,min=2"`
	Lat     *string `json:"lat" validate:"omitempty,min=2"`
	Long    *string `json:"long" validate:"omitempty,min=2"`
}
