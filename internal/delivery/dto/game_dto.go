package dto

import (
	"time"
)

type UpdateGameRequest struct {
	PublisherId    *string    `bson:"publisher_id" json:"publisher_id" binding:"required"`
	DeveloperId    *string    `bson:"developer_id" json:"developer_id" binding:"required"`
	EmulationId    *string    `bson:"emulation_id" json:"emulation_id"`
	OriginalSystem *string    `bson:"original_system" json:"original_system" binding:"omitempty,min=1,max=250"`
	Title          *string    `bson:"title" json:"title" binding:"required,min=2,max=100"`
	Description    *string    `bson:"description" json:"description" binding:"max=1000"`
	ReleaseDate    *time.Time `bson:"release_date" json:"release_date"`
	Price          *float32   `bson:"price" json:"price" binding:"required,min=0"`
	Category       []string   `bson:"category" json:"category" binding:"omitempty,min=1,dive,min=1"`
	UpdatedAt      *time.Time `bson:"updated_at" json:"updated_at"`
}
