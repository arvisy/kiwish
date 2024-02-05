package handler

import (
	"context"
	"ms-gateway/model"
	"ms-gateway/pb"
	"net/http"
	"strconv"
	"strings"

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

func (h *SellerHandler) AddProduct(c echo.Context) error {
	var input model.Product
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}

	sellerID := c.Get("id").(string)
	input.SellerID, _ = strconv.Atoi(sellerID)

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
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
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

func (h *SellerHandler) GetProductsBySeller(c echo.Context) error {
	id := c.Param(":id")
	sellerID, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}

	in := pb.GetProductsRequest{
		SellerId: int32(sellerID),
	}

	resp, err := h.sellerGRPC.GetProductsBySeller(context.Background(), &in)
	if err != nil {
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}

	var result []*model.Product
	for _, v := range resp.Products {
		r := model.Product{
			ID:          int(v.Productid),
			SellerID:    int(v.SellerId),
			Name:        v.Name,
			Price:       v.Price,
			Stock:       int(v.Stock),
			Category_id: int(v.CategoryId),
			Discount:    int(v.Discount),
		}
		result = append(result, &r)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"products": result,
	})
}

func (h *SellerHandler) GetProductsByCategory(c echo.Context) error {
	category := c.Param("category")
	category = strings.ToLower(category)

	in := pb.GetProductByCategoryRequest{
		CategoryName: category,
	}

	resp, err := h.sellerGRPC.GetProductsByCategory(context.Background(), &in)
	if err != nil {
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}

	var result []*model.Product
	for _, v := range resp.Products {
		r := model.Product{
			ID:          int(v.Productid),
			SellerID:    int(v.SellerId),
			Name:        v.Name,
			Price:       v.Price,
			Stock:       int(v.Stock),
			Category_id: int(v.CategoryId),
			Discount:    int(v.Discount),
		}
		result = append(result, &r)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"products": result,
	})
}

func (h *SellerHandler) GetProductByID(c echo.Context) error {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}

	in := pb.GetProductByIDRequest{
		ProductId: int32(productID),
	}

	resp, err := h.sellerGRPC.GetProductByID(context.Background(), &in)
	if err != nil {
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}

	return c.JSON(http.StatusOK, model.Product{
		ID:          int(resp.Productid),
		SellerID:    int(resp.SellerId),
		Name:        resp.Name,
		Price:       resp.Price,
		Stock:       int(resp.Stock),
		Category_id: int(resp.CategoryId),
		Discount:    int(resp.Discount),
	})
}

func (h *SellerHandler) DeleteProduct(c echo.Context) error {
	id1 := c.Param("id")
	productID, err := strconv.Atoi(id1)
	if err != nil {
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}

	id2 := c.Get("id").(string)
	sellerID, _ := strconv.Atoi(id2)

	in := pb.DeleteProductRequest{
		Productid: int32(productID),
		SellerId:  int32(sellerID),
	}

	_, err = h.sellerGRPC.DeleteProduct(context.Background(), &in)
	if err != nil {
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Product successfully deleted!",
	})
}

func (h *SellerHandler) UpdateProduct(c echo.Context) error {
	id1 := c.Param("id")
	productID, err := strconv.Atoi(id1)
	if err != nil {
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}

	id2 := c.Get("id").(string)
	sellerID, _ := strconv.Atoi(id2)

	var input model.ProductInput
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}

	// seller can't modify product & seller id
	// if product & seller id doesn't match, query will return error
	in := pb.UpdateProductRequest{
		Productid:  int32(productID),
		SellerId:   int32(sellerID),
		Name:       input.Name,
		Price:      input.Price,
		CategoryId: int32(input.Category_id),
		Discount:   int32(input.Discount),
	}

	_, err = h.sellerGRPC.UpdateProduct(context.Background(), &in)
	if err != nil {
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Product successfully deleted!",
	})
}
