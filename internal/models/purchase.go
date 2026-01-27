package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Purchase struct {
	ID            bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	User_id       bson.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	Purchase_date time.Time     `bson:"purchase_date" json:"purchase_date"`
	Total_amount  float32       `bson:"total_amount" json:"total_amount"`
	Items         []Item        `bson:"items" json:"items"`
	Payment       Payment       `bson:"payment" json:"payment"`
}

type Item struct {
	Game_id  bson.ObjectID `bson:"game_id" json:"game_id"`
	Price_at float32       `bson:"price_at" json:"price_at"`
}

type Payment struct {
	Payment_method string    `bson:"payment_method" json:"payment_method"`
	Payment_status string    `bson:"payment_status" json:"payment_status"`
	Paid_at        time.Time `bson:"paid_at" json:"paid_at"`
}
