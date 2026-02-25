package mongodb

import (
	"context"

	"github.com/BlackHole55/software-store-final/internal/delivery/dto"
	"github.com/BlackHole55/software-store-final/internal/domain"
	"github.com/BlackHole55/software-store-final/internal/repositories/mongodb/dao"
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
	game.UserId = id
	doc := dao.FromGameDomain(game)

	res, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}
	if oid, ok := res.InsertedID.(bson.ObjectID); ok {
		game.ID = oid.Hex()
	}
	return nil
}

func (r *GameRepository) GetAll(ctx context.Context) ([]*domain.Game, error) {
	var docs []*dao.GameDoc

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &docs)
	if err != nil {
		return nil, err
	}
	games := make([]*domain.Game, len(docs))
	for i, d := range docs {
		games[i] = d.ToDomain()
	}
	return games, err
}

func (r *GameRepository) GetAllVerified(ctx context.Context) ([]*domain.Game, error) {
	var docs []*dao.GameDoc

	filter := bson.M{"is_verified": true}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &docs)
	if err != nil {
		return nil, err
	}
	games := make([]*domain.Game, len(docs))
	for i, d := range docs {
		games[i] = d.ToDomain()
	}
	return games, err

}

func (r *GameRepository) GetById(ctx context.Context, id string) (*domain.Game, error) {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var doc dao.GameDoc
	filter := bson.M{"_id": objID}

	err = r.collection.FindOne(ctx, filter).Decode(&doc)

	if err != nil {
		return nil, domain.ErrorNotFound
	}

	return doc.ToDomain(), err
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

	var docs []dao.GameDoc
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, err
	}
	games := make([]domain.Game, len(docs))
	for i, d := range docs {
		games[i] = *d.ToDomain()
	}
	return games, nil
}

func (r *GameRepository) GetByUserId(ctx context.Context, userId string) ([]*domain.Game, error) {
	objID, err := bson.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	var docs []*dao.GameDoc

	filter := bson.M{"user_id": objID}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &docs)
	if err != nil {
		return nil, err
	}
	games := make([]*domain.Game, len(docs))
	for i, d := range docs {
		games[i] = d.ToDomain()
	}
	return games, err
}

func (r *GameRepository) Update(ctx context.Context, id string, updatedGame *domain.Game) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	doc := dao.FromGameDomain(updatedGame)

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"title":           doc.Title,
			"description":     doc.Description,
			"price":           doc.Price,
			"original_system": doc.OriginalSystem,
			"publisher_id":    doc.PublisherId,
			"developer_id":    doc.DeveloperId,
			"emulation_id":    doc.EmulationId,
			"category":        doc.Category,
			"is_verified":     doc.IsVerified,
			"updated_at":      doc.UpdatedAt,
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
	var docs []*dao.GameDoc

	filter := bson.M{"is_verified": true, "title": bson.M{"$regex": title, "$options": "i"}}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &docs); err != nil {
		return nil, err
	}

	games := make([]*domain.Game, len(docs))
	for i, d := range docs {
		games[i] = d.ToDomain()
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

func (r *GameRepository) GetStats(ctx context.Context) (*dto.GameStatsDTO, error) {
	pipeline := mongo.Pipeline{
		{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: nil},
				{Key: "total_games", Value: bson.D{{Key: "$sum", Value: 1}}},
				{Key: "total_revenue", Value: bson.D{{Key: "$sum", Value: "$price"}}},
				{Key: "avg_price", Value: bson.D{{Key: "$avg", Value: "$price"}}},
				{Key: "min_price", Value: bson.D{{Key: "$min", Value: "$price"}}},
				{Key: "max_price", Value: bson.D{{Key: "$max", Value: "$price"}}},
			}},
		},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []struct {
		TotalGames   int     `bson:"total_games"`
		TotalPrice   float64 `bson:"total_price"`
		AvgGamePrice float64 `bson:"avg_price"`
		MinGamePrice float64 `bson:"min_price"`
		MaxGamePrice float64 `bson:"max_price"`
	}

	if err := cursor.All(ctx, &results); err != nil || len(results) == 0 {
		return &dto.GameStatsDTO{}, err
	}

	res := results[0]
	return &dto.GameStatsDTO{
		TotalGames:   res.TotalGames,
		TotalPrice:   res.TotalPrice,
		AvgGamePrice: res.AvgGamePrice,
		MinGamePrice: res.MinGamePrice,
		MaxGamePrice: res.MaxGamePrice,
	}, nil
}
