package handlers

import (
	"interview-bank/database"
	"interview-bank/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateInterviewQuestionHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var questions []models.InterviewQuestion
		if err := c.ShouldBindJSON(&questions); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}
		if err := database.CreateInterviewQuestion(db, questions); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create interview questions"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Interview questions created successfully"})
	}
}

func UpdateInterviewQuestionsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var questions []models.InterviewQuestion
		if err := c.ShouldBindJSON(&questions); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}
		if err := database.UpadteInterviewQuestions(db, questions); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update interview questions"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Interview questions updated successfully"})
	}
}

func GetInterviewQuestionsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user information from the request context, assuming it is set during the authentication process.
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}

		questions, err := database.GetInterviewQuestions(db, userID.(int))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve interview questions"})
			return
		}

		var responseQuestions []gin.H
		for _, q := range questions {
			responseQuestions = append(responseQuestions, gin.H{
				"date":         q.Date,
				"company":      q.Company,
				"level":        q.Level,
				"question":     q.Question,
				"satisfaction": q.Satisfaction,
			})
		}

		c.JSON(http.StatusOK, responseQuestions)
	}
}

func DeleteInterviewQuestionsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid structure"})
			return
		}
		var questions []string
		if err := c.ShouldBindJSON(&questions); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := database.DeleteInterviewQuestions(db, userID, questions); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Interview questions deleted successfully"})
	}
}

func GetAllInterviewQuestionsWithSearchHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}

		company := c.Query("company")
		level := c.Query("level")

		questions, err := database.GetAllInterviewQuestionsWithSearch(db, userID.(uint), company, level)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve interview questions"})
			return
		}

		var responseQuestions []gin.H
		for _, q := range questions {
			responseQuestions = append(responseQuestions, gin.H{
				"date":         q.Date,
				"company":      q.Company,
				"level":        q.Level,
				"question":     q.Question,
				"satisfaction": q.Satisfaction,
			})
		}

		c.JSON(http.StatusOK, responseQuestions)
	}
}
