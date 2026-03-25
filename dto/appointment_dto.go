package dto

type CreateAppointmentRequest struct {
	ServiceID    uint   `json:"serviceId" binding:"required"`
	Date         string `json:"date"`
	Time         string `json:"time"`
}

type UpdateAppointmentRequest struct {
	ServiceID    uint   `json:"serviceId"`
	Date         string `json:"date"`
	Time         string `json:"time"`
	Status       string `json:"status"`
}
