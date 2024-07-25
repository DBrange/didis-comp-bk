package repository

import (
	"context"
	"fmt"
	"sync"
	"time"

	enum_models "github.com/DBrange/didis-comp-bk/cmd/api/models"
	models "github.com/DBrange/didis-comp-bk/domains/repository/models/role"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/role/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) InitialiseRole(ctx context.Context) error {

	currentDate := time.Now().UTC()

	allRoles := []any{
		dao.RoleDAOReq{
			Name:      enum_models.ROLE_FREE,
			RoleType:  enum_models.ROLE_TYPE_USER,
			CreatedAt: currentDate,
			UpdatedAt: currentDate,
		},
		dao.RoleDAOReq{
			Name:      enum_models.ROLE_BASIC,
			RoleType:  enum_models.ROLE_TYPE_USER,
			CreatedAt: currentDate,
			UpdatedAt: currentDate,
		},
		dao.RoleDAOReq{
			Name:      enum_models.ROLE_ADMIN,
			RoleType:  enum_models.ROLE_TYPE_USER,
			CreatedAt: currentDate,
			UpdatedAt: currentDate,
		},
	}

	fmt.Printf("%v", allRoles)

	opts := options.InsertMany().SetOrdered(true)
	_, err := r.roleColl.InsertMany(ctx, allRoles, opts)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return fmt.Errorf("%w: error user scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}
		return fmt.Errorf("error when inserting user: %w", err)
	}

	return nil
}

func (r *Repository) GetRoleByNameAndType(ctx context.Context, roleName, roleType string) (*models.Role, error) {
	var role models.Role

	filter := bson.M{"name": roleName, "role_type": roleType}

	err := r.roleColl.FindOne(ctx, filter).Decode(&role)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for the role: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the role: %w", err)
	}

	return &role, nil
}

func (r *Repository) getRoleByNameAndTypeConcurrently(sessCtx mongo.SessionContext, roleName string, roleType string, wg *sync.WaitGroup, roleCh chan<- *roleResult) {
	defer wg.Done()
	role, err := r.GetRoleByNameAndType(sessCtx, roleName, roleType)
	roleCh <- &roleResult{Role: role, Err: err}
}
