package model

import "time"

const (
	ORDER_STATUS_UNPAID   = "UNPAID"
	ORDER_STATUS_PACKED   = "PACKED"
	ORDER_STATUS_SHIPPED  = "SHIPPED"
	ORDER_STATUS_COMPLETE = "COMPLETE"
)

type Order struct {
	OrderID       string `bson:"_id,omitempty"`
	UserID        int64  `bson:"user_id"`
	OrderDetailID string `bson:"order_detail_id"`
	Status        string `bson:"status"`
	PaymentID     string `bson:"payment_id"`
	// kurir
	TotalPrice float64   `bson:"total_price"`
	CreatedAt  time.Time `bson:"created_at"`
	UpdatedAt  time.Time `bson:"updated_at"`
}
