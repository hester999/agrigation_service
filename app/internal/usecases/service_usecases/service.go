package service_usecases

import (
	"app/internal/model"
	"github.com/google/uuid"
	"log"
	"strings"
	"time"
)

type ServiceUsecases struct {
	repo   Service
	logger *log.Logger
}

func NewService(repo Service, logger *log.Logger) *ServiceUsecases {
	return &ServiceUsecases{
		repo:   repo,
		logger: logger,
	}
}

//type UsecasesService interface {
//	Create(request model.CreateRequest) (model.CreateResponse, error)
//	Update(request model.UpdateRequest) (model.UpdateResponse, error)
//	GetByID(id string) (model.Service, error)
//	GetAll() (model.GetAllResponse, error)
//	GetTotalPrice(userID, serviceName string, from, to time.Time) (int, error)
//	DeleteByID(id string) error
//}

func (s *ServiceUsecases) Create(request model.CreateRequest) (model.CreateResponse, error) {

	id, err := uuid.NewRandom()

	if err != nil {
		s.logger.Println("generate uuid error", err)
		return model.CreateResponse{}, err
	}
	start, err := s.normalizeData(request.StartDate)
	if err != nil {
		s.logger.Println("parse start date error", err)
		return model.CreateResponse{}, err
	}

	newService := model.Service{
		ID:        id.String(),
		Name:      request.Name,
		Price:     request.Price,
		UserID:    request.UserID,
		StartDate: start,
		EndDate:   start.AddDate(0, request.Duration, 0).UTC(),
		CreatedAt: time.Now().UTC(),
	}

	res, err := s.repo.Create(newService)
	if err != nil {
		s.logger.Println("create service error:", err)
		return model.CreateResponse{}, err
	}
	return res, nil
}

func (s *ServiceUsecases) Update(id string, request *model.UpdateRequest) (model.UpdateResponse, error) {
	res, err := s.repo.GetByID(id)
	if err != nil {
		s.logger.Println("get service error: ", err)
		return model.UpdateResponse{}, err
	}

	err = s.compareUpdate(&res, request)
	if err != nil {
		return model.UpdateResponse{}, err
	}
	//start, err := s.normalizeData(request.StartDate)
	//if err != nil {
	//	s.logger.Println("parse start date error", err)
	//	return model.UpdateResponse{}, err
	//}

	//res.Name = request.Name
	//res.Price = request.Price
	//res.EndDate = start.AddDate(0, request.Duration, 0).UTC()
	res.CreatedAt = time.Now().UTC()

	result, err := s.repo.Update(res)
	if err != nil {
		s.logger.Printf("update service error: %v", err)
		return model.UpdateResponse{}, err
	}
	return result, nil

}

func (s *ServiceUsecases) compareUpdate(old *model.Service, new *model.UpdateRequest) error {
	var updatedStartDate time.Time
	useUpdatedStart := false

	if new.Name != nil {
		old.Name = *new.Name
	}
	if new.Price != nil {
		old.Price = *new.Price
	}
	if new.StartDate != nil {
		start, err := s.normalizeData(*new.StartDate)
		if err != nil {
			s.logger.Println("parse start date error", err)
			return err
		}
		updatedStartDate = start
		old.StartDate = start
		useUpdatedStart = true
	}

	if new.Duration != nil {
		var baseStart time.Time
		if useUpdatedStart {
			baseStart = updatedStartDate
		} else {
			baseStart = old.StartDate
		}
		old.EndDate = baseStart.AddDate(0, *new.Duration, 0).UTC()
	}

	return nil
}

func (s *ServiceUsecases) GetByID(id string) (model.Service, error) {
	res, err := s.repo.GetByID(id)
	if err != nil {
		s.logger.Println("get service error: ", err)
		return model.Service{}, err
	}
	return res, nil
}

func (s *ServiceUsecases) GetAll() (model.GetAllResponse, error) {
	res, err := s.repo.GetAll()
	if err != nil {
		s.logger.Println("get service error: ", err)
		return res, err // так как res не nil вернет []
	}
	return res, nil
}

func (s *ServiceUsecases) GetTotalPrice(userID, serviceName, from, to string) (int, error) {

	fromDate, err := s.normalizeData(from)
	if err != nil {
		s.logger.Println("parse from date error", err)
		return 0, err
	}
	toDate, err := s.normalizeData(to)
	if err != nil {
		s.logger.Println("parse to date error", err)
		return 0, err
	}

	res, err := s.repo.GetTotalPrice(userID, serviceName, fromDate, toDate)
	if err != nil {
		s.logger.Println("get service error: ", err)
		return 0, err
	}
	return res, nil

}

func (s *ServiceUsecases) DeleteByID(id string) error {

	err := s.repo.DeleteByID(id)
	if err != nil {
		s.logger.Println("delete service error: ", err)
		return err
	}
	return nil
}

func (s *ServiceUsecases) normalizeData(data string) (time.Time, error) {
	data = strings.TrimSpace(data)
	return time.Parse("2006-01-02", data)
}
