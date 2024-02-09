package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetAllNotification struct {
	Notifications []Notification
}

type Notification struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ReceiverID  int64              `bson:"receiver_id"`
	Subject     string             `bson:"subject"`
	Description string             `bson:"description"`
	Status      string             `bson:"status"`
	CreatedAt   time.Time          `bson:"created_at"`
}
