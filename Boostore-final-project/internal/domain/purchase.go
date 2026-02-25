package domain

import (
	"time"
)

type Purchase struct {
	ID          string  `json:"_id"`
	UserId      string  `json:"user_id" binding:"required"`
	TotalAmount float32 `json:"total_amount"`
	Items       []Item  `json:"items" binding:"required,min=1"`
	Payment     Payment `json:"payment"`
}

type Item struct {
	GameId  string  `json:"game_id"`
	PriceAt float32 `json:"price_at"`
}

type Payment struct {
	PaymentMethod string    `json:"payment_method"`
	PaymentStatus string    `json:"payment_status"`
	PaidAt        time.Time `json:"paid_at"`
}
