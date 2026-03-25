package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"IkonKutz.API/dto"
	"IkonKutz.API/initializers"
	"IkonKutz.API/models"
	"IkonKutz.API/services"
	"IkonKutz.API/utils"
	"github.com/gin-gonic/gin"
)

// GetAppointments is intended for admins.
// Admins can see all appointments.
func GetAppointments(c *gin.Context) {
	var appointments []models.Appointment

	if err := initializers.DB.Order("created_at desc").Find(&appointments).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to fetch appointments")
		return
	}

	utils.OK(c, "Appointments fetched successfully", appointments)
}

// GetMyAppointments returns only the logged-in user's appointments.
func GetMyAppointments(c *gin.Context) {
	userIDValue, _ := c.Get("userID")
	userID := userIDValue.(uint)

	var appointments []models.Appointment
	if err := initializers.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&appointments).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to fetch your appointments")
		return
	}

	utils.OK(c, "Your appointments fetched successfully", appointments)
}

// GetAppointment returns one appointment.
// Admins can fetch any appointment.
// Customers can only fetch their own.
func GetAppointment(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid appointment ID")
		return
	}

	var appointment models.Appointment
	if err := initializers.DB.First(&appointment, uint(parsedID)).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Appointment not found")
		return
	}

	roleValue, _ := c.Get("role")
	userIDValue, _ := c.Get("userID")

	role := roleValue.(string)
	userID := userIDValue.(uint)

	if role != "admin" && appointment.UserID != userID {
		utils.Error(c, http.StatusForbidden, "You can only view your own appointments")
		return
	}

	utils.OK(c, "Appointment fetched successfully", appointment)
}

// CreateAppointment creates a booking for the current authenticated user.
func CreateAppointment(c *gin.Context) {
	userIDValue, _ := c.Get("userID")
	userID := userIDValue.(uint)

	var user models.User
	if err := initializers.DB.First(&user, userID).Error; err != nil {
		utils.Error(c, http.StatusUnauthorized, "User not found")
		return
	}

	var body dto.CreateAppointmentRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	if body.ServiceID == 0 || strings.TrimSpace(body.Date) == "" || strings.TrimSpace(body.Time) == "" {
		utils.Error(c, http.StatusBadRequest, "ServiceId, date and time are required")
		return
	}

	tx := initializers.DB.Begin()
	if tx.Error != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to start database transaction")
		return
	}
	defer tx.Rollback()

	service, err := services.FindService(tx, body.ServiceID)
	if err != nil {
		utils.Error(c, http.StatusNotFound, "Service not found")
		return
	}

	slot, err := services.FindSlotByDateTime(tx, body.Date, body.Time)
	if err != nil {
		utils.Error(c, http.StatusNotFound, "Slot not found for the given date and time")
		return
	}

	if err := services.ValidateSlotForBooking(slot); err != nil {
		switch {
		case errors.Is(err, services.ErrSlotLocked):
			utils.Error(c, http.StatusBadRequest, "Slot is locked")
		case errors.Is(err, services.ErrSlotAlreadyBooked):
			utils.Error(c, http.StatusBadRequest, "Slot is already booked")
		default:
			utils.Error(c, http.StatusBadRequest, "Slot is not available")
		}
		return
	}

	appointment := models.Appointment{
		UserID:       user.ID,
		CustomerName: user.Name,
		ServiceID:    service.ID,
		ServiceName:  service.Name,
		SlotID:       slot.ID,
		Date:         slot.Date,
		Time:         slot.Time,
		Price:        service.Price,
		Status:       "confirmed",
	}

	if err := tx.Create(&appointment).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to create appointment")
		return
	}

	if err := services.MarkSlotBooked(tx, slot); err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to reserve slot")
		return
	}

	if err := tx.Commit().Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to finalize appointment creation")
		return
	}

	utils.Created(c, "Appointment created successfully", appointment)
}

// UpdateAppointment:
// - admins can update any appointment
// - customers can update only their own appointment
func UpdateAppointment(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid appointment ID")
		return
	}

	var body dto.UpdateAppointmentRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	if body.ServiceID == 0 || strings.TrimSpace(body.Date) == "" || strings.TrimSpace(body.Time) == "" {
		utils.Error(c, http.StatusBadRequest, "ServiceId, date and time are required")
		return
	}

	normalizedStatus, err := services.NormalizeStatus(body.Status)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Status must be either confirmed or cancelled")
		return
	}

	roleValue, _ := c.Get("role")
	userIDValue, _ := c.Get("userID")

	role := roleValue.(string)
	userID := userIDValue.(uint)

	tx := initializers.DB.Begin()
	if tx.Error != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to start database transaction")
		return
	}
	defer tx.Rollback()

	var appointment models.Appointment
	if err := tx.First(&appointment, uint(parsedID)).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Appointment not found")
		return
	}

	if role != "admin" && appointment.UserID != userID {
		utils.Error(c, http.StatusForbidden, "You can only update your own appointments")
		return
	}

	service, err := services.FindService(tx, body.ServiceID)
	if err != nil {
		utils.Error(c, http.StatusNotFound, "Service not found")
		return
	}

	currentSlot, err := services.FindSlotByID(tx, appointment.SlotID)
	if err != nil {
		utils.Error(c, http.StatusNotFound, "Current slot linked to appointment was not found")
		return
	}

	targetSlot, err := services.FindSlotByDateTime(tx, body.Date, body.Time)
	if err != nil {
		utils.Error(c, http.StatusNotFound, "Target slot not found for the given date and time")
		return
	}

	isSameSlot := currentSlot.ID == targetSlot.ID

	if !isSameSlot {
		if err := services.ValidateSlotForBooking(targetSlot); err != nil {
			switch {
			case errors.Is(err, services.ErrSlotLocked):
				utils.Error(c, http.StatusBadRequest, "Target slot is locked")
			case errors.Is(err, services.ErrSlotAlreadyBooked):
				utils.Error(c, http.StatusBadRequest, "Target slot is already booked")
			default:
				utils.Error(c, http.StatusBadRequest, "Target slot is not available")
			}
			return
		}
	}

	finalStatus := appointment.Status
	if normalizedStatus != "" {
		finalStatus = normalizedStatus
	}

	if !isSameSlot {
		if err := services.MarkSlotAvailable(tx, currentSlot); err != nil {
			utils.Error(c, http.StatusInternalServerError, "Failed to release old slot")
			return
		}

		if finalStatus == "confirmed" {
			if err := services.MarkSlotBooked(tx, targetSlot); err != nil {
				utils.Error(c, http.StatusInternalServerError, "Failed to reserve target slot")
				return
			}
		}
	} else {
		if finalStatus == "cancelled" && currentSlot.IsBooked {
			if err := services.MarkSlotAvailable(tx, currentSlot); err != nil {
				utils.Error(c, http.StatusInternalServerError, "Failed to release slot")
				return
			}
		}

		if finalStatus == "confirmed" && !currentSlot.IsBooked {
			if currentSlot.IsLocked {
				utils.Error(c, http.StatusBadRequest, "Current slot is locked")
				return
			}
			if err := services.MarkSlotBooked(tx, currentSlot); err != nil {
				utils.Error(c, http.StatusInternalServerError, "Failed to reserve slot")
				return
			}
		}
	}

	appointment.ServiceID = service.ID
	appointment.ServiceName = service.Name
	appointment.Price = service.Price
	appointment.SlotID = targetSlot.ID
	appointment.Date = targetSlot.Date
	appointment.Time = targetSlot.Time
	appointment.Status = finalStatus

	if err := tx.Save(&appointment).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to update appointment")
		return
	}

	if err := tx.Commit().Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to finalize appointment update")
		return
	}

	utils.OK(c, "Appointment updated successfully", appointment)
}

// CancelAppointment:
// - admins can cancel any appointment
// - customers can cancel only their own
func CancelAppointment(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid appointment ID")
		return
	}

	roleValue, _ := c.Get("role")
	userIDValue, _ := c.Get("userID")

	role := roleValue.(string)
	userID := userIDValue.(uint)

	tx := initializers.DB.Begin()
	if tx.Error != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to start database transaction")
		return
	}
	defer tx.Rollback()

	var appointment models.Appointment
	if err := tx.First(&appointment, uint(parsedID)).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Appointment not found")
		return
	}

	if role != "admin" && appointment.UserID != userID {
		utils.Error(c, http.StatusForbidden, "You can only cancel your own appointments")
		return
	}

	slot, err := services.FindSlotByID(tx, appointment.SlotID)
	if err != nil {
		utils.Error(c, http.StatusNotFound, "Linked slot not found")
		return
	}

	appointment.Status = "cancelled"

	if err := tx.Save(&appointment).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to cancel appointment")
		return
	}

	if slot.IsBooked {
		if err := services.MarkSlotAvailable(tx, slot); err != nil {
			utils.Error(c, http.StatusInternalServerError, "Failed to free slot")
			return
		}
	}

	if err := tx.Commit().Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to finalize cancellation")
		return
	}

	utils.OK(c, "Appointment cancelled successfully", appointment)
}

// DeleteAppointment:
// - admins can delete any appointment
// - customers can delete only their own
func DeleteAppointment(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid appointment ID")
		return
	}

	roleValue, _ := c.Get("role")
	userIDValue, _ := c.Get("userID")

	role := roleValue.(string)
	userID := userIDValue.(uint)

	tx := initializers.DB.Begin()
	if tx.Error != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to start database transaction")
		return
	}
	defer tx.Rollback()

	var appointment models.Appointment
	if err := tx.First(&appointment, uint(parsedID)).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Appointment not found")
		return
	}

	if role != "admin" && appointment.UserID != userID {
		utils.Error(c, http.StatusForbidden, "You can only delete your own appointments")
		return
	}

	slot, err := services.FindSlotByID(tx, appointment.SlotID)
	if err == nil && slot.IsBooked {
		if err := services.MarkSlotAvailable(tx, slot); err != nil {
			utils.Error(c, http.StatusInternalServerError, "Failed to free slot")
			return
		}
	}

	if err := tx.Delete(&appointment).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to delete appointment")
		return
	}

	if err := tx.Commit().Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to finalize appointment deletion")
		return
	}

	utils.OK(c, "Appointment deleted successfully", nil)
}
