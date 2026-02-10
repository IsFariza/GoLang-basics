package dao

import (
	"time"

	"github.com/BlackHole55/software-store-final/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type GameDoc struct {
	ID             bson.ObjectID  `bson:"_id,omitempty"`
	PublisherId    bson.ObjectID  `bson:"publisher_id"`
	DeveloperId    bson.ObjectID  `bson:"developer_id"`
	EmulationId    *bson.ObjectID `bson:"emulation_id,omitempty"`
	UserId         bson.ObjectID  `bson:"user_id"`
	OriginalSystem string         `bson:"original_system"`
	Title          string         `bson:"title"`
	Description    string         `bson:"description"`
	ReleaseDate    *time.Time     `bson:"release_date"`
	Price          *float32       `bson:"price"`
	IsVerified     bool           `bson:"is_verified"`
	Category       []string       `bson:"category"`
	CreatedAt      time.Time      `bson:"created_at"`
	UpdatedAt      *time.Time     `bson:"updated_at"`
}

type PopulatedGameDoc struct {
	ID             bson.ObjectID `bson:"_id"`
	Publisher      CompanyDoc    `bson:"publisher"`
	Developer      CompanyDoc    `bson:"developer"`
	Emulation      *EmulationDoc `bson:"emulation"`
	OriginalSystem string        `bson:"original_system"`
	Title          string        `bson:"title"`
	Description    string        `bson:"description"`
	ReleaseDate    *time.Time    `bson:"release_date"`
	Price          *float32      `bson:"price"`
	IsVerified     bool          `bson:"is_verified"`
	Category       []string      `bson:"category"`
	CreatedAt      time.Time     `bson:"created_at"`
	UpdatedAt      *time.Time    `bson:"updated_at"`
}

func FromGameDomain(g *domain.Game) *GameDoc {
	doc := &GameDoc{
		OriginalSystem: g.OriginalSystem,
		Title:          g.Title,
		Description:    g.Description,
		ReleaseDate:    g.ReleaseDate,
		Price:          g.Price,
		IsVerified:     g.IsVerified,
		Category:       g.Category,
		CreatedAt:      g.CreatedAt,
		UpdatedAt:      g.UpdatedAt,
	}
	if oid, err := bson.ObjectIDFromHex(g.ID); err == nil {
		doc.ID = oid
	}
	if oid, err := bson.ObjectIDFromHex(g.PublisherId); err == nil {
		doc.PublisherId = oid
	}
	if oid, err := bson.ObjectIDFromHex(g.DeveloperId); err == nil {
		doc.DeveloperId = oid
	}
	if oid, err := bson.ObjectIDFromHex(g.UserId); err == nil {
		doc.UserId = oid
	}
	if g.EmulationId != "" {
		if oid, err := bson.ObjectIDFromHex(g.EmulationId); err == nil {
			doc.EmulationId = &oid
		}
	}
	return doc
}

func (d GameDoc) ToDomain() *domain.Game {
	game := &domain.Game{
		ID:             d.ID.Hex(),
		PublisherId:    d.PublisherId.Hex(),
		DeveloperId:    d.DeveloperId.Hex(),
		UserId:         d.UserId.Hex(),
		OriginalSystem: d.OriginalSystem,
		Title:          d.Title,
		Description:    d.Description,
		ReleaseDate:    d.ReleaseDate,
		Price:          d.Price,
		IsVerified:     d.IsVerified,
		Category:       d.Category,
		CreatedAt:      d.CreatedAt,
		UpdatedAt:      d.UpdatedAt,
	}
	if d.EmulationId != nil {
		game.EmulationId = d.EmulationId.Hex()
	}
	return game
}
func (d PopulatedGameDoc) ToDomain() *domain.PopulatedGame {
	game := &domain.PopulatedGame{
		ID:             d.ID.Hex(),
		Publisher:      d.Publisher.ToDomain(),
		Developer:      d.Developer.ToDomain(),
		OriginalSystem: d.OriginalSystem,
		Title:          d.Title,
		Description:    d.Description,
		ReleaseDate:    d.ReleaseDate,
		Price:          d.Price,
		IsVerified:     d.IsVerified,
		Category:       d.Category,
		CreatedAt:      d.CreatedAt,
		UpdatedAt:      d.UpdatedAt,
	}
	if d.Emulation != nil {
		game.Emulation = d.Emulation.ToDomain()
	}

	return game
}
