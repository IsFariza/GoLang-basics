package mongodb

import (
	"context"

	"github.com/BlackHole55/software-store-final/internal/domain"
	"github.com/BlackHole55/software-store-final/internal/repositories/mongodb/dao"
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

func (r *EmulationRepository) Create(ctx context.Context, emulation *domain.Emulation) error {
	doc := dao.FromEmulationDomain(emulation)
	res, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}

	objId, ok := res.InsertedID.(bson.ObjectID)
	if ok {
		emulation.ID = objId.Hex()
	}

	return nil
}

func (r *EmulationRepository) GetAll(ctx context.Context) ([]*domain.Emulation, error) {
	var docs []*dao.EmulationDoc

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &docs)
	if err != nil {
		return nil, err
	}

	emulations := make([]*domain.Emulation, len(docs))
	for i, d := range docs {
		emulations[i] = d.ToDomain()
	}
	return emulations, err
}

func (r *EmulationRepository) GetById(ctx context.Context, id string) (*domain.Emulation, error) {
	var doc dao.EmulationDoc

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}

	err = r.collection.FindOne(ctx, filter).Decode(&doc)

	if err != nil {
		return nil, domain.ErrorNotFound
	}

	return doc.ToDomain(), nil
}

func (r *EmulationRepository) Update(ctx context.Context, id string, updatedEmulation *domain.Emulation) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	doc := dao.FromEmulationDomain(updatedEmulation)

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"name": doc.Name}}

	res, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return domain.ErrorNotFound
	}

	return nil
}

func (r *EmulationRepository) Delete(ctx context.Context, id string) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}

	res, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return domain.ErrorNotFound
	}

	return nil
}
