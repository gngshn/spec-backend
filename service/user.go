package service

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gngshn/spec-backend/model"
	"github.com/gngshn/spec-backend/service/crud"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

const defaultPassword = "123456"

/* need move to config */
var Secret = []byte("sl-dkfj#$sd#$jfs43#$#")

func CheckAdmin(c echo.Context) error {
	claims := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)
	if !claims["admin"].(bool) {
		return echo.NewHTTPError(http.StatusUnauthorized, "Only admin allowed")
	}
	return nil
}

func GetLoginToken(user *model.User) (string, error) {
	if user.Username == "" || user.RealPassword == "" {
		return "", echo.NewHTTPError(http.StatusBadRequest, "Username and password can not be empty")
	}
	userFind := new(model.User)
	err := user.GetColl().Find(crud.Ctx, bson.M{"username": user.Username}).One(userFind)
	if err != nil {
		if user.Username != "admin" {
			return "", echo.NewHTTPError(http.StatusBadRequest, "Username or password is incorrent")
		}
		err = CreateUser(user)
		if err != nil {
			return "", echo.NewHTTPError(http.StatusInternalServerError, "Login fail")
		}
		userFind = user
	} else {
		user.NeedChangePassword = userFind.NeedChangePassword
	}
	if user.NeedChangePassword {
		return "", nil
	}
	err = bcrypt.CompareHashAndPassword([]byte(userFind.Password), []byte(user.RealPassword))
	if err != nil {
		return "", echo.NewHTTPError(http.StatusBadRequest, "Username or password is incorrent")
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = userFind.Username
	claims["admin"] = userFind.Admin
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix()

	return token.SignedString(Secret)
}

func ChangePassword(changePasswordDto *model.ChangePasswordDto) error {
	userFind := new(model.User)
	err := userFind.GetColl().Find(crud.Ctx, bson.M{"username": changePasswordDto.Username}).One(userFind)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Username or password is error")
	}
	err = bcrypt.CompareHashAndPassword([]byte(userFind.Password), []byte(changePasswordDto.OldPassword))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Username or password is error")
	}
	userFind.Password, err = bcrypt.GenerateFromPassword([]byte(changePasswordDto.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "update password fail")
	}
	if userFind.RealPassword == defaultPassword {
		userFind.NeedChangePassword = true
	} else {
		userFind.NeedChangePassword = false
	}
	crud.UpdateOne(userFind)

	return nil
}

func CreateUser(user *model.User) error {
	if user.Username == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Username can not be empty")
	}
	userFind := new(model.User)
	err := user.GetColl().Find(crud.Ctx, bson.M{"username": user.Username}).One(userFind)
	if err == nil {
		return echo.NewHTTPError(http.StatusConflict, "Username has been exist")
	}
	user.Password, err = bcrypt.GenerateFromPassword([]byte(user.RealPassword), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Create user fail")
	}
	if user.RealPassword == defaultPassword {
		user.NeedChangePassword = true
	}
	if user.Username == "admin" {
		user.Admin = true
	} else {
		user.Admin = false
	}

	return crud.Create(user)
}

func UpdateOneUser(user *model.User) error {
	if user.RealPassword != defaultPassword {
		user.NeedChangePassword = true
	}
	return crud.UpdateOne(user)
}
