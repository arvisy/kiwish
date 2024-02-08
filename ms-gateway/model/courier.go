package model

type CourierRequest struct {
	NoResi  string `json:"no_resi"`
	Company string `json:"company"`
}

type ConfirmOrderID struct {
	OrderID int `json:"order_id"`
}

// swagger
type History struct {
	Date        string `json:"date"`
	Description string `json:"desc"`
}

type Courier struct {
	AWB         string    `json:"awb"`
	Company     string    `json:"company"`
	Status      string    `json:"status"`
	Date        string    `json:"date"`
	Fee         float64   `json:"fee"`
	Origin      string    `json:"origin"`
	Destination string    `json:"destination"`
	History     []History `json:"history"`
}
