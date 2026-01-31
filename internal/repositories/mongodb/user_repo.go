package mongodb

import (
	"context"

	"github.com/BlackHole55/software-store-final/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{
		collection: client.Database("softwarestore").Collection("users"),
	}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	_, err := r.collection.InsertOne(ctx, user)

	return err
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*domain.User, error) {
	var users []*domain.User

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &users)
	return users, err
}

func (r *UserRepository) GetById(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}

	err = r.collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return nil, domain.ErrorNotFound
	}

	return &user, err
}

func (r *UserRepository) Update(ctx context.Context, id string, updatedUser *domain.User) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updatedUser}

	res, err := r.collection.UpdateOne(ctx, filter, update)
	if res.MatchedCount == 0 {
		return domain.ErrorNotFound
	}

	return err
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}

	res, err := r.collection.DeleteOne(ctx, filter)
	if res.DeletedCount == 0 {
		return domain.ErrorNotFound
	}

	return err
}
