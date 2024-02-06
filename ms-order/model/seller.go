package model

type Seller struct {
	ID      int64  `bson:"_id,omitempty"`
	Name    string `bson:"name"`
	Address string `bson:"address"`
}
