package services

import (
	"context"
	"ms-order/helpers"
	"ms-order/model"
	orderpb "ms-order/pb"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s Service) OrderCreate(ctx context.Context, in *orderpb.OrderCreateRequest) (*orderpb.OrderCreateResponse, error) {
	// TODO: minus shipment
	order := &model.Order{
		ID: primitive.NewObjectID(),
		User: model.User{
			ID:      in.User.Id,
			Name:    in.User.Name,
			Address: in.User.Address,
		},
		Seller: model.Seller{
			ID:      in.Seller.Id,
			Name:    in.Seller.Name,
			Address: in.Seller.Address,
		},
		Payment: model.Payment{
			Method: in.PaymentMethod,
		},
		CreatedAt: time.Now(),
		Status:    model.ORDER_STATUS_UNPAID,
	}
	for _, p := range in.Products {
		order.Products = append(order.Products, model.Product{
			ID:          int(p.Id),
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Quantity:    p.Quantity,
		})
	}
	order.Subtotal = helpers.CalculateSubtotal(order.Products)
	order.Total = helpers.CalculateTotal(order)

	invoice, err := s.client.Payment.CreateInvoice(order.ID.Hex(), order.Total, &order.Payment.Method)
	if err != nil {
		return nil, ErrInternal(err, s.log)
	}
	order.Payment.InvoiceID = *invoice.Id
	order.Payment.InvoiceURL = invoice.InvoiceUrl
	order.Payment.Status = string(invoice.Status)

	// insert repo
	err = s.repo.Order.CreateOrder(order)
	if err != nil {
		return nil, ErrInternal(err, s.log)
	}

	response := &orderpb.OrderCreateResponse{
		Id: order.ID.Hex(),
		User: &orderpb.OrderCreateResponse_User{
			Id:      order.User.ID,
			Name:    order.User.Name,
			Address: order.User.Address,
		},
		Seller: &orderpb.OrderCreateResponse_Seller{
			Id:      order.Seller.ID,
			Name:    order.Seller.Name,
			Address: order.Seller.Address,
		},
		Payment: &orderpb.OrderCreateResponse_Payment{
			InvoiceId:  order.Payment.InvoiceID,
			InvoiceUrl: order.Payment.InvoiceURL,
			Method:     order.Payment.Method,
			Status:     order.Payment.Status,
		},
		Shipment:  &orderpb.OrderCreateResponse_Shipment{},
		Subtotal:  order.Subtotal,
		Total:     order.Total,
		CreatedAt: timestamppb.New(order.CreatedAt),
	}

	for _, p := range order.Products {
		response.Products = append(response.Products, &orderpb.OrderCreateResponse_Product{
			Id:          int64(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Quantity:    p.Quantity,
		})
	}

	return response, nil
}
