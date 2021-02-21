package service

import (
	"context"
	"sync"

	"github.com/gngshn/spec-backend/model"
	"github.com/gngshn/spec-backend/model/dao"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var chipsColl *qmgo.Collection
var once sync.Once
var ctx context.Context

func getChipsColl() *qmgo.Collection {
	once.Do(func() {
		chipsColl = dao.GetDB().Collection("chips")
		ctx = context.Background()
	})
	return chipsColl
}

func CreateChip(chip *model.Chip) error {
	chip.ID = primitive.NilObjectID
	res, err := getChipsColl().InsertOne(ctx, chip)
	if err != nil {
		return err
	}
	chip.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func FindAllChips(pagination *model.Pagination) (err error) {
	pagination.Total, err = getChipsColl().Find(ctx, bson.M{}).Count()
	if err != nil {
		return
	}
	err = getChipsColl().Find(ctx, bson.M{}).Skip(pagination.Skip).Limit(pagination.Limit).All(&pagination.Data)
	if err != nil {
		return
	}
	return
}

func FindOneChip(id primitive.ObjectID) (*model.Chip, error) {
	chip := new(model.Chip)
	err := getChipsColl().Find(ctx, bson.M{"_id": id}).One(&chip)
	if err != nil {
		return nil, err
	}
	return chip, nil
}

func UpdateOneChip(id primitive.ObjectID, chip *model.Chip) error {
	chip.ID = id
	return getChipsColl().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": chip})
}

func DeleteOneChip(id primitive.ObjectID) error {
	return getChipsColl().Remove(ctx, bson.M{"_id": id})
}
