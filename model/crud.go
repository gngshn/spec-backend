package model

import (
	"errors"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Crud interface {
	GetColl() *qmgo.Collection
	GetID() primitive.ObjectID
	SetID(id primitive.ObjectID)
}

func CreateCrud(name string) (Crud, error) {
	switch name {
	case "chip":
		return new(Chip), nil
	case "mod":
		return new(Mod), nil
	default:
		return nil, errors.New("No such model exist")
	}
}

func CreateCruds(name string) (interface{}, error) {
	switch name {
	case "chip":
		return []Chip{}, nil
	case "mod":
		return []Mod{}, nil
	default:
		return nil, errors.New("No such model exist")
	}
}
