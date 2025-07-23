package dto

import "time"

type ServiceDTO struct {
	ID        string    `json:"id" example:"a1b2c3d4"`
	Name      string    `json:"service_name" example:"Yandex_plus"`
	Price     int       `json:"price" example:"999"`
	UserID    string    `json:"user_id" example:"user-123"`
	StartDate time.Time `json:"start_date" example:"2025-07-25T00:00:00Z"`
	EndDate   time.Time `json:"end_date" example:"2025-10-25T00:00:00Z"`
	CreatedAt time.Time `json:"created_at" example:"2025-07-20T10:00:00Z"`
}
