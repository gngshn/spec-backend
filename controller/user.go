package controller

import (
	"net/http"

	"github.com/gngshn/spec-backend/model"
	"github.com/gngshn/spec-backend/service"
	"github.com/labstack/echo/v4"
)

func login(c echo.Context) error {
	user := new(model.User)
	err := c.Bind(user)
	if err != nil {
		return err
	}
	token, err := service.GetLoginToken(user)
	if err != nil {
		return err
	}
	if user.NeedChangePassword {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"token":              "",
			"needChangePassword": true,
		})
	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"token":              token,
			"needChangePassword": false,
		})
	}
}

func changePassword(c echo.Context) error {
	changePasswordDto := new(model.ChangePasswordDto)
	err := c.Bind(changePasswordDto)
	if err != nil {
		return err
	}
	err = service.ChangePassword(changePasswordDto)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "successs",
	})
}

func createUser(c echo.Context) error {
	return create(c, "user")
}

func findSomeUsers(c echo.Context) error {
	return findSome(c, "user")
}

func findOneUser(c echo.Context) error {
	return findOne(c, "user")
}

func updateOneUser(c echo.Context) error {
	return updateOne(c, "user")
}

func deleteOneUser(c echo.Context) error {
	return deleteOne(c, "user")
}

func AdminCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := service.CheckAdmin(c); err != nil {
			return err
		}
		return next(c)
	}
}

func addUsers(g *echo.Group) {
	users := g.Group("/users")
	users.Use(AdminCheck)
	users.POST("", createUser)
	users.GET("", findSomeUsers)
	users.GET("/:id", findOneUser)
	users.PUT("/:id", updateOneUser)
	users.DELETE("/:id", deleteOneUser)
}
