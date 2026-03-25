package models

import "gorm.io/gorm"

type Appointment struct {
	gorm.Model
	CustomerName string  `json:"customer_name" binding:"required"`
	ServiceId    uint    `json:"service_id" binding:"required"`
	ServiceName  string  `json:"serviceName"`
	Date         string  `json:"date`
	Time         string  `json:"time"`
	Price        float64 `json:"price"`
	Status       string  `json:"status"`
}
