package model

import "time"

type Cart struct {
	ID        string    `bson:"_id,omitempty"`
	UserID    int64     `bson:"user_id"`
	ProductID int64     `bson:"product_id"`
	Quantity  float64   `bson:"quantity"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
