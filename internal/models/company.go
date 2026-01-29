package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Company struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
	Country     string        `bson:"country" json:"country"`
	Contacts    Contacts      `bson:"contacts" json:"contacts"`
}

type Contacts struct {
	Email   string `bson:"email" json:"email"`
	Phone   string `bson:"phone" json:"phone"`
	Website string `bson:"website" json:"website"`
}
