package controllers

import (
	"net/http"
	"strings"

	"IkonKutz.API/dto"
	"IkonKutz.API/initializers"
	"IkonKutz.API/models"
	"IkonKutz.API/utils"
	"github.com/gin-gonic/gin"
)

// Register creates a new customer account.
// New public registrations always get the "customer" role.
func Register(c *gin.Context) {
	var body dto.RegisterRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	body.Name = strings.TrimSpace(body.Name)
	body.Email = strings.TrimSpace(strings.ToLower(body.Email))
	body.Password = strings.TrimSpace(body.Password)

	if body.Name == "" || body.Email == "" || body.Password == "" {
		utils.Error(c, http.StatusBadRequest, "Name, email and password are required")
		return
	}

	var existingUser models.User
	if err := initializers.DB.Where("email = ?", body.Email).First(&existingUser).Error; err == nil {
		utils.Error(c, http.StatusBadRequest, "Email is already in use")
		return
	}

	passwordHash, err := utils.HashPassword(body.Password)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	user := models.User{
		Name:         body.Name,
		Email:        body.Email,
		PasswordHash: passwordHash,
		Role:         "customer",
	}

	if err := initializers.DB.Create(&user).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Role, user.Email)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	utils.Created(c, "User registered successfully", gin.H{
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
		"token": token,
	})
}

// Login authenticates a user and returns a JWT.
func Login(c *gin.Context) {
	var body dto.LoginRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	body.Email = strings.TrimSpace(strings.ToLower(body.Email))
	body.Password = strings.TrimSpace(body.Password)

	if body.Email == "" || body.Password == "" {
		utils.Error(c, http.StatusBadRequest, "Email and password are required")
		return
	}

	var user models.User
	if err := initializers.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		utils.Error(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	if err := utils.CheckPasswordHash(body.Password, user.PasswordHash); err != nil {
		utils.Error(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Role, user.Email)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	utils.OK(c, "Login successful", gin.H{
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
		"token": token,
	})
}

// Me returns the currently authenticated user.
func Me(c *gin.Context) {
	userIDValue, exists := c.Get("userID")
	if !exists {
		utils.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	userID, ok := userIDValue.(uint)
	if !ok {
		utils.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var user models.User
	if err := initializers.DB.First(&user, userID).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "User not found")
		return
	}

	utils.OK(c, "Current user fetched successfully", gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
	})
}

// Logout is stateless with JWT unless you add token blacklisting.
// For v1, this is simply a client-side token discard endpoint.
func Logout(c *gin.Context) {
	utils.OK(c, "Logout successful", gin.H{
		"message": "Remove the token on the client side",
	})
}
