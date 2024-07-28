package dto

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type GetUserByIDDTORes struct {
	ID         string       `json:"_id"`
	FirstName  string       `json:"first_name"`
	LastName   string       `json:"last_name"`
	Username   *string      `json:"username"`
	Birthdate  *time.Time   `json:"birthdate"`
	Password   *string      `json:"password"`
	Email      string       `json:"email"`
	Phone      *string      `json:"phone"`
	Image      *string      `json:"image"`
	Genre      models.GENRE `json:"genre"`
	LocationID string      `json:"location_id"`
}
