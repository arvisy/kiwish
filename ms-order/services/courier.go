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

func (s Service) GetCourierPrice(ctx context.Context, in *pb.GetCourierPriceRequest) (*pb.GetCourierPriceResponse, error) {
	res, err := s.client.Courier.GetPrice(in.Origin, in.Destination, in.Company)
	if err != nil {
		return nil, ErrInternal(err, s.log)
	}

	response := &pb.GetCourierPriceResponse{
		Origin: &pb.GetCourierPriceResponse_Detail{
			CityId:     res.Rajaongkir.OriginDetails.CityID,
			ProvinceId: res.Rajaongkir.OriginDetails.ProvinceID,
			Province:   res.Rajaongkir.OriginDetails.Province,
			Type:       res.Rajaongkir.OriginDetails.Type,
			CityName:   res.Rajaongkir.OriginDetails.CityName,
			PostalCode: res.Rajaongkir.OriginDetails.PostalCode,
		},
		Destination: &pb.GetCourierPriceResponse_Detail{
			CityId:     res.Rajaongkir.DestinationDetails.CityID,
			ProvinceId: res.Rajaongkir.DestinationDetails.ProvinceID,
			Province:   res.Rajaongkir.DestinationDetails.Province,
			Type:       res.Rajaongkir.DestinationDetails.Type,
			CityName:   res.Rajaongkir.DestinationDetails.CityName,
			PostalCode: res.Rajaongkir.DestinationDetails.PostalCode,
		},
		Cost: &pb.GetCourierPriceResponse_Cost{},
	}

	for _, r := range res.Rajaongkir.Results {
		response.Cost.Company = r.Code
		for _, val := range r.Costs {
			if in.Service == val.Service {
				response.Cost.Service = val.Service
				response.Cost.Price = val.Cost[0].Value
			}
		}
	}

	return response, nil
}

func (s Service) AddCourierInfo(ctx context.Context, in *pb.AddCourierInfoRequest) (*pb.CourierResponse, error) {
	order, err := s.repo.Order.GetByID(in.OrderId)
	if err != nil {
		return nil, err
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
	order, err := s.repo.Order.GetByID(in.OrderId)
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
	order, err := s.repo.Order.GetByID(in.OrderId)
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
