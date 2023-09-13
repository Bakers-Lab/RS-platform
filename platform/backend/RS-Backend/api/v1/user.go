package v1

import (
	"RS-Backend/config"
	"RS-Backend/models"
	"RS-Backend/services"

	"github.com/gin-gonic/gin"
)

// RegisterPayload defines the data structure for user registration
type RegisterPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginPayload defines the data structure for user login
type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// TokenResponse defines the structure of JWT token response
type TokenResponse struct {
	Token string `json:"token"`
}

// Register handles user registration
func Register(c *gin.Context) {
	var data RegisterPayload
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := services.HashPassword(data.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to encrypt password"})
		return
	}

	user := models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: hashedPassword,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(201, gin.H{"message": "User registered successfully"})
}

// Login handles user login and returns JWT token
func Login(c *gin.Context) {
	var data LoginPayload
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", data.Email).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	if !services.CheckPasswordHash(data.Password, user.Password) {
		c.JSON(400, gin.H{"error": "Incorrect password"})
		return
	}

	token, err := services.GenerateToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate JWT token"})
		return
	}

	c.JSON(200, TokenResponse{Token: token})
}

// GetUsers returns a list of all users
func GetUsers(c *gin.Context) {
	var users []models.User

	if err := config.DB.Select("id, name, email").Find(&users).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error fetching users from database"})
		return
	}

	c.JSON(200, users)
}

// Additional handlers for creating, updating, deleting users, etc. can be added here.
