package handler

import (
	"context"
	"fmt"
	"ms-gateway/dto"
	"ms-gateway/model"
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

// /orders/:id [put]
func (h OrderHandler) AddCourierinfo(c echo.Context) error {
	sellerID := c.Get("id").(string)

	orderID := c.Param("id")
	_, err := strconv.Atoi(orderID) // check if param is not a digit
	if err != nil {
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input",
		})
	}

	var input model.CourierRequest
	if err := c.Bind(&input); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}

	in := pb.AddCourierInfoRequest{
		Awb:      input.NoResi,
		Company:  input.Company,
		OrderId:  orderID,
		SellerId: sellerID,
	}

	resp, err := h.orderGRPC.AddCourierInfo(context.TODO(), &in)
	if err != nil {
		return echo.NewHTTPError(500, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h OrderHandler) TrackCourierShipment(c echo.Context) error {
	orderID := c.Param("id")
	_, err := strconv.Atoi(orderID) // check if param is not a digit
	if err != nil {
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input",
		})
	}

	in := pb.TrackCourierShipmentRequest{
		OrderId: orderID,
	}

	resp, err := h.orderGRPC.TrackCourierShipment(context.TODO(), &in)
	if err != nil {
		return echo.NewHTTPError(500, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
