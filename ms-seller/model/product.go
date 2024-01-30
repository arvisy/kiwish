package model

type Product struct {
	ID          uint    `json:"id,omitempty"`
	SellerID    uint    `json:"seller_id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Stock       int     `json:"stock"`
	Category_id int     `json:"category_id"`
	Discount    int     `json:"discount"`
}
