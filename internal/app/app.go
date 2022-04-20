package app

import (
	"github.com/DanilKl4/crud-gorm/internal/handlers"
	"github.com/DanilKl4/crud-gorm/internal/repository"
	"github.com/labstack/echo/v4"
)

func Run(port string) error {
	repo, err := repository.New()
	if err != nil {
		return err
	}
	hand := handlers.New(repo)

	e := echo.New()

	e.GET("/message/:id", hand.Get)
	e.POST("/message", hand.Create)
	e.PUT("/message/:id", hand.Update)
	e.DELETE("/message/:id", hand.Delete)

	return e.Start(":" + port)
}
