package handler

import (
	"context"
	"ms-gateway/dto"
	"ms-gateway/pb"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	userGRPC   pb.UserServiceClient
	sellerGRPC pb.SellerServiceClient
	orderGRPC  pb.OrderServiceClient
}

func NewOrderHandler(userGRPC pb.UserServiceClient, sellerGRPC pb.SellerServiceClient, orderGRPC pb.OrderServiceClient) *OrderHandler {
	return &OrderHandler{
		userGRPC:   userGRPC,
		sellerGRPC: sellerGRPC,
		orderGRPC:  orderGRPC,
	}
}

func (h OrderHandler) CreateOrder(c echo.Context) error {
	var input dto.ReqCreateOrderDirect
	if err := c.Bind(&input); err != nil {
		return err
	}

	userid, _ := strconv.ParseInt(c.Get("id").(string), 10, 64)

	user, err := h.userGRPC.GetCustomer(context.Background(), &pb.GetCustomerRequest{
		Id: c.Get("id").(string),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error get user entity")
	}

	useraddr, err := h.userGRPC.GetUserAddress(context.Background(), &pb.GetUserAddressRequest{
		UserId: c.Get("id").(string),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error get user address")
	}

	product, err := h.sellerGRPC.GetProductByID()

	req := pb.OrderCreateRequest{
		User: &pb.OrderCreateRequest_User{
			Id:      userid,
			Name:    user.Name,
			Address: useraddr.Address,
			City:    useraddr.City,
		},
		Seller: &pb.OrderCreateRequest_Seller{
			// Id: ,
		},
	}

	return nil
}
