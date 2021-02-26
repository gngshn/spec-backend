package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type RegField struct {
	Msb         string `json:"msb" bson:"msb"`
	Lsb         string `json:"lsb" bson:"lsb"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Access      string `json:"access" bson:"access"`
	Reset       string `json:"reset" bson:"reset"`
}

type Register struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Offset      string             `json:"offset" bson:"offset"`
	Fields      []RegField         `json:"field" bson:"field"`
}
