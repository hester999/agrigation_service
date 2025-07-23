package service_handler

import (
	"app/internal/model"
)

type Service interface {
	Create(request model.CreateRequest) (model.CreateResponse, error)
	Update(id string, request *model.UpdateRequest) (model.UpdateResponse, error)
	GetByID(id string) (model.Service, error)
	GetAll(limit, offset int) (model.GetAllResponse, error)
	GetTotalPrice(userID, serviceName, from, to string) (int, error)
	DeleteByID(id string) error
	Replace(id string, request model.CreateRequest) (model.CreateResponse, error)
}
