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
		panic(err)
	}

	return nil
}

func (u *UserRepository) SetAddressCustomer(user model.User, addressID int) error {
	query := "UPDATE users SET address_id=$1 WHERE id=$2"
	_, err := u.DB.Exec(query, addressID, user.Id)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (u *UserRepository) UpdateAddress(addressID int, address model.Address) error {
	query := "UPDATE address SET address=$1, regency=$2, city=$3 WHERE id=$4"
	_, err := u.DB.Exec(query, address.Address, address.Regency, address.City, addressID)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (u *UserRepository) AddAddress(address model.Address) (int, error) {
	var id int
	err := u.DB.QueryRow("INSERT INTO address(address, regency, city) VALUES($1, $2, $3) RETURNING id", address.Address, address.Regency, address.City).Scan(&id)
	if err != nil {
		log.Fatal(err)
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
		panic(err)
	}

	return nil
}

func (u *UserRepository) Delete(userID int) error {
	_, err := u.DB.Exec("DELETE FROM users WHERE id=$1", userID)
	if err != nil {
		panic(err)
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
