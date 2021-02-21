package controller

import (
	"github.com/labstack/echo/v4"
)

func AddController(e *echo.Echo) {
	admin := e.Group("/api/v1/admin")
	addChips(admin)
}
