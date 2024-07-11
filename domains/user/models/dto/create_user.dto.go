package dto

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type CreateUserDTOReq struct {
	ID          string         `json:"id"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	Username    *string        `json:"username"`
	Birthdate   *time.Time     `json:"birthdate"`
	Password    *string        `json:"password"`
	Email       string         `json:"email"`
	Phone       *string        `json:"phone"`
	Image       *string        `json:"image"`
	Active      bool           `json:"active"`
	AccessLevel *int16         `json:"access_level"`
	Genre       []models.GENRE `json:"genre" validate:"dive,genre"`
	Role        []models.ROLE  `json:"role"`
	LocationID  *string        `json:"location_id"`
	ScheduleID  *string        `json:"schedule_id"`
	PaymentID   *string        `json:"payment_id"`
}

// func (u *CreateUserDTOReq) SetTimeStamp() {
// 	currentDate := time.Now()
// 	if u.CreatedAt.IsZero() {
// 		u.CreatedAt = currentDate
// 	}
// 	u.UpdatedAt = currentDate
// }

// func (u *CreateUserDTOReq) RenewUpdate() {
// 	u.UpdatedAt = time.Now()
// }
