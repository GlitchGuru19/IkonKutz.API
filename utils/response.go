package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// function to send a success response
func Success(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, gin.H{
		"message": message,
		"data":    data,
	})
}

// fucntion to send an error response
func Error(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"error": message,
	})
}

// function to send a 200 OK response
func OK(c *gin.Context, message string, data interface{}) {
	// This is just a wrapper around the Success function for convenience.
	Success(c, http.StatusOK, message, data)
}

// function to create a 201 Created response
func Created(c *gin.Context, message string, data interface{}) {
	Success(c, http.StatusCreated, message, data)
}
