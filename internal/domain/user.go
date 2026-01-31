package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	Username  string        `bson:"username" json:"username" binding:"required"`
	Email     string        `bson:"email" json:"email" binding:"required"`
	Password  []byte        `bson:"password_hash" json:"-" binding:"required"`
	Role      string        `bson:"role,omitempty" json:"role"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt *time.Time    `bson:"updated_at" json:"updated_at"`
}
