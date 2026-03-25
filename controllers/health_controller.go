package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health is a simple endpoint to check if the API is running and healthy.
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "API is healthy",
	})
}
