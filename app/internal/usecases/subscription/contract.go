package subscription

import (
	"app/internal/model"
	"time"
)

//go:generate mockgen -source=contract.go -destination=./mock/mock.go -package=mocks

type subscriptionRepo interface {
	Create(service model.Subscription) (model.CreateResponse, error)
	Update(service model.Subscription) (model.UpdateResponse, error)
	GetByID(id string) (model.Subscription, error)
	GetAll(limit, offset int) (model.GetAllResponse, error)
	DeleteByID(id string) error
	GetTotalPrice(userID, serviceName string, from, to time.Time) (int, error)
}
