package repository

import (
	"context"
	"ms-notification/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationRepo struct {
	coll *mongo.Collection
}

func (r NotificationRepo) Create(notif *model.Notification) error {
	_, err := r.coll.InsertOne(context.Background(), notif)
	if err != nil {
		return err
	}
	return nil
}

func (r NotificationRepo) GetByID(notifid string, receiverid int64) (*model.Notification, error) {
	oid, err := primitive.ObjectIDFromHex(notifid)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "$and", Value: bson.A{
		bson.M{"_id": oid},
		bson.M{"receiver_id": notifid},
	}}}

	var result model.Notification
	err = r.coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r NotificationRepo) GetAll(receiverid int64) ([]model.Notification, error) {
	filter := bson.M{"receiver_id": receiverid}

	var result = []model.Notification{}
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

func (r NotificationRepo) GetAllUnread(receiverid int64) ([]*model.Notification, error) {
	filter := bson.M{"$and": bson.A{
		bson.M{"receiver_id": receiverid},
		bson.M{"status": model.NOTIFICATION_UNREAD},
	}}

	var result = []*model.Notification{}
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

func (r NotificationRepo) Update(notif *model.Notification) error {
	filter := bson.D{{Key: "_id", Value: notif.ID}}
	update := bson.M{
		"$set": bson.M{
			"status": notif.Status,
		}}
	_, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
