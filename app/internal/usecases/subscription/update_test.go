package subscription_test

import (
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

func TestSubscriptionUsecases_UpdatePriceOnly(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMocksubscriptionRepo(ctrl)
	logger := log.New(os.Stdout, "[test] ", log.LstdFlags)
	usecase := usecases.NewSubscriptionUsecases(mockRepo, logger)

	price := 10
	p := &price

	input := model.UpdateRequest{
		Name:      nil,
		Price:     p,
		StartDate: nil,
		Duration:  nil,
	}

	start := time.Now()
	end := start.AddDate(0, 1, 0)
	created := time.Now().UTC()

	mockRepo.EXPECT().GetByID("sub-123").Return(model.Subscription{
		ID:        "sub-123",
		Name:      "test",
		Price:     100,
		UserID:    "user-123",
		StartDate: start,
		EndDate:   end,
		CreatedAt: created,
	}, nil)

	mockRepo.EXPECT().
		Update(gomock.AssignableToTypeOf(model.Subscription{})).
		DoAndReturn(func(sub model.Subscription) (model.UpdateResponse, error) {
			assert.Equal(t, 10, sub.Price)
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

	res, err := usecase.Update("sub-123", &input)
	assert.NoError(t, err)
	assert.Equal(t, 10, res.Price)
}

func TestSubscriptionUsecases_UpdateNameOnly(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMocksubscriptionRepo(ctrl)
	logger := log.New(os.Stdout, "[test] ", log.LstdFlags)
	usecase := usecases.NewSubscriptionUsecases(mockRepo, logger)

	newName := "new name"
	p := &newName

	input := model.UpdateRequest{
		Name:      p,
		Price:     nil,
		StartDate: nil,
		Duration:  nil,
	}

	start := time.Now()
	end := start.AddDate(0, 1, 0)
	created := time.Now().UTC()

	mockRepo.EXPECT().GetByID("sub-123").Return(model.Subscription{
		ID:        "sub-123",
		Name:      "old name",
		Price:     100,
		UserID:    "user-123",
		StartDate: start,
		EndDate:   end,
		CreatedAt: created,
	}, nil)

	mockRepo.EXPECT().
		Update(gomock.AssignableToTypeOf(model.Subscription{})).
		DoAndReturn(func(sub model.Subscription) (model.UpdateResponse, error) {
			assert.Equal(t, "new name", sub.Name)
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

	res, err := usecase.Update("sub-123", &input)
	assert.NoError(t, err)
	assert.Equal(t, "new name", res.Name)
}

func TestSubscriptionUsecases_UpdateStartDateAndDuration(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMocksubscriptionRepo(ctrl)
	logger := log.New(os.Stdout, "[test] ", log.LstdFlags)
	usecase := usecases.NewSubscriptionUsecases(mockRepo, logger)

	startStr := "07-2025"
	duration := 3
	input := model.UpdateRequest{
		Name:      nil,
		Price:     nil,
		StartDate: &startStr,
		Duration:  &duration,
	}

	expectedStart := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
	expectedEnd := expectedStart.AddDate(0, 3, 0)

	mockRepo.EXPECT().GetByID("sub-123").Return(model.Subscription{
		ID:        "sub-123",
		Name:      "test",
		Price:     100,
		UserID:    "user-123",
		StartDate: time.Now(),
		EndDate:   time.Now(),
		CreatedAt: time.Now().UTC(),
	}, nil)

	mockRepo.EXPECT().
		Update(gomock.AssignableToTypeOf(model.Subscription{})).
		DoAndReturn(func(sub model.Subscription) (model.UpdateResponse, error) {
			assert.Equal(t, expectedStart, sub.StartDate)
			assert.Equal(t, expectedEnd, sub.EndDate)
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

	res, err := usecase.Update("sub-123", &input)
	assert.NoError(t, err)
	assert.Equal(t, expectedStart, res.StartDate)
	assert.Equal(t, expectedEnd, res.EndDate)
}
