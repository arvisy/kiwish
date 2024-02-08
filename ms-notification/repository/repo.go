package repository

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNotFound = errors.New("record not found")
)

type MongoRepository struct {
	Notif NotificationRepo
}

func NewMongo(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		Notif: NotificationRepo{coll: db.Collection("notification")},
	}
}
