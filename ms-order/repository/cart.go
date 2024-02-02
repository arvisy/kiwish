package repository

import (
	"context"
	"errors"
	"fmt"
	"ms-order/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CartRepository struct {
	coll *mongo.Collection
}

func (r CartRepository) Create(ctx context.Context, cart *model.Cart) error {
	result, err := r.coll.InsertOne(ctx, cart)
	if err != nil {
		return err
	}

	objid, ok := result.InsertedID.(*primitive.ObjectID)
	if !ok {
		return fmt.Errorf("inserted id is not object id")
	}

	cart.ID = objid.String()
	return nil
}

func (r CartRepository) GetAll(ctx context.Context, userid int64) ([]*model.Cart, error) {
	options := options.Find().SetSort(bson.D{{Key: "updated_at", Value: -1}})
	filter := bson.M{"user_id": userid}
	cursor, err := r.coll.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}

	var carts []*model.Cart
	err = cursor.All(ctx, &carts)
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (r CartRepository) GetByID(ctx context.Context, cartid string, userid int64) (*model.Cart, error) {
	filter := bson.D{
		{Key: "_id", Value: cartid},
		{Key: "user_id", Value: userid},
	}
	var cart = &model.Cart{}
	err := r.coll.FindOne(ctx, filter).Decode(cart)
	if err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return cart, nil
}

func (r CartRepository) Update(ctx context.Context, cart *model.Cart) error {
	filter := bson.D{
		{Key: "_id", Value: cart.ID},
		{Key: "user_id", Value: cart.UserID},
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "quantity", Value: cart.Quantity},
			{Key: "created_at", Value: cart.CreatedAt},
		}}}
	_, err := r.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r CartRepository) DeleteOne(ctx context.Context, cartid string, userid int64) error {
	filter := bson.D{
		{Key: "_id", Value: cartid},
		{Key: "user_id", Value: userid},
	}

	_, err := r.coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (r CartRepository) DeleteAll(ctx context.Context, userid int64) error {
	filter := bson.D{
		{Key: "user_id", Value: userid},
	}

	result, err := r.coll.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return ErrNotFound
	}

	return nil
}
