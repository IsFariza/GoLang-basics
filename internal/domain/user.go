package domain

import (
	"time"
)

type User struct {
	ID           string     `json:"_id"`
	Username     string     `json:"username" binding:"required"`
	Email        string     `json:"email" binding:"required,email"`
	PasswordHash []byte     `json:"-"`
	Role         string     `json:"role,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	Library      []UserGame `json:"library"`
}

type UserGame struct {
	GameId        string    `json:"game_id"`
	AddedAt       time.Time `json:"added_at"`
	PlaytimeHours float32   `json:"playtime_hours"`
}
