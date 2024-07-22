package dao

import "time"

// "reflect"

type UpdateUserDAOReq struct {
	// ID         string    `bson:"_id"`
	FirstName *string   `bson:"first_name,omitempty"`
	LastName  *string   `bson:"last_name,omitempty"`
	Username  *string   `bson:"username,omitempty"`
	Phone     *string   `bson:"phone,omitempty"`
	Image     *string   `bson:"image,omitempty"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func (u *UpdateUserDAOReq) RenewUpdate() {
	u.UpdatedAt = time.Now().UTC()
}

// func (u *UpdateUserDAOReq) AreAllFieldsNil() bool {
// 	v := reflect.ValueOf(u).Elem()
// 	for i := 0; i < v.NumField(); i++ {
// 		fieldValue := v.Field(i)
// 		// Verificamos si el campo es nil
// 		if !fieldValue.IsNil() {
// 			return false
// 		}
// 	}

// 	return true
// }
