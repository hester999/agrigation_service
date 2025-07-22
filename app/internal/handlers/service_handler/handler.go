package service_handler

import (
	"log"
)

type ServiceHandler struct {
	usecase Service
	logger  *log.Logger
}

type ErrDTO struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewServiceRepo(usecases Service, log *log.Logger) *ServiceHandler {
	return &ServiceHandler{
		usecase: usecases,
		logger:  log,
	}
}
