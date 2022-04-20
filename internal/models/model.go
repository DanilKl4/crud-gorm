package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Id      int    `gorm:"primaryKey;autoIncrement:true"`
	Name    string `gorm:"size:20;not null"`
	Message string `gorm:"size:200;not null"`
}
