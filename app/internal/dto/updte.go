package dto

import "time"

type UpdateDTO struct {
	Name      *string `json:"service_name" example:"YouTube Premium"`
	Price     *int    `json:"price" example:"599"`
	StartDate *string `json:"start_date" example:"2025-08-01"`
	Duration  *int    `json:"duration" example:"6"`
}

type UpdateResponseDTO struct {
	ID        string    `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Name      string    `json:"service_name" example:"YouTube Premium"`
	Price     int       `json:"price" example:"599"`
	UserID    string    `json:"user_id" example:"user_001"`
	StartDate time.Time `json:"start_date" example:"2025-08-01T00:00:00Z"`
	EndDate   time.Time `json:"end_date" example:"2026-02-01T00:00:00Z"`
	CreatedAt time.Time `json:"created_at" example:"2025-07-23T15:04:05Z"`
}
