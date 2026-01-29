package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Game struct {
	ID                 bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	Publisher_id       bson.ObjectID `bson:"publisher_id" json:"publisher_id"`
	Developer_id       bson.ObjectID `bson:"developer_id" json:"developer_id"`
	Original_system_id bson.ObjectID `bson:"original_system_id" json:"original_system_id"`
	Emulation_id       bson.ObjectID `bson:"emulation_id" json:"emulation_id"`
	Title              string        `bson:"title" json:"title"`
	Description        string        `bson:"description" json:"description"`
	Release_date       time.Time     `bson:"release_date" json:"release_date"`
	Price              float32       `bson:"price" json:"price"`
	Created_at         time.Time     `bson:"created_at" json:"created_at"`
	Updated_at         time.Time     `bson:"updated_at" json:"updated_at"`
	Category           bson.ObjectID `bson:"category" json:"category"`
	IsVerified         bool          `bson:"is_verified" json:"is_verified"`
}
