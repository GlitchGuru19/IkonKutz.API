package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name         string `json:"name"`
	Email        string `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash string `json:"-"`
	Role         string `json:"role" gorm:"default:customer not null"` // e.g., "admin", "customer"
}
