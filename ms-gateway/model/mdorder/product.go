package mdorder

type Product struct {
	ID       int     `bson:"_id,omitempty"`
	Name     string  `bson:"name"`
	Price    float64 `bson:"price"`
	Quantity int64   `bson:"quantity"`
}
