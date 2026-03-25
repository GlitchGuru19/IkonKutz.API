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

// GetAppointments returns all appointments by newest first.
func GetAppointments(c *gin.Context) {
	var appointments []models.Appointment

	// Fetch appointments ordered by creation date in descending order (newest first).
	if err := initializers.DB.Order("created_at desc").Find(&appointments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch appointments",
		})
		return
	}

	// Return the appointments in the response.
	c.JSON(http.StatusOK, gin.H{
		"message": "Appointments fetched successfully",
		"data":    appointments,
	})
}

// GetAppointment returns one appointment by ID.
func GetAppointment(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid appointment ID",
		})
		return
	}

	var appointment models.Appointment

	if err := initializers.DB.First(&appointment, uint(parsedID)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Appointment not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Appointment fetched successfully",
		"data":    appointment,
	})
}

// CreateAppointment creates a new booking.
// For now, we only validate that the referenced service exists.
// Later batches can add slot validation and double-booking protection.
func CreateAppointment(c *gin.Context) {
	var body dto.CreateAppointmentRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if strings.TrimSpace(body.CustomerName) == "" || body.ServiceID == 0 || strings.TrimSpace(body.Date) == "" || strings.TrimSpace(body.Time) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "CustomerName, serviceId, date and time are required",
		})
		return
	}

	var service models.Service

	// We fetch the service so we can copy useful data like name and price
	// into the appointment record. This keeps appointment history stable
	// even if the service later changes.
	if err := initializers.DB.First(&service, body.ServiceID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Service not found",
		})
		return
	}

	appointment := models.Appointment{
		CustomerName: strings.TrimSpace(body.CustomerName),
		ServiceId:    service.ID,
		ServiceName:  service.Name,
		Date:         strings.TrimSpace(body.Date),
		Time:         strings.TrimSpace(body.Time),
		Price:        service.Price,
		Status:       "confirmed",
	}

	if err := initializers.DB.Create(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create appointment",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Appointment created successfully",
		"data":    appointment,
	})
}

// UpdateAppointment updates an appointment by ID.
// We also refresh serviceName and price from the selected service.
func UpdateAppointment(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid appointment ID",
		})
		return
	}

	var body dto.UpdateAppointmentRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if strings.TrimSpace(body.CustomerName) == "" || body.ServiceID == 0 || strings.TrimSpace(body.Date) == "" || strings.TrimSpace(body.Time) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "CustomerName, serviceId, date and time are required",
		})
		return
	}

	var appointment models.Appointment
	if err := initializers.DB.First(&appointment, uint(parsedID)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Appointment not found",
		})
		return
	}

	var service models.Service
	if err := initializers.DB.First(&service, body.ServiceID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Service not found",
		})
		return
	}

	appointment.CustomerName = strings.TrimSpace(body.CustomerName)
	appointment.ServiceId = service.ID
	appointment.ServiceName = service.Name
	appointment.Date = strings.TrimSpace(body.Date)
	appointment.Time = strings.TrimSpace(body.Time)
	appointment.Price = service.Price

	// If no status is sent, keep the existing one.
	if strings.TrimSpace(body.Status) != "" {
		appointment.Status = strings.TrimSpace(body.Status)
	}

	if err := initializers.DB.Save(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update appointment",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Appointment updated successfully",
		"data":    appointment,
	})
}

// CancelAppointment only changes the status field to cancelled.
func CancelAppointment(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid appointment ID",
		})
		return
	}

	var appointment models.Appointment

	if err := initializers.DB.First(&appointment, uint(parsedID)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Appointment not found",
		})
		return
	}

	appointment.Status = "cancelled"

	if err := initializers.DB.Save(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to cancel appointment",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Appointment cancelled successfully",
		"data":    appointment,
	})
}

// DeleteAppointment soft-deletes an appointment.
func DeleteAppointment(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid appointment ID",
		})
		return
	}

	var appointment models.Appointment

	if err := initializers.DB.First(&appointment, uint(parsedID)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Appointment not found",
		})
		return
	}

	if err := initializers.DB.Delete(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete appointment",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Appointment deleted successfully",
	})
}
