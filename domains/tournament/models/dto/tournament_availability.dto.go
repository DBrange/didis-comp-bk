package dto

type TournamentAvailabilityDTO struct {
	AvailableCourts int `json:"available_courts"`
	AverageHours    int `json:"average_hours"`
}
