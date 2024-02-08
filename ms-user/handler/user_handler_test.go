package handler

import (
	"context"
	"database/sql"
	pb "ms-user/pb"
	"ms-user/repository"
	"regexp"
	"testing"

	"github.com/go-playground/assert/v2"
	_ "github.com/mattn/go-sqlite3"
	"github.com/redis/go-redis/v9"
)

func TestRegister_Success(t *testing.T) {
	testDB, err := setupTestDB()
	if err != nil {
		t.Fatalf("Error setting up test database: %v", err)
	}
	defer testDB.Close()

	userRepository := repository.NewUserRepository(testDB)

	Redis := redis.Client{}

	userHandler := NewUserHandler(*userRepository, &Redis)

	testCases := []struct {
		name     string
		request  *pb.RegisterRequest
		expected *pb.RegisterResponse
		err      error
	}{
		{
			name: "Successful registration",
			request: &pb.RegisterRequest{
				Name:     "John Doe",
				Email:    "john@example.com",
				Password: "password123",
			},
			expected: &pb.RegisterResponse{
				Name:  "John Doe",
				Email: "john@example.com",
			},
			err: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			response, err := userHandler.Register(context.Background(), tc.request)

			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.expected, response)
		})
	}
}

func TestIsValidEmail(t *testing.T) {
	testCases := []struct {
		email    string
		expected bool
	}{
		{"john@example.com", true},       // Valid email
		{"jane.doe@example.co.uk", true}, // Valid email
		{"invalid.email@", false},        // Invalid email (missing domain)
		{"@example.com", false},          // Invalid email (missing username)
		{"john@localhost", false},        // Invalid email (invalid domain)
		{"john.doe@example", false},      // Invalid email (missing top-level domain)
		{"john.doe@example.", false},     // Invalid email (empty top-level domain)
		{"john@example", false},          // Invalid email (missing top-level domain)
		{"john@example.c", false},        // Invalid email (top-level domain with less than 2 characters)
		{"john@.com", false},             // Invalid email (empty domain and top-level domain)
		{"@.com", false},                 // Invalid email (empty username, domain, and top-level domain)
	}

	for _, tc := range testCases {
		t.Run(tc.email, func(t *testing.T) {
			result := isValidEmail(tc.email)
			if result != tc.expected {
				t.Errorf("Expected isValidEmail(%s) to be %v, but got %v", tc.email, tc.expected, result)
			}
		})
	}
}

func setupTestDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY,
        name TEXT NOT NULL,
        email TEXT NOT NULL,
        password TEXT NOT NULL,
        role_id INTEGER NOT NULL
    );`)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
