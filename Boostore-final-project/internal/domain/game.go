package domain

import (
	"time"
)

type Game struct {
	ID             string     `json:"_id"`
	PublisherId    string     `json:"publisher_id" binding:"required"`
	DeveloperId    string     `json:"developer_id" binding:"required"`
	EmulationId    string     `json:"emulation_id"`
	UserId         string     `json:"user_id"`
	OriginalSystem string     `json:"original_system" binding:"omitempty,min=1,max=250"`
	Title          string     `json:"title" binding:"required,min=2,max=100"`
	Description    string     `json:"description" binding:"max=1000"`
	ReleaseDate    *time.Time `json:"release_date"`
	Price          *float32   `json:"price" binding:"required,min=0"`
	IsVerified     bool       `json:"is_verified"`
	Category       []string   `json:"category" binding:"omitempty,min=1,dive,min=1"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

type PopulatedGame struct {
	ID             string     `json:"_id"`
	Publisher      *Company   `json:"publisher"`
	Developer      *Company   `json:"developer"`
	Emulation      *Emulation `json:"emulation"`
	OriginalSystem string     `json:"original_system"`
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	ReleaseDate    *time.Time `json:"release_date"`
	Price          *float32   `json:"price"`
	IsVerified     bool       `json:"is_verified"`
	Category       []string   `json:"category"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}
