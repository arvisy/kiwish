package model

type Courier struct {
	AWB         string  `json:"awb"`
	Company     string  `json:"company"`
	Status      string  `json:"status"`
	Date        string  `json:"date"`
	Fee         float64 `json:"fee"`
	Origin      string  `json:"origin"`
	Destination string  `json:"destination"`
	History     History `json:"history"`
}
