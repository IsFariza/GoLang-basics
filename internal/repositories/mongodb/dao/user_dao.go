package dao

import (
	"time"

	"github.com/BlackHole55/software-store-final/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserDoc struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	Username     string        `bson:"username"`
	Email        string        `bson:"email"`
	PasswordHash []byte        `bson:"password_hash"`
	Role         string        `bson:"role,omitempty"`
	CreatedAt    time.Time     `bson:"created_at"`
	UpdatedAt    *time.Time    `bson:"updated_at"`
	Library      []UserGameDoc `bson:"library"`
}

type UserGameDoc struct {
	GameId        bson.ObjectID `bson:"game_id"`
	AddedAt       time.Time     `bson:"added_at"`
	PlaytimeHours float32       `bson:"playtime_hours"`
}

func FromUserDomain(u *domain.User) *UserDoc {
	libraryDocs := make([]UserGameDoc, len(u.Library))
	for i, item := range u.Library {
		gID, _ := bson.ObjectIDFromHex(item.GameId)
		libraryDocs[i] = UserGameDoc{
			GameId:        gID,
			AddedAt:       item.AddedAt,
			PlaytimeHours: item.PlaytimeHours,
		}
	}

	doc := &UserDoc{
		Username:     u.Username,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		Role:         u.Role,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
		Library:      libraryDocs,
	}

	if oid, err := bson.ObjectIDFromHex(u.ID); err == nil {
		doc.ID = oid
	}
	return doc
}

func (d UserDoc) ToDomain() *domain.User {
	library := make([]domain.UserGame, len(d.Library))
	for i, item := range d.Library {
		library[i] = domain.UserGame{
			GameId:        item.GameId.Hex(),
			AddedAt:       item.AddedAt,
			PlaytimeHours: item.PlaytimeHours,
		}
	}

	return &domain.User{
		ID:           d.ID.Hex(),
		Username:     d.Username,
		Email:        d.Email,
		PasswordHash: d.PasswordHash,
		Role:         d.Role,
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
		Library:      library,
	}
}
