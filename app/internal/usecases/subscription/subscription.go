package subscription

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type SubscriptionUsecases struct {
	repo   subscriptionRepo
	logger *log.Logger
}

func NewSubscriptionUsecases(repo subscriptionRepo, logger *log.Logger) *SubscriptionUsecases {
	return &SubscriptionUsecases{
		repo:   repo,
		logger: logger,
	}
}

func (s *SubscriptionUsecases) normalizeData(data string) (time.Time, error) {
	data = strings.TrimSpace(data)

	if t, err := time.Parse("2006-01-02", data); err == nil {
		return t, nil
	}

	if t, err := time.Parse("01-2006", data); err == nil {

		return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC), nil
	}

	return time.Time{}, fmt.Errorf("cannot parse date: %s", data)
}
