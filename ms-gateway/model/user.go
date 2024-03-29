package model

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	RoleID    int    `json:"role_id"`
	AddressID int    `json:"address_id"`
}

type Address struct {
	Id      int    `json:"id,omitempty"`
	Address string `json:"address"`
	Regency string `json:"regency"`
	City    string `json:"city"`
}

type Role struct {
	Id   int    `json:"id"`
	Name string `json:"role"`
}

// response for swagger
type RegisterResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
