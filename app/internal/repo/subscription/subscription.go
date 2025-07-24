package subscription

import (
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type SubscriptionRepo struct {
	db     *sqlx.DB
	logger *log.Logger
}

func NewSubscriptionRepo(db *sqlx.DB, logger *log.Logger) *SubscriptionRepo {
	return &SubscriptionRepo{
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
