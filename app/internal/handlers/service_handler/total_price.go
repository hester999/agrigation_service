package service_handler

import (
	"app/internal/apperr"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

func (s *ServiceHandler) TotalPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	id := mux.Vars(r)["id"]

	req := struct {
		ServiceName string `json:"service_name"`
		From        string `json:"from"`
		To          string `json:"to"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		s.logger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrDTO{
			Code:    http.StatusBadRequest,
			Message: "invalid JSON",
		})
		return
	}

	err = s.validateTotalPrice(req.ServiceName, req.From, req.To)
	if err != nil {
		s.logger.Println("validateTotalPrice error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrDTO{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	res, err := s.usecase.GetTotalPrice(id, req.ServiceName, req.From, req.To)

	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrDTO{
				Code:    http.StatusNotFound,
				Message: "user not found",
			})
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		TotalPrice int `json:"total_price"`
	}{res})
}

func (s *ServiceHandler) validateTotalPrice(name, from, to string) error {
	if name == "" {
		return apperr.ErrNameIsRequired
	}

	if from == "" {
		return apperr.ErrDataIsRequired
	}
	if to == "" {
		return apperr.ErrDataIsRequired
	}
	return nil
}
