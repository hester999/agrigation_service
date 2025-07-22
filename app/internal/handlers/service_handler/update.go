package service_handler

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

func (s *ServiceHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	id := mux.Vars(r)["id"]
	req := struct {
		Name      *string `json:"service_name"`
		Price     *int    `json:"price"`
		StartDate *string `json:"start_date"`
		Duration  *int    `json:"duration"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		s.logger.Println("error decoding update request", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrDTO{
			Message: "Bad Request",
			Code:    http.StatusBadRequest,
		})
		return
	}

	err = s.validationUpdate(req.Name, req.Price, req.Duration)
	if err != nil {
		s.logger.Println("error validating update request", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrDTO{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	modeUpdate := model.UpdateRequest{
		Name:      req.Name,
		Price:     req.Price,
		StartDate: req.StartDate,
		Duration:  req.Duration,
	}

	res, err := s.usecase.Update(id, &modeUpdate)
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			s.logger.Println("service not found")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrDTO{
				Message: "service not found",
				Code:    http.StatusNotFound,
			})
			return
		}
		s.logger.Println("internal error", err)
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
	}{res.ID, res.Name, res.Price, res.UserID, res.StartDate, res.EndDate, res.CreatedAt}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}

func (s *ServiceHandler) validationUpdate(name *string, price, duration *int) error {
	if name != nil && utf8.RuneCountInString(*name) > 50 {
		return apperr.ErrNameTooLong
	}
	if price != nil && *price <= 0 {
		return apperr.ErrInvalidPrice
	}
	if duration != nil && *duration <= 0 {
		return apperr.ErrInvalidDuration
	}
	return nil
}
