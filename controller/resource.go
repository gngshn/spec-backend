package controller

import (
	"net/http"
	"strconv"

	"github.com/gngshn/spec-backend/model"
	"github.com/gngshn/spec-backend/service/crud"
	"github.com/jinzhu/inflection"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getRes(c echo.Context) string {
	resources := c.Param("resources")
	return inflection.Singular(resources)
}

func createResource(c echo.Context) error {
	resource, err := model.CreateCrud(getRes(c))
	if err != nil {
		return err
	}
	err = c.Bind(resource)
	if err != nil {
		return err
	}
	err = crud.Create(resource)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, resource)
}

func findSomeResources(c echo.Context) error {
	skip, _ := strconv.ParseInt(c.QueryParam("skip"), 0, 64)
	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 0, 64)
	if limit > 100 || limit == 0 {
		limit = 100
	}
	filter := bson.M{}
	ic, err := primitive.ObjectIDFromHex(c.QueryParam(("ic")))
	if err == nil {
		filter["ic"] = ic
	}

	resource, err := model.CreateCrud(getRes(c))
	if err != nil {
		return err
	}
	resources, err := model.CreateCruds(getRes(c))
	if err != nil {
		return err
	}
	total := crud.Count(resource, filter)
	if skip >= total {
		return c.JSON(http.StatusOK, createPagination(total, skip, limit, resources))
	}
	err = crud.FindSome(resource, skip, limit, filter, &resources)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, createPagination(total, skip, limit, resources))
}

func findOneResource(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return err
	}
	resource, err := model.CreateCrud(getRes(c))
	if err != nil {
		return err
	}
	resource.SetID(id)
	err = crud.FindOne(resource)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resource)
}

func updateOneResource(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return err
	}
	resource, err := model.CreateCrud(getRes(c))
	if err != nil {
		return err
	}
	err = c.Bind(resource)
	if err != nil {
		return err
	}
	resource.SetID(id)
	err = crud.UpdateOne(resource)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resource)
}

func deleteOneResource(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return err
	}
	resource, err := model.CreateCrud(getRes(c))
	if err != nil {
		return err
	}
	resource.SetID(id)
	err = crud.DeleteOne(resource)
	if err != nil {
		return err
	}
	return c.JSONBlob(http.StatusOK, []byte(`{"status": "success"}`))
}

func addResources(g *echo.Group) {
	resource := g.Group("/generic")
	resource.POST("/:resources", createResource)
	resource.GET("/:resources", findSomeResources)
	resource.GET("/:resources/:id", findOneResource)
	resource.PUT("/:resources/:id", updateOneResource)
	resource.DELETE("/:resources/:id", deleteOneResource)
}
