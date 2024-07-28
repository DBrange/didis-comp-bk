package repository

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/tournament_registration/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) CreateTournamentRegistration(ctx context.Context, tournamentRegistrationInfoDAO *dao.CreateTournamentRegistrationDAOReq) error {
	if err := r.VerifyCompetitorAlreadyResgisteredInTournament(ctx, tournamentRegistrationInfoDAO); err != nil {
		return err
	}

	tournamentRegistrationInfoDAO.SetTimeStamp()

	_, err := r.tournamentRegistrationColl.InsertOne(ctx, tournamentRegistrationInfoDAO)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return fmt.Errorf("%w: error tournamentRegistration scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return fmt.Errorf("error when inserting tournamentRegistration: %w", err)
	}

	return nil
}

func (r *Repository) GetTournamentRegistrationByID(ctx context.Context, tournamentRegistrationID string) (*dao.GetTournamentRegistrationByIDDAORes, error) {
	var tournamentRegistration dao.GetTournamentRegistrationByIDDAORes

	tournamentRegistrationOID, err := r.ConvertToObjectID(tournamentRegistrationID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *tournamentRegistrationOID}

	err = r.tournamentRegistrationColl.FindOne(ctx, filter).Decode(&tournamentRegistration)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for tournamentRegistration: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the tournamentRegistration: %w", err)
	}

	return &tournamentRegistration, nil
}

func (r *Repository) DeleteTournamentRegistration(ctx context.Context, tournamentRegistrationID string) error {
	err := r.SetDeletedAt(ctx, r.tournamentRegistrationColl, tournamentRegistrationID, "tournamentRegistration")
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) VerifyCompetitorAlreadyResgisteredInTournament(ctx context.Context, tournamentRegistrationInfoDAO *dao.CreateTournamentRegistrationDAOReq) error {
	filter := bson.M{
		"tournament_id": tournamentRegistrationInfoDAO.TournamentID,
		"competitor_id": tournamentRegistrationInfoDAO.CompetitorID,
		// "$or": []bson.M{
		// 	{"deleted_at": bson.M{"$exists": false}},
		// 	{"deleted_at": nil},
		// },
	}

	var documentFinded *dao.CreateTournamentRegistrationDAOReq

	err := r.tournamentRegistrationColl.FindOne(ctx, filter).Decode(&documentFinded)
	if err == nil {
		return fmt.Errorf("error relation in tournamentRegistration already exists: %w", err)
	}

	return nil
}

func (r *Repository) GetAllCompetitorInTournament(ctx context.Context, tournamentOID *primitive.ObjectID, limit, page int) ([]dao.GetTournamentRegistrationByIDDAORes, error) {
	// tournamentOID, err := r.ConvertToObjectID(tournamentID)
	// if err != nil {
	// 	return nil, err
	// }

	filter := bson.M{"tournament_id": *tournamentOID}

	skip := (page - 1) * limit

	opts := options.Find().
		SetLimit(int64(limit)).
		SetSkip(int64(skip))

	cursor, err := r.tournamentRegistrationColl.Find(ctx, filter, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for tournamentRegistration: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the tournamentRegistration: %w", err)
	}

	defer cursor.Close(ctx)

	var tournamentRegistrations []dao.GetTournamentRegistrationByIDDAORes
	
	err = cursor.All(ctx, &tournamentRegistrations)
	if err != nil {
		return nil, fmt.Errorf("error when decoding tournamentRegistrations: %w", err)
	}

	if len(tournamentRegistrations) == 0 {
		return nil, fmt.Errorf("%w: no tournament registrations found for tournament ID: %s on page %d", customerrors.ErrNotFound, tournamentOID.Hex(), page)
	}

	return tournamentRegistrations, nil
}
