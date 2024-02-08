package repository

import (
	"context"
	"fmt"
	"ms-seller/model"
)

func (us *postgresRepository) CreateSeller(input *model.Seller) (*model.Seller, error) {
	query := `INSERT INTO sellers (id, name, address_id, last_active) 
	VALUES ($1, $2, $3, $4) RETURNING id`

	row := us.DB.QueryRow(query,
		input.ID,
		input.Name,
		input.AddressID,
		input.LastActive,
	)

	err := row.Scan(&input.ID)
	if err != nil {
		return nil, fmt.Errorf("create seller: %w", err)
	}

	return input, nil
}

func (us *postgresRepository) CreateAddress(input *model.Address) (*model.Address, error) {
	query := `INSERT INTO address (address, regency, city)
	VALUES ($1, $2, $3) RETURNING id;`

	row := us.DB.QueryRow(query,
		input.Name,
		input.Regency,
		input.City,
	)

	err := row.Scan(&input.ID)
	if err != nil {
		return nil, fmt.Errorf("create address: %w", err)
	}

	return input, nil
}

func (po *postgresRepository) ReadAllSellers() ([]*model.Seller, error) {
	query := `SELECT * FROM sellers`

	rows, err := po.DB.QueryContext(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sellers []*model.Seller

	for rows.Next() {
		var seller model.Seller
		err := rows.Scan(
			&seller.ID,
			&seller.Name,
			&seller.AddressID,
			&seller.LastActive,
		)
		if err != nil {
			return nil, fmt.Errorf("read all sellers: %w", err)
		}
		sellers = append(sellers, &seller)
	}

	return sellers, nil
}

func (po *postgresRepository) ReadSellerID(sellerID int) (*model.SellerDetail, error) {
	query := `SELECT s.id, s.name, s.last_active, a.address, a.regency, a.city 
			FROM sellers s
			INNER JOIN address a ON s.address_id = a.id
			WHERE s.id = $1;`

	var seller model.SellerDetail
	row := po.DB.QueryRow(query, sellerID)

	err := row.Scan(
		&seller.ID,
		&seller.Name,
		&seller.LastActive,
		&seller.Address.Name,
		&seller.Address.Regency,
		&seller.Address.City,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get seller by ID: %v", err)
	}

	return &seller, nil
}

func (po *postgresRepository) ReadSellerName(sellerName string) (*model.SellerDetail, error) {
	query := `SELECT s.id, s.name, s.last_active, a.address, a.regency, a.city 
			FROM sellers s
			INNER JOIN address a ON s.address_id = a.id
			WHERE s.name = $1;`

	var seller model.SellerDetail
	row := po.DB.QueryRow(query, sellerName)

	err := row.Scan(
		&seller.ID,
		&seller.Name,
		&seller.LastActive,
		&seller.Address.Name,
		&seller.Address.Regency,
		&seller.Address.City,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get seller by ID: %v", err)
	}

	return &seller, nil
}

func (po *postgresRepository) UpdateAddress(input *model.Address) error {
	query := `UPDATE address SET
				address = $1,
				regency = $2,
				city = $3
			WHERE id = $4;`

	_, err := po.DB.Exec(query, input.Name, input.Regency, input.City, input.ID)
	if err != nil {
		return err
	}

	return nil
}

func (po *postgresRepository) UpdateName(sellerName string, sellerID int) error {
	query := `UPDATE sellers SET 
				name = $1
			WHERE id = $2;
			`

	_, err := po.DB.Exec(query, sellerName, sellerID)
	if err != nil {
		return err
	}
	return nil
}

func (po *postgresRepository) UpdateAddressID(addressID int, sellerID int) error {
	query := `UPDATE sellers SET 
				address_id = $1,
			WHERE id = $2;
			`

	_, err := po.DB.Exec(query, addressID, sellerID)
	if err != nil {
		return err
	}
	return nil
}

func (po *postgresRepository) UpdateActivity(timestamp string, sellerID int) error {
	query := `UPDATE sellers SET 
				last_active = $1,
			WHERE id = $2;
			`

	_, err := po.DB.Exec(query, timestamp, sellerID)
	if err != nil {
		return err
	}
	return nil
}
