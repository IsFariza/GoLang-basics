package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID         bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	Username   string        `bson:"username" json:"username"`
	Email      string        `bson:"email" json:"email"`
	Password   []byte        `bson:"password_hash" json:"-"`
	Role       string        `bson:"role,omitempty" json:"role"`
	Created_at time.Time     `bson:"created_at" json:"created_at"`
	Updated_at time.Time     `bson:"updated_at" json:"updated_at"`
}
