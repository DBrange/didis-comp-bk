package repository

import (
	"context"
	"fmt"

	repo_dto "github.com/DBrange/didis-comp-bk/internal/repository/models/user/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateUser(ctx context.Context, user *repo_dto.CreateUserDTO) error {
	_, err := r.user_coll.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("%w: duplicate key error for user: %s", customerrors.ErrUserDuplicateKey, err.Error()) // hay que retornar esto a routes
	}
	return nil
}

func (r *Repository) GetUserByID(ctx context.Context, id string) (*repo_dto.GetUserByIDDTO, error) {
	var user repo_dto.GetUserByIDDTO

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("%w: error when searching for user: %s", customerrors.ErrUserInvalidID, err.Error())
	}

	filter := bson.M{"_id": oid}

	err = r.user_coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for user: %s", customerrors.ErrUserNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for user: %w", err)
	}

	return &user, nil
}
