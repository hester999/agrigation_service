package service_repo

import "app/internal/model"

func (s *ServiceRepo) Create(service model.Service) (model.CreateResponse, error) {

	query := `INSERT INTO services(id, name, price, user_id, start_date, end_date, created_at) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id, name, price, user_id, start_date, end_date, created_at`

	var tmp DTO

	err := s.db.Get(&tmp, query, service.ID, service.Name, service.Price, service.UserID, service.StartDate, service.EndDate, service.CreatedAt)

	if err != nil {
		s.logger.Println("create error:", err)
		return model.CreateResponse{}, err
	}

	res := model.CreateResponse{
		ID:        tmp.ID,
		Name:      tmp.Name,
		Price:     tmp.Price,
		UserID:    tmp.UserID,
		StartDate: tmp.StartDate,
		EndDate:   tmp.EndDate,
		CreatedAt: tmp.CreatedAt,
	}
	return res, nil
}
