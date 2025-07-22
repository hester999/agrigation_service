package service_handler

import (
	"app/internal/apperr"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

type ResponseDTO struct {
	Data []ServiceDTO `json:"data"`
}

func (s *ServiceHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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
	resp := ServiceDTO{
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

func (s *ServiceHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res, err := s.usecase.GetAll()
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			s.logger.Println("services not found", err)
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrDTO{
				Message: "services not found",
				Code:    http.StatusNotFound,
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
		Data: make([]ServiceDTO, 0, len(res.Data)),
	}

	for _, service := range res.Data {
		resp.Data = append(resp.Data, ServiceDTO{
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
