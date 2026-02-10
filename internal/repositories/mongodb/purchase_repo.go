package mongodb

import (
	"context"

	"github.com/BlackHole55/software-store-final/internal/domain"
	"github.com/BlackHole55/software-store-final/internal/repositories/mongodb/dao"
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
	doc := dao.FromPurchaseDomain(purchase)
	res, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}
	if oid, ok := res.InsertedID.(bson.ObjectID); ok {
		purchase.ID = oid.Hex()
	}
	return err
}

func (r *PurchaseRepo) GetAll(ctx context.Context) ([]*domain.Purchase, error) {
	var docs []*dao.PurchaseDoc
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &docs)
	if err != nil {
		return nil, err
	}
	purchases := make([]*domain.Purchase, len(docs))
	for i, d := range docs {
		purchases[i] = d.ToDomain()
	}
	return purchases, err
}

func (r *PurchaseRepo) GetById(ctx context.Context, id string) (*domain.Purchase, error) {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var doc dao.PurchaseDoc

	filter := bson.M{"_id": objID}

	err = r.collection.FindOne(ctx, filter).Decode(&doc)
	if err != nil {
		return nil, domain.ErrorNotFound
	}
	return doc.ToDomain(), err
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
