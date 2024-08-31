package dto

import "time"

type CompetitorsInMatchDTO struct {
	ID            string   `json:"_id"`
	CompetitorIDs []string `json:"competitor_ids"`
}

type CourtsDTO struct {
	AvailableCourts     int                               `json:"available_courts"`
	DailyAvailabilities []*GetDailyAvailabilityByIDDTORes `json:"daily_availabilities"`
}

type MatchDateDTOReq struct {
	ID   string    `json:"_id"`
	Date *time.Time `json:"date"`
}
