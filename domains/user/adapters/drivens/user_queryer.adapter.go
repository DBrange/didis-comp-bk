package adapters

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/cmd/api/assets"
	ports "github.com/DBrange/didis-comp-bk/domains/repository/ports/drivers"
	"github.com/DBrange/didis-comp-bk/domains/user/adapters/mappers"
	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserQueryerAdapter struct {
	drivers ports.ForManagingUser
}

func NewUserQueryerAdapter(drivers ports.ForManagingUser) *UserQueryerAdapter {
	return &UserQueryerAdapter{
		drivers: drivers,
	}
}

func (a *UserQueryerAdapter) CreateUser(ctx context.Context, userDTO *user_dto.CreateUserDTOReq) error {
	mappedUser := mappers.CreateUserDTOReqtoDAO(userDTO)

	err := a.drivers.CreateUser(ctx, mappedUser)
	if err != nil {
		return err
	}

	return nil
}

func (a *UserQueryerAdapter) GetUserByID(ctx context.Context, id string) (*user_dto.GetUserByIDDTO, error) {
	userDTO, err := a.drivers.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	mappedUser := mappers.GetUserByIDDAOtoDTO(userDTO)

	return mappedUser, nil
}

func (a *UserQueryerAdapter) UpdateUser(ctx context.Context, userID string, newUser *user_dto.UpdateUserDTOReq) error {
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("invalid id format: %w", err)
	}
	
	filter := bson.M{"_id": oid}
	update, err := assets.StructToBsonMap(newUser)
	if err != nil {
		return err
	}

	return a.drivers.UpdateUser(ctx, filter, update)
}
