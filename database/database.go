package database

import (
	"interview-bank/models"

	"github.com/jinzhu/gorm"
)

func CreateInterviewQuestion(db *gorm.DB, questions []models.InterviewQuestion) error {
	err := db.Create(&questions).Error
	return err
}

func GetInterviewQuestions(db *gorm.DB, userID int) ([]models.InterviewQuestion, error) {
	var questions []models.InterviewQuestion
	err := db.Table("interview_questions").
		Select("date, company, level, question, satisfaction").
		Where("user_id = ?", userID).
		Find(&questions).Error

	return questions, err
}

func UpadteInterviewQuestions(db *gorm.DB, questions []models.InterviewQuestion) error {
	err := db.Save(&questions).Error
	return err
}

func DeleteInterviewQuestions(db *gorm.DB, UserID int, questions []string) error {
	err := db.Where("user_id = ?", UserID).Delete(&models.InterviewQuestion{}, questions).Error
	return err
}

func GetAllInterviewQuestionsWithSearch(db *gorm.DB, userID uint, company, level string) ([]models.InterviewQuestion, error) {
	var questions []models.InterviewQuestion

	query := db.Where("user_id = ?", userID)
	if company != "" {
		query = query.Where("company LIKE ?", "%"+company+"%")
	}
	if level != "" {
		query = query.Where("level = ?", level)
	}

	err := query.Find(&questions).Error
	return questions, err
}
