package dto

type CreateAppointmentRequest struct {
	CustomerName string `json:"customerName" binding:"required"`
	ServiceID    uint   `json:"serviceId" binding:"required"`
	Date         string `json:"date"`
	Time         string `json:"time"`
}

type UpdateAppointmentRequest struct {
	CustomerName string `json:"customerName"`
	ServiceID    uint   `json:"serviceId"`
	Date         string `json:"date"`
	Time         string `json:"time"`
	Status       string `json:"status"`
}
