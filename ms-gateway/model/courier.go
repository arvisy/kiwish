package model

type CourierRequest struct {
	NoResi  string `bson:"no_resi"`
	Company string `bson:"company"`
}
