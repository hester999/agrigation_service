package service_usecases

import "app/internal/model"

func (s *ServiceUsecases) GetByID(id string) (model.Service, error) {
	res, err := s.repo.GetByID(id)
	if err != nil {
		s.logger.Println("get service error: ", err)
		return model.Service{}, err
	}
	return res, nil
}

func (s *ServiceUsecases) GetAll(limit, offset int) (model.GetAllResponse, error) {
	if limit == 0 {
		limit = 10
	}
	res, err := s.repo.GetAll(limit, offset)
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
