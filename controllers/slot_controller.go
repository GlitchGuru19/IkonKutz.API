package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"IkonKutz.API/dto"
	"IkonKutz.API/initializers"
	"IkonKutz.API/models"
	"IkonKutz.API/utils"
	"github.com/gin-gonic/gin"
)

func GetSlots(c *gin.Context) {
	var slots []models.Slot

	db := initializers.DB.Model(&models.Slot{})

	if date := strings.TrimSpace(c.Query("date")); date != "" {
		db = db.Where("date = ?", date)
	}

	if booked := strings.TrimSpace(c.Query("booked")); booked != "" {
		if booked == "true" {
			db = db.Where("is_booked = ?", true)
		} else if booked == "false" {
			db = db.Where("is_booked = ?", false)
		}
	}

	if locked := strings.TrimSpace(c.Query("locked")); locked != "" {
		if locked == "true" {
			db = db.Where("is_locked = ?", true)
		} else if locked == "false" {
			db = db.Where("is_locked = ?", false)
		}
	}

	if err := db.Order("date asc, time asc").Find(&slots).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to fetch slots")
		return
	}

	utils.OK(c, "Slots fetched successfully", slots)
}

func GetSlot(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid slot ID")
		return
	}

	var slot models.Slot
	if err := initializers.DB.First(&slot, uint(parsedID)).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Slot not found")
		return
	}

	utils.OK(c, "Slot fetched successfully", slot)
}

func CreateSlot(c *gin.Context) {
	var body dto.CreateSlotRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	if strings.TrimSpace(body.Date) == "" || strings.TrimSpace(body.Time) == "" {
		utils.Error(c, http.StatusBadRequest, "Date and time are required")
		return
	}

	var existing models.Slot
	if err := initializers.DB.Where("date = ? AND time = ?", strings.TrimSpace(body.Date), strings.TrimSpace(body.Time)).First(&existing).Error; err == nil {
		utils.Error(c, http.StatusBadRequest, "A slot with the same date and time already exists")
		return
	}

	slot := models.Slot{
		Date:     strings.TrimSpace(body.Date),
		Time:     strings.TrimSpace(body.Time),
		IsBooked: false,
		IsLocked: body.IsLocked,
	}

	if err := initializers.DB.Create(&slot).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to create slot")
		return
	}

	utils.Created(c, "Slot created successfully", slot)
}

func UpdateSlot(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid slot ID")
		return
	}

	var body dto.UpdateSlotRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	if strings.TrimSpace(body.Date) == "" || strings.TrimSpace(body.Time) == "" {
		utils.Error(c, http.StatusBadRequest, "Date and time are required")
		return
	}

	var slot models.Slot
	if err := initializers.DB.First(&slot, uint(parsedID)).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Slot not found")
		return
	}

	if slot.IsBooked && (slot.Date != strings.TrimSpace(body.Date) || slot.Time != strings.TrimSpace(body.Time)) {
		utils.Error(c, http.StatusBadRequest, "Booked slots cannot be moved to a different date or time")
		return
	}

	var existing models.Slot
	if err := initializers.DB.Where("date = ? AND time = ? AND id <> ?", strings.TrimSpace(body.Date), strings.TrimSpace(body.Time), slot.ID).First(&existing).Error; err == nil {
		utils.Error(c, http.StatusBadRequest, "Another slot with the same date and time already exists")
		return
	}

	slot.Date = strings.TrimSpace(body.Date)
	slot.Time = strings.TrimSpace(body.Time)
	slot.IsLocked = body.IsLocked

	if err := initializers.DB.Save(&slot).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to update slot")
		return
	}

	utils.OK(c, "Slot updated successfully", slot)
}

func LockSlot(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid slot ID")
		return
	}

	var slot models.Slot
	if err := initializers.DB.First(&slot, uint(parsedID)).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Slot not found")
		return
	}

	slot.IsLocked = true

	if err := initializers.DB.Save(&slot).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to lock slot")
		return
	}

	utils.OK(c, "Slot locked successfully", slot)
}

func UnlockSlot(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid slot ID")
		return
	}

	var slot models.Slot
	if err := initializers.DB.First(&slot, uint(parsedID)).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Slot not found")
		return
	}

	slot.IsLocked = false

	if err := initializers.DB.Save(&slot).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to unlock slot")
		return
	}

	utils.OK(c, "Slot unlocked successfully", slot)
}

func DeleteSlot(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid slot ID")
		return
	}

	var slot models.Slot
	if err := initializers.DB.First(&slot, uint(parsedID)).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "Slot not found")
		return
	}

	if slot.IsBooked {
		utils.Error(c, http.StatusBadRequest, "Booked slots cannot be deleted")
		return
	}

	if err := initializers.DB.Delete(&slot).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to delete slot")
		return
	}

	utils.OK(c, "Slot deleted successfully", nil)
}
