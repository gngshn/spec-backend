package controller

import (
	"net/http"
	"strings"

	"github.com/gngshn/spec-backend/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getJWT() echo.MiddlewareFunc {
	jwtConfig := middleware.DefaultJWTConfig
	jwtConfig.ErrorHandler = func(err error) error {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login!")
	}
	jwtConfig.Skipper = func(c echo.Context) bool {
		if strings.HasSuffix(c.Request().URL.Path, "/login") ||
			strings.HasSuffix(c.Request().URL.Path, "/change-password") {
			return true
		}
		return false
	}
	jwtConfig.SigningKey = service.Secret
	return middleware.JWTWithConfig(jwtConfig)
}

func AddController(e *echo.Echo) {
	admin := e.Group("/api/v1/admin")
	admin.Use(getJWT())
	admin.POST("/login", login)
	admin.POST("/change-password", changePassword)
	addResources(admin)
	addUsers(admin)
}
