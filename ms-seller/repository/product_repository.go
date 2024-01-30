package repository

import (
	"context"
	"database/sql"
	"fmt"
	"ms-seller/model"
)

type ProductRepository interface {
	Create(input *model.Product) (*model.Product, error)
	ReadAll(sellerID int) ([]*model.Product, error)
	ReadID(productID int, sellerID int) (*model.Product, error)
	Delete(productID int, sellerID int) error
	Update(input *model.Product) error
}

type postgresRepository struct {
	DB *sql.DB
}

func NewPostgresRepository(db *sql.DB) *postgresRepository {
	return &postgresRepository{
		DB: db,
	}
}

func (us *postgresRepository) Create(input *model.Product) (*model.Product, error) {
	query := `INSERT INTO products (seller_id, name, price, stock, category_id, discount) 
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	row := us.DB.QueryRow(query,
		input.SellerID,
		input.Name,
		input.Price,
		input.Stock,
		input.Category_id,
		input.Discount,
	)

	err := row.Scan(&input.ID)
	if err != nil {
		return nil, fmt.Errorf("create product: %w", err)
	}

	return input, nil
}

func (us *postgresRepository) ReadAll(sellerID int) ([]*model.Product, error) {
	query := `SELECT * FROM products WHERE seller_id = $1`

	rows, err := us.DB.QueryContext(context.Background(), query, sellerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*model.Product

	for rows.Next() {
		var product model.Product
		err := rows.Scan(
			&product.ID,
			&product.SellerID,
			&product.Name,
			&product.Price,
			&product.Stock,
			&product.Category_id,
			&product.Discount,
		)
		if err != nil {
			return nil, fmt.Errorf("read all product: %w", err)
		}
		products = append(products, &product)
	}

	return products, nil
}

func (us *postgresRepository) ReadID(productID int, sellerID int) (*model.Product, error) {
	query := `SELECT * FROM products WHERE id = $1 AND seller_id = $2;`

	var product model.Product
	row := us.DB.QueryRow(query, productID, sellerID)

	err := row.Scan(
		&product.ID,
		&product.SellerID,
		&product.Name,
		&product.Price,
		&product.Stock,
		&product.Category_id,
		&product.Discount,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get product by ID: %v", err)
	}

	return &product, nil
}

func (us *postgresRepository) Delete(productID int, sellerID int) error {
	query := `DELETE FROM products WHERE id = $1 AND seller_id = $2`

	result, err := us.DB.Exec(query, productID, sellerID)
	if err != nil {
		return err
	}
	if rows, _ := result.RowsAffected(); rows == 0 { // TODO: fix not found error, RowsAffected() not working
		return fmt.Errorf("no product deleted")
	}

	return nil
}

// TODO: tambah fitur update beberapa fild saja
func (t *postgresRepository) Update(input *model.Product) error {
	query := `UPDATE products SET 
				name = $1,
				price = $2,
				stock = $3,
				category_id = $4,
				discount = $5
			WHERE id = $6 AND seller_id = $7;
			`

	_, err := t.DB.Exec(query, input.Name, input.Price, input.Stock, input.Category_id, input.Discount, input.ID, input.SellerID)
	if err != nil {
		return err
	}
	return nil
}
