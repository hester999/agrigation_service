package service_repo

import (
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
