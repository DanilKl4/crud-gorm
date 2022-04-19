package main

import (
	"github.com/DanilKl4/crud-gorm/pkg/api"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/message/:id", api.Get)
	e.POST("/message", api.Create)
	e.PUT("/message/:id", api.Update)
	e.DELETE("/message/:id", api.Delete)
	e.Logger.Fatal(e.Start(":5000"))
}
