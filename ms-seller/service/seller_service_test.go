package service

import (
	"context"
	"database/sql"
	pb "ms-seller/pb"
	"ms-seller/repository"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestGetProductByID_Success(t *testing.T) {
	testDB, err := setupTestDB()
	if err != nil {
		t.Fatalf("Error setting up test database: %v", err)
	}
	defer testDB.Close()

	sellerRepository := repository.NewPostgresRepository(testDB)

	sellerHandler := NewSellerService(sellerRepository)

	_, err = testDB.Exec(`INSERT INTO products (seller_id, name, price, stock, category_id) VALUES (?, ?, ?, ?, ?)`,
		1, "Baju", 10.0, 12, 1)
	if err != nil {
		t.Fatalf("Error inserting test data into products table: %v", err)
	}

	testCases := []struct {
		name     string
		request  *pb.GetProductByIDRequest
		expected *pb.ProductResponse
		err      error
	}{
		{
			name: "Successful get product by ID",
			request: &pb.GetProductByIDRequest{
				ProductId: 1,
			},
			expected: &pb.ProductResponse{
				Productid:  1,
				SellerId:   1,
				Name:       "Baju",
				Price:      10.0,
				Stock:      12,
				CategoryId: 1,
				Discount:   0,
			},
			err: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			response, err := sellerHandler.GetProductByID(context.Background(), tc.request)

			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.expected, response)
		})
	}
}

func TestGetSellerByID_Success(t *testing.T) {
	testDB, err := setupTestDB()
	if err != nil {
		t.Fatalf("Error setting up test database: %v", err)
	}
	defer testDB.Close()

	_, err = testDB.Exec(`INSERT INTO address(id, address, regency, city) VALUES (?, ?, ?, ?)`,
		1, "Jakarta", "Jakarta", "Jakarta")
	if err != nil {
		t.Fatalf("Error inserting test data into sellers table: %v", err)
	}

	_, err = testDB.Exec(`INSERT INTO sellers(id, name, address_id, last_active) VALUES (?, ?, ?, ?)`,
		1, "kratos", 1, "0001-01-01T00:00:00Z")
	if err != nil {
		t.Fatalf("Error inserting test data into sellers table: %v", err)
	}

	sellerRepository := repository.NewPostgresRepository(testDB)
	sellerHandler := NewSellerService(sellerRepository)

	testCases := []struct {
		name     string
		request  *pb.GetSellerByIDRequest
		expected *pb.SellerDetailResponse
	}{
		{
			name: "Successful get seller by ID",
			request: &pb.GetSellerByIDRequest{
				SellerId: 1,
			},
			expected: &pb.SellerDetailResponse{
				SellerId:       1,
				Name:           "kratos",
				LastActive:     "0001-01-01T00:00:00Z",
				AddressName:    "Jakarta",
				AddressRegency: "Jakarta",
				AddressCity:    "Jakarta",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			response, err := sellerHandler.GetSellerByID(context.Background(), tc.request)

			if err != nil {
				t.Fatalf("Error calling GetSellerByID: %v", err)
			}

			if response == nil {
				t.Fatal("Expected non-nil response")
			}

			if !reflect.DeepEqual(response, tc.expected) {
				t.Errorf("Response does not match expected: got %v, want %v", response, tc.expected)
			}
		})
	}
}

func setupTestDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS sellers (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL UNIQUE,
			address_id INTEGER,
			last_active DATE
		);
		CREATE TABLE IF NOT EXISTS address (
			id INTEGER PRIMARY KEY,
			address TEXT NOT NULL,
			regency TEXT NOT NULL,
			city TEXT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY,
			seller_id INTEGER NOT NULL,
			name TEXT NOT NULL UNIQUE,
			price DECIMAL NOT NULL,
			stock INTEGER NOT NULL,
			category_id INTEGER NOT NULL,
			discount INTEGER NOT NULL DEFAULT 0,
			FOREIGN KEY (seller_id) REFERENCES sellers (id),
			FOREIGN KEY (category_id) REFERENCES categories (id)
		);
		CREATE TABLE IF NOT EXISTS categories (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL UNIQUE
		);
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
