package dao

import (
	"time"

	"github.com/BlackHole55/software-store-final/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ReviewDoc struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	UserId    bson.ObjectID `bson:"user_id"`
	Username  string        `bson:"username"`
	GameId    bson.ObjectID `bson:"game_id"`
	Rating    *int          `bson:"rating"`
	Content   string        `bson:"content"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt *time.Time    `bson:"updated_at"`
}

func FromReviewDomain(r *domain.Review) *ReviewDoc {
	doc := &ReviewDoc{
		Username:  r.Username,
		Rating:    r.Rating,
		Content:   r.Content,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
	if oid, err := bson.ObjectIDFromHex(r.ID); err == nil {
		doc.ID = oid
	}
	if oid, err := bson.ObjectIDFromHex(r.UserId); err == nil {
		doc.UserId = oid
	}
	if oid, err := bson.ObjectIDFromHex(r.GameId); err == nil {
		doc.GameId = oid
	}
	return doc
}

func (d ReviewDoc) ToDomain() *domain.Review {
	return &domain.Review{
		ID:        d.ID.Hex(),
		UserId:    d.UserId.Hex(),
		Username:  d.Username,
		GameId:    d.GameId.Hex(),
		Rating:    d.Rating,
		Content:   d.Content,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}
