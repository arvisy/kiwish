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

func (r OrderRepo) GetByID(orderid string, userid int64, role string) (*model.Order, error) {
	oid, err := primitive.ObjectIDFromHex(orderid)
	if err != nil {
		return nil, err
	}

	var filter bson.D
	switch role {
	case "2":
		filter = bson.D{{Key: "$and", Value: bson.A{
			bson.M{"_id": oid},
			bson.M{"user._id": userid},
		}}}
	case "3":
		filter = bson.D{{Key: "$and", Value: bson.A{
			bson.M{"_id": oid},
			bson.M{"seller._id": userid},
		}}}
	}

	var result model.Order
	err = r.coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r OrderRepo) GetAll(userid int64, status string, role string) ([]model.Order, error) {
	var filter primitive.D
	switch role {
	case "2":
		filter = bson.D{{Key: "$and", Value: bson.A{
			bson.M{"user._id": userid},
			bson.M{"status": status},
		}}}
	case "3":
		filter = bson.D{{Key: "$and", Value: bson.A{
			bson.M{"seller._id": userid},
			bson.M{"status": status},
		}}}
	}

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

func (r OrderRepo) Update(order *model.Order) error {
	filter := bson.D{{Key: "_id", Value: order.ID}}
	update := bson.M{
		"$set": bson.M{
			"status":                   order.Status,
			"payment.status":           order.Payment.Status,
			"shipment.status":          order.Shipment.Status,
			"shipment.no_resi":         order.Shipment.NoResi,
			"confirmation.status":      order.Confirmation.Status,
			"confirmation.description": order.Confirmation.Description,
		}}
	_, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r OrderRepo) UpdateShipmentResiStatus(order *model.Order) (*model.Order, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: order.ID}}

	update := bson.M{ // update resi status
		"$set": bson.M{
			"Shipment": order.Shipment,
		}}

	_, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return order, nil
}
