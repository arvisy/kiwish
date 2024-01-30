package repository

import (
	"database/sql"
	"errors"
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

func (u *UserRepository) AddAddress(address model.Address) error {
	_, err := u.DB.Exec("INSERT INTO address(id, user_id, address, regency, city) VALUES($1, $2, $3, $4, $5)", address.Id, address.UserID, address.Address, address.Regency, address.City)
	if err != nil {
		panic(err)
	}

	return nil
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
	_, err := u.DB.Exec(query)
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
