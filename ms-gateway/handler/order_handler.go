package handler

import (
	sellerpb "ms-gateway/pb"
	orderpb "ms-order/pb"
	userpb "ms-user/pb"

	"github.com/labstack/echo/v4"
)

/*
	payload {
		payment_method
		product_id
		// kurir
	}

	func order direct
		// validation
		// get login user
		// fetch produk with seller
		//

*/

type OrderHandler struct {
	userGRPC   userpb.UserServiceClient
	sellerGRPC sellerpb.SellerServiceClient
	orderGRPC  orderpb.OrderServiceClient
}

func NewOrderHandler(userGRPC userpb.UserServiceClient, sellerGRPC sellerpb.SellerServiceClient, orderGRPC orderpb.OrderServiceClient) *OrderHandler {
	return &OrderHandler{
		userGRPC:   userGRPC,
		sellerGRPC: sellerGRPC,
		orderGRPC:  orderGRPC,
	}
}

func (h OrderHandler) CreateOrderDirect(ctx echo.Context) error {
	// validation

	// sellerID := c.Get("id").(string)
	// input.SellerID, _ = strconv.Atoi(sellerID)
	return nil
}
