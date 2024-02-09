package mdorder

type TrackingInfo struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	Summary Summary   `json:"summary"`
	Detail  Detail    `json:"detail"`
	History []History `json:"history"`
}

type Summary struct {
	AWB     string `json:"awb"`
	Courier string `json:"courier"`
	Service string `json:"service"` // reguler, nextday
	Status  string `json:"status"`  // DELIVERED, PENDING
	Date    string `json:"date"`
	Amount  string `json:"amount"` // ongkir
	Weight  string `json:"weight"` // per kg
}

type Detail struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Shipper     string `json:"shipper"`
	Receiver    string `json:"receiver"`
}

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
