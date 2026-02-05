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

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &companies)

	return companies, err
}

func (r *CompanyRepository) GetById(ctx context.Context, id string) (*domain.Company, error) {
	var company domain.Company

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}

	err = r.collection.FindOne(ctx, filter).Decode(&company)

	if err != nil {
		return nil, domain.ErrorNotFound
	}

	return &company, err

}

func (r *CompanyRepository) GetVerified(ctx context.Context) ([]*domain.Company, error) {
	var companies []*domain.Company

	cursor, err := r.collection.Find(ctx, bson.M{"is_verified": true})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &companies); err != nil {
		return nil, err
	}
	return companies, nil
}

func (r *CompanyRepository) Verify(ctx context.Context, id string) error {
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

func (r *CompanyRepository) Unverify(ctx context.Context, id string) error {
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

func (r *CompanyRepository) Update(ctx context.Context, id string, updatedCompany *domain.Company) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}

	update := bson.M{
		"$set": bson.M{
			"name":             updatedCompany.Name,
			"description":      updatedCompany.Description,
			"country":          updatedCompany.Country,
			"is_verified":      updatedCompany.IsVerified,
			"contacts.email":   updatedCompany.Contacts.Email,
			"contacts.phone":   updatedCompany.Contacts.Phone,
			"contacts.website": updatedCompany.Contacts.Website,
			"updated_at":       updatedCompany.UpdatedAt,
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
