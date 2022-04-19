package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Id      int `gorm:"primaryKey;autoIncrement:true"`
	Name    string
	Message string
}
