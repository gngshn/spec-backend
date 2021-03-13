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

func create(c echo.Context, res string) error {
	resource, err := model.CreateCrud(res)
	if err != nil {
		return err
	}
	err = c.Bind(resource)
	if err != nil {
		return err
	}
	err = resource.CheckRefine(true)
	if err != nil {
		return err
	}
	err = crud.Create(resource)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, resource)

}

func findSome(c echo.Context, res string) error {
	filter := bson.M{}
	var sort = []string{}
	var skip int64 = 0
	var limit int64 = 0
	for k, v := range c.QueryParams() {
		switch k {
		case "$skip":
			skip, _ = strconv.ParseInt(c.QueryParam(k), 0, 64)
		case "$limit":
			limit, _ = strconv.ParseInt(c.QueryParam(k), 0, 64)
			if limit > 1000 || limit == 0 {
				limit = 1000
			}
		case "$sort":
			sort = append(sort, c.QueryParam(k))
		default:
			fid, err := primitive.ObjectIDFromHex(v[0])
			if err == nil {
				if fid == primitive.NilObjectID {
					filter[k] = bson.M{"$exists": false}
				} else {
					filter[k] = fid
				}
			} else {
				filter[k] = v[0]
			}
		}
	}
	resource, err := model.CreateCrud(res)
	if err != nil {
		return err
	}
	resources, err := model.CreateCruds(res)
	if err != nil {
		return err
	}
	total := crud.Count(resource, filter)
	if skip >= total {
		return c.JSON(http.StatusOK, createPagination(total, skip, limit, resources))
	}
	err = crud.FindSome(resource, filter, sort, skip, limit, &resources)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, createPagination(total, skip, limit, resources))
}

func findOne(c echo.Context, res string) error {

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return err
	}
	resource, err := model.CreateCrud(res)
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

func updateOne(c echo.Context, res string) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return err
	}
	resource, err := model.CreateCrud(res)
	if err != nil {
		return err
	}
	err = c.Bind(resource)
	if err != nil {
		return err
	}
	err = resource.CheckRefine(false)
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

func deleteOne(c echo.Context, res string) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return err
	}
	resource, err := model.CreateCrud(res)
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

func getRes(c echo.Context) string {
	resources := c.Param("resources")
	return inflection.Singular(resources)
}

func createResource(c echo.Context) error {
	return create(c, getRes(c))
}

func findSomeResources(c echo.Context) error {
	return findSome(c, getRes(c))
}

func findOneResource(c echo.Context) error {
	return findOne(c, getRes((c)))
}

func updateOneResource(c echo.Context) error {
	return updateOne(c, getRes(c))
}

func deleteOneResource(c echo.Context) error {
	return deleteOne(c, getRes(c))
}

func addResources(g *echo.Group) {
	resource := g.Group("/generic")
	resource.POST("/:resources", createResource)
	resource.GET("/:resources", findSomeResources)
	resource.GET("/:resources/:id", findOneResource)
	resource.PUT("/:resources/:id", updateOneResource)
	resource.DELETE("/:resources/:id", deleteOneResource)
}
