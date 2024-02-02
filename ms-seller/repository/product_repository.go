package repository

import (
	"context"
	"fmt"
	"ms-seller/model"
)

func (us *postgresRepository) CreateProduct(input *model.Product) (*model.Product, error) {
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

// tambahin category name
func (po *postgresRepository) ReadAllProducts(sellerID int) ([]*model.Product, error) {
	query := `SELECT * FROM products WHERE seller_id = $1`

	rows, err := po.DB.QueryContext(context.Background(), query, sellerID)
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

// tambahin category name
func (po *postgresRepository) ReadProductID(productID int) (*model.Product, error) {
	query := `SELECT * FROM products WHERE id = $1;`

	var product model.Product
	row := po.DB.QueryRow(query, productID)

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

// tambahin category name
func (po *postgresRepository) ReadProductCategory(categoryName string) ([]*model.Product, error) {
	query := `SELECT products.id, products.seller_id, products.name, products.price, products.stock, products.category_id, products.discount FROM products
			JOIN categories ON products.category_id = categories.id
			WHERE categories."name" = $1`

	rows, err := po.DB.QueryContext(context.Background(), query, categoryName)
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

func (po *postgresRepository) DeleteProduct(productID int, sellerID int) error {
	query := `DELETE FROM products WHERE id = $1 AND seller_id = $2`

	result, err := po.DB.Exec(query, productID, sellerID)
	if err != nil {
		return err
	}
	if rows, _ := result.RowsAffected(); rows == 0 { // TODO: fix not found error, RowsAffected() not working
		return fmt.Errorf("no product deleted")
	}

	return nil
}

// TODO: tambah fitur update beberapa fild saja
func (po *postgresRepository) UpdateProduct(input *model.Product) error {
	query := `UPDATE products SET 
				name = $1,
				price = $2,
				stock = $3,
				category_id = $4,
				discount = $5
			WHERE id = $6 AND seller_id = $7;
			`

	_, err := po.DB.Exec(query, input.Name, input.Price, input.Stock, input.Category_id, input.Discount, input.ID, input.SellerID)
	if err != nil {
		return err
	}
	return nil
}
