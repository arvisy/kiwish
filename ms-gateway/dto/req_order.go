package dto

type ReqCreateOrderDirect struct {
	ProductID     int32  `json:"product_id"`
	PaymentMethod string `json:"payment_method"`
	Quantity      int64  `json:"quantity"`
}
