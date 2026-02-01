package mongodb

import (
	"context"

	"github.com/BlackHole55/software-store-final/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type PurchaseRepo struct {
	collection *mongo.Collection
}

func NewPurchaseRepo(client *mongo.Client) *PurchaseRepo {
	return &PurchaseRepo{
		collection: client.Database("softwarestore").Collection("purchases")}
}

func (r *PurchaseRepo) Create(ctx context.Context, purchase *domain.Purchase) error {
	_, err := r.collection.InsertOne(ctx, purchase)
	return err
}

func (r *PurchaseRepo) GetAll(ctx context.Context) ([]*domain.Purchase, error) {
	var purchases []*domain.Purchase
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &purchases)
	return purchases, err
}

func (r *PurchaseRepo) GetById(ctx context.Context, id string) (*domain.Purchase, error) {
	var purchase domain.Purchase

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"_id", objID}}
	err = r.collection.FindOne(ctx, filter).Decode(&purchase)
	if err != nil {
		return nil, domain.ErrorNotFound
	}
	return &purchase, err
}

func (r *PurchaseRepo) Delete(ctx context.Context, id string) error {
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
