package mongodb

import (
	"context"

	"github.com/BlackHole55/software-store-final/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CompanyRepository struct {
	collection *mongo.Collection
}

func NewCompanyRepository(client *mongo.Client) *CompanyRepository {
	return &CompanyRepository{
		collection: client.Database("softwarestore").Collection("companies"),
	}
}

func (r *CompanyRepository) Create(ctx context.Context, company *domain.Company) error {

	_, err := r.collection.InsertOne(ctx, company)

	return err
}

func (r *CompanyRepository) GetAll(ctx context.Context) ([]*domain.Company, error) {
	var companies []*domain.Company

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &companies)

	return companies, err
}

func (r *CompanyRepository) GetById(ctx context.Context, id string) (*domain.Company, error) {
	var company *domain.Company

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return company, err
	}

	filter := bson.M{"_id": objID}

	err = r.collection.FindOne(ctx, filter).Decode(&company)

	if err != nil {
		return company, domain.ErrorNotFound
	}

	return company, err

}

func (r *CompanyRepository) Update(ctx context.Context, id string, updatedCompany *domain.Company) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updatedCompany}

	res, err := r.collection.UpdateOne(ctx, filter, update)
	if res.MatchedCount == 0 {
		return domain.ErrorNotFound
	}

	return err
}

func (r *CompanyRepository) Delete(ctx context.Context, id string) error {
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
