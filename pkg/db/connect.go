package db

import (
	"fmt"
	"github.com/DanilKl4/crud-gorm/pkg/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func connect() (*gorm.DB, error) {
	er := godotenv.Load("C:\\Users\\danii\\crud-gorm\\configs\\variables.env")
	if er != nil {
		return nil, er
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv("USER"),
		os.Getenv("PASS"),
		os.Getenv("HOST"),
		os.Getenv("NAME"),
	)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.Message{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
