package dto

import "time"

type CreateUserDTO struct {
	FirstName string     `json:"first_name" bson:"first_name" validate:"required"`
	LastName  string     `json:"last_name" bson:"last_name" validate:"required"`
  CreatedAt time.Time  `json:"created_at" bson:"created_at" validate:"required"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at" validate:"required"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" bson:"deleted_at"`
}

func (u *CreateUserDTO) SetTimeStamp() {
	currentDate := time.Now()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = currentDate
	}
	u.UpdatedAt = currentDate
}

func (u *CreateUserDTO) RenewUpdate() {
	u.UpdatedAt = time.Now()
}
