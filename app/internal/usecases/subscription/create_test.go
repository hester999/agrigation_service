package subscription_test

import (
	"app/internal/model"
	"app/internal/usecases/subscription"
	"app/internal/usecases/subscription/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"log"
	"os"
	"testing"
	"time"
)

func TestSubscriptionUsecases_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMocksubscriptionRepo(ctrl)
	logger := log.New(os.Stdout, "[test] ", log.LstdFlags)
	usecase := subscription.NewSubscriptionUsecases(mockRepo, logger)

	req := model.CreateRequest{
		Name:      "Netflix",
		Price:     899,
		UserID:    "user-123",
		StartDate: "07-2025", // проверим преобразование в normalizeData
		Duration:  3,
	}

	expectedStart := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
	expectedEnd := expectedStart.AddDate(0, 3, 0)

	mockRepo.
		EXPECT().
		Create(gomock.AssignableToTypeOf(model.Subscription{})).
		DoAndReturn(func(sub model.Subscription) (model.CreateResponse, error) {
			assert.Equal(t, "Netflix", sub.Name)
			assert.Equal(t, 899, sub.Price)
			assert.Equal(t, "user-123", sub.UserID)
			assert.Equal(t, expectedStart, sub.StartDate)
			assert.Equal(t, expectedEnd, sub.EndDate)
			assert.NotEmpty(t, sub.ID)
			assert.WithinDuration(t, time.Now(), sub.CreatedAt, time.Second)

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

	resp, err := usecase.Create(req)
	assert.NoError(t, err)
	assert.Equal(t, "Netflix", resp.Name)
	assert.Equal(t, 899, resp.Price)
	assert.Equal(t, "user-123", resp.UserID)
	assert.Equal(t, expectedStart, resp.StartDate)
	assert.Equal(t, expectedEnd, resp.EndDate)
}
