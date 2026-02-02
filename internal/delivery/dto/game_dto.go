package dto

import (
	"time"
)

type UpdateGameRequest struct {
	PublisherId    *string    `json:"publisher_id"`
	DeveloperId    *string    `json:"developer_id"`
	EmulationId    *string    `json:"emulation_id"`
	OriginalSystem *string    `json:"original_system" binding:"omitempty,min=1,max=250"`
	Title          *string    `json:"title" binding:"omitempty,min=2,max=100"`
	Description    *string    `json:"description" binding:"max=1000"`
	ReleaseDate    *time.Time `json:"release_date"`
	Price          *float32   `json:"price" binding:"omitempty,min=0"`
	Category       []string   `json:"category" binding:"omitempty,min=1,dive,min=1"`
	UpdatedAt      *time.Time `json:"updated_at"`
}
