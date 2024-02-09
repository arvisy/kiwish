package mdorder

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	ORDER_STATUS_UNPAID   = "UNPAID"
	ORDER_STATUS_PACKED   = "PACKED"
	ORDER_STATUS_SHIPPED  = "SHIPPED"
	ORDER_STATUS_COMPLETE = "COMPLETE"
	ORDER_STATUS_CANCELED = "CANCELED"
)

type GetAllOrder struct {
	Orders []Order
}

type CreateOrder struct {
	Order Order
}

type Order struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	User         User               `bson:"user"`
	Seller       Seller             `bson:"seller"`
	Shipment     Shipment           `bson:"shipment"`
	Payment      Payment            `bson:"payment"`
	Products     []Product          `bson:"products"`
	Confirmation Confirmation       `json:"confirmation"`
	Subtotal     float64            `bson:"subtotal"`
	Total        float64            `bson:"total"`
	Status       string             `bson:"status"`
	CreatedAt    time.Time          `bson:"created_at"`
}

const (
	ORDER_CONFIRMATION_ACCEPTED = "ACCEPTED"
	ORDER_CONFIRMATION_REJECTED = "REJECTED"
	ORDER_CONFIRMATION_WAITING  = "WAITING"
)

type Confirmation struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}
