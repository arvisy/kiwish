package repository

import (
	"context"
	"ms-order/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r OrderRepo) GetOrder(orderID string) (*model.Order, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: orderID}}

	var result model.Order

	//Passing the bson.D{{}} as the filter matches  documents in the collection
	err := r.coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r OrderRepo) GetAllForCustomer(userid int64) ([]model.Order, error) {
	filter := bson.D{{Key: "user._id", Value: userid}}

	var result = []model.Order{}
	cursor, err := r.coll.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.Background(), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r OrderRepo) GetAllForSeller(sellerid int64) ([]model.Order, error) {
	filter := bson.D{{Key: "seller._id", Value: sellerid}}

	var result = []model.Order{}
	cursor, err := r.coll.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.Background(), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r OrderRepo) UpdateShipmentResiStatus(order *model.Order) (*model.Order, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: order.ID}}

	update := bson.M{ // update resi status
		"$set": bson.M{
			"Shipment": order.Shipment,
		}}

	// Passing the bson.D{{}} as the filter matches  documents in the collection
	_, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return order, nil
}
