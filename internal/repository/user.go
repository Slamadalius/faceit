package repository

import (
	"context"
	"time"

	"github.com/Slamadalius/faceit/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type user struct {
	ID        primitive.ObjectID `bson:"_id"`
	FirstName string             `bson:"first_name"`
	LastName  string             `bson:"last_name"`
	Nickname  string             `bson:"nickname"`
	Password  string             `bson:"password"`
	Email     string             `bson:"email"`
	Country   string             `bson:"country"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type userRepository struct {
	client *mongo.Client
}

func NewUserRepository(client *mongo.Client) entity.UserRepository {
	return &userRepository{
		client: client,
	}
}

func (r *userRepository) Insert(ctx context.Context, entityUser entity.User) (err error) {
	coll := r.client.Database("faceit").Collection("users")

	document := user{}
	document.ID = primitive.NewObjectID()
	document.mapEntity(entityUser)
	document.CreatedAt = time.Now()

	_, err = coll.InsertOne(ctx, document)
	return
}

func (r *userRepository) Update(ctx context.Context, userID string, entityUser entity.User) (err error) {
	coll := r.client.Database("faceit").Collection("users")

	objectID, _ := primitive.ObjectIDFromHex(userID)

	update := bson.M{
		"$set": bson.M{
			"first_name": entityUser.FirstName,
			"last_name":  entityUser.LastName,
			"nickname":   entityUser.Nickname,
			"password":   entityUser.Password,
			"email":      entityUser.Email,
			"country":    entityUser.Country,
		},
	}

	_, err = coll.UpdateByID(ctx, objectID, update)
	return
}

func (r *userRepository) Delete(ctx context.Context, userID string) (err error) {
	coll := r.client.Database("faceit").Collection("users")

	objectID, _ := primitive.ObjectIDFromHex(userID)

	_, err = coll.DeleteOne(ctx, bson.M{"_id": objectID})
	return
}

func (u *user) mapEntity(entityUser entity.User) {
	u.FirstName = entityUser.FirstName
	u.LastName = entityUser.LastName
	u.Nickname = entityUser.Nickname
	u.Password = entityUser.Password
	u.Email = entityUser.Email
	u.Country = entityUser.Country
}
