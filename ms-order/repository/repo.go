package repository

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNotFound = errors.New("record not found")
)

type MongoRepository struct {
	Order OrderRepo
}

func NewMongo(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		Order: OrderRepo{coll: db.Collection("order")},
	}
}
