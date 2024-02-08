package helper

import (
	"context"

	"github.com/xendit/xendit-go/v4"
	"github.com/xendit/xendit-go/v4/common"
	"github.com/xendit/xendit-go/v4/invoice"
)

func GetInvoice(invoiceID string) (*invoice.Invoice, *common.XenditSdkError) {
	xenditClient := xendit.NewClient("xnd_development_J9kh4VwOMvHRgxiknN5tiCb5tVyHOO3OaOKQm6gkhtuqjova7nUAPxsoCxXiFRS")

	resp, _, err := xenditClient.
		InvoiceApi.
		GetInvoiceById(context.Background(), invoiceID).
		Execute()

	if err != nil {
		return nil, err
	}
	return resp, nil
}
