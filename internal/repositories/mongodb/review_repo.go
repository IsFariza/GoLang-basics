package mongodb

import (
	"context"

	"github.com/BlackHole55/software-store-final/internal/domain"
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
	_, err := r.collection.InsertOne(ctx, review)
	return err
}

func (r *ReviewRepository) GetAll(ctx context.Context) ([]*domain.Review, error) {
	var reviews []*domain.Review

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &reviews)
	return reviews, err
}

func (r *ReviewRepository) GetById(ctx context.Context, id string) (*domain.Review, error) {
	var review domain.Review

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}

	err = r.collection.FindOne(ctx, filter).Decode(&review)

	if err != nil {
		return nil, domain.ErrorNotFound
	}
	return &review, nil
}

func (r *ReviewRepository) GetReviewsByGameId(ctx context.Context, gameId string) ([]*domain.Review, error) {
	var reviews []*domain.Review

	objID, err := bson.ObjectIDFromHex(gameId)
	if err != nil {
		return nil, err
	}

	cursor, err := r.collection.Find(ctx, bson.M{"game_id": objID})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	err = cursor.All(ctx, &reviews)
	return reviews, err
}

func (r *ReviewRepository) Update(ctx context.Context, id string, updateReview *domain.Review) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"content":    updateReview.Content,
			"rating":     updateReview.Rating,
			"updated_at": updateReview.UpdatedAt,
		},
	}

	res, err := r.collection.UpdateOne(ctx, filter, update)
	if res.MatchedCount == 0 {
		return domain.ErrorNotFound
	}
	return err
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
