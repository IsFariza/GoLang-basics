package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Game struct {
	ID             bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	PublisherId    bson.ObjectID `bson:"publisher_id" json:"publisher_id"`
	DeveloperId    bson.ObjectID `bson:"developer_id" json:"developer_id"`
	EmulationId    bson.ObjectID `bson:"emulation_id" json:"emulation_id"`
	OriginalSystem string        `bson:"original_system" json:"original_system"`
	Title          string        `bson:"title" json:"title" binding:"required"`
	Description    string        `bson:"description" json:"description"`
	ReleaseDate    time.Time     `bson:"release_date" json:"release_date"`
	Price          float32       `bson:"price" json:"price"`
	IsVerified     bool          `bson:"is_verified" json:"is_verified"`
	Category       []string      `bson:"category" json:"category"`
	CreatedAt      time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time     `bson:"updated_at" json:"updated_at"`
}
