package dto

type UserRelationsToDeleteDTO struct {
	LocationID string `json:"location_id"`
	PaymentID  string `json:"payments_id"`
	ScheduleID string `json:"schedule_id"`
}

// func (a *UserRelationsToDeleteDTO) LocationID() string {
// 	return a.locationID
// }

// func (a *UserRelationsToDeleteDTO) PaymentID() string {
// 	return a.paymentID
// }

// func (a *UserRelationsToDeleteDTO) ScheduleID() string {
// 	return a.scheduleID
// }
