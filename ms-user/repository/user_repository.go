package repository

import (
	"database/sql"
	"errors"
	"log"
	"ms-user/model"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) *UserRepository {
	return &UserRepository{DB: DB}
}

func (u *UserRepository) AddUser(user model.User) error {
	_, err := u.DB.Exec("INSERT INTO users(name, email, password, role_id) VALUES($1, $2, $3, $4)", user.Name, user.Email, user.Password, user.RoleID)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) SetAddressCustomer(user model.User, addressID int) error {
	query := "UPDATE users SET address_id=$1 WHERE id=$2"
	_, err := u.DB.Exec(query, addressID, user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetAddressID(userID int) (int, error) {
	var addressID int
	err := u.DB.QueryRow("SELECT address_id FROM users WHERE id=$1", userID).Scan(&addressID)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	return addressID, nil
}

func (u *UserRepository) UpdateAddress(addressID int, address model.Address) error {
	query := "UPDATE address SET address=$1, regency=$2, city=$3 WHERE id=$4"
	_, err := u.DB.Exec(query, address.Address, address.Regency, address.City, addressID)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) AddAddress(address model.Address) (int, error) {
	var id int
	err := u.DB.QueryRow("INSERT INTO address(address, regency, city) VALUES($1, $2, $3) RETURNING id", address.Address, address.Regency, address.City).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *UserRepository) GetCustomer(user *model.User) error {
	query := "SELECT id, name, email, password, role_id FROM users WHERE id = $1"
	err := u.DB.QueryRow(query, user.Id).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.RoleID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}

	return nil
}

func (u *UserRepository) UpdateCustomer(userID int, customer model.User) error {
	query := "UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4"
	_, err := u.DB.Exec(query, customer.Name, customer.Email, customer.Password, userID)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) Delete(userID int) error {
	_, err := u.DB.Exec("DELETE FROM users WHERE id=$1", userID)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) FindByCredentials(email, password string) (*model.User, error) {
	var user model.User

	err := u.DB.QueryRow("SELECT id, name, email, password, role_id FROM users WHERE email=$1", email).
		Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.RoleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetUserAdmin(userID int) (*model.User, error) {
	query := "SELECT id, name, email, password, role_id FROM users WHERE id=$1"
	row := u.DB.QueryRow(query, userID)

	var user model.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.RoleID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetAllCustomerAdmin() ([]*model.User, error) {
	query := "SELECT id, name, email, password, role_id FROM users"
	rows, err := u.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.RoleID)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepository) UpdateCustomerAdmin(userID int, customer model.User) error {
	query := "UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4"
	_, err := u.DB.Exec(query, customer.Name, customer.Email, customer.Password, userID)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) DeleteCustomer(userID int) error {
	_, err := u.DB.Exec("DELETE FROM users WHERE id=$1", userID)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetSellerAdmin(sellerID int) (*model.User, error) {
	query := "SELECT id, name, email, password, role_id FROM users WHERE id=$1 AND role_id=$2"
	row := u.DB.QueryRow(query, sellerID, 3)

	var user model.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.RoleID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetAllSellerAdmin() ([]*model.User, error) {
	query := "SELECT id, name, email, password, role_id FROM users WHERE role_id=$1"
	rows, err := u.DB.Query(query, 3)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.RoleID)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepository) DeleteSellerAdmin(sellerID int) error {
	query := "UPDATE users SET role_id=$1 WHERE id=$2"
	_, err := u.DB.Exec(query, 2, sellerID)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) CreateSeller(customerID int) error {
	query := "UPDATE users SET role_id=$1 WHERE id=$2"
	_, err := u.DB.Exec(query, 3, customerID)
	if err != nil {
		return err
	}

	return nil
}
