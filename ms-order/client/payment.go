package client

import (
	"context"
	"ms-order/config"

	"github.com/xendit/xendit-go/v4"
	"github.com/xendit/xendit-go/v4/invoice"
)

type PaymentClient struct {
	cfg config.Config
}

func (s *PaymentClient) CreateInvoice(orderid string, amount float64, method *string) (*invoice.Invoice, error) {
	createInvoiceRequest := *invoice.NewCreateInvoiceRequest(orderid, amount)
	createInvoiceRequest.SetCurrency("IDR")
	if method != nil {
		createInvoiceRequest.SetPaymentMethods([]string{*method})
	}
	xenditClient := xendit.NewClient(s.cfg.Xendit.ApiKey)

	resp, _, err := xenditClient.InvoiceApi.CreateInvoice(context.Background()).
		CreateInvoiceRequest(createInvoiceRequest).
		Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
