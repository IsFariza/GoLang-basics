package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Review struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserId    bson.ObjectID `bson:"user_id" json:"user_id"`
	Username  string        `bson:"username" json:"username"`
	GameId    bson.ObjectID `bson:"game_id" json:"game_id"`
	Rating    *int          `bson:"rating" json:"rating" binding:"required"`
	Content   string        `bson:"content" json:"content"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt *time.Time    `bson:"updated_at" json:"updated_at"`
}
