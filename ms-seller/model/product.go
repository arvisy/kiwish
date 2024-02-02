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
