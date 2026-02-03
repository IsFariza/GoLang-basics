package dto

import "time"

type UserLibraryItemDTO struct {
	GameId        string    `json:"game_id"`
	Title         string    `json:"title"`
	AddedAt       time.Time `json:"added_at"`
	PlaytimeHours float32   `json:"playtime_hours"`
}
