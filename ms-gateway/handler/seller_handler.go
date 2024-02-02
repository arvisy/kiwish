package handler

import (
	"context"
	"ms-gateway/model"
	"ms-gateway/pb"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SellerHandler struct {
	sellerGRPC pb.SellerServiceClient
}

func NewSellerHandler(grpc pb.SellerServiceClient) *SellerHandler {
	return &SellerHandler{
		sellerGRPC: grpc,
	}
}

func (h *SellerHandler) AddProduct(c echo.Context) {
	// userID := c.Get("id").(string)

	var input model.Product
	if err := c.Bind(&input); err != nil {
		echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err
		})
		return
	}

	in := pb.AddProductRequest{
		SellerId:   int32(input.SellerID),
		Name:       input.Name,
		Price:      input.Price,
		Stock:      int32(input.Stock),
		CategoryId: int32(input.Category_id),
		Discount:   int32(input.Discount),
	}

	resp, err := h.sellerGRPC.AddProduct(context.Background(), &in)
	if err != nil {

		return
	}

	c.JSON(http.StatusOK, echo.Map{
		"message": "Product successfully added!",
		"task": model.Product{
			SellerID:    int(resp.SellerId),
			Name:        resp.Name,
			Price:       resp.Price,
			Stock:       int(resp.Stock),
			Category_id: int(resp.CategoryId),
			Discount:    int(resp.Discount),
		},
	})
}
