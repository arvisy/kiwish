package dto

type ReqCreateOrderDirect struct {
	Products []struct {
		ProductID int `json:"product_id"`
		Quantity  int `json:"quantity"`
	} `json:"products"`
	SellerID int `json:"seller_id"`
	Shipment struct {
		Company string `json:"company"` // kalo bisa pake REG case sensitif
		Service string `json:"service"` // Kalo bisa pake jne | tiki case sensitif
	} `json:"shipment"`
	PaymentMethod string `json:"payment_method"` // kalo bisa pake OVO | DANA case sensitive
}
