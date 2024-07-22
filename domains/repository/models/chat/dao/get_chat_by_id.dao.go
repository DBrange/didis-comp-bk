package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetChatByIDDAORes struct {
	ID           primitive.ObjectID   `bson:"_id"`
	Name         string               `bson:"name"`
	ChatType     models.CHAT          `bson:"chat_type"`
	Status       models.CHAT_STATUS   `bson:"status"`
	MatchID      primitive.ObjectID   `bson:"organizer_id"`
	Participants []primitive.ObjectID `bson:"participants"`
	CreatedAt    time.Time            `bson:"created_at"`
	UpdatedAt    time.Time            `bson:"updated_at"`
	DeletedAt    *time.Time           `bson:"deleted_at,omitempty"`
}
