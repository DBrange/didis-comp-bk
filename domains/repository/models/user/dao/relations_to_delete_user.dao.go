package dao

type UserRelationsToDeleteDAO struct {
	LocationID string `bson:"location_id"`
	PaymentID  string `bson:"payments_id"`
	ScheduleID string `bson:"schedule_id"`
}

// func (a *UserRelationsToDeleteDAO) LocationID() string {
// 	return a.locationID
// }

// func (a *UserRelationsToDeleteDAO) PaymentID() string {
// 	return a.paymentID
// }

// func (a *UserRelationsToDeleteDAO) ScheduleID() string {
// 	return a.scheduleID
// }
