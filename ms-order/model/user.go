package model

type User struct {
	ID      int64  `bson:"_id,omitempty"`
	Name    string `bson:"name"`
	Address string `bson:"address"`
	City    string `bson:"city"`
}
