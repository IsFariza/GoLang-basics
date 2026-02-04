package mongodb

import (
	"context"
	"fmt"

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

func (r *GameRepository) Create(ctx context.Context, game *domain.Game, id string) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	game.UserId = objID

	_, err = r.collection.InsertOne(ctx, game)

	return err
}

func (r *GameRepository) GetAll(ctx context.Context) ([]*domain.Game, error) {
	var games []*domain.Game

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &games)
	return games, err
}

func (r *GameRepository) GetAllVerified(ctx context.Context) ([]*domain.Game, error) {
	var games []*domain.Game

	filter := bson.M{"is_verified": true}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &games)
	return games, err

}

func (r *GameRepository) GetById(ctx context.Context, id string) (*domain.Game, error) {
	var game domain.Game

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}

	err = r.collection.FindOne(ctx, filter).Decode(&game)

	if err != nil {
		return nil, domain.ErrorNotFound
	}

	return &game, err
}

func (r *GameRepository) GetByIds(ctx context.Context, ids []string) ([]domain.Game, error) {
	var objectIDs []bson.ObjectID
	for _, id := range ids {
		if objID, err := bson.ObjectIDFromHex(id); err == nil {
			objectIDs = append(objectIDs, objID)
		}
	}

	filter := bson.M{"_id": bson.M{"$in": objectIDs}}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var games []domain.Game
	if err := cursor.All(ctx, &games); err != nil {
		return nil, err
	}

	return games, nil
}

func (r *GameRepository) GetByUserId(ctx context.Context, userId string) ([]*domain.Game, error) {
	var games []*domain.Game

	objID, err := bson.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"user_id": objID}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &games)
	return games, err
}

func (r *GameRepository) Update(ctx context.Context, id string, updatedGame *domain.Game) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	fmt.Printf("DEBUG: Categories received for update: %v\n", updatedGame.Category)

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"title":           updatedGame.Title,
			"description":     updatedGame.Description,
			"price":           updatedGame.Price,
			"original_system": updatedGame.OriginalSystem,
			"publisher_id":    updatedGame.PublisherId,
			"developer_id":    updatedGame.DeveloperId,
			"emulation_id":    updatedGame.EmulationId,
			"category":        updatedGame.Category,
			"is_verified":     updatedGame.IsVerified,
			"updated_at":      updatedGame.UpdatedAt,
		},
	}

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

	res, err := r.collection.DeleteOne(ctx, filter)
	if res.DeletedCount == 0 {
		return domain.ErrorNotFound
	}

	return err
}

func (r *GameRepository) Verify(ctx context.Context, id string) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"is_verified": true}}

	res, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return domain.ErrorNotFound
	}
	return nil
}

func (r *GameRepository) Unverify(ctx context.Context, id string) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"is_verified": false}}

	res, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return domain.ErrorNotFound
	}
	return nil
}

func (r *GameRepository) SearchByTitle(ctx context.Context, title string) ([]*domain.Game, error) {
	var games []*domain.Game

	filter := bson.M{"is_verified": true, "title": bson.M{"$regex": title, "$options": "i"}}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &games); err != nil {
		return nil, err
	}

	return games, nil
}

func (r *GameRepository) InitIndices(ctx context.Context) error {
	indices := []mongo.IndexModel{
		{Keys: bson.D{{Key: "is_verified", Value: 1}, {Key: "title", Value: 1}}},

		{Keys: bson.D{{Key: "user_id", Value: 1}}},

		{Keys: bson.D{{Key: "is_verified", Value: 1}}},
	}

	_, err := r.collection.Indexes().CreateMany(ctx, indices)
	return err
}
