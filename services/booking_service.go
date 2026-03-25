// This file contains the core business logic for handling bookings, such as validating services and slots,
package services

import (
	"errors"
	"strings"

	"IkonKutz.API/models"
	"gorm.io/gorm"
)

// Define custom errors for better error handling in the service layer.
var (
	ErrServiceNotFound   = errors.New("service not found")
	ErrSlotNotFound      = errors.New("slot not found")
	ErrSlotLocked        = errors.New("slot is locked")
	ErrSlotAlreadyBooked = errors.New("slot is already booked")
	ErrInvalidStatus     = errors.New("invalid appointment status")
)

// FindService checks if a service with the given ID exists and returns it.
func FindService(db *gorm.DB, serviceID uint) (*models.Service, error) {
	var service models.Service
	if err := db.First(&service, serviceID).Error; err != nil {
		return nil, ErrServiceNotFound
	}
	return &service, nil
}

// FindSlotByDateTime checks if a slot exists for the given date and time.
func FindSlotByDateTime(db *gorm.DB, date string, time string) (*models.Slot, error) {
	var slot models.Slot
	err := db.Where("date = ? AND time = ?", strings.TrimSpace(date), strings.TrimSpace(time)).First(&slot).Error
	if err != nil {
		return nil, ErrSlotNotFound
	}
	return &slot, nil
}

// FindSlotByID checks if a slot with the given ID exists and returns it.
func FindSlotByID(db *gorm.DB, slotID uint) (*models.Slot, error) {
	var slot models.Slot
	if err := db.First(&slot, slotID).Error; err != nil {
		return nil, ErrSlotNotFound
	}
	return &slot, nil
}

// ValidateSlotForBooking checks if the slot is available for booking.
func ValidateSlotForBooking(slot *models.Slot) error {
	if slot.IsLocked {
		return ErrSlotLocked
	}
	if slot.IsBooked {
		return ErrSlotAlreadyBooked
	}
	return nil
}

// MarkSlotBooked sets the IsBooked field of the slot to true and saves it to the database.
func MarkSlotBooked(tx *gorm.DB, slot *models.Slot) error {
	slot.IsBooked = true
	return tx.Save(slot).Error
}

// MarkSlotLocked sets the IsLocked field of the slot to true and saves it to the database.
func MarkSlotAvailable(tx *gorm.DB, slot *models.Slot) error {
	slot.IsBooked = false
	return tx.Save(slot).Error
}

// NormalizeStatus converts a status string to a standardized format and validates it.
func NormalizeStatus(status string) (string, error) {
	s := strings.TrimSpace(strings.ToLower(status))

	switch s {
	case "":
		return "", nil
	case "confirmed":
		return "confirmed", nil
	case "cancelled":
		return "cancelled", nil
	default:
		return "", ErrInvalidStatus
	}
}
