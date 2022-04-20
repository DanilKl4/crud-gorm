package db

import (
	"fmt"
	"github.com/DanilKl4/crud-gorm/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func Init() (*gorm.DB, error) {
	con, err := getEnv()
	if err != nil {
		return nil, err
	}

	db, err := connect(con)
	if err != nil {
		return nil, err
	}

	db, err = migrate(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getEnv() (string, error) {
	err := godotenv.Load("C:\\Users\\danii\\crud-gorm\\configs\\variables.env")
	if err != nil {
		return "", err
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv("USER"),
		os.Getenv("PASS"),
		os.Getenv("HOST"),
		os.Getenv("NAME"),
	)

	return connectionString, err
}

func connect(con string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(con), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func migrate(db *gorm.DB) (*gorm.DB, error) {
	err := db.AutoMigrate(&models.Message{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
