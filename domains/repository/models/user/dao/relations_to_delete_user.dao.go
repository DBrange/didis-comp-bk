package dao

type UserRelationsToDeleteDAOReq struct {
	LocationID     string `bson:"location_id,omitempty"`
}

// func (a *UserRelationsToDeleteDAO) LocationID() string {
// 	return a.locationID
// }

// func (a *UserRelationsToDeleteDAO) PaymentID() string {
// 	return a.paymentID
// }

// func (a *UserRelationsToDeleteDAO) AvailabilityID() string {
// 	return a.availabilityID
// }
