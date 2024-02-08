package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

// need to close db w/db.Close() later
func Open(config PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", config.String())
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	db.SetMaxOpenConns(10)

	db.SetMaxIdleConns(5)

	return db, nil

}

// get db from env
func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL"),
	}
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s database=%s sslmode=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SSLMode)
}

// func Config() *pgxpool.Config {
// 	const defaultMaxConns = int32(4)
// 	const defaultMinConns = int32(0)
// 	const defaultMaxConnLifetime = time.Hour
// 	const defaultMaxConnIdleTime = time.Minute * 30
// 	const defaultHealthCheckPeriod = time.Minute
// 	const defaultConnectTimeout = time.Second * 5

// 	// Your own Database URL
// 	const DATABASE_URL string = "postgres://postgres:12345678@localhost:5432/postgres?"

// 	dbConfig, err := pgxpool.ParseConfig(DATABASE_URL)
// 	if err != nil {
// 		log.Fatal("Failed to create a config, error: ", err)
// 	}

// 	dbConfig.MaxConns = defaultMaxConns
// 	dbConfig.MinConns = defaultMinConns
// 	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
// 	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
// 	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
// 	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

// 	dbConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
// 		log.Println("Before acquiring the connection pool to the database!!")
// 		return true
// 	}

// 	dbConfig.AfterRelease = func(c *pgx.Conn) bool {
// 		log.Println("After releasing the connection pool to the database!!")
// 		return true
// 	}

// 	dbConfig.BeforeClose = func(c *pgx.Conn) {
// 		log.Println("Closed the connection pool to the database!!")
// 	}

// 	return dbConfig
// }
