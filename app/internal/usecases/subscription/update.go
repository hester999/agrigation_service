package subscription

import (
	"app/internal/model"
	"time"
)

func (s *SubscriptionUsecases) Update(id string, request *model.UpdateRequest) (model.UpdateResponse, error) {
	subscription, err := s.repo.GetByID(id)
	if err != nil {
		s.logger.Println("get subscription error:", err)
		return model.UpdateResponse{}, err
	}

	err = s.compareUpdate(&subscription, request)
	if err != nil {
		s.logger.Println("compare update error:", err)
		return model.UpdateResponse{}, err
	}

	subscription.CreatedAt = time.Now().UTC()

	updated, err := s.repo.Update(subscription)
	if err != nil {
		s.logger.Println("update subscription error:", err)
		return model.UpdateResponse{}, err
	}
	return updated, nil
}

func (s *SubscriptionUsecases) compareUpdate(old *model.Subscription, new *model.UpdateRequest) error {
	if new.Name != nil {
		old.Name = *new.Name
	}
	if new.Price != nil {
		old.Price = *new.Price
	}
	if new.StartDate != nil && new.Duration != nil {
		start, err := s.normalizeData(*new.StartDate)
		if err != nil {
			s.logger.Println("parse start date error:", err)
			return err
		}
		old.StartDate = start
		old.EndDate = start.AddDate(0, *new.Duration, 0).UTC()
	}

	if new.Duration != nil {
		newEndDate := old.EndDate.AddDate(0, *new.Duration, 0).UTC()
		old.EndDate = newEndDate
	}
	return nil
}
