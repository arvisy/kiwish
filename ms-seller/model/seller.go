package model

type Seller struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name"`
	AddressID  int    `json:"address_id"`
	LastActive string `json:"last_active"`
}

type SellerDetail struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	LastActive string  `json:"last_active"`
	Address    Address `json:"address_id"`
}
