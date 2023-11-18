package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}

type InterviewQuestion struct {
	gorm.Model
	UserID       uint      `json:"-" gorm:"not null"`
	Company      string    `json:"company" gorm:"not null"`
	Date         time.Time `json:"date" gorm:"not null;type:date"`
	Level        string    `json:"level" gorm:"not null"`
	Question     string    `json:"question" gorm:"not null"`
	Satisfaction int       `json:"satisfaction" gorm:"not null"`
}

// CreateUserRecord creates a user record in the database
func (user *User) CreateUserRecord(db *gorm.DB) error {
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// HashPassword encrypts user password
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword checks user password
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
