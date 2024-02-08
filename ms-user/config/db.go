package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
}

func ConnectPostgresDB() *sql.DB {
	LoadEnv()

	Host := os.Getenv("DB_HOST")
	Port := os.Getenv("DB_PORT")
	User := os.Getenv("DB_USER")
	Password := os.Getenv("DB_PASSWORD")
	Database := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", Host, User, Password, Database, Port)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("Failed to connect to the database")
	}

	db.SetMaxOpenConns(10)

	db.SetMaxIdleConns(5)

	return db
}
