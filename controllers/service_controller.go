package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"IkonKutz.API/dto"
	"IkonKutz.API/initializers"
	"IkonKutz.API/models"
	"github.com/gin-gonic/gin"
)

// GetServices returns all services in descending creation order.
func GetServices(c *gin.Context) {
	var services []models.Service

	if err := initializers.DB.Order("created_at desc").Find(&services).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch services",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Services fetched successfully",
		"data":    services,
	})
}

// GetService returns one service by ID.
func GetService(c *gin.Context) {
	id := c.Param("id")

	// Convert the route parameter to an unsigned integer because GORM IDs are uint.
	parsedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid service ID",
		})
		return
	}

	var service models.Service

	if err := initializers.DB.First(&service, uint(parsedID)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Service not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Service fetched successfully",
		"data":    service,
	})
}

// CreateService creates a new service record.
func CreateService(c *gin.Context) {
	var body dto.CreateServiceRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// Basic validation to keep bad data out of the database.
	if strings.TrimSpace(body.Name) == "" || body.Price <= 0 || body.DurationMinutes <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name, price and durationMinutes are required",
		})
		return
	}

	service := models.Service{
		Name:            strings.TrimSpace(body.Name),
		Price:           body.Price,
		DurationMinutes: body.DurationMinutes,
		Description:     strings.TrimSpace(body.Description),
	}

	if err := initializers.DB.Create(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create service",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Service created successfully",
		"data":    service,
	})
}

// UpdateService updates an existing service by ID.
func UpdateService(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid service ID",
		})
		return
	}

	var body dto.UpdateServiceRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if strings.TrimSpace(body.Name) == "" || body.Price <= 0 || body.DurationMinutes <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name, price and durationMinutes are required",
		})
		return
	}

	var service models.Service

	if err := initializers.DB.First(&service, uint(parsedID)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Service not found",
		})
		return
	}

	// Assign new values to the existing record.
	service.Name = strings.TrimSpace(body.Name)
	service.Price = body.Price
	service.DurationMinutes = body.DurationMinutes
	service.Description = strings.TrimSpace(body.Description)

	if err := initializers.DB.Save(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update service",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Service updated successfully",
		"data":    service,
	})
}

// DeleteService soft-deletes a service by ID.
func DeleteService(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid service ID",
		})
		return
	}

	var service models.Service

	// First check that the record exists before deleting.
	if err := initializers.DB.First(&service, uint(parsedID)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Service not found",
		})
		return
	}

	// Because Service uses gorm.Model, this performs a soft delete.
	if err := initializers.DB.Delete(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete service",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Service deleted successfully",
	})
}
