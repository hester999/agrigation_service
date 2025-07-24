package subscription

import (
	"app/internal/apperr"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

// Delete @Summary      Удаление подписки
// @Description  Удаляет подписку по ID
// @Tags         Subscribtions
// @Produce      json
// @Param        id path string true "ID подписки"
// @Success      200 "Подписка успешно удалена"
// @Failure      404 {object} dto.ErrDTO404 "Подписка не найдена"
// @Failure      500 {object} dto.ErrDTO500 "Внутренняя ошибка"
// @Router       /api/v1/subscriptions/{id} [delete]
func (s *SubscriptionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	id := mux.Vars(r)["id"]

	err := s.usecase.DeleteByID(id)
	if err != nil {
		if errors.Is(err, apperr.ErrNotFound) {
			s.logger.Println("delete error:", err)
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrDTO{
				Message: "service not found",
				Code:    http.StatusNotFound,
			})
			return
		}
		s.logger.Println("delete error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrDTO{
			Message: "internal server error",
			Code:    http.StatusInternalServerError,
		})
		return
	}
	w.WriteHeader(http.StatusOK)
}
