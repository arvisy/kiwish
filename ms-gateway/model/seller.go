package model

type Product struct {
	ID          int     `json:"id,omitempty"`
	SellerID    int     `json:"seller_id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Stock       int     `json:"stock"`
	Category_id int     `json:"category_id"`
	Discount    int     `json:"discount"`
}

type ProductInput struct {
	Name        string  `json:"name"` //add validation
	Price       float32 `json:"price"`
	Stock       int     `json:"stock"`
	Category_id int     `json:"category_id"`
	Discount    int     `json:"discount"`
}

type SellerID struct {
	ID int `json:"id"`
}

type SellerName struct {
	Name string `json:"name"`
}

type SellerIDName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Seller struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name"`
	LastActive string `json:"last_active"`
}

type SellerDetail struct {
	ID         int     `json:"id,omitempty"`
	Name       string  `json:"name"`
	LastActive string  `json:"last_active"`
	Address    Address `json:"address"`
}
