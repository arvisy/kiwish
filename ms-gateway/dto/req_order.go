package dto

type ReqCreateOrderDirect struct {
	Products []struct {
		ProductID int `json:"product_id"`
		Quantity  int `json:"quantity"`
	} `json:"products"`

	

	Shipment struct {
	} `json:"shipment"`
	PaymentMethod string `json:"payment_method"`
}
