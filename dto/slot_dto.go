package dto

type CreateSlotRequest struct {
	Date     string `json:"date"`
	Time     string `json:"time"`
	IsLocked bool   `json:"isLocked"`
}

type UpdateSlotRequest struct {
	Date     string `json:"date"`
	Time     string `json:"time"`
	IsLocked bool   `json:"isLocked"`
}
