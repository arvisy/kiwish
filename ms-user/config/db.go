package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	Host     = os.Getenv("DB_HOST")
	Port     = os.Getenv("DB_PORT")
	User     = os.Getenv("DB_USER")
	Password = os.Getenv("DB_PASSWORD")
	Database = os.Getenv("DB_NAME")
	SSLMode  = os.Getenv("DB_SSL")
)

func ConnectPostgresDB() *sql.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", Host, User, Password, Database, Port, SSLMode)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("Failed to connect to the database")
	}

	db.SetMaxOpenConns(10)

	db.SetMaxIdleConns(5)

	return db
}
