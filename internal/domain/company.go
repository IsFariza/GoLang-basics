package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Company struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name        string        `bson:"name" json:"name" binding:"required"`
	Description string        `bson:"description" json:"description"`
	Country     string        `bson:"country" json:"country"`
	Contacts    Contacts      `bson:"contacts" json:"contacts"`
	IsVerified  bool          `bson:"is_verified" json:"is_verified"`
	CreatedAt   time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt   *time.Time    `bson:"updated_at" json:"updated_at"`
}

type Contacts struct {
	Email   string `bson:"email" json:"email"`
	Phone   string `bson:"phone" json:"phone"`
	Website string `bson:"website" json:"website"`
}
