package repository

import (
	"database/sql"
	"errors"
	"ms-user/model"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) *UserRepository {
	return &UserRepository{DB: DB}
}

func (u *UserRepository) AddUser(user model.User) error {
	_, err := u.DB.Exec("INSERT INTO users(id, name, email, password, role_id) VALUES($1, $2, $3, $4, $5)", user.Id, user.Name, user.Email, user.Password)
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

func (u *UserRepository) Get(user model.User) error {
	rows, err := u.DB.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	} else {
		for rows.Next() {
			rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		}
	}

	return nil
}

func (u *UserRepository) Update(user *model.User) model.User {
	_, err := u.DB.Exec("UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4")
	if err != nil {
		panic(err)
	}

	return *user
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

	err := u.DB.QueryRow("SELECT * FROM users WHERE email=$1 AND password=$2", email, password).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
