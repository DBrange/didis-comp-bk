package dto

type ProfileRelationsToDeleteDTO struct {
	LocationID     string `json:"location_id"`
	PaymentID      string `json:"payments_id"`
	AvailabilityID string `json:"availability_id"`
}
