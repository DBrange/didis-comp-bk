package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/follower/dao"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) CreateFollower(ctx context.Context, followerDAO *dao.CreateFollowerDAOReq) error {
	followerDAO.SetTimeStamp()

	_, err := r.followerColl.InsertOne(ctx, followerDAO)
	if err != nil {
		if writeErr, ok := err.(mongo.WriteException); ok {
			for _, we := range writeErr.WriteErrors {
				if we.Code == 14 {
					return fmt.Errorf("%w: error follower scheme type: %s", customerrors.ErrSchemaViolation, err.Error())
				}
			}
		}

		return fmt.Errorf("error when inserting follower: %w", err)
	}

	return nil
}

func (r *Repository) VerifyFollowerExistsRelation(ctx context.Context, followerDAO *dao.CreateFollowerDAOReq) error {
	filter := bson.M{"from": followerDAO.From}

	if followerDAO.ToUser != nil {
		filter["to_user"] = followerDAO.ToUser
	} else {
		filter["to_competitor"] = followerDAO.ToCompetitor
	}

	projection := bson.M{"_id": 1}

	opts := options.FindOne().SetProjection(projection)

	var documentFinded struct{}

	if err := r.followerColl.FindOne(ctx, filter, opts).Decode(&documentFinded); err == nil {
		return fmt.Errorf("%w: follower already exists", customerrors.ErrAuthorizationHeader)
	} else if err != mongo.ErrNoDocuments {
		return fmt.Errorf("error when checking for existing follower: %w", err)
	}

	return nil
}

func (r *Repository) FollowOrUnfollow(ctx context.Context, followerDAO *dao.CreateFollowerDAOReq) error {
	filter := bson.M{"from": followerDAO.From, "deleted_at": bson.M{"$exists": false}}

	// Agregar filtro por to_user o to_competitor según corresponda
	if followerDAO.ToUser != nil {
		filter["to_user"] = followerDAO.ToUser
	} else {
		filter["to_competitor"] = followerDAO.ToCompetitor
	}

	projection := bson.M{"_id": 1, "deleted_at": 1}

	opts := options.FindOne().SetProjection(projection)

	var documentFound struct {
		ID        primitive.ObjectID `bson:"_id"`
		DeletedAt *time.Time         `bson:"deleted_at,omitempty"`
	}

	err := r.followerColl.FindOne(ctx, filter, opts).Decode(&documentFound)

	if err == nil {
		// Si el documento existe y no tiene deleted_at, eliminamos al seguidor
		if documentFound.DeletedAt == nil {
			// Actualizar el documento, agregando deleted_at para "eliminar" el seguidor
			return r.SetDeletedAt(ctx, r.followerColl, documentFound.ID.Hex(), "follower")
		} else {
			// Si el documento tiene deleted_at (relación "borrada"), creamos un nuevo seguidor
			return r.CreateFollower(ctx, followerDAO)
		}
	} else if err == mongo.ErrNoDocuments {
		// Si no se encuentra ningún documento, creamos un nuevo seguidor
		return r.CreateFollower(ctx, followerDAO)
	} else {
		// En caso de otro tipo de error
		return fmt.Errorf("error when checking for existing follower: %w", err)
	}
}


func (r *Repository) GetFollowerByID(ctx context.Context, followerID string) (*dao.GetFollowerByIDDAORes, error) {
	var follower dao.GetFollowerByIDDAORes

	followerOID, err := r.ConvertToObjectID(followerID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": *followerOID}

	err = r.followerColl.FindOne(ctx, filter).Decode(&follower)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for follower: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the follower: %w", err)
	}

	return &follower, nil
}

func (r *Repository) DeleteFollower(ctx context.Context, followerID string) error {
	err := r.SetDeletedAt(ctx, r.followerColl, followerID, "follower")
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetCompetitorsFollowed(ctx context.Context, userOID *primitive.ObjectID, name string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) ([]*dao.GetCompetitorFollowedDAORes, error) {
	var followers []*dao.GetCompetitorFollowedDAORes

	if name == "" {
		return followers, nil
	}

	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{
			"from": userOID,
		}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "users",
			"localField":   "to_user",
			"foreignField": "_id",
			"as":           "user",
		}}},
		bson.D{{Key: "$unwind", Value: "$user"}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "competitor_users",
			"localField":   "user._id",
			"foreignField": "user_id",
			"as":           "competitor_user",
		}}},
		bson.D{{Key: "$unwind", Value: "$competitor_user"}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "competitor_users",
			"localField":   "user._id",
			"foreignField": "user_id",
			"as":           "competitor_user",
		}}},
		bson.D{{Key: "$unwind", Value: "$competitor_user"}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "competitors",
			"localField":   "competitor_user.competitor_id",
			"foreignField": "_id",
			"as":           "competitor",
		}}},
		bson.D{{Key: "$unwind", Value: "$competitor"}},
		bson.D{{Key: "$match", Value: bson.M{
			"competitor.sport": sport,
		}}},
		bson.D{{Key: "$limit", Value: 10}},
		bson.D{{Key: "$group", Value: bson.M{
			"_id": "$competitor_user.competitor_id",
		}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "competitor_users",
			"localField":   "_id",
			"foreignField": "competitor_id",
			"as":           "all_competitor_users",
		}}},
		bson.D{{Key: "$unwind", Value: "$all_competitor_users"}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "users",
			"localField":   "all_competitor_users.user_id",
			"foreignField": "_id",
			"as":           "all_users",
		}}},
		bson.D{{Key: "$unwind", Value: "$all_users"}},
		bson.D{{Key: "$group", Value: bson.M{
			"_id": "$_id",
			"users": bson.M{
				"$push": bson.M{
					"_id":        "$all_users._id",
					"first_name": "$all_users.first_name",
					"last_name":  "$all_users.last_name",
					"image":      "$all_users.image",
				},
			},
		}}},
		bson.D{{Key: "$project", Value: bson.M{
			"_id":   "$_id",
			"users": 1,
		}}},
	}

	pipeline = r.getParticipantsOfCategoryNameFilter(pipeline, name,true)
	pipeline = r.singlesOrDoublesCategoryFilter(pipeline, competitorType)

	cursor, err := r.followerColl.Aggregate(ctx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("%w: error when searching for follower: %s", customerrors.ErrNotFound, err.Error())
		}
		return nil, fmt.Errorf("error when searching for the follower: %w", err)
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &followers); err != nil {
		return nil, fmt.Errorf("error when decoding follower: %w", err)
	}
	for _, f := range followers {

		fmt.Printf("asas %+v", f)
	}
	return followers, nil
}

func (r *Repository) singlesOrDoublesFollowersFilter(pipeline mongo.Pipeline, competitorType models.COMPETITOR_TYPE) mongo.Pipeline {
if competitorType == models.COMPETITOR_TYPE_SINGLE {
		// Solo 1 user y ningún guest_user
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.M{
			"users":       bson.M{"$size": 1},
			"guest_users": bson.M{"$size": 0}, // Asegura que no haya guest_users
		}}})
	} else {
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.M{
			"$or": []bson.M{
				{"users": bson.M{"$size": 2}},              // 2 usuarios
				{"guest_users": bson.M{"$size": 2}},        // 2 guest_users
				{"$and": []bson.M{                         // 1 usuario y 1 guest_user
					{"users": bson.M{"$size": 1}},
					{"guest_users": bson.M{"$size": 1}},
				}},
			},
		}}})
	}

	return pipeline
}


func (r *Repository) GetNumberFollowers(ctx context.Context, userOID *primitive.ObjectID) (int, error) {
	// Definir el filtro para que el campo "to_user" coincida con el "userOID"
	filter := bson.M{
		"to_user":    userOID,
		"deleted_at": bson.M{"$exists": false}, // Ignorar los registros eliminados
	}

	// Contar los documentos que coinciden con el filtro
	count, err := r.followerColl.CountDocuments(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("error when counting followers: %w", err)
	}

	// Retornar el conteo
	return int(count), nil
}

func (r *Repository) GetUserFollowers(
	ctx context.Context,
	userOID *primitive.ObjectID,
	name string,
	limit int,
	lastCreatedAt *time.Time,
) (*dao.GetUserFollowersDAORes, error) {
	// Etapa de $match para filtrar seguidores del usuario
	matchStage := bson.D{{Key: "$match", Value: bson.M{
		"to_user":    userOID,
		"deleted_at": bson.M{"$exists": false},
	}}}
	if lastCreatedAt != nil {
		matchStage = bson.D{{Key: "$match", Value: bson.M{
			"to_user":    userOID,
			"created_at": bson.M{"$lt": lastCreatedAt},
		}}}
	}

	// Pipeline para obtener seguidores paginados
	pipeline := mongo.Pipeline{
		matchStage,
		bson.D{{Key: "$sort", Value: bson.M{"created_at": -1}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "users",
			"localField":   "from",
			"foreignField": "_id",
			"as":           "users",
		}}},
		bson.D{{Key: "$unwind", Value: "$users"}},
		bson.D{{Key: "$project", Value: bson.M{
			"users._id":        1,
			"users.first_name": 1,
			"users.last_name":  1,
			"users.image":      1,
			"users.username":   1,
			"created_at":       "$created_at",
		}}},
	}

	// Aplicar filtro de nombre si se proporciona
	if name != "" {
		pipeline = r.getParticipantsOfCategoryNameFilter(pipeline, name, false)
	}

	// Clonar el pipeline para el conteo total antes de aplicar el límite
	countPipeline := make(mongo.Pipeline, len(pipeline))
	copy(countPipeline, pipeline)

	// Agregar etapa de límite para la paginación
	pipeline = append(pipeline, bson.D{{Key: "$limit", Value: limit}})

	// Ejecutar la agregación para obtener seguidores paginados
	cursor, err := r.followerColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error when searching for followers: %w", err)
	}
	defer cursor.Close(ctx)

	var results []struct {
		User      *dao.GetUserCompetitorFollowedDAORes `bson:"users"`
		CreatedAt *time.Time                           `bson:"created_at"`
	}

	if err = cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("error when decoding followers: %w", err)
	}

	// Agregar etapa de conteo al pipeline de conteo
	countPipeline = append(countPipeline, bson.D{{Key: "$count", Value: "total"}})

	// Ejecutar la agregación para obtener el total filtrado
	countCursor, err := r.followerColl.Aggregate(ctx, countPipeline)
	if err != nil {
		return nil, fmt.Errorf("error when counting total followers: %w", err)
	}
	defer countCursor.Close(ctx)

	var totalResult []struct {
		Total int `bson:"total"`
	}
	if err = countCursor.All(ctx, &totalResult); err != nil {
		return nil, fmt.Errorf("error when decoding total followers: %w", err)
	}

	total := 0
	if len(totalResult) > 0 {
		total = totalResult[0].Total
	}

	if len(results) == 0 {
		return &dao.GetUserFollowersDAORes{
			LastCreatedAt: nil,
			Followers:     []*dao.GetUserCompetitorFollowedDAORes{},
			Total:         total,
		}, nil
	}

	lastCreatedAtRes := results[len(results)-1].CreatedAt
	followers := make([]*dao.GetUserCompetitorFollowedDAORes, len(results))
	for i, res := range results {
		followers[i] = res.User
	}

	userFollowers := &dao.GetUserFollowersDAORes{
		LastCreatedAt: lastCreatedAtRes,
		Followers:     followers,
		Total:         total,
	}

	return userFollowers, nil
}

func (r *Repository) IsFollowing(ctx context.Context, fromOID, userToOID *primitive.ObjectID) (bool, error) {
	// Crea el filtro para verificar si existe una relación de seguimiento
	filter := bson.M{
		"from":       fromOID,                  // El ID del usuario que sigue
		"to_user":    userToOID,                // El ID del usuario al que sigue
		"deleted_at": bson.M{"$exists": false}, // Verifica que no esté marcado como eliminado
	}

	// Intenta encontrar un documento que cumpla con el filtro
	count, err := r.followerColl.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}

	// Si el recuento es mayor a 0, significa que la relación existe
	return count > 0, nil
}

