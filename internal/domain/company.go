package domain

import (
	"time"
)

type Company struct {
	ID          string     `json:"_id"`
	Name        string     `json:"name" binding:"required"`
	Description string     `json:"description"`
	Country     string     `json:"country"`
	Contacts    Contacts   `json:"contacts"`
	IsVerified  bool       `json:"is_verified"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type Contacts struct {
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
}
