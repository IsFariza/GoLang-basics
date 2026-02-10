package dao

import (
	"time"

	"github.com/BlackHole55/software-store-final/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CompanyDoc struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Description string        `bson:"description"`
	Country     string        `bson:"country"`
	Contacts    ContactsDoc   `bson:"contacts"`
	IsVerified  bool          `bson:"is_verified"`
	CreatedAt   time.Time     `bson:"created_at"`
	UpdatedAt   *time.Time    `bson:"updated_at"`
}

type ContactsDoc struct {
	Email   string `bson:"email"`
	Phone   string `bson:"phone"`
	Website string `bson:"website"`
}

func FromCompanyDomain(c *domain.Company) *CompanyDoc {
	doc := &CompanyDoc{
		Name:        c.Name,
		Description: c.Description,
		Country:     c.Country,
		Contacts: ContactsDoc{
			Email:   c.Contacts.Email,
			Phone:   c.Contacts.Phone,
			Website: c.Contacts.Website,
		},
		IsVerified: c.IsVerified,
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
	}
	if c.ID != "" {
		if objID, err := bson.ObjectIDFromHex(c.ID); err == nil {
			doc.ID = objID
		}
	}
	return doc
}

func (d CompanyDoc) ToDomain() *domain.Company {
	return &domain.Company{
		ID:          d.ID.Hex(),
		Name:        d.Name,
		Description: d.Description,
		Country:     d.Country,
		Contacts: domain.Contacts{
			Email:   d.Contacts.Email,
			Phone:   d.Contacts.Phone,
			Website: d.Contacts.Website,
		},
		IsVerified: d.IsVerified,
		CreatedAt:  d.CreatedAt,
		UpdatedAt:  d.UpdatedAt,
	}
}
