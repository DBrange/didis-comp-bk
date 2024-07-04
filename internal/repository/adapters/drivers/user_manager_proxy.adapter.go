package adapters

import (
	"context"

	repo_dto "github.com/DBrange/didis-comp-bk/internal/repository/models/user/dto"
	"github.com/DBrange/didis-comp-bk/internal/repository/repository"
)

type UserMangerProxyAdapter struct {
	ctx        context.Context
	repository *repository.Repository
}

func NewUserMangerProxyAdapter(ctx context.Context, repository *repository.Repository) *UserMangerProxyAdapter {
	return &UserMangerProxyAdapter{
		ctx:        ctx,
		repository: repository,
	}
}

func (a *UserMangerProxyAdapter) CreateUser(user *repo_dto.CreateUserDTO) error {
	err := a.repository.CreateUser(a.ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (a *UserMangerProxyAdapter) GetUserByID(ctx context.Context, id string) (*repo_dto.GetUserByIDDTO, error) {
	return a.repository.GetUserByID(ctx, id)
}
