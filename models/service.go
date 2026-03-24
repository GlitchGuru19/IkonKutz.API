package models

import "gorm.io/gorm"

type Service struct {
	gorm.Model

	Name            string  `json:"name"`
	Price           float64 `json:"price"`
	DurationMinutes int     `json:"durationMinutes"`
	Description     string  `json:"description"`
}