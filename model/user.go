package model

import (
	"github.com/gngshn/spec-backend/model/dao"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                 primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username           string             `json:"username" bson:"username"`
	RealPassword       string             `json:"password,omitempty" bson:"-"`
	Password           []byte             `json:"-" bson:"password"`
	Admin              bool               `json:"-" bson:"admin"`
	NeedChangePassword bool               `json:"-" bson:"needChangePassword"`
}

type ChangePasswordDto struct {
	Username    string `json:"username"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

func (user *User) GetColl() *qmgo.Collection {
	return dao.GetDB().Collection("users")
}

func (user *User) GetID() primitive.ObjectID {
	return user.ID
}

func (user *User) SetID(id primitive.ObjectID) {
	user.ID = id
}
