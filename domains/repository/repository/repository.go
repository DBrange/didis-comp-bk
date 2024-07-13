package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	user_coll     *mongo.Collection
	location_coll *mongo.Collection
}

func NewRepository( user_coll *mongo.Collection, location_coll *mongo.Collection) (*Repository, error) {
	repository := &Repository{
		user_coll:     user_coll,
		location_coll: location_coll,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := repository.EnsureIndexes(ctx); err != nil {
		return nil, err
	}

	return repository, nil
}

func (r *Repository) EnsureIndexes(ctx context.Context) error {
	collections := map[*mongo.Collection][]mongo.IndexModel{
		r.user_coll: {
			{Keys: bson.D{{Key: "email", Value: 1}}, Options: options.Index().SetUnique(true)},
			{Keys: bson.D{{Key: "username", Value: 1}}, Options: options.Index().SetUnique(true)},
		},
		// r.location_coll: {
		//     {Keys: bson.D{{Key: "unique_field", Value: 1}}, Options: options.Index().SetUnique(true)},
		// },
	}

	for coll, indexes := range collections {
		if err := r.createIndexes(ctx, coll, indexes); err != nil {
			return fmt.Errorf("failed to create indexes for collection: %w", err)
		}
	}

	return nil
}

func (r *Repository) createIndexes(ctx context.Context, coll *mongo.Collection, indexes []mongo.IndexModel) error {
	_, err := coll.Indexes().CreateMany(ctx, indexes)
	if err != nil {
		return fmt.Errorf("failed to create indexes: %w", err)
	}
	return nil
}
