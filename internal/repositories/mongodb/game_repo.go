package mongodb

import (
	"context"

	"github.com/BlackHole55/software-store-final/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type GameRepository struct {
	collection *mongo.Collection
}

func NewGameRepository(client *mongo.Client) *GameRepository {
	return &GameRepository{
		collection: client.Database("softwarestore").Collection("games"),
	}
}

func (r *GameRepository) Create(ctx context.Context, game domain.Game) error {
	_, err := r.collection.InsertOne(ctx, game)

	return err
}

func (r *GameRepository) GetAll(ctx context.Context) ([]domain.Game, error) {
	var games []domain.Game

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &games)
	return games, err
}

func (r *GameRepository) GetById(ctx context.Context, id string) (domain.Game, error) {
	var game domain.Game

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return game, err
	}

	filter := bson.M{"_id": objID}

	err = r.collection.FindOne(ctx, filter).Decode(&game)

	return game, err
}

// TODO: make update
func (r *GameRepository) Update(ctx context.Context, id string, updates domain.Game) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{
		"title":        updates.Title,
		"description":  updates.Description,
		"release_date": updates.ReleaseDate,
		"price":        updates.Price,
		"updated_at":   updates.UpdatedAt,
		"is_verified":  false,
	}}

	// TODO: handle not found
	res, err := r.collection.UpdateOne(ctx, filter, update)
	if res.MatchedCount == 0 {
		return domain.ErrorNotFound
	}

	return err
}

func (r *GameRepository) Delete(ctx context.Context, id string) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}

	// TODO: handle not found
	_, err = r.collection.DeleteOne(ctx, filter)

	return err
}
