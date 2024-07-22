package repository

import (
	"context"
	"fmt"

	api_assets "github.com/DBrange/didis-comp-bk/cmd/api/assets"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) CreateLocation(ctx context.Context, locationInfoDAO *dao.CreateLocationDAOReq) (string, error) {
	locationInfoDAO.SetTimeStamp()

	result, err := r.locationColl.InsertOne(ctx, locationInfoDAO)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("%w: error duplicate key for location: %s", customerrors.ErrDuplicateKey, err.Error())
		}

		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return "", fmt.Errorf("%w: error location scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return "", fmt.Errorf("error when inserting location: %w", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *Repository) GetLocationByID(ctx context.Context, locationID string) (*dao.GetLocationByIDDAORes, error) {
	var location dao.GetLocationByIDDAORes

	locationOID, err := r.ConvertToObjectID(locationID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *locationOID}

	err = r.locationColl.FindOne(ctx, filter).Decode(&location)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for location: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the location: %w", err)
	}

	return &location, nil
}

func (r *Repository) UpdateLocation(ctx context.Context, locationID string, locationInfoDAO *dao.UpdateLocationDAOReq) error {
	locationOID, err := r.ConvertToObjectID(locationID)
	if err != nil {
		return err
	}

	locationInfoDAO.RenewUpdate()

	filter := bson.M{"_id": *locationOID}
	update, err := api_assets.StructToBsonMap(locationInfoDAO)
	if err != nil {
		return err
	}

	result, err := r.locationColl.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": update},
	)
	if err != nil {
		return fmt.Errorf("error updating location: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: no location found with id: %s", customerrors.ErrNotFound, locationID)
	}

	return nil
}

func (r *Repository) DeleteLocation(ctx context.Context, locationID string) error {
	err := r.setDeletedAt(ctx, r.locationColl, locationID, "location")
	if err != nil {
		return err
	}

	return nil
}

// para ver como se haria si solo quisiera datos espectificos, y ver por log, como es la respuesta
// func (r *Repository) GetLocationByID(ctx context.Context, id string) (*dao.GetLocationByIDDAORes, error) {
//     oid, err := primitive.ObjectIDFromHex(id)
//     if err != nil {
//         return nil, fmt.Errorf("%w: error when searching for location: %s", customerrors.ErrInvalidID, err.Error())
//     }

//     filter := bson.M{"_id": oid}

//     projection := bson.M{
//         "state": 1,
//         "city": 1,
//     }

//     opts := options.FindOne().SetProjection(projection)

//     // Usar bson.M para recibir los datos crudos
//     var rawResult bson.M
//     err = r.locationColl.FindOne(ctx, filter, opts).Decode(&rawResult)
//     if err != nil {
//         if err == mongo.ErrNoDocuments {
//             return nil, fmt.Errorf("%w: error when searching for location: %s", customerrors.ErrNotFound, err.Error())
//         }
//         return nil, fmt.Errorf("error when searching for location: %w", err)
//     }

//     // Imprimir o loguear los datos crudos
//     fmt.Printf("Datos crudos de MongoDB: %+v\n", rawResult)

//     // Ahora puedes decodificar los datos crudos en tu estructura
//     var location dao.GetLocationByIDDAORes
//     bsonBytes, _ := bson.Marshal(rawResult)
//     bson.Unmarshal(bsonBytes, &location)

//     return &location, nil
// }
