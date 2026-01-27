package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Category struct {
	ID         bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name       string        `bson:"name" json:"name"`
	Created_at time.Time     `bson:"created_at" json:"created_at"`
	Updated_at time.Time     `bson:"updated_at" json:"updated_at"`
}
