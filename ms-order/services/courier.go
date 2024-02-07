package services

import (
	"context"
	"fmt"
	"ms-order/helpers"
	"ms-order/model"
	"ms-order/pb"
	"strconv"
	"strings"
)

func (s Service) AddShipmentInfo(ctx context.Context, in *pb.AddCourierInfoRequest) (*pb.CourierResponse, error) {
	order, err := s.repo.Order.GetOrder(in.OrderId)
	if err != nil {
		return nil, err
	}

	if strconv.Itoa(int(order.Seller.ID)) != in.SellerId {
		return nil, fmt.Errorf("order id invalid")
	}

	// company = jne, pos, jnt_cargo,sicepat, tiki, anteraja, ninja, lion
	// jet, dll
	info, err := helpers.TrackPackage(in.Awb, strings.ToLower(in.Company))
	if err != nil {
		return nil, err
	}

	inputShipment := model.Shipment{
		NoResi:  info.Data.Summary.AWB,
		Company: order.Shipment.Company,
		Service: order.Shipment.Service,
		Status:  info.Data.Summary.Status,
		Price:   order.Shipment.Price,
	}

	order.Shipment = inputShipment

	_, err = s.repo.Order.UpdateShipmentResiStatus(order)
	if err != nil {
		return nil, err
	}

	var list []*pb.HistoryResponse

	for _, v := range info.Data.History {
		u := pb.HistoryResponse{
			Date:        v.Date,
			Description: v.Description,
		}
		list = append(list, &u)
	}

	resp := pb.CourierResponse{
		Awb:         info.Data.Summary.AWB,
		Company:     info.Data.Summary.Courier,
		Status:      info.Data.Summary.Status,
		Date:        info.Data.Summary.Date,
		Origin:      info.Data.Detail.Origin,
		Destination: info.Data.Detail.Destination,
		History:     list,
	}

	return &resp, nil
}

func (s Service) TrackCourierShipment(ctx context.Context, in *pb.TrackCourierShipmentRequest) (*pb.CourierResponse, error) {
	order, err := s.repo.Order.GetOrder(in.OrderId)
	if err != nil {
		return nil, err
	}

	// company = jne, pos, jnt_cargo,sicepat, tiki, anteraja, ninja, lion
	// jet, dll
	info, err := helpers.TrackPackage(order.Shipment.NoResi, strings.ToLower(order.Shipment.Company))
	if err != nil {
		return nil, err
	}

	if order.Shipment.Status != "DELIVERED" && info.Data.Summary.Status == "DELIVERED" {
		inputShipment := model.Shipment{
			NoResi:  info.Data.Summary.AWB,
			Company: order.Shipment.Company,
			Service: order.Shipment.Service,
			Status:  info.Data.Summary.Status,
			Price:   order.Shipment.Price,
		}

		order.Shipment = inputShipment

		_, err = s.repo.Order.UpdateShipmentResiStatus(order)
		if err != nil {
			return nil, err
		}
	}

	var list []*pb.HistoryResponse

	for _, v := range info.Data.History {
		u := pb.HistoryResponse{
			Date:        v.Date,
			Description: v.Description,
		}
		list = append(list, &u)
	}

	resp := pb.CourierResponse{
		Awb:         info.Data.Summary.AWB,
		Company:     info.Data.Summary.Courier,
		Status:      info.Data.Summary.Status,
		Date:        info.Data.Summary.Date,
		Origin:      info.Data.Detail.Origin,
		Destination: info.Data.Detail.Destination,
		History:     list,
	}

	return &resp, nil
}

func (s Service) CustomerConfirmOrder(ctx context.Context, in *pb.ConfirmOrderRequest) (*pb.ConfirmOrderResponse, error) {
	order, err := s.repo.Order.GetOrder(in.OrderId)
	if err != nil {
		return nil, err
	}

	if strconv.Itoa(int(order.User.ID)) != in.CustomerId {
		return nil, fmt.Errorf("order id invalid")
	}

	if order.Shipment.Status != "DELIVERED" {
		inputShipment := model.Shipment{
			NoResi:  order.Shipment.NoResi,
			Company: order.Shipment.Company,
			Service: order.Shipment.Service,
			Status:  "DELIVERED",
			Price:   order.Shipment.Price,
		}

		order.Shipment = inputShipment

		_, err = s.repo.Order.UpdateShipmentResiStatus(order)
		if err != nil {
			return nil, err
		}
	}

	order.Status = model.ORDER_STATUS_COMPLETE

	_, err = s.repo.Order.UpdateOrderStatus(order)
	if err != nil {
		return nil, err
	}

	return &pb.ConfirmOrderResponse{}, nil
}
