package mdorder

type Shipment struct {
	NoResi  string  `bson:"no_resi"`
	Company string  `bson:"company"`
	Service string  `bson:"service"`
	Status  string  `bson:"status"`
	Price   float64 `bson:"float"`
}
