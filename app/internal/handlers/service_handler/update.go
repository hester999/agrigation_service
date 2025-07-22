package service_handler

import (
	"app/internal/apperr"
	"app/internal/model"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

func (s *ServiceHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	id := mux.Vars(r)["id"]
	req := struct {
		Name      *string `json:"name"`
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
				Message: "Not Found",
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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}

//func (s *ServiceHandler) validationUpdate(name, startDate string, price, duration int) error {
//	if name == "" {
//		return apperr.ErrNameIsRequired
//	}
//	if startDate == "" {
//		return apperr.ErrDataIsRequired
//	}
//	if duration <= "" {
//		return apperr.ErrDataIsRequired
//	}
//}
