package dto

import "time"

type CreateRequest struct {
	Name      string `json:"service_name" example:"Yandex_plus"`
	Price     int    `json:"price" example:"1499"`
	UserID    string `json:"user_id" example:"81f4b2ae-4af4-4e15-8c85-7a324b6f0c58"`
	StartDate string `json:"start_date" example:"2025-07"`
	Duration  int    `json:"duration" example:"6"`
}

type CreateResponse struct {
	ID        string    `json:"id" example:"8cd063d1-1f1b-4199-a3fa-f9aab8983166"`
	Name      string    `json:"service_name" example:"Yandex_plus"`
	Price     int       `json:"price" example:"1499"`
	UserID    string    `json:"user_id" example:"81f4b2ae-4af4-4e15-8c85-7a324b6f0c58"`
	StartDate time.Time `json:"start_date" example:"2025-07-01T00:00:00Z"`
	EndDate   time.Time `json:"end_date" example:"2026-01-01T00:00:00Z"`
	CreatedAt time.Time `json:"created_at" example:"2025-07-22T17:36:55.517136Z"`
}

type ResponseDTO struct {
	Data []ServiceDTO `json:"data"`
}
