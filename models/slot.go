package models

import "gorm.io/gorm"

type Slot struct {
	gorm.Model

	Date     string `json:"date" gorm:"uniqueIndex;idx_slots_date_time"`
	Time     string `json:"time" gorm:"uniqueIndex;idx_slots_date_time"`
	IsBooked bool   `json:"isBooked" gorm:"default:false"`
	IsLocked bool   `json:"isLocked" gorm:"default:false"`
}
