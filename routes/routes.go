package routes

import (
	"interview-bank/database"
	"interview-bank/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures and returns the Gin router
func SetupRouter() *gin.Engine {
	// Initialize the database connection
	db, err := database.InitDB()
	if err != nil {
		panic("Error initializing database: " + err.Error())
	}
	defer db.Close()
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome To This Website")
	})
	// Authentication routes
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/signup", handlers.Signup)
		authGroup.POST("/login", handlers.Login)
	}

	// Interview question routes
	interviewGroup := router.Group("/interviews")
	{
		interviewGroup.POST("/create", handlers.CreateInterviewQuestionHandler(db))
		interviewGroup.PUT("/update", handlers.UpdateInterviewQuestionsHandler(db))
		interviewGroup.GET("/get", handlers.GetInterviewQuestionsHandler(db))
		interviewGroup.DELETE("/delete/:user_id", handlers.DeleteInterviewQuestionsHandler(db))
		interviewGroup.GET("/search", handlers.GetAllInterviewQuestionsWithSearchHandler(db))
	}

	return router
}
