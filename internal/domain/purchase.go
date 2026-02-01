package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Purchase struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserId      bson.ObjectID `bson:"user_id" json:"user_id" binding:"required"`
	TotalAmount float32       `bson:"total_amount" json:"total_amount"`
	Items       []Item        `bson:"items" json:"items" binding:"required,min=1"`
	Payment     Payment       `bson:"payment" json:"payment"`
}

type Item struct {
	GameId  bson.ObjectID `bson:"game_id" json:"game_id"`
	PriceAt float32       `bson:"price_at" json:"price_at"`
}

type Payment struct {
	PaymentMethod string    `bson:"payment_method" json:"payment_method"`
	PaymentStatus string    `bson:"payment_status" json:"payment_status"`
	PaidAt        time.Time `bson:"paid_at" json:"paid_at"`
}
