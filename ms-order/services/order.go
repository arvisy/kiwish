package services

import (
	"context"
	"errors"
	"fmt"
	"ms-order/helpers"
	"ms-order/model"
	orderpb "ms-order/pb"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s Service) OrderCreate(ctx context.Context, in *orderpb.OrderCreateRequest) (*orderpb.OrderCreateResponse, error) {
	order := &model.Order{
		ID: primitive.NewObjectID(),
		User: model.User{
			ID:      in.User.Id,
			Name:    in.User.Name,
			Address: in.User.Address,
			City:    in.User.City,
		},
		Seller: model.Seller{
			ID:      in.Seller.Id,
			Name:    in.Seller.Name,
			Address: in.Seller.Address,
			City:    in.Seller.City,
		},
		Shipment: model.Shipment{
			Company: in.Shipment.Company,
			Service: in.Shipment.Service,
		},

		Payment: model.Payment{
			Method: in.PaymentMethod,
		},
		Status:    model.ORDER_STATUS_UNPAID,
		CreatedAt: time.Now(),
	}
	for _, p := range in.Products {
		order.Products = append(order.Products, model.Product{
			ID:       int(p.Id),
			Name:     p.Name,
			Price:    p.Price,
			Quantity: p.Quantity,
		})
	}

	// get shipment price
	origin, ok := helpers.LIST_KOTA[order.User.City]
	if !ok {
		return nil, ErrInvalidArgument(fmt.Errorf("origin city not valid: %v", order.User.City))
	}

	destination, ok := helpers.LIST_KOTA[order.Seller.City]
	if !ok {
		return nil, ErrInvalidArgument(fmt.Errorf("destination city not valid: %v", order.Seller.City))
	}

	res, err := s.client.Courier.GetPrice(origin, destination, order.Shipment.Company)
	if err != nil {
		return nil, ErrInternal(err, s.log)
	}

	for _, r := range res.Rajaongkir.Results {
		if r.Code == order.Shipment.Company {
			for _, val := range r.Costs {
				if order.Shipment.Service == val.Service {
					order.Shipment.Price = val.Cost[0].Value
				}
			}
		}
	}
	if order.Shipment.Price == 0 {
		return nil, ErrNotFound(fmt.Sprintf("shipment service not found: %v - %v", order.Shipment.Company, order.Shipment.Service))
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
		Order: &orderpb.Order{
			Id:     order.ID.Hex(),
			Status: order.Status,
			User: &orderpb.Order_User{
				Id:      order.User.ID,
				Name:    order.User.Name,
				Address: order.User.Address,
				City:    order.User.City,
			},
			Seller: &orderpb.Order_Seller{
				Id:      order.Seller.ID,
				Name:    order.Seller.Name,
				Address: order.Seller.Address,
				City:    order.Seller.City,
			},
			Payment: &orderpb.Order_Payment{
				InvoiceId:  order.Payment.InvoiceID,
				InvoiceUrl: order.Payment.InvoiceURL,
				Method:     order.Payment.Method,
				Status:     order.Payment.Status,
			},
			Shipment: &orderpb.Order_Shipment{
				Company: order.Shipment.Company,
				Service: order.Shipment.Service,
				Price:   order.Shipment.Price,
			},
			Confirmation: &orderpb.Order_Confirmation{
				Status:      model.ORDER_CONFIRMATION_WAITING,
				Description: "",
			},
			Subtotal:  order.Subtotal,
			Total:     order.Total,
			CreatedAt: timestamppb.New(order.CreatedAt),
		},
	}

	for _, p := range order.Products {
		response.Order.Products = append(response.Order.Products, &orderpb.Order_Product{
			Id:       int64(p.ID),
			Name:     p.Name,
			Price:    p.Price,
			Quantity: p.Quantity,
		})
	}

	return response, nil
}

func (s Service) OrderGetAll(ctx context.Context, in *orderpb.OrderGetAllRequest) (*orderpb.OrderGetAllResponse, error) {
	orders, err := s.repo.Order.GetAll(in.Userid, in.Status, in.Role)
	if err != nil {
		return nil, err
	}

	var response = &orderpb.OrderGetAllResponse{}
	for idx, order := range orders {
		response.Orders = append(response.Orders, &orderpb.Order{
			Id:     order.ID.Hex(),
			Status: order.Status,
			Confirmation: &orderpb.Order_Confirmation{
				Status:      order.Confirmation.Status,
				Description: order.Confirmation.Description,
			},
			User: &orderpb.Order_User{
				Id:      order.User.ID,
				Name:    order.User.Name,
				Address: order.User.Address,
				City:    order.User.City,
			},
			Seller: &orderpb.Order_Seller{
				Id:      order.Seller.ID,
				Name:    order.Seller.Name,
				Address: order.Seller.Address,
				City:    order.Seller.City,
			},
			Payment: &orderpb.Order_Payment{
				InvoiceId:  order.Payment.InvoiceID,
				InvoiceUrl: order.Payment.InvoiceURL,
				Method:     order.Payment.Method,
				Status:     order.Payment.Status,
			},
			Shipment: &orderpb.Order_Shipment{
				Company: order.Shipment.Company,
				Service: order.Shipment.Service,
				Price:   order.Shipment.Price,
				NoResi:  order.Shipment.NoResi,
				Status:  order.Shipment.Status,
			},
			Subtotal:  order.Subtotal,
			Total:     order.Total,
			CreatedAt: timestamppb.New(order.CreatedAt),
		})

		for _, p := range order.Products {
			response.Orders[idx].Products = append(response.Orders[idx].Products, &orderpb.Order_Product{
				Id:       int64(p.ID),
				Name:     p.Name,
				Price:    p.Price,
				Quantity: p.Quantity,
			})
		}
	}

	return response, nil
}

func (s Service) OrderGetById(ctx context.Context, in *orderpb.OrderGetByIdRequest) (*orderpb.OrderGetByIdResponse, error) {
	order, err := s.repo.Order.GetByID(in.Id, in.Userid, in.Role)
	if err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, ErrNotFound("order not found")
		default:
			return nil, ErrInternal(err, s.log)
		}
	}

	var response = &orderpb.OrderGetByIdResponse{
		Order: &orderpb.Order{
			Id: order.ID.Hex(),
			User: &orderpb.Order_User{
				Id:      order.User.ID,
				Name:    order.User.Name,
				Address: order.User.Address,
				City:    order.User.City,
			},
			Seller: &orderpb.Order_Seller{
				Id:      order.Seller.ID,
				Name:    order.Seller.Name,
				Address: order.Seller.Address,
				City:    order.Seller.City,
			},
			Shipment: &orderpb.Order_Shipment{
				NoResi:  order.Shipment.NoResi,
				Company: order.Shipment.Company,
				Service: order.Shipment.Service,
				Status:  order.Shipment.Status,
				Price:   order.Shipment.Price,
			},
			Payment: &orderpb.Order_Payment{
				InvoiceId:  order.Payment.InvoiceID,
				InvoiceUrl: order.Payment.InvoiceURL,
				Method:     order.Payment.Method,
				Status:     order.Payment.Status,
			},
			Confirmation: &orderpb.Order_Confirmation{
				Status:      order.Confirmation.Status,
				Description: order.Confirmation.Description,
			},
			Subtotal:  order.Subtotal,
			Total:     order.Total,
			Status:    order.Status,
			CreatedAt: timestamppb.New(order.CreatedAt),
		},
	}

	for _, p := range order.Products {
		response.Order.Products = append(response.Order.Products, &orderpb.Order_Product{
			Id:       int64(p.ID),
			Name:     p.Name,
			Price:    p.Price,
			Quantity: p.Quantity,
		})
	}

	return response, nil
}

func (s Service) OrderConfirmationAccept(ctx context.Context, in *orderpb.OrderConfirmationAcceptRequest) (*orderpb.OrderConfirmationAcceptResponse, error) {
	order, err := s.repo.Order.GetByID(in.Id, in.Userid, in.Role)
	if err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, ErrNotFound("order not found")
		default:
			return nil, ErrInternal(err, s.log)
		}
	}

	if order.Status == model.ORDER_STATUS_SHIPPED || order.Status == model.ORDER_STATUS_COMPLETE || order.Status == model.ORDER_STATUS_CANCELED {
		return nil, ErrInvalidArgument(fmt.Errorf("you order is being proccess, you cannnot change it"))
	}

	order.Confirmation.Status = model.ORDER_CONFIRMATION_ACCEPTED
	order.Confirmation.Description = "OK"
	order.Status = model.ORDER_STATUS_SHIPPED

	err = s.repo.Order.Update(order)
	if err != nil {
		return nil, ErrInternal(err, s.log)
	}

	return &orderpb.OrderConfirmationAcceptResponse{
		Id:      order.ID.Hex(),
		Message: "OK",
	}, nil
}

func (s Service) OrderConfirmationCancel(ctx context.Context, in *orderpb.OrderConfirmationCancelRequest) (*orderpb.OrderConfirmationCancelResponse, error) {
	order, err := s.repo.Order.GetByID(in.Id, in.Userid, in.Role)
	if err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, ErrNotFound("order not found")
		default:
			return nil, ErrInternal(err, s.log)
		}
	}

	if order.Status == model.ORDER_STATUS_SHIPPED || order.Status == model.ORDER_STATUS_COMPLETE || order.Status == model.ORDER_STATUS_CANCELED {
		return nil, ErrInvalidArgument(fmt.Errorf("you order is being proccess, you cannnot change it"))
	}

	order.Confirmation.Status = model.ORDER_CONFIRMATION_REJECTED
	order.Confirmation.Description = in.Description
	order.Status = model.ORDER_STATUS_CANCELED

	err = s.repo.Order.Update(order)
	if err != nil {
		return nil, ErrInternal(err, s.log)
	}

	return &orderpb.OrderConfirmationCancelResponse{
		Id:      order.ID.Hex(),
		Message: "OK",
	}, nil
}

func (s Service) OrderUpdate(ctx context.Context, in *orderpb.OrderUpdateRequest) (*orderpb.OrderUpdateResponse, error) {
	order, err := s.repo.Order.GetByID(in.Id, in.Userid, in.Role)
	if err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, ErrNotFound("order not found")
		default:
			return nil, ErrInternal(err, s.log)
		}
	}

	if in.OrderStatus != "" {
		order.Status = in.OrderStatus
	}

	if in.PaymentStatus != "" {
		order.Payment.Status = in.PaymentStatus
	}

	if in.ShipmentResi != "" {
		order.Shipment.NoResi = in.ShipmentResi
	}

	if in.ShipmentStatus != "" {
		order.Shipment.Status = in.ShipmentStatus
	}

	if in.ConfirmationStatus != "" {
		order.Confirmation.Status = in.ConfirmationStatus
	}

	if in.Description != "" {
		order.Confirmation.Description = in.Description
	}

	err = s.repo.Order.Update(order)
	if err != nil {
		return nil, ErrInternal(err, s.log)
	}

	var response = &orderpb.OrderUpdateResponse{
		Order: &orderpb.Order{
			Id: order.ID.Hex(),
			User: &orderpb.Order_User{
				Id:      order.User.ID,
				Name:    order.User.Name,
				Address: order.User.Address,
				City:    order.User.City,
			},
			Seller: &orderpb.Order_Seller{
				Id:      order.Seller.ID,
				Name:    order.Seller.Name,
				Address: order.Seller.Address,
				City:    order.Seller.City,
			},
			Shipment: &orderpb.Order_Shipment{
				NoResi:  order.Shipment.NoResi,
				Company: order.Shipment.Company,
				Service: order.Shipment.Service,
				Status:  order.Shipment.Status,
				Price:   order.Shipment.Price,
			},
			Payment: &orderpb.Order_Payment{
				InvoiceId:  order.Payment.InvoiceID,
				InvoiceUrl: order.Payment.InvoiceURL,
				Method:     order.Payment.Method,
				Status:     order.Payment.Status,
			},
			Confirmation: &orderpb.Order_Confirmation{
				Status:      order.Confirmation.Status,
				Description: order.Confirmation.Description,
			},
			Subtotal:  order.Subtotal,
			Total:     order.Total,
			Status:    order.Status,
			CreatedAt: timestamppb.New(order.CreatedAt),
		},
	}

	for _, p := range order.Products {
		response.Order.Products = append(response.Order.Products, &orderpb.Order_Product{
			Id:       int64(p.ID),
			Name:     p.Name,
			Price:    p.Price,
			Quantity: p.Quantity,
		})
	}

	return response, nil
}
