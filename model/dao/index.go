package dao

import (
	"context"
	"log"
	"sync"

	"github.com/qiniu/qmgo"
)

type Mgo struct {
	cli  *qmgo.Client
	db   *qmgo.Database
	coll *qmgo.Collection
}

var mgo *Mgo
var once sync.Once

func GetMgo() *Mgo {
	once.Do(func() {
		mgoOpen()
	})
	return mgo
}

func GetDB() *qmgo.Database {
	return GetMgo().db
}

func mgoOpen() (*Mgo, error) {
	client, err := qmgo.NewClient(context.Background(),
		&qmgo.Config{Uri: "mongodb://localhost:27017"})
	if err != nil {
		return nil, err
	}
	mgo = new(Mgo)
	mgo.cli = client
	mgo.db = client.Database("spec")
	mgo.coll = mgo.db.Collection("chips")
	return mgo, nil
}

func (mgo *Mgo) MgoClose() {
	if err := mgo.cli.Close(context.Background()); err != nil {
		log.Fatal("Close mongodb fail")
	}
}
