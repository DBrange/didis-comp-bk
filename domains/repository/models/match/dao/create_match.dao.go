package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateMatchDAOReq struct {
	Sport        models.SPORT        `bson:"sport"`
	RoundID      *primitive.ObjectID `bson:"round_id"`
	Result       string              `bson:"result"`
	Winner       *primitive.ObjectID `bson:"winner"`
	TournamentID *primitive.ObjectID `bson:"tournament_id"`
	Position     int                 `bson:"position"`
	Date         *time.Time           `bson:"date"`
	// Votes        map[string]string  `bson:"votes"`
	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}

func (u *CreateMatchDAOReq) SetTimeStamp() {
	currentDate := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = currentDate
	}
	u.UpdatedAt = currentDate
}

func (u *CreateMatchDAOReq) RenewUpdate() {
	u.UpdatedAt = time.Now().UTC()
}
