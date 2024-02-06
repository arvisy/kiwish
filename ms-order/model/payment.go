package model

type Payment struct {
	ID         string `bson:"_id,omitempty"`
	InvoiceID  string `bson:"invoice_id"`
	InvoiceURL string `bson:"invoice_URL"`
	Method     string `bson:"method"`
	Status     string `bson:"status"`
}
