package model

type OrderDetail struct {
	ID         string  `bson:"_id,omitempty"`
	ProductID  int64   `bson:"product_id"`
	Quantity   int64   `bson:"quantity"`
	Subtotal   float64 `bson:"subtotal"`    // biaya produk only
	TotalPrice float64 `bson:"total_price"` // biaya dengan ongkir
	// Kurir
}
