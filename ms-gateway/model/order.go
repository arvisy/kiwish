package model

const (
	ORDER_STATUS_UNPAID   = "UNPAID"
	ORDER_STATUS_PACKED   = "PACKED"
	ORDER_STATUS_SHIPPED  = "SHIPPED"
	ORDER_STATUS_COMPLETE = "COMPLETE"
)

const (
	ORDER_CONFIRMATION_ACCEPTED = "ACCEPTED"
	ORDER_CONFIRMATION_REJECTED = "REJECTED"
	ORDER_CONFIRMATION_WAITING  = "WAITING"
)

// import (
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// type Order struct {
// 	ID        primitive.ObjectID `bson:"_id,omitempty"`
// 	User      User               `bson:"user"`
// 	Seller    Seller             `bson:"seller"`
// 	Shipment  Shipment           `bson:"shipment"`
// 	Payment   Payment            `bson:"payment"`
// 	Products  []Product          `bson:"products"`
// 	Subtotal  float64            `bson:"subtotal"`
// 	Total     float64            `bson:"total"`
// 	Status    string             `bson:"status"`
// 	CreatedAt time.Time          `bson:"created_at"`
// }
