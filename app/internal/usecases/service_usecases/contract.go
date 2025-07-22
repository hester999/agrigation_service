package service_usecases

import (
	"app/internal/model"
	"time"
)

type Service interface {
	Create(service model.Service) (model.CreateResponse, error)
	Update(service model.Service) (model.UpdateResponse, error)
	GetByID(id string) (model.Service, error)
	GetAll() (model.GetAllResponse, error)
	DeleteByID(id string) error
	GetTotalPrice(userID, serviceName string, from, to time.Time) (int, error)
}
