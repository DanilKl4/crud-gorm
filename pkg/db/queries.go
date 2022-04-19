package db

import (
	"errors"
	"github.com/DanilKl4/crud-gorm/pkg/models"
)

func CreateMessage(m models.Message) (int64, error) {
	db, err := connect()
	if err != nil {
		return 0, err
	}
	result := db.Create(&m)
	if result.RowsAffected == 0 {
		return 0, errors.New("message not created")
	}
	return result.RowsAffected, nil
}

func GetMessage(id int) (models.Message, error) {
	db, err := connect()
	if err != nil {
		return models.Message{}, err
	}
	var message models.Message
	result := db.Find(&message, id)
	if result.RowsAffected == 0 {
		return models.Message{}, errors.New("message not found")
	}
	return message, nil
}

func UpdateMessage(id int, m models.Message) (int64, error) {
	db, err := connect()
	if err != nil {
		return 0, err
	}
	var message models.Message
	result := db.Model(&message).Where("id = ?", id).Updates(m)
	if result.RowsAffected == 0 {
		return 0, errors.New("message not update")
	}
	return result.RowsAffected, nil
}

func DeleteMessage(id int) (int64, error) {
	db, err := connect()
	if err != nil {
		return 0, err
	}
	var message models.Message
	result := db.Where("id = ?", id).Delete(&message)
	if result.RowsAffected == 0 {
		return 0, errors.New("message not delete")
	}
	return result.RowsAffected, nil
}
