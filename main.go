package main

import (
	"github.com/gngshn/spec-backend/controller"
	"github.com/gngshn/spec-backend/model/dao"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	mgo := dao.GetMgo()
	defer mgo.MgoClose()

	e := echo.New()
	e.Use(middleware.CORS())
	controller.AddController(e)
	e.Start(":8080")
}
