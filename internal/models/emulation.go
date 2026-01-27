package model

import "go.mongodb.org/mongo-driver/v2/bson"

type Emulation struct {
	ID   bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name string        `bson:"name" json:"name"`
}
