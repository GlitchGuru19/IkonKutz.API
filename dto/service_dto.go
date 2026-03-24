package dto

type CreateServiceRequest struct {
	Name            string  `json:"name"`
	Price           float64 `json:"price"`
	DurationMinutes int     `json:"durationMinutes"`
	Description     string  `json:"description"`
}

type UpdateServiceRequest struct {
	Name            string  `json:"name"`
	Price           float64 `json:"price"`
	DurationMinutes int     `json:"durationMinutes"`
	Description     string  `json:"description"`
}
