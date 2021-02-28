package model

import (
	"github.com/gngshn/spec-backend/model/dao"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegField struct {
	Bits        [2]uint8 `json:"bits" bson:"bits"`
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	Access      string   `json:"access" bson:"access"`
	Reset       string   `json:"reset" bson:"reset"`
}

type Register struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Offset      uint32             `json:"offset" bson:"offset"`
	Parent      primitive.ObjectID `json:"parent" bson:"parent"`
	Fields      []RegField         `json:"fields" bson:"fields"`
}

func (register *Register) GetColl() *qmgo.Collection {
	return dao.GetDB().Collection("registers")
}

func (register *Register) GetID() primitive.ObjectID {
	return register.ID
}

func (register *Register) SetID(id primitive.ObjectID) {
	register.ID = id
}
