package repository

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) CreateLocation(ctx context.Context, user *dao.CreateLocationDAOReq) (string, error) {

	user.SetTimeStamp()

	result, err := r.location_coll.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("%w: error inserting location: %s", customerrors.ErrLocationInsertionFailed, err.Error())
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}
