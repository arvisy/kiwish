package config

import (
	"database/sql"
)

func ConnectPostgresDB() *sql.DB {
	dsn := "host=127.0.0.1 user=postgres password=postgres dbname=user-service port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("Failed to connect to the database")
	}

	return db
}
