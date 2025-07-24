package subscription

import (
	"app/internal/model"
)

type subscriptionUsecases interface {
	Create(request model.CreateRequest) (model.CreateResponse, error)
	Update(id string, request *model.UpdateRequest) (model.UpdateResponse, error)
	GetByID(id string) (model.Subscription, error)
	GetAll(limit, offset int) (model.GetAllResponse, error)
	GetTotalPrice(userID, subscriptionName, from, to string) (int, error)
	DeleteByID(id string) error
	Replace(id string, request model.CreateRequest) (model.CreateResponse, error)
}
