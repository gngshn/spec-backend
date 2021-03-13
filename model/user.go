package model

import (
	"context"
	"net/http"

	"github.com/gngshn/spec-backend/model/dao"
	"github.com/labstack/echo/v4"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const DefaultPassword = "12345678"

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

func (user *User) CheckRefine(isCreate bool) error {
	var err error
	if user.Username == "" || user.RealPassword == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Username and password can not be empty")
	}
	if isCreate {
		userFind := new(User)
		err = user.GetColl().Find(context.Background(), bson.M{"username": user.Username}).One(userFind)
		if err == nil {
			return echo.NewHTTPError(http.StatusConflict, "Username has been exist")
		}
	}
	user.Password, err = bcrypt.GenerateFromPassword([]byte(user.RealPassword), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Create user fail")
	}
	if user.RealPassword == DefaultPassword {
		user.NeedChangePassword = true
	}
	if user.Username == "admin" {
		user.Admin = true
	} else {
		user.Admin = false
	}
	return nil
}
