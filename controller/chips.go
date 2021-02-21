package controller

import (
	"net/http"
	"strconv"

	"github.com/gngshn/spec-backend/model"
	"github.com/gngshn/spec-backend/service"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createChip(c echo.Context) error {
	chip := new(model.Chip)
	err := c.Bind(chip)
	if err != nil {
		return err
	}
	err = service.CreateChip(chip)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, chip)
}

func findAllChips(c echo.Context) error {
	chips := []model.Chip{}
	pagination := new(model.Pagination)
	pagination.Skip, _ = strconv.ParseInt(c.QueryParam("skip"), 0, 64)
	pagination.Limit, _ = strconv.ParseInt(c.QueryParam("limit"), 0, 64)
	if pagination.Limit > 10 || pagination.Limit == 0 {
		pagination.Limit = 10
	}
	pagination.Data = chips
	err := service.FindAllChips(pagination)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, pagination)
}

func findOneChip(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return err
	}
	chip, err := service.FindOneChip(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, chip)
}

func updateOneChip(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return err
	}
	chip := new(model.Chip)
	err = c.Bind(chip)
	if err != nil {
		return err
	}
	err = service.UpdateOneChip(id, chip)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, chip)
}

func deleteOneChip(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return err
	}
	err = service.DeleteOneChip(id)
	if err != nil {
		return err
	}
	return c.JSONBlob(http.StatusOK, []byte("{status: \"success\"}"))
}

func addChips(g *echo.Group) {
	g.POST("/chips", createChip)
	g.GET("/chips", findAllChips)
	g.GET("/chips/:id", findOneChip)
	g.PUT("/chips/:id", updateOneChip)
	g.DELETE("/chips/:id", deleteOneChip)
}
