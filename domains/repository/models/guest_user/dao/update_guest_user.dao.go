package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type UpdateGuestUserDAOReq struct {
	FirstName *string       `bson:"first_nameomitempty,"`
	LastName  *string       `bson:"last_name,omitempty"`
	Email     *string       `bson:"email,omitempty"`
	Image     *string       `bson:"image,omitempty"`
	Genre     *models.GENRE `bson:"genre,omitempty"`
	UpdatedAt time.Time     `bson:"updated_at,omitempty"`
}

func (u *UpdateGuestUserDAOReq) RenewUpdate() {
	u.UpdatedAt = time.Now().UTC()
}
