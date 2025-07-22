package model

import "time"

type Service struct {
	ID        string
	Name      string
	Price     int
	UserID    string
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
}

// Create
type CreateRequest struct {
	Name      string
	Price     int
	UserID    string
	StartDate string
	Duration  int //количество месяцев на которое выдается подписка
}

type CreateResponse struct {
	ID        string
	Name      string
	Price     int
	UserID    string
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
}

////////////////////////////////

// Update
type UpdateRequest struct {
	Name      *string
	Price     *int
	StartDate *string
	Duration  *int //количество месяцев на которое выдается подписка
}
type UpdateResponse struct {
	ID        string
	Name      string
	Price     int
	UserID    string
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
}

// ///////////////////////////////////////
type GetAllResponse struct {
	Data []Service
}
