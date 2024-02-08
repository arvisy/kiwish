package repository

import (
	"database/sql"
	"ms-seller/model"
)

// TODO: read by activity date range
type SellerRepository interface {
	// products
	CreateProduct(input *model.Product) (*model.Product, error)
	ReadAllProducts(sellerID int) ([]*model.Product, error)            // public
	ReadProductID(productID int) (*model.Product, error)               // public
	ReadProductCategory(categoryName string) ([]*model.Product, error) //public
	DeleteProduct(productID int, sellerID int) error
	UpdateProduct(input *model.Product) error

	// seller
	CreateSeller(input *model.Seller) (*model.Seller, error)
	CreateAddress(input *model.Address) (*model.Address, error)
	UpdateAddressID(addressID int, sellerID int) error
	ReadAllSellers() ([]*model.Seller, error)                      // public
	ReadSellerID(sellerID int) (*model.SellerDetail, error)        // public
	ReadSellerName(sellerName string) (*model.SellerDetail, error) // public
	UpdateAddress(input *model.Address) error
	UpdateName(sellerName string, sellerID int) error
	UpdateActivity(timestamp string, sellerID int) error // call when order finished
}

type postgresRepository struct {
	DB *sql.DB
}

func NewPostgresRepository(db *sql.DB) *postgresRepository {
	return &postgresRepository{
		DB: db,
	}
}
