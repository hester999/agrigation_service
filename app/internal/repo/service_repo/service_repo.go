package service_repo

import (
	"app/internal/apperr"
	"app/internal/model"
	"database/sql"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type ServiceRepo struct {
	db     *sqlx.DB
	logger *log.Logger
}

func NewServiceRepo(db *sqlx.DB, logger *log.Logger) *ServiceRepo {
	return &ServiceRepo{
		db:     db,
		logger: logger,
	}
}

type DTO struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Price     int       `db:"price"`
	UserID    string    `db:"user_id"`
	StartDate time.Time `db:"start_date"`
	EndDate   time.Time `db:"end_date"`
	CreatedAt time.Time `db:"created_at"`
}

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

func (s *ServiceRepo) Update(service model.Service) (model.UpdateResponse, error) {

	query := `UPDATE services SET id = $1, name = $2, price = $3, user_id = $4, start_date = $5, end_date = $6, created_at = $7 WHERE id = $1 RETURNING id, name, price, user_id, start_date, end_date, created_at`

	var tmp DTO
	err := s.db.Get(&tmp, query, service.ID, service.Name, service.Price, service.UserID, service.StartDate, service.EndDate, service.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Println("service not found")
			return model.UpdateResponse{}, apperr.ErrNotFound
		}
		s.logger.Println("update err:", err)
		return model.UpdateResponse{}, err
	}

	res := model.UpdateResponse{
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

func (s *ServiceRepo) GetByID(id string) (model.Service, error) {

	query := `SELECT id, name, price, user_id, start_date, end_date, created_at from services WHERE id = $1 `

	var tmp DTO
	err := s.db.Get(&tmp, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Println("GetByID error:", err)
			return model.Service{}, apperr.ErrNotFound
		}
		s.logger.Println("GetByID error:", err)
		return model.Service{}, err
	}
	res := model.Service{
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

func (s *ServiceRepo) GetAll() (model.GetAllResponse, error) {

	query := `SELECT id, name, price, user_id, start_date, end_date, created_at from services;`

	var tmp []DTO

	err := s.db.Select(&tmp, query)

	if err != nil {
		s.logger.Println("GetAll error:", err)
		return model.GetAllResponse{Data: make([]model.Service, 0)}, err
	}

	if len(tmp) == 0 {
		s.logger.Println("Not found services")
		return model.GetAllResponse{Data: make([]model.Service, 0)}, apperr.ErrNotFound
	}

	res := model.GetAllResponse{
		Data: make([]model.Service, 0, len(tmp)),
	}

	for _, service := range tmp {
		res.Data = append(res.Data, model.Service{
			ID:        service.ID,
			Name:      service.Name,
			Price:     service.Price,
			UserID:    service.UserID,
			StartDate: service.StartDate,
			EndDate:   service.EndDate,
			CreatedAt: service.CreatedAt,
		})
	}
	return res, nil
}

func (s *ServiceRepo) DeleteByID(id string) error {
	query := `DELETE FROM services WHERE id = $1`

	_, err := s.db.Exec(query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Println("Delete error:", err)
			return apperr.ErrNotFound
		}
		s.logger.Println("Delete error:", err)
		return apperr.ErrNotFound

	}
	return nil
}

func (s *ServiceRepo) GetTotalPrice(userID, serviceName string, from, to time.Time) (int, error) {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	query := builder.
		Select("COALESCE(SUM(price), 0)").
		From("services").
		Where(squirrel.And{
			squirrel.GtOrEq{"start_date": from},
			squirrel.LtOrEq{"start_date": to},
		})

	if userID != "" {
		query = query.Where(squirrel.Eq{"user_id": userID})
	}
	if serviceName != "" {
		query = query.Where(squirrel.Eq{"name": serviceName})
	}

	sqlStr, args, err := query.ToSql()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Println("GetTotalPrice exec error:", err)
			return 0, apperr.ErrNotFound
		}
		s.logger.Println("GetTotalPrice build error:", err)
		return 0, err
	}

	var total int
	err = s.db.Get(&total, sqlStr, args...)
	if err != nil {
		s.logger.Println("GetTotalPrice exec error:", err)
		return 0, err
	}

	return total, nil
}
