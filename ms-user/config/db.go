package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectPostgresDB() *sql.DB {
	dsn := "host=127.0.0.1 user=postgres password=12345 dbname=user_service port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("Failed to connect to the database")
	}

	return db
}
