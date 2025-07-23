package service_usecases

import (
	"app/internal/model"
	"github.com/google/uuid"
	"time"
)

func (s *ServiceUsecases) Create(request model.CreateRequest) (model.CreateResponse, error) {

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

	newService := model.Service{
		ID:        id.String(),
		Name:      request.Name,
		Price:     request.Price,
		UserID:    request.UserID,
		StartDate: start,
		EndDate:   start.AddDate(0, request.Duration, 0).UTC(),
		CreatedAt: time.Now().UTC(),
	}

	res, err := s.repo.Create(newService)
	if err != nil {
		s.logger.Println("create service error:", err)
		return model.CreateResponse{}, err
	}
	return res, nil
}
