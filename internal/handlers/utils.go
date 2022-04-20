package handlers

import (
	"errors"
	"github.com/DanilKl4/crud-gorm/internal/models"
)

func getVars(name string, msg string) (models.Message, error) {
	if name == "" || msg == "" {
		return models.Message{}, errors.New("not enough arguments")
	}

	m := models.Message{
		Name:    name,
		Message: msg,
	}

	return m, nil
}
