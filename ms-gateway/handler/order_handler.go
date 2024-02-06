package handler

import (
	sellerpb "ms-gateway/pb"
	orderpb "ms-order/pb"
	userpb "ms-user/pb"
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

// func (h OrderHandler) CreateOrderDirect(c echo.Context) error {
// 	// validation
// 	var input dto.ReqCreateOrderDirect
// 	if err := c.Bind(&input); err != nil {
// 		return err
// 	}

// 	useridstr := c.Get("id").(string)
// 	rescustomer, err := h.userGRPC.GetCustomer(context.Background(), &userpb.GetCustomerRequest{
// 		Id: useridstr,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	// get address

// 	// get product
// 	resproduct, err := h.sellerGRPC.GetProductByID(context.Background(), &sellerpb.GetProductByIDRequest{
// 		ProductId: input.ProductID,
// 	})

// 	// calculate total price
// 	totalprice := helper.CalculatePrice(input.Quantity, resproduct.Price)

// 	//

// 	return nil
// }
