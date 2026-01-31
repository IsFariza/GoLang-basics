package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Game struct {
	ID             bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	PublisherId    bson.ObjectID `bson:"publisher_id" json:"publisher_id" binding:"required"`
	DeveloperId    bson.ObjectID `bson:"developer_id" json:"developer_id" binding:"required"`
	EmulationId    bson.ObjectID `bson:"emulation_id" json:"emulation_id"`
	OriginalSystem string        `bson:"original_system" json:"original_system" binding:"omitempty,min=1,max=250"`
	Title          string        `bson:"title" json:"title" binding:"required,min=2,max=100"`
	Description    string        `bson:"description" json:"description" binding:"max=1000"`
	ReleaseDate    *time.Time    `bson:"release_date" json:"release_date"`
	Price          *float32      `bson:"price" json:"price" binding:"omitempty,min=0"`
	IsVerified     bool          `bson:"is_verified" json:"is_verified"`
	Category       []string      `bson:"category" json:"category" binding:"omitempty,min=1,dive,min=1"`
	CreatedAt      time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt      *time.Time    `bson:"updated_at" json:"updated_at"`
}
