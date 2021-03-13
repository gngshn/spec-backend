package model

import (
	"github.com/gngshn/spec-backend/model/dao"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chip struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
}

func (chip *Chip) GetColl() *qmgo.Collection {
	return dao.GetDB().Collection("chips")
}

func (chip *Chip) GetID() primitive.ObjectID {
	return chip.ID
}

func (chip *Chip) SetID(id primitive.ObjectID) {
	chip.ID = id
}

func (chip *Chip) CheckRefine(_ bool) error {
	return nil
}
