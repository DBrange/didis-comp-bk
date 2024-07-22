package repository

import (
	"context"
	"fmt"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) withTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) error) error {
	session, err := r.client.StartSession()
	if err != nil {
		return fmt.Errorf("%w: failed to start session: %s", customerrors.ErrStartSessionFailed, err.Error())
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		return nil, fn(sessCtx)
	})

	if err != nil {
		return fmt.Errorf("%w: transaction failed: %s", customerrors.ErrTransaction, err.Error())
	}

	return nil
}

func (r *Repository) ConvertToObjectID(ID string) (*primitive.ObjectID, error) {
	if ID == "" {
		return nil, fmt.Errorf("invalid availability id format: %w", customerrors.ErrInvalidID)
	}
	OID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, fmt.Errorf("%w: invalid availability id format: %s", customerrors.ErrInvalidID, err.Error())
	}

	return &OID, nil
}
