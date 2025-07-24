package subscription

import (
	"app/internal/apperr"
	"app/internal/model"
	"errors"
	"time"
)

func (s *SubscriptionUsecases) Replace(id string, request model.CreateRequest) (model.CreateResponse, error) {
	start, err := s.normalizeData(request.StartDate)
	if err != nil {
		s.logger.Println("error parsing start date:", err)
		return model.CreateResponse{}, err
	}

	newSubscription := model.Subscription{
		ID:        id,
		Name:      request.Name,
		Price:     request.Price,
		UserID:    request.UserID,
		StartDate: start,
		EndDate:   start.AddDate(0, request.Duration, 0).UTC(),
		CreatedAt: time.Now().UTC(),
	}

	_, err = s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			s.logger.Println("service not found, creating new one with id:", id)
			res, err := s.repo.Create(newSubscription)
			if err != nil {
				s.logger.Println("error creating service:", err)
				return model.CreateResponse{}, err
			}
			return res, nil
		}
		s.logger.Println("error checking service existence:", err)
		return model.CreateResponse{}, err
	}

	updated, err := s.repo.Update(newSubscription)
	if err != nil {
		s.logger.Println("error updating service:", err)
		return model.CreateResponse{}, err
	}

	res := model.CreateResponse{
		ID:        updated.ID,
		Name:      updated.Name,
		Price:     updated.Price,
		UserID:    updated.UserID,
		StartDate: updated.StartDate,
		EndDate:   updated.EndDate,
		CreatedAt: updated.CreatedAt,
	}
	return res, nil
}
