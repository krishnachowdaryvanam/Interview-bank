package handlers

import (
	"interview-bank/auth"
	"interview-bank/database"
	"interview-bank/models"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// LoginPayload login body
// LoginPayload is a struct that contains the fields for a user's login credentials
type LoginPayload struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse token response
// LoginResponse is a struct that contains the fields for a user's login response
type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshtoken"`
}

// Signup is a function that handles user signup
func Signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs",
		})
		c.Abort()
		return
	}

	err = user.HashPassword(user.Password)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"Error": "Error Hashing Password",
		})
		c.Abort()
		return
	}

	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Error initializing database:", err)
		c.JSON(500, gin.H{
			"Error": "Error Connecting to Database",
		})
		c.Abort()
		return
	}

	defer db.Close()

	err = user.CreateUserRecord(db)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Error Creating User",
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"Message": "Successfully Registered",
	})
}

// Login is a function that handles user login
func Login(c *gin.Context) {
	var payload LoginPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Inputs"})
		c.Abort()
		return
	}

	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Error initializing database:", err)
		c.JSON(500, gin.H{"error": "Error Connecting to Database"})
		c.Abort()
		return
	}
	defer db.Close()

	var user models.User
	result := db.Where("email = ?", payload.Email).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(401, gin.H{"error": "Invalid User Credentials"})
		c.Abort()
		return
	}

	if err := user.CheckPassword(payload.Password); err != nil {
		log.Println(err)
		c.JSON(401, gin.H{"error": "Invalid User Credentials"})
		c.Abort()
		return
	}

	jwtWrapper := auth.JwtWrapper{
		SecretKey:         "verysecretkey",
		Issuer:            "AuthService",
		ExpirationMinutes: 1,
		ExpirationHours:   12,
	}

	signedToken, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "Error Signing Token"})
		c.Abort()
		return
	}

	signedRefreshToken, err := jwtWrapper.RefreshToken(user.Email)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "Error Signing Token"})
		c.Abort()
		return
	}

	tokenResponse := LoginResponse{
		Token:        signedToken,
		RefreshToken: signedRefreshToken,
	}

	c.JSON(200, tokenResponse)
}
