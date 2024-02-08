package handler

import (
	"context"
	"fmt"
	"ms-gateway/dto"
	"ms-gateway/helper"
	"ms-gateway/model"
	"ms-gateway/pb"
	"net/http"
	"strconv"
	"strings"

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
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	useraddr, err := h.userGRPC.GetUserAddress(context.Background(), &pb.GetUserAddressRequest{
		UserId: c.Get("id").(string),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	seller, err := h.sellerGRPC.GetSellerByID(context.Background(), &pb.GetSellerByIDRequest{
		SellerId: int32(input.SellerID),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	var productsreq []*pb.OrderCreateRequest_Product
	for _, input := range input.Products {
		res, err := h.sellerGRPC.GetProductByID(context.Background(), &pb.GetProductByIDRequest{
			ProductId: int32(input.ProductID),
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		if res.Stock < int32(input.Quantity) {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("out of stock for product id %v ", res.Productid))
		}

		productsreq = append(productsreq, &pb.OrderCreateRequest_Product{
			Id:       int64(res.Productid),
			Name:     res.Name,
			Price:    float64(res.Price),
			Quantity: int64(input.Quantity),
		})
	}

	req := pb.OrderCreateRequest{
		User: &pb.OrderCreateRequest_User{
			Id:      userid,
			Name:    user.Name,
			Address: useraddr.Address,
			City:    useraddr.City,
		},
		Seller: &pb.OrderCreateRequest_Seller{
			Id:      int64(seller.SellerId),
			Name:    seller.Name,
			Address: seller.AddressName,
			City:    seller.AddressCity,
		},

		Products: productsreq,
		Shipment: &pb.OrderCreateRequest_Shipment{
			Company: input.Shipment.Company,
			Service: input.Shipment.Service,
		},
		PaymentMethod: input.PaymentMethod,
	}

	res, err := h.orderGRPC.OrderCreate(context.Background(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"order":   res.Order,
		"message": "order created",
	})
}

func (h OrderHandler) GetAllOrder(c echo.Context) error {
	userid, _ := strconv.ParseInt(c.Get("id").(string), 10, 64)
	role := c.Get("role").(string)
	paramstatus := strings.ToUpper(c.QueryParam("status"))

	if paramstatus != "" && !helper.ValidOrderStatus(paramstatus) {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid status query param")
	}

	res, err := h.orderGRPC.OrderGetAll(context.Background(), &pb.OrderGetAllRequest{
		Role:   role,
		Userid: userid,
		Status: paramstatus,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if len(res.Orders) == 0 {
		res.Orders = []*pb.Order{}
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"orders": res.Orders,
	})
}

func (h OrderHandler) GetByIdOrder(c echo.Context) error {
	userid, _ := strconv.ParseInt(c.Get("id").(string), 10, 64)
	role := c.Get("role").(string)
	orderid := c.Param("id")

	res, err := h.orderGRPC.OrderGetById(context.Background(), &pb.OrderGetByIdRequest{
		Id:     orderid,
		Role:   role,
		Userid: userid,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"order": res.Order,
	})
}

func (h OrderHandler) ConfirmOrder(c echo.Context) error {
	userid, _ := strconv.ParseInt(c.Get("id").(string), 10, 64)
	role := c.Get("role").(string)
	orderid := c.Param("id")

	res, err := h.orderGRPC.OrderConfirmationAccept(context.Background(), &pb.OrderConfirmationAcceptRequest{
		Id:     orderid,
		Userid: userid,
		Role:   role,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"order_id": res.Id,
		"message":  res.Message,
	})
}

func (h OrderHandler) RejectOrder(c echo.Context) error {
	userid, _ := strconv.ParseInt(c.Get("id").(string), 10, 64)
	role := c.Get("role").(string)
	orderid := c.Param("id")

	res, err := h.orderGRPC.OrderConfirmationCancel(context.Background(), &pb.OrderConfirmationCancelRequest{
		Id:     orderid,
		Userid: userid,
		Role:   role,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"order_id": res.Id,
		"message":  res.Message,
	})
}

// @Summary      Add Courier Info
// @Description  Seller add courier info to order
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Param ID path int true "Order ID"
// @Param		 data body model.CourierRequest true "The input courier struct"
// @Success      200  {object}  model.Courier
// @Failure      400  {object}  helper.Message
// @Failure      401  {object}  helper.Message
// @Failure      500  {object}  helper.Message
// @Router       /courier/:id [Put]
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

// @Summary      Track Shipment
// @Description  User or Seller can track courier status
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Param ID path int true "Order ID"
// @Success      200  {object}  model.Courier
// @Failure      400  {object}  helper.Message
// @Failure      401  {object}  helper.Message
// @Failure      500  {object}  helper.Message
// @Router       /courier/:id [Get]
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

func (h OrderHandler) CustomerConfirmOrder(c echo.Context) error {
	customerID := c.Get("id").(string)

	var input model.ConfirmOrderID
	if err := c.Bind(&input); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}
	in := pb.ConfirmOrderRequest{
		OrderId:    strconv.Itoa(input.OrderID),
		CustomerId: customerID,
	}

	_, err := h.orderGRPC.CustomerConfirmOrder(context.TODO(), &in)
	if err != nil {
		return echo.NewHTTPError(500, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "order finished",
	})
}
