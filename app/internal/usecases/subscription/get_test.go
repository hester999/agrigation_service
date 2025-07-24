package subscription_test

import (
	"app/internal/model"
	usecases "app/internal/usecases/subscription"
	"app/internal/usecases/subscription/mock"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"log"
	"os"
	"testing"
	"time"
)

func TestSubscriptionUsecases_GetAll_ZeroLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMocksubscriptionRepo(ctrl)
	logger := log.New(os.Stdout, "[test] ", log.LstdFlags)
	usecase := usecases.NewSubscriptionUsecases(mockRepo, logger)

	expected := model.GetAllResponse{
		Data: []model.Subscription{
			{ID: "1", Name: "more.tv"},
			{ID: "2", Name: "yandex_plus"},
		},
	}

	mockRepo.EXPECT().
		GetAll(10, 0).
		Return(expected, nil)

	result, err := usecase.GetAll(0, 0)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestSubscriptionUsecases_GetTotalPrice_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMocksubscriptionRepo(ctrl)
	logger := log.New(os.Stdout, "[test] ", log.LstdFlags)
	uc := usecases.NewSubscriptionUsecases(mockRepo, logger)

	from := "2025-01-01"
	to := "2025-07-01"
	fromDate, _ := time.Parse("2006-01-02", from)
	toDate, _ := time.Parse("2006-01-02", to)

	mockRepo.EXPECT().
		GetTotalPrice("user-1", "Netflix", fromDate, toDate).
		Return(1999, nil)

	total, err := uc.GetTotalPrice("user-1", "Netflix", from, to)
	assert.NoError(t, err)
	assert.Equal(t, 1999, total)
}

func TestSubscriptionUsecases_GetTotalPrice_InvalidFromDate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMocksubscriptionRepo(ctrl)
	logger := log.New(os.Stdout, "[test] ", log.LstdFlags)
	uc := usecases.NewSubscriptionUsecases(mockRepo, logger)

	// Некорректная дата
	from := "invalid-date"
	to := "2025-07-01"

	total, err := uc.GetTotalPrice("user-1", "Netflix", from, to)

	assert.Error(t, err)
	assert.Equal(t, 0, total)
}

func TestSubscriptionUsecases_GetTotalPrice_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMocksubscriptionRepo(ctrl)
	logger := log.New(os.Stdout, "[test] ", log.LstdFlags)
	uc := usecases.NewSubscriptionUsecases(mockRepo, logger)

	from := "2025-01-01"
	to := "2025-07-01"
	fromDate, _ := time.Parse("2006-01-02", from)
	toDate, _ := time.Parse("2006-01-02", to)

	mockRepo.EXPECT().
		GetTotalPrice("user-1", "Netflix", fromDate, toDate).
		Return(0, errors.New("db error"))

	total, err := uc.GetTotalPrice("user-1", "Netflix", from, to)

	assert.Error(t, err)
	assert.Equal(t, 0, total)
}
