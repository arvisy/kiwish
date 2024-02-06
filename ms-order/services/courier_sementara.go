package services

import (
	"context"
	orderpb "ms-order/pb"
)

func (s Service) GetCourierPrice(ctx context.Context, in *orderpb.GetCourierPriceRequest) (*orderpb.GetCourierPriceResponse, error) {
	res, err := s.client.Courier.GetPrice(in.Origin, in.Destination, in.Company)
	if err != nil {
		return nil, ErrInternal(err, s.log)
	}

	response := &orderpb.GetCourierPriceResponse{
		Origin: &orderpb.GetCourierPriceResponse_Detail{
			CityId:     res.Rajaongkir.OriginDetails.CityID,
			ProvinceId: res.Rajaongkir.OriginDetails.ProvinceID,
			Province:   res.Rajaongkir.OriginDetails.Province,
			Type:       res.Rajaongkir.OriginDetails.Type,
			CityName:   res.Rajaongkir.OriginDetails.CityName,
			PostalCode: res.Rajaongkir.OriginDetails.PostalCode,
		},
		Destination: &orderpb.GetCourierPriceResponse_Detail{
			CityId:     res.Rajaongkir.DestinationDetails.CityID,
			ProvinceId: res.Rajaongkir.DestinationDetails.ProvinceID,
			Province:   res.Rajaongkir.DestinationDetails.Province,
			Type:       res.Rajaongkir.DestinationDetails.Type,
			CityName:   res.Rajaongkir.DestinationDetails.CityName,
			PostalCode: res.Rajaongkir.DestinationDetails.PostalCode,
		},
		Cost: &orderpb.GetCourierPriceResponse_Cost{},
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
