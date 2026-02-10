package domain

import (
	"time"
)

type Review struct {
	ID        string     `json:"_id"`
	UserId    string     `json:"user_id"`
	Username  string     `json:"username"`
	GameId    string     `json:"game_id"`
	Rating    *int       `json:"rating" binding:"required"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
