package dto

import (
	"reflect"
)

type UpdateUserDTOReq struct {
	// ID         string    `json:"id,omitempty" validate:"omitempty"`
	FirstName  *string    `json:"first_name,omitempty" validate:"omitempty,min=2"`
	LastName   *string    `json:"last_name,omitempty" validate:"omitempty,min=2"`
	Username   *string    `json:"username,omitempty" validate:"omitempty,min=3"`
	Phone      *string    `json:"phone,omitempty" validate:"omitempty,e164"`
	Image      *string    `json:"image,omitempty" validate:"omitempty,url"`
}

func (u *UpdateUserDTOReq) AreAllFieldsNil() bool {
	v := reflect.ValueOf(u).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		// Verificamos si el campo es nil
		if !fieldValue.IsNil() {
			return false
		}
	}

	return true
}
