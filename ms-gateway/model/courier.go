package model

type CourierRequest struct {
	NoResi  string `json:"no_resi"`
	Company string `json:"company"`
}

type ConfirmOrderID struct {
	OrderID int `json:"order_id"`
}
