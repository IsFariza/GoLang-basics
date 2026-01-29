package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Review struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserId    bson.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	GameId    bson.ObjectID `bson:"game_id_id,omitempty" json:"game_id"`
	Rating    float32       `bson:"rating" json:"rating"`
	Content   string        `bson:"content" json:"content"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`
}
