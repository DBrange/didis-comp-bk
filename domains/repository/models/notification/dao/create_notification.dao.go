package dao

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateNotificationDAOReq struct {
	Receiber primitive.ObjectID           `bson:"receiber"`
	Priority models.NOTIFICATION_PRIORITY `bson:"priority"`
	Type     models.NOTIFICATION_TYPE     `bson:"type"`
	State    models.NOTIFICATION_STATE    `bson:"state"`
common.CreateBaseDAO `bson:",inline"`
}
