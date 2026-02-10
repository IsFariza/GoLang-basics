package mongodb

import (
	"context"

	"github.com/BlackHole55/software-store-final/internal/domain"
	"github.com/BlackHole55/software-store-final/internal/repositories/mongodb/dao"
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
	doc := dao.FromUserDomain(user)
	res, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}
	if oid, ok := res.InsertedID.(bson.ObjectID); ok {
		user.ID = oid.Hex()
	}
	return nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*domain.User, error) {
	var docs []*dao.UserDoc

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &docs)
	if err != nil {
		return nil, err
	}
	users := make([]*domain.User, len(docs))
	for i, d := range docs {
		users[i] = d.ToDomain()
	}
	return users, nil
}

func (r *UserRepository) GetById(ctx context.Context, id string) (*domain.User, error) {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var doc dao.UserDoc

	filter := bson.M{"_id": objID}
	err = r.collection.FindOne(ctx, filter).Decode(&doc)

	if err != nil {
		return nil, domain.ErrorNotFound
	}

	return doc.ToDomain(), nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var doc dao.UserDoc

	filter := bson.M{"email": email}

	err := r.collection.FindOne(ctx, filter).Decode(&doc)

	if err != nil {
		return nil, domain.ErrorNotFound
	}

	return doc.ToDomain(), nil
}

func (r *UserRepository) Update(ctx context.Context, id string, updatedUser *domain.User) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	doc := dao.FromUserDomain(updatedUser)

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": doc}

	res, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return domain.ErrorNotFound
	}

	return nil
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

func (r *UserRepository) ChangeRoleToModerator(ctx context.Context, id string) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"role": "moderator"}}

	res, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return domain.ErrorNotFound
	}
	return nil

}

func (r *UserRepository) ChangeRoleToUser(ctx context.Context, id string) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"role": "user"}}

	res, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return domain.ErrorNotFound
	}
	return nil
}
