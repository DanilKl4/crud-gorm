package repository

import (
	"errors"
	"github.com/DanilKl4/crud-gorm/internal/models"
	"github.com/DanilKl4/crud-gorm/pkg/db"
	"gorm.io/gorm"
)

const (
	NotCreated = "message not created"
	NotFound   = "message not found"
	NotUpdate  = "message not update"
	NotDelete  = "message not delete"
)

func New() (*Repository, error) {
	db, err := db.Init()
	if err != nil {
		return nil, err
	}
	return &Repository{db}, nil
}

type Repository struct {
	db *gorm.DB
}

func (r *Repository) CreateMessage(m models.Message) (int64, error) {
	result := r.db.Create(&m)

	if result.RowsAffected == 0 {
		return 0, errors.New(NotCreated)
	}

	return result.RowsAffected, nil
}

func (r *Repository) GetMessage(id int) (models.Message, error) {
	var message models.Message

	if r.db.Find(&message, id).RowsAffected == 0 {
		return models.Message{}, errors.New(NotFound)
	}

	return message, nil
}

func (r *Repository) UpdateMessage(id int, m models.Message) (int64, error) {
	var message models.Message

	result := r.db.Model(&message).Where("id = ?", id).Updates(m)

	if result.RowsAffected == 0 {
		return 0, errors.New(NotUpdate)
	}

	return result.RowsAffected, nil
}

func (r *Repository) DeleteMessage(id int) (int64, error) {
	var message models.Message

	result := r.db.Where("id = ?", id).Delete(&message)

	if result.RowsAffected == 0 {
		return 0, errors.New(NotDelete)
	}

	return result.RowsAffected, nil
}
