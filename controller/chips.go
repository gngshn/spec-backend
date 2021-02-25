package controller

import (
	"net/http"
	"strconv"

	"github.com/gngshn/spec-backend/model"
	"github.com/gngshn/spec-backend/service/crud"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createChip(c echo.Context) error {
	chip := new(model.Chip)
	err := c.Bind(chip)
	if err != nil {
		return err
	}
	err = crud.Create(chip)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, chip)
}

func findSomeChips(c echo.Context) error {
	skip, _ := strconv.ParseInt(c.QueryParam("skip"), 0, 64)
	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 0, 64)
	if limit > 10 || limit == 0 {
		limit = 10
	}
	total := crud.Count(&model.Chip{})
	if skip >= total {
		return c.JSON(http.StatusOK, createPagination(total, skip, limit, []model.Chip{}))
	}
	chips := []model.Chip{}
	err := crud.FindSome(&model.Chip{}, skip, limit, &chips)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, createPagination(total, skip, limit, chips))
}

func findOneChip(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return err
	}
	chip := &model.Chip{ID: id}
	err = crud.FindOne(chip)
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
	chip.ID = id
	err = crud.UpdateOne(chip)
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
	chip := &model.Chip{ID: id}
	err = crud.DeleteOne(chip)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

func addChips(g *echo.Group) {
	g.POST("/chips", createChip)
	g.GET("/chips", findSomeChips)
	g.GET("/chips/:id", findOneChip)
	g.PUT("/chips/:id", updateOneChip)
	g.DELETE("/chips/:id", deleteOneChip)
}
