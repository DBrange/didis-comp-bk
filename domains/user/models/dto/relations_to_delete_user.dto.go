package dto

type UserRelationsToDeleteDTO struct {
	LocationID string `json:"location_id"`
	PaymentID  string `json:"payments_id"`
	ScheduleID string `json:"schedule_id"`
}
