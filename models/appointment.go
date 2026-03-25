package models

import "gorm.io/gorm"

type Appointment struct {
	gorm.Model

	UserID       uint    `json:"userId"`
	CustomerName string  `json:"customerName"`
	ServiceID    uint    `json:"serviceId"`
	ServiceName  string  `json:"serviceName"`
	SlotID       uint    `json:"slotId"`
	Date         string  `json:"date"`
	Time         string  `json:"time"`
	Price        float64 `json:"price"`
	Status       string  `json:"status"`
}
