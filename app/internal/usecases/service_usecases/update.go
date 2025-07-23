package service_usecases

import (
	"app/internal/model"
	"time"
)

func (s *ServiceUsecases) Update(id string, request *model.UpdateRequest) (model.UpdateResponse, error) {
	service, err := s.repo.GetByID(id)
	if err != nil {
		s.logger.Println("get service error:", err)
		return model.UpdateResponse{}, err
	}

	err = s.compareUpdate(&service, request)
	if err != nil {
		s.logger.Println("compare update error:", err)
		return model.UpdateResponse{}, err
	}

	service.CreatedAt = time.Now().UTC()

	updated, err := s.repo.Update(service)
	if err != nil {
		s.logger.Println("update service error:", err)
		return model.UpdateResponse{}, err
	}
	return updated, nil
}

func (s *ServiceUsecases) compareUpdate(old *model.Service, new *model.UpdateRequest) error {
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
	return nil
}
