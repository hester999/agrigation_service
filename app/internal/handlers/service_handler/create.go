package service_handler

import (
	"app/internal/apperr"
	"app/internal/model"
	"encoding/json"
	"net/http"
	"time"
	"unicode/utf8"
)

func (s *ServiceHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	req := struct {
		Name      string `json:"service_name"`
		Price     int    `json:"price"`
		UserID    string `json:"user_id"`
		StartDate string `json:"start_date"`
		Duration  int    `json:"duration"` //количество месяцев на которое выдается подписка
	}{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		s.logger.Println("Error in decoding request body", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrDTO{
			Message: "invalid JSON",
			Code:    http.StatusBadRequest,
		})
		return
	}

	err = s.validateCreate(req.Name, req.Price, req.UserID, req.StartDate, req.Duration)
	if err != nil {
		s.logger.Println("Error in validating request body", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrDTO{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	modelReq := model.CreateRequest{
		Name:      req.Name,
		Price:     req.Price,
		UserID:    req.UserID,
		StartDate: req.StartDate,
		Duration:  req.Duration,
	}

	res, err := s.usecase.Create(modelReq)
	if err != nil {
		s.logger.Println("Error in creating request", err)
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

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}

func (s *ServiceHandler) validateCreate(name string, price int, userID string, startDate string, duration int) error {
	if utf8.RuneCountInString(name) > 50 {
		return apperr.ErrNameTooLong
	}

	if name == "" {
		return apperr.ErrNameIsRequired
	}
	if price <= 0 {
		return apperr.ErrInvalidPrice
	}

	if userID == "" {
		return apperr.ErrUserIsRequired
	}

	if startDate == "" {
		return apperr.ErrDataIsRequired
	}

	if duration <= 0 {
		return apperr.ErrInvalidDuration
	}
	return nil
}
