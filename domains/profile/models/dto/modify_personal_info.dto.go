package dto

import "reflect"

type ModifyPersonalInfoDTOReq struct {
	FirstName *string                           `json:"first_name,omitempty" validate:"omitempty,min=2"`
	LastName  *string                           `json:"last_name,omitempty" validate:"omitempty,min=2"`
	Username  *string                           `json:"username,omitempty" validate:"omitempty,min=3"`
	Phone     *string                           `json:"phone,omitempty" validate:"omitempty,e164"`
	Birthdate     *string                           `json:"birthdate,omitempty" validate:"omitempty,date"`
	Image     *string                           `json:"image,omitempty" validate:"omitempty,url"`
	Location  *ModifyPersonalInfoLocationDTOReq `json:"location,omitempty"`
}

func (u *ModifyPersonalInfoDTOReq) AreAllFieldsNil() bool {
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

type ModifyPersonalInfoLocationDTOReq struct {
	ID      string  `json:"id"`
	State   *string `json:"state,omitempty" validate:"omitempty,min=2"`
	Country *string `json:"country,omitempty" validate:"omitempty,min=2"`
	City    *string `json:"city,omitempty" validate:"omitempty,min=2"`
	Lat     *string `json:"lat,omitempty"`
	Long    *string `json:"long,omitempty"`
}
