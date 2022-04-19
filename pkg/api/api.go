package api

import (
	"fmt"
	"github.com/DanilKl4/crud-gorm/pkg/db"
	"github.com/DanilKl4/crud-gorm/pkg/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func Get(c echo.Context) error {
	id := c.Param("id")
	val, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	m, e := db.GetMessage(val)
	if e != nil {
		return c.String(http.StatusBadRequest, e.Error())
	}
	return c.String(http.StatusOK, "name: "+m.Name+" : "+m.Message)
}

func Create(c echo.Context) error {
	name := c.QueryParam("name")
	message := c.QueryParam("message")
	if name == "" || message == "" {
		return c.String(http.StatusBadRequest, "not enough arguments")
	}
	m := models.Message{
		Name:    name,
		Message: message,
	}
	res, err := db.CreateMessage(m)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, fmt.Sprintf("created rows: %v", res))
}

func Update(c echo.Context) error {
	id := c.Param("id")
	val, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	name := c.QueryParam("name")
	message := c.QueryParam("message")
	if name == "" && message == "" {
		return c.String(http.StatusBadRequest, "not enough arguments")
	}
	m := models.Message{
		Name:    name,
		Message: message,
	}
	res, e := db.UpdateMessage(val, m)
	if e != nil {
		return c.String(http.StatusBadRequest, e.Error())
	}
	return c.String(http.StatusOK, fmt.Sprintf("updated rows: %v", res))
}

func Delete(c echo.Context) error {
	id := c.Param("id")
	val, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	res, e := db.DeleteMessage(val)
	if e != nil {
		return c.String(http.StatusBadRequest, e.Error())
	}
	return c.String(http.StatusOK, fmt.Sprintf("deleted rows: %v", res))
}
