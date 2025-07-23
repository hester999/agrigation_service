package service_handler

import (
	"log"
	"time"
)

type ServiceHandler struct {
	usecase Service
	logger  *log.Logger
}

type ErrDTO struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type ErrDTOArr struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

type ServiceDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"service_name"`
	Price     int       `json:"price"`
	UserID    string    `json:"user_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	CreatedAt time.Time `json:"created_at"`
}

func NewServiceRepo(usecases Service, log *log.Logger) *ServiceHandler {
	return &ServiceHandler{
		usecase: usecases,
		logger:  log,
	}
}
