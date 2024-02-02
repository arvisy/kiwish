package model

type Address struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name"`
	Regency string `json:"regency"`
	City    string `json:"city"`
}
