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
	client                *mongo.Client
	userColl              *mongo.Collection
	locationColl          *mongo.Collection
	availabilityColl      *mongo.Collection
	roleColl              *mongo.Collection
	organizerColl         *mongo.Collection
	leagueColl            *mongo.Collection
	tournamentColl        *mongo.Collection
	potColl               *mongo.Collection
	tournamentGroupColl   *mongo.Collection
	doubleEliminationColl *mongo.Collection
	roundColl             *mongo.Collection
	competitorColl        *mongo.Collection
	competitorStatsColl   *mongo.Collection
	singleColl            *mongo.Collection
	doubleColl            *mongo.Collection
	teamColl              *mongo.Collection
	matchColl             *mongo.Collection
	chatColl              *mongo.Collection
	chatMessageColl       *mongo.Collection
	guestPlayerColl       *mongo.Collection
	notificationColl      *mongo.Collection

	competitorMatchColl        *mongo.Collection // Intermediate table
	competitorUserColl         *mongo.Collection // Intermediate table
	followerColl               *mongo.Collection // Intermediate table
	guestCompetitorColl        *mongo.Collection // Intermediate table
	leagueRegistrationColl     *mongo.Collection // Intermediate table
	opinionColl                *mongo.Collection // Intermediate table
	tournamentRegistrationColl *mongo.Collection // Intermediate table
	userChatColl               *mongo.Collection // Intermediate table
}

func NewRepository(

	client *mongo.Client,
	userColl *mongo.Collection,
	locationColl *mongo.Collection,
	availabilityColl *mongo.Collection,
	roleColl *mongo.Collection,
	organizerColl *mongo.Collection,
	leagueColl *mongo.Collection,
	tournamentColl *mongo.Collection,
	potColl *mongo.Collection,
	tournamentGroupColl *mongo.Collection,
	doubleEliminationColl *mongo.Collection,
	roundColl *mongo.Collection,
	competitorColl *mongo.Collection,
	competitorStatsColl *mongo.Collection,
	singleColl *mongo.Collection,
	doubleColl *mongo.Collection,
	teamColl *mongo.Collection,
	matchColl *mongo.Collection,
	chatColl *mongo.Collection,
	chatMessageColl *mongo.Collection,
	guestPlayerColl *mongo.Collection,
	notificationColl *mongo.Collection,

	competitorMatchColl *mongo.Collection, // Intermediate table
	competitorUserColl *mongo.Collection, // Intermediate table
	followerColl *mongo.Collection, // Intermediate table
	guestCompetitorColl *mongo.Collection, // Intermediate table
	leagueRegistrationColl *mongo.Collection, // Intermediate table
	opinionColl *mongo.Collection, // Intermediate table
	tournamentRegistrationColl *mongo.Collection, // Intermediate table
	userChatColl *mongo.Collection, // Intermediate table

) (*Repository, error) {

	repository := &Repository{
		client:                client,
		userColl:              userColl,
		locationColl:          locationColl,
		availabilityColl:      availabilityColl,
		roleColl:              roleColl,
		organizerColl:         organizerColl,
		leagueColl:            leagueColl,
		tournamentColl:        tournamentColl,
		potColl:               potColl,
		tournamentGroupColl:   tournamentGroupColl,
		doubleEliminationColl: doubleEliminationColl,
		roundColl:             roundColl,
		competitorColl:        competitorColl,
		competitorStatsColl:   competitorStatsColl,
		singleColl:            singleColl,
		doubleColl:            doubleColl,
		teamColl:              teamColl,
		matchColl:             matchColl,
		chatColl:              chatColl,
		chatMessageColl:       chatMessageColl,
		guestPlayerColl:       guestPlayerColl,
		notificationColl:      notificationColl,

		competitorMatchColl:        competitorMatchColl,        // Intermediate table
		competitorUserColl:         competitorUserColl,         // Intermediate table
		followerColl:               followerColl,               // Intermediate table
		guestCompetitorColl:        guestCompetitorColl,        // Intermediate table
		leagueRegistrationColl:     leagueRegistrationColl,     // Intermediate table
		opinionColl:                opinionColl,                // Intermediate table
		tournamentRegistrationColl: tournamentRegistrationColl, // Intermediate table
		userChatColl:               userChatColl,               // Intermediate table
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
		r.userColl: {
			{Keys: bson.D{{Key: "email", Value: 1}}, Options: options.Index().SetUnique(true)},
			{Keys: bson.D{{Key: "username", Value: 1}}, Options: options.Index().SetUnique(true)},
		},
		// r.locationColl: {
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
