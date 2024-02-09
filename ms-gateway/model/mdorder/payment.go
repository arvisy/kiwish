package mdorder

type Payment struct {
	InvoiceID  string `bson:"_id,omitempty"`
	InvoiceURL string `bson:"invoice_URL"`
	Method     string `bson:"method"`
	Status     string `bson:"status"`
}
