package subscription_test

import (
	"app/internal/apperr"
	"app/internal/model"
	usecases "app/internal/usecases/subscription"
	mocks "app/internal/usecases/subscription/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"log"
	"os"
	"testing"
	"time"
)

func TestSubscriptionUsecases_Replace_CreateIfNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMocksubscriptionRepo(ctrl)
	logger := log.New(os.Stdout, "[test] ", log.LstdFlags)
	usecase := usecases.NewSubscriptionUsecases(mockRepo, logger)

	req := model.CreateRequest{
		Name:      "Netflix",
		Price:     900,
		UserID:    "user-123",
		StartDate: "07-2025",
		Duration:  2,
	}

	expectedStart := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
	expectedEnd := expectedStart.AddDate(0, req.Duration, 0)

	mockRepo.EXPECT().
		GetByID("sub-123").
		Return(model.Subscription{}, apperr.ErrNotFound)

	mockRepo.EXPECT().
		Create(gomock.AssignableToTypeOf(model.Subscription{})).
		DoAndReturn(func(sub model.Subscription) (model.CreateResponse, error) {
			assert.Equal(t, expectedStart, sub.StartDate)
			assert.Equal(t, expectedEnd, sub.EndDate)
			assert.Equal(t, "Netflix", sub.Name)
			assert.Equal(t, 900, sub.Price)
			assert.Equal(t, "user-123", sub.UserID)

			return model.CreateResponse{
				ID:        sub.ID,
				Name:      sub.Name,
				Price:     sub.Price,
				UserID:    sub.UserID,
				StartDate: sub.StartDate,
				EndDate:   sub.EndDate,
				CreatedAt: sub.CreatedAt,
			}, nil
		})

	resp, err := usecase.Replace("sub-123", req)

	assert.NoError(t, err)
	assert.Equal(t, "Netflix", resp.Name)
	assert.Equal(t, 900, resp.Price)
	assert.Equal(t, "user-123", resp.UserID)
	assert.Equal(t, expectedStart, resp.StartDate)
	assert.Equal(t, expectedEnd, resp.EndDate)
}

func TestSubscriptionUsecases_Replace_UpdateIfExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMocksubscriptionRepo(ctrl)
	logger := log.New(os.Stdout, "[test] ", log.LstdFlags)
	usecase := usecases.NewSubscriptionUsecases(mockRepo, logger)

	req := model.CreateRequest{
		Name:      "Netflix",
		Price:     900,
		UserID:    "user-123",
		StartDate: "07-2025",
		Duration:  2,
	}

	existing := model.Subscription{
		ID:        "sub-123",
		Name:      "Old",
		Price:     100,
		UserID:    "user-123",
		StartDate: time.Now(),
		EndDate:   time.Now(),
		CreatedAt: time.Now().UTC(),
	}

	expectedStart := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
	expectedEnd := expectedStart.AddDate(0, req.Duration, 0)

	mockRepo.EXPECT().
		GetByID("sub-123").
		Return(existing, nil)

	mockRepo.EXPECT().
		Update(gomock.AssignableToTypeOf(existing)).
		DoAndReturn(func(sub model.Subscription) (model.UpdateResponse, error) {
			assert.Equal(t, expectedStart, sub.StartDate)
			assert.Equal(t, expectedEnd, sub.EndDate)
			assert.Equal(t, "Netflix", sub.Name)
			assert.Equal(t, 900, sub.Price)
			assert.Equal(t, "user-123", sub.UserID)

			return model.UpdateResponse{
				ID:        sub.ID,
				Name:      sub.Name,
				Price:     sub.Price,
				UserID:    sub.UserID,
				StartDate: sub.StartDate,
				EndDate:   sub.EndDate,
				CreatedAt: sub.CreatedAt,
			}, nil
		})

	resp, err := usecase.Replace("sub-123", req)

	assert.NoError(t, err)
	assert.Equal(t, "sub-123", resp.ID)
	assert.Equal(t, "Netflix", resp.Name)
	assert.Equal(t, 900, resp.Price)
	assert.Equal(t, "user-123", resp.UserID)
	assert.Equal(t, expectedStart, resp.StartDate)
	assert.Equal(t, expectedEnd, resp.EndDate)
}
