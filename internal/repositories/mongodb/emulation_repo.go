package mongodb

import (
	"context"

	"github.com/BlackHole55/software-store-final/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type EmulationRepository struct {
	collection *mongo.Collection
}

func NewEmulationRepository(client *mongo.Client) *EmulationRepository {
	return &EmulationRepository{
		collection: client.Database("softwarestore").Collection("emulations"),
	}
}

func (r *EmulationRepository) Create(ctx context.Context, emulation domain.Emulation) error {
	_, err := r.collection.InsertOne(ctx, emulation)

	return err
}

func (r *EmulationRepository) GetAll(ctx context.Context) ([]domain.Emulation, error) {
	var emulations []domain.Emulation

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &emulations)
	return emulations, err
}

func (r *EmulationRepository) GetById(ctx context.Context, id string) (domain.Emulation, error) {
	var emulation domain.Emulation

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return emulation, err
	}

	filter := bson.M{"_id": objID}

	err = r.collection.FindOne(ctx, filter).Decode(&emulation)

	if err != nil {
		return emulation, domain.ErrorNotFound
	}

	return emulation, err
}

func (r *EmulationRepository) Update(ctx context.Context, id string, updatedEmulation domain.Emulation) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updatedEmulation}

	res, err := r.collection.UpdateOne(ctx, filter, update)
	if res.MatchedCount == 0 {
		return domain.ErrorNotFound
	}

	return err
}

func (r *EmulationRepository) Delete(ctx context.Context, id string) error {
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
