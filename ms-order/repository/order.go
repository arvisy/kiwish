package repository

import (
	"context"
	"ms-order/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepo struct {
	coll *mongo.Collection
}

func (r OrderRepo) CreateOrder(order *model.Order) error {
	_, err := r.coll.InsertOne(context.Background(), order)
	if err != nil {
		return err
	}
	return nil
}
