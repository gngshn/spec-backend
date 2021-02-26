package model

import (
	"github.com/gngshn/spec-backend/model/dao"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mod struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	RegOffset   string             `json:"regOffset" bson:"regOffset"`
	Chip        primitive.ObjectID `json:"chip" bson:"chip"`
	Parent      primitive.ObjectID `json:"parent,omitempty" bson:"parent,omitempty"`
}

func (mod *Mod) GetColl() *qmgo.Collection {
	return dao.GetDB().Collection("mods")
}

func (mod *Mod) GetID() primitive.ObjectID {
	return mod.ID
}

func (mod *Mod) SetID(id primitive.ObjectID) {
	mod.ID = id
}
