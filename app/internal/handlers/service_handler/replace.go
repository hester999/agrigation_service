package service_handler

import (
	"app/internal/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

// Replace @Summary      Полная замена подписки
// @Description  Полностью заменяет услугу по ID. Все поля обязательны для передачи.
//
//	Возвращает обновлённую сущность.
//
// @Tags         Services
// @Accept       json
// @Produce      json
// @Param        id   path      string           true  "ID услуги"
// @Param        body body      dto.CreateRequest true  "Полные данные новой подписки"
// @Success      200  {object}  dto.CreateResponse
// @Failure      400  {object}  dto.ErrDTO400
// @Failure      500  {object}  dto.ErrDTO500
// @Router       /services/{id} [put]
func (s *ServiceHandler) Replace(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	id := mux.Vars(r)["id"]
	req := struct {
		Name      string `json:"service_name"`
		Price     int    `json:"price"`
		UserID    string `json:"user_id"`
		StartDate string `json:"start_date"`
		Duration  int    `json:"duration"` //количество месяцев на которое выдается подписка
	}{}

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		s.logger.Println("Error decoding request body", err)
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

	res, err := s.usecase.Replace(id, modelReq)

	if err != nil {
		s.logger.Println("Error updating request", err)
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
