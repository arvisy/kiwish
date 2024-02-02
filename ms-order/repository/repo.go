package repository

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNotFound = errors.New("record not found")
)

type MongoRepository struct {
	Cart CartRepository
}

func NewMongo(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		Cart: CartRepository{coll: db.Collection("cart")},
	}
}
