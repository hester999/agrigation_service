package service_handler

import (
	"app/internal/apperr"
	"encoding/json"
	"errors"
	"net/http"
)

// TotalPrice @Summary      Calculate total service price
// @Description  Returns total sum of services by user and optional service name in date range
// @Tags         Services
// @Accept       json
// @Produce      json
// @Param        totalPriceRequest body dto.TotalRequestDTO true "Total price request"
// @Success      200  {object}  dto.ResponseTotalDTO
// @Failure      400  {object}  dto.ErrDTO400
// @Failure      404  {object}  dto.ErrDTO404
// @Failure      500  {object}  dto.ErrDTO500
// @Router       /services/total/{id} [post]
func (s *ServiceHandler) TotalPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	req := struct {
		ID          string `json:"id"`
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

	err = s.validateTotalPrice(req.ID, req.ServiceName, req.From, req.To)
	if err != nil {
		s.logger.Println("validateTotalPrice error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrDTO{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	res, err := s.usecase.GetTotalPrice(req.ID, req.ServiceName, req.From, req.To)

	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrDTO{
				Code:    http.StatusNotFound,
				Message: "not found",
			})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrDTO{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		})
		return

	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		TotalPrice int `json:"total_price"`
	}{res})
}

func (s *ServiceHandler) validateTotalPrice(id, name, from, to string) error {
	if id == "" {
		return apperr.ErrDataIsRequired
	}
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
