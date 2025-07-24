package subscription

import (
	"app/internal/model"
	"github.com/google/uuid"
	"time"
)

func (s *SubscriptionUsecases) Create(request model.CreateRequest) (model.CreateResponse, error) {

	id, err := uuid.NewRandom()

	if err != nil {
		s.logger.Println("generate uuid error", err)
		return model.CreateResponse{}, err
	}
	start, err := s.normalizeData(request.StartDate)
	if err != nil {
		s.logger.Println("parse start date error", err)
		return model.CreateResponse{}, err
	}

	subscription := model.Subscription{
		ID:        id.String(),
		Name:      request.Name,
		Price:     request.Price,
		UserID:    request.UserID,
		StartDate: start,
		EndDate:   start.AddDate(0, request.Duration, 0).UTC(),
		CreatedAt: time.Now().UTC(),
	}

	res, err := s.repo.Create(subscription)

	if err != nil {
		s.logger.Println("create subscription error:", err)
		return model.CreateResponse{}, err
	}
	return res, nil
}
