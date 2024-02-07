package services

import (
	"context"
	"fmt"
	"ms-order/helpers"
	"ms-order/model"
	orderpb "ms-order/pb"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
			ID:          int(p.Id),
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Quantity:    p.Quantity,
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
		Id:     order.ID.Hex(),
		Status: order.Status,
		User: &orderpb.OrderCreateResponse_User{
			Id:      order.User.ID,
			Name:    order.User.Name,
			Address: order.User.Address,
			City:    order.User.City,
		},
		Seller: &orderpb.OrderCreateResponse_Seller{
			Id:      order.Seller.ID,
			Name:    order.Seller.Name,
			Address: order.Seller.Address,
			City:    order.Seller.City,
		},
		Payment: &orderpb.OrderCreateResponse_Payment{
			InvoiceId:  order.Payment.InvoiceID,
			InvoiceUrl: order.Payment.InvoiceURL,
			Method:     order.Payment.Method,
			Status:     order.Payment.Status,
		},
		Shipment: &orderpb.OrderCreateResponse_Shipment{
			Company: order.Shipment.Company,
			Service: order.Shipment.Service,
			Price:   order.Shipment.Price,
		},
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

func (s Service) OrderGetAllForCustomer(ctx context.Context, in *orderpb.OrderGetAllForCustomerRequest) (*orderpb.OrderGetAllForCustomerResponse, error) {
	orders, err := s.repo.Order.GetAllForCustomer(in.Userid)
	if err != nil {
		return nil, err
	}

	var response = &orderpb.OrderGetAllForCustomerResponse{}
	for idx, order := range orders {
		response.Orders = append(response.Orders, &orderpb.OrderGetAllForCustomerResponse_Orders{
			Id:     order.ID.Hex(),
			Status: order.Status,
			User: &orderpb.OrderGetAllForCustomerResponse_Orders_User{
				Id:      order.User.ID,
				Name:    order.User.Name,
				Address: order.User.Address,
				City:    order.User.City,
			},
			Seller: &orderpb.OrderGetAllForCustomerResponse_Orders_Seller{
				Id:      order.Seller.ID,
				Name:    order.Seller.Name,
				Address: order.Seller.Address,
				City:    order.Seller.City,
			},
			Payment: &orderpb.OrderGetAllForCustomerResponse_Orders_Payment{
				InvoiceId:  order.Payment.InvoiceID,
				InvoiceUrl: order.Payment.InvoiceURL,
				Method:     order.Payment.Method,
				Status:     order.Payment.Status,
			},
			Shipment: &orderpb.OrderGetAllForCustomerResponse_Orders_Shipment{
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
			response.Orders[idx].Products = append(response.Orders[idx].Products, &orderpb.OrderGetAllForCustomerResponse_Orders_Product{
				Id:          int64(p.ID),
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
				Quantity:    p.Quantity,
			})
		}
	}

	return response, nil
}

func (s Service) OrderGetAllForSeller(ctx context.Context, in *orderpb.OrderGetAllForSellerRequest) (*orderpb.OrderGetAllForSellerResponse, error) {
	orders, err := s.repo.Order.GetAllForCustomer(in.Sellerid)
	if err != nil {
		return nil, err
	}

	var response = &orderpb.OrderGetAllForSellerResponse{}
	for idx, order := range orders {
		response.Orders = append(response.Orders, &orderpb.OrderGetAllForSellerResponse_Orders{
			Id:     order.ID.Hex(),
			Status: order.Status,
			User: &orderpb.OrderGetAllForSellerResponse_Orders_User{
				Id:      order.User.ID,
				Name:    order.User.Name,
				Address: order.User.Address,
				City:    order.User.City,
			},
			Seller: &orderpb.OrderGetAllForSellerResponse_Orders_Seller{
				Id:      order.Seller.ID,
				Name:    order.Seller.Name,
				Address: order.Seller.Address,
				City:    order.Seller.City,
			},
			Payment: &orderpb.OrderGetAllForSellerResponse_Orders_Payment{
				InvoiceId:  order.Payment.InvoiceID,
				InvoiceUrl: order.Payment.InvoiceURL,
				Method:     order.Payment.Method,
				Status:     order.Payment.Status,
			},
			Shipment: &orderpb.OrderGetAllForSellerResponse_Orders_Shipment{
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
			response.Orders[idx].Products = append(response.Orders[idx].Products, &orderpb.OrderGetAllForSellerResponse_Orders_Product{
				Id:          int64(p.ID),
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
				Quantity:    p.Quantity,
			})
		}
	}

	return response, nil
}

func (s Service) OrderUpdate(ctx context.Context, in *orderpb.OrderUpdateRequest) (*orderpb.OrderUpdateResponse, error) {
	return &orderpb.OrderUpdateResponse{}, nil
}
