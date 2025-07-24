package subscription

import (
	"log"
	"time"
)

type SubscriptionHandler struct {
	usecase subscriptionUsecases
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

type SubscriptionDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"service_name"`
	Price     int       `json:"price"`
	UserID    string    `json:"user_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	CreatedAt time.Time `json:"created_at"`
}

func NewSubscriptionHandler(usecases subscriptionUsecases, log *log.Logger) *SubscriptionHandler {
	return &SubscriptionHandler{
		usecase: usecases,
		logger:  log,
	}
}
