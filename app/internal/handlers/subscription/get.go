package subscription

import (
	"app/internal/apperr"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ResponseDTO struct {
	Data []SubscriptionDTO `json:"data"`
}

// GetByID @Summary      Получить подписку по ID
// @Description  Возвращает подписку по её идентификатору
// @Tags         Subscribtions
// @Produce      json
// @Param        id path string true "ID подписки"
// @Success      200 {object} dto.ServiceDTO
// @Failure      400 {object} dto.ErrDTO400 "ID не передан"
// @Failure      404 {object} dto.ErrDTO404 "Подписка не найдена"
// @Failure      500 {object} dto.ErrDTO500 "Внутренняя ошибка"
// @Router       /api/v1/subscriptions/{id} [get]
func (s *SubscriptionHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	id := mux.Vars(r)["id"]

	if id == "" {
		s.logger.Println("missing id")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrDTO{
			Message: "id is required",
			Code:    http.StatusBadRequest,
		})
		return
	}

	res, err := s.usecase.GetByID(id)
	if err != nil {
		s.logger.Println("error get by id", err)
		if errors.Is(err, apperr.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrDTO{
				Message: "service not found",
				Code:    http.StatusNotFound,
			})
			return
		}
		s.logger.Println("error get by id", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrDTO{
			Message: "internal server error",
			Code:    http.StatusInternalServerError,
		})
		return
	}
	resp := SubscriptionDTO{
		res.ID,
		res.Name,
		res.Price,
		res.UserID,
		res.StartDate,
		res.EndDate,
		res.CreatedAt,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// GetAll  @Summary      Получить все подписки
// @Description  Возвращает список всех подписок с пагинацией
// @Tags         Subscribtions
// @Produce      json
// @Param        limit  query int false "Ограничение количества результатов"
// @Param        offset query int false "Смещение для пагинации"
// @Success      200 {object} dto.ResponseDTO
// @Failure      400 {object} subscription.ErrDTO "Некорректные limit или offset"
// @Failure      404 {object} subscription.ErrDTOArr "Подписки не найдены"
// @Failure      500 {object} subscription.ErrDTO "Внутренняя ошибка"
// @Router       /api/v1/subscriptions [get]
func (s *SubscriptionHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var limit, offset int
	var err error

	limitStr := r.URL.Query().Get("limit")
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			s.logger.Println("invalid limit", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrDTO{
				Message: "internal server error",
				Code:    http.StatusInternalServerError,
			})
			return
		}
	}

	offsetStr := r.URL.Query().Get("offset")
	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			s.logger.Println("invalid offset", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrDTO{
				Message: "internal server error",
				Code:    http.StatusInternalServerError,
			})
			return
		}
	}

	err = s.validateLimitOffset(limit, offset)
	if err != nil {
		s.logger.Println("invalid offset,limit", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrDTO{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}
	res, err := s.usecase.GetAll(limit, offset)
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			s.logger.Println("services not found", err)
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrDTOArr{
				Message: "services not found",
				Code:    http.StatusNotFound,
				Data:    []interface{}{},
			})
			return
		}
		s.logger.Println("error get all", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrDTO{
			Message: "internal server error",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	resp := ResponseDTO{
		Data: make([]SubscriptionDTO, 0, len(res.Data)),
	}

	for _, service := range res.Data {
		resp.Data = append(resp.Data, SubscriptionDTO{
			ID:        service.ID,
			Name:      service.Name,
			Price:     service.Price,
			UserID:    service.UserID,
			StartDate: service.StartDate,
			EndDate:   service.EndDate,
			CreatedAt: service.CreatedAt,
		})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (s *SubscriptionHandler) validateLimitOffset(limit, offset int) error {
	if limit < 0 {
		return apperr.ErrInvalidLimit
	}
	if offset < 0 {
		return apperr.ErrInvalidOffset
	}
	return nil
}
