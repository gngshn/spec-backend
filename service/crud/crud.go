package crud

import (
	"context"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Crud interface {
	GetColl() *qmgo.Collection
	GetID() primitive.ObjectID
	SetID(id primitive.ObjectID)
}

var ctx context.Context = context.Background()

func Create(crud Crud) error {
	crud.SetID(primitive.NilObjectID)
	res, err := crud.GetColl().InsertOne(ctx, crud)
	if err != nil {
		return err
	}
	crud.SetID(res.InsertedID.(primitive.ObjectID))
	return nil
}

func Count(crud Crud) int64 {
	count, err := crud.GetColl().Find(ctx, bson.M{}).Count()
	if err != nil {
		return 0
	}
	return count
}

func FindSome(crud Crud, skip int64, limit int64, data interface{}) error {
	err := crud.GetColl().Find(ctx, bson.M{}).Skip(skip).Limit(limit).All(data)
	if err != nil {
		return err
	}
	return nil
}

func FindOne(crud Crud) error {
	id := crud.GetID()
	err := crud.GetColl().Find(ctx, bson.M{"_id": id}).One(crud)
	if err != nil {
		return err
	}
	return nil
}

func UpdateOne(crud Crud) error {
	return crud.GetColl().UpdateOne(ctx, bson.M{"_id": crud.GetID()}, bson.M{"$set": crud})
}

func DeleteOne(crud Crud) error {
	return crud.GetColl().Remove(ctx, bson.M{"_id": crud.GetID()})
}
