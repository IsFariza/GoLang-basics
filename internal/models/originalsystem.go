package model

import "go.mongodb.org/mongo-driver/v2/bson"

type OriginalSystem struct {
	ID   bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name string        `bson:"name" jsom:"name"`
}
