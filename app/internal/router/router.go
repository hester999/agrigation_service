package router

import (
	"app/internal/handlers/subscription"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(serviceSubscriptions *subscription.SubscriptionHandler) *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/subscriptions", serviceSubscriptions.CreateHandler).Methods(http.MethodPost)
	api.HandleFunc("/subscriptions/{id}", serviceSubscriptions.UpdateHandler).Methods(http.MethodPatch)
	api.HandleFunc("/subscriptions/{id}", serviceSubscriptions.Replace).Methods(http.MethodPut)
	api.HandleFunc("/subscriptions/total", serviceSubscriptions.TotalPrice).Methods(http.MethodPost)
	api.HandleFunc("/subscriptions", serviceSubscriptions.GetAll).Methods(http.MethodGet)
	api.HandleFunc("/subscriptions/{id}", serviceSubscriptions.GetByID).Methods(http.MethodGet)
	api.HandleFunc("/subscriptions/{id}", serviceSubscriptions.Delete).Methods(http.MethodDelete)

	return r
}
