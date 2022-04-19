package main

import (
	"github.com/DanilKl4/crud-gorm/pkg/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/message/:id", handlers.Get)
	e.POST("/message", handlers.Create)
	e.PUT("/message/:id", handlers.Update)
	e.DELETE("/message/:id", handlers.Delete)
	e.Logger.Fatal(e.Start(":5000"))
}
