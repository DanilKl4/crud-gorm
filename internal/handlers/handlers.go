package handlers

import (
	"fmt"
	"github.com/DanilKl4/crud-gorm/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func New(r *repository.Repository) *Handlers {
	return &Handlers{r}
}

type Handlers struct {
	r *repository.Repository
}

func (h *Handlers) Get(c echo.Context) error {
	val, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	m, err := h.r.GetMessage(val)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "name: "+m.Name+" : "+m.Message)
}

func (h *Handlers) Create(c echo.Context) error {
	m, err := getVars(c.QueryParam("name"), c.QueryParam("message"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	res, err := h.r.CreateMessage(m)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, fmt.Sprintf("created rows: %v", res))
}

func (h *Handlers) Update(c echo.Context) error {
	val, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	m, err := getVars(c.QueryParam("name"), c.QueryParam("message"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	res, e := h.r.UpdateMessage(val, m)
	if e != nil {
		return c.String(http.StatusBadRequest, e.Error())
	}

	return c.String(http.StatusOK, fmt.Sprintf("updated rows: %v", res))
}

func (h *Handlers) Delete(c echo.Context) error {
	val, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	res, e := h.r.DeleteMessage(val)
	if e != nil {
		return c.String(http.StatusBadRequest, e.Error())
	}

	return c.String(http.StatusOK, fmt.Sprintf("deleted rows: %v", res))
}
