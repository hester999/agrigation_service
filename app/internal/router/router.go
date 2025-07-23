package router

import (
	"app/internal/handlers/service_handler"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(serviceHandler *service_handler.ServiceHandler) *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/services", serviceHandler.CreateHandler).Methods(http.MethodPost)
	api.HandleFunc("/services/{id}", serviceHandler.UpdateHandler).Methods(http.MethodPatch)
	api.HandleFunc("/services/{id}", serviceHandler.Replace).Methods(http.MethodPut)
	api.HandleFunc("/services/total", serviceHandler.TotalPrice).Methods(http.MethodPost)
	api.HandleFunc("/services", serviceHandler.GetAll).Methods(http.MethodGet)
	api.HandleFunc("/services/{id}", serviceHandler.GetByID).Methods(http.MethodGet)
	api.HandleFunc("/services/{id}", serviceHandler.Delete).Methods(http.MethodDelete)

	return r
}
