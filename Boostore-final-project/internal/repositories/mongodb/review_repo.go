package mongodb

import (
	"context"

	"github.com/BlackHole55/software-store-final/internal/domain"
	"github.com/BlackHole55/software-store-final/internal/repositories/mongodb/dao"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ReviewRepository struct {
	collection *mongo.Collection
}

func NewReviewRepository(client *mongo.Client) *ReviewRepository {
	return &ReviewRepository{
		collection: client.Database("softwarestore").Collection("reviews"),
	}
}

func (r *ReviewRepository) Create(ctx context.Context, review *domain.Review) error {
	doc := dao.FromReviewDomain(review)
	res, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}
	if oid, ok := res.InsertedID.(bson.ObjectID); ok {
		review.ID = oid.Hex()
	}
	return nil
}

func (r *ReviewRepository) GetAll(ctx context.Context) ([]*domain.Review, error) {
	var docs []*dao.ReviewDoc

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &docs)
	if err != nil {
		return nil, err
	}
	reviews := make([]*domain.Review, len(docs))
	for i, d := range docs {
		reviews[i] = d.ToDomain()
	}
	return reviews, err
}

func (r *ReviewRepository) GetById(ctx context.Context, id string) (*domain.Review, error) {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var doc dao.ReviewDoc

	filter := bson.M{"_id": objID}

	err = r.collection.FindOne(ctx, filter).Decode(&doc)

	if err != nil {
		return nil, domain.ErrorNotFound
	}
	return doc.ToDomain(), nil
}

func (r *ReviewRepository) GetReviewsByGameId(ctx context.Context, gameId string) ([]*domain.Review, error) {
	objID, err := bson.ObjectIDFromHex(gameId)
	if err != nil {
		return nil, err
	}

	var docs []*dao.ReviewDoc

	cursor, err := r.collection.Find(ctx, bson.M{"game_id": objID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &docs)
	if err != nil {
		return nil, err
	}
	reviews := make([]*domain.Review, len(docs))
	for i, d := range docs {
		reviews[i] = d.ToDomain()
	}
	return reviews, err
}

func (r *ReviewRepository) Update(ctx context.Context, id string, updateReview *domain.Review) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	doc := dao.FromReviewDomain(updateReview)

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"content":    doc.Content,
			"rating":     doc.Rating,
			"updated_at": doc.UpdatedAt,
		},
	}

	res, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return domain.ErrorNotFound
	}
	return nil
}

func (r *ReviewRepository) Delete(ctx context.Context, id string, userId string, userRole string) error {
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
