package service_usecases

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type ServiceUsecases struct {
	repo   Service
	logger *log.Logger
}

func NewService(repo Service, logger *log.Logger) *ServiceUsecases {
	return &ServiceUsecases{
		repo:   repo,
		logger: logger,
	}
}

func (s *ServiceUsecases) normalizeData(data string) (time.Time, error) {
	data = strings.TrimSpace(data)

	if t, err := time.Parse("2006-01-02", data); err == nil {
		return t, nil
	}

	if t, err := time.Parse("01-2006", data); err == nil {

		return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC), nil
	}

	return time.Time{}, fmt.Errorf("cannot parse date: %s", data)
}
