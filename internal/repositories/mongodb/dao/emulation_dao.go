package dao

import (
	"github.com/BlackHole55/software-store-final/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type EmulationDoc struct {
	ID   bson.ObjectID `bson:"_id,omitempty"`
	Name string        `bson:"name"`
}

func FromEmulationDomain(e *domain.Emulation) *EmulationDoc {
	doc := &EmulationDoc{
		Name: e.Name,
	}

	if e.ID != "" {
		if objID, err := bson.ObjectIDFromHex(e.ID); err == nil {
			doc.ID = objID
		}
	}

	return doc
}
func (d EmulationDoc) ToDomain() *domain.Emulation {
	return &domain.Emulation{
		ID:   d.ID.Hex(),
		Name: d.Name,
	}
}
