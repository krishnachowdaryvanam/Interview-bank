package database

import (
	"interview-bank/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
)

func InitDB() (*gorm.DB, error) {
	conn := "host=localhost port=5432 user=postgres password=postgres dbname=Personal sslmode=disable"
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{}, &models.InterviewQuestion{})
	logrus.Info("Successfully connected to database")
	return db, nil
}
