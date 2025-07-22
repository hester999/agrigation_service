package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connection(connectionStr string) (*sqlx.DB, error) {
	fmt.Println("Connecting to database ...")
	//dsn := os.Getenv("DATABASE_URL")
	//if dsn == "" {
	//	dsn = "postgres://admin:123@localhost:5432/vk?sslmode=disable"
	//}

	db, err := sqlx.Open("postgres", connectionStr)
	if err != nil {
		return nil, fmt.Errorf("could not connect to postgres database: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("could not ping postgres database: %w", err)
	}
	return db, nil
}
