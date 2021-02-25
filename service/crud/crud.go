package crud

import (
	"context"

	"github.com/gngshn/spec-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ctx context.Context = context.Background()

func Create(crud model.Crud) error {
	crud.SetID(primitive.NilObjectID)
	res, err := crud.GetColl().InsertOne(ctx, crud)
	if err != nil {
		return err
	}
	crud.SetID(res.InsertedID.(primitive.ObjectID))
	return nil
}

func Count(crud model.Crud) int64 {
	count, err := crud.GetColl().Find(ctx, bson.M{}).Count()
	if err != nil {
		return 0
	}
	return count
}

func FindSome(crud model.Crud, skip int64, limit int64, data interface{}) error {
	err := crud.GetColl().Find(ctx, bson.M{}).Skip(skip).Limit(limit).All(data)
	if err != nil {
		return err
	}
	return nil
}

func FindOne(crud model.Crud) error {
	id := crud.GetID()
	err := crud.GetColl().Find(ctx, bson.M{"_id": id}).One(crud)
	if err != nil {
		return err
	}
	return nil
}

func UpdateOne(crud model.Crud) error {
	return crud.GetColl().UpdateOne(ctx, bson.M{"_id": crud.GetID()}, bson.M{"$set": crud})
}

func DeleteOne(crud model.Crud) error {
	return crud.GetColl().RemoveId(ctx, crud.GetID())
}
