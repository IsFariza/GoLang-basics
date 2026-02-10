package mongodb

import (
	"context"

	"github.com/BlackHole55/software-store-final/internal/domain"
	"github.com/BlackHole55/software-store-final/internal/repositories/mongodb/dao"
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
	doc := dao.FromCompanyDomain(company)

	res, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}
	if objID, ok := res.InsertedID.(bson.ObjectID); ok {
		company.ID = objID.Hex()
	}
	return nil
}

func (r *CompanyRepository) GetAll(ctx context.Context) ([]*domain.Company, error) {
	var docs []*dao.CompanyDoc

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &docs)
	if err != nil {
		return nil, err
	}
	companies := make([]*domain.Company, len(docs))
	for i, d := range docs {
		companies[i] = d.ToDomain()
	}
	return companies, nil
}

func (r *CompanyRepository) GetById(ctx context.Context, id string) (*domain.Company, error) {
	var doc dao.CompanyDoc

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, domain.ErrorNotFound
	}

	filter := bson.M{"_id": objID}

	err = r.collection.FindOne(ctx, filter).Decode(&doc)

	if err != nil {
		return nil, domain.ErrorNotFound
	}

	return doc.ToDomain(), err

}

func (r *CompanyRepository) GetVerified(ctx context.Context) ([]*domain.Company, error) {
	var docs []*dao.CompanyDoc

	cursor, err := r.collection.Find(ctx, bson.M{"is_verified": true})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &docs); err != nil {
		return nil, err
	}
	companies := make([]*domain.Company, len(docs))
	for i, d := range docs {
		companies[i] = d.ToDomain()
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
	doc := dao.FromCompanyDomain(updatedCompany)

	filter := bson.M{"_id": objID}

	update := bson.M{
		"$set": bson.M{
			"name":             doc.Name,
			"description":      doc.Description,
			"country":          doc.Country,
			"is_verified":      doc.IsVerified,
			"contacts.email":   doc.Contacts.Email,
			"contacts.phone":   doc.Contacts.Phone,
			"contacts.website": doc.Contacts.Website,
			"updated_at":       doc.UpdatedAt,
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
