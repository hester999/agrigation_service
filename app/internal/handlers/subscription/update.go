package subscription

import (
	"app/internal/apperr"
	"app/internal/model"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"unicode/utf8"
)

// UpdateHandler @Summary      Обновление подписки
// @Description  Позволяет обновить один или несколько параметров подписки (название, цену, дату начала, длительность).
//
//	Если передаётся `start_date`, то обязательно указывать `duration` — и наоборот.
//
// @Tags         Subscribtions
// @Accept       json
// @Produce      json
// @Param        id   path      string       true  "ID услуги"
// @Param        body body      dto.UpdateDTO    true  "Обновляемые поля"
// @Success      200  {object}  dto.UpdateResponseDTO
// @Failure      400  {object}  dto.ErrDTO400
// @Failure      404  {object}  dto.ErrDTO404
// @Failure      500  {object}  dto.ErrDTO500
// @Router       /subscriptions/{id} [patch]
func (s *SubscriptionHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id := mux.Vars(r)["id"]

	req := struct {
		Name      *string `json:"service_name"`
		Price     *int    `json:"price"`
		StartDate *string `json:"start_date"`
		Duration  *int    `json:"duration"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.logger.Println("error decoding update request", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrDTO{
			Message: "Bad Request",
			Code:    http.StatusBadRequest,
		})
		return
	}

	if err := s.validationUpdate(req.Name, req.Price, req.StartDate, req.Duration); err != nil {
		s.logger.Println("validation error:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrDTO{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	update := model.UpdateRequest{
		Name:      req.Name,
		Price:     req.Price,
		StartDate: req.StartDate,
		Duration:  req.Duration,
	}

	res, err := s.usecase.Update(id, &update)
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrDTO{
				Message: "Subscription not found",
				Code:    http.StatusNotFound,
			})
			return
		}
		s.logger.Println("internal error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrDTO{
			Message: "Internal Server Error",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	resp := struct {
		ID        string    `json:"id"`
		Name      string    `json:"service_name"`
		Price     int       `json:"price"`
		UserID    string    `json:"user_id"`
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		CreatedAt time.Time `json:"created_at"`
	}{
		res.ID, res.Name, res.Price, res.UserID, res.StartDate, res.EndDate, res.CreatedAt,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (s *SubscriptionHandler) validationUpdate(name *string, price *int, startDate *string, duration *int) error {
	if name != nil && utf8.RuneCountInString(*name) > 50 {
		return apperr.ErrNameTooLong
	}
	if price != nil && *price <= 0 {
		return apperr.ErrInvalidPrice
	}
	if duration != nil && *duration <= 0 {
		return apperr.ErrInvalidDuration
	}
	if (startDate != nil && duration == nil) || (startDate == nil && duration != nil) {
		return apperr.ErrStartDateDurationPair
	}
	return nil
}
