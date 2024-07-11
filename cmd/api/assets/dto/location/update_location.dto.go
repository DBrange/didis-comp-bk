package dto

type UpdateLocationDTO struct {
	ID      string  `json:"id,omitempty" validate:"required"`
	State   *string `json:"state,omitempty" validate:"omitempty,min=2"`
	Country *string `json:"country,omitempty" validate:"omitempty,min=2"`
	City    *string `json:"city,omitempty" validate:"omitempty,min=2"`
	Lat     *string `json:"lat,omitempty" validate:"omitempty,min=2"`
	Long    *string `json:"long,omitempty" validate:"omitempty,min=2"`
}
