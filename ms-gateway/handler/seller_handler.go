package handler

import (
	"context"
	"fmt"
	"ms-gateway/model"
	"ms-gateway/pb"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SellerHandler struct {
	sellerGRPC pb.SellerServiceClient
}

func NewSellerHandler(grpc pb.SellerServiceClient) *SellerHandler {
	return &SellerHandler{
		sellerGRPC: grpc,
	}
}

// @Summary      Add Product
// @Description  Seller can add a product
// @Tags         Seller
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Param		 data body model.ProductInput true "The input payment struct"
// @Success      201  {object}  model.Product
// @Failure      400  {object}  helpers.message
// @Failure      401  {object}  helpers.message
// @Failure      500  {object}  helpers.message
// @Router       /products [Post]
func (h *SellerHandler) AddProduct(c echo.Context) error {
	var input model.Product
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input",
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
		fmt.Println(err)
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"product": model.Product{
			ID:          int(resp.Productid),
			SellerID:    int(resp.SellerId),
			Name:        resp.Name,
			Price:       resp.Price,
			Stock:       int(resp.Stock),
			Category_id: int(resp.CategoryId),
			Discount:    int(resp.Discount),
		},
	})
}

// @Summary      Get Product By Seller ID
// @Description  Get products by seller ID path
// @Tags         Seller
// @Accept       json
// @Produce      json
// @Param ID path int true "Seller ID"
// @Success      200  {object}  []model.Product
// @Failure      400  {object}  helpers.message
// @Failure      401  {object}  helpers.message
// @Failure      500  {object}  helpers.message
// @Router       /products/seller/:id [Get]
func (h *SellerHandler) GetProductsBySeller(c echo.Context) error {
	id := c.Param("id")
	fmt.Println(id)
	sellerID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}

	in := pb.GetProductsRequest{
		SellerId: int32(sellerID),
	}

	resp, err := h.sellerGRPC.GetProductsBySeller(context.Background(), &in)
	if err != nil {
		fmt.Println(err)
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

	if result == nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "No products available",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"products": result,
	})
}

// @Summary      Get Product By Category
// @Description  Get products by category ID
// @Tags         Seller
// @Accept       json
// @Produce      json
// @Param ID path string true "category name"
// @Success      200  {object}  []model.Product
// @Failure      400  {object}  helpers.message
// @Failure      401  {object}  helpers.message
// @Failure      500  {object}  helpers.message
// @Router       /products/category/:category [Get]
func (h *SellerHandler) GetProductsByCategory(c echo.Context) error {
	category := c.Param("category")

	in := pb.GetProductByCategoryRequest{
		CategoryName: category,
	}

	resp, err := h.sellerGRPC.GetProductsByCategory(context.Background(), &in)
	if err != nil {
		fmt.Println(err)
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

// @Summary      Get Product By ID
// @Description  Get products by product ID
// @Tags         Seller
// @Accept       json
// @Produce      json
// @Param ID path int true "Product ID"
// @Success      200  {object}  model.Product
// @Failure      400  {object}  helpers.message
// @Failure      401  {object}  helpers.message
// @Failure      500  {object}  helpers.message
// @Router       /products/:id [Get]
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
		return echo.NewHTTPError(404, echo.Map{
			"message": "product not found", // add custom err later
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
		Stock:      int32(input.Stock),
		CategoryId: int32(input.Category_id),
		Discount:   int32(input.Discount),
	}

	resp, err := h.sellerGRPC.UpdateProduct(context.Background(), &in)
	if err != nil {
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"product": model.Product{
			ID:          int(resp.Productid),
			SellerID:    int(resp.SellerId),
			Name:        resp.Name,
			Price:       resp.Price,
			Stock:       int(resp.Stock),
			Category_id: int(resp.CategoryId),
			Discount:    int(resp.Discount),
		},
	})
}

// seller
func (h *SellerHandler) AddSeller(c echo.Context) error {
	var input model.Seller
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}

	in := pb.AddSellerRequest{
		SellerId: int32(input.ID),
		Name:     input.Name,
	}

	resp, err := h.sellerGRPC.AddSeller(context.Background(), &in)
	if err != nil {
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Seller successfully added!",
		"seller": model.SellerIDName{
			ID:   int(resp.SellerId),
			Name: resp.Name,
		},
	})
}

// @Summary      Get All Sellers
// @Description  Get All Sellers
// @Tags         Seller
// @Accept       json
// @Produce      json
// @Success      200  {object}  []model.Seller
// @Failure      400  {object}  helpers.message
// @Failure      401  {object}  helpers.message
// @Failure      500  {object}  helpers.message
// @Router       /sellers [Get]
func (h *SellerHandler) GetAllSellers(c echo.Context) error {
	resp, err := h.sellerGRPC.GetAllSellers(context.Background(), &emptypb.Empty{})
	if err != nil {
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}

	var result []*model.Seller
	for _, v := range resp.Sellers {
		s := model.Seller{
			ID:         int(v.SellerId),
			Name:       v.Name,
			LastActive: v.LastActive,
		}
		result = append(result, &s)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"sellers": result,
	})
}

// @Summary      Get Seller By ID
// @Description  Get Seller by ID
// @Tags         Seller
// @Accept       json
// @Produce      json
// @Param ID path int true "Seller ID"
// @Success      200  {object}  model.SellerDetail
// @Failure      400  {object}  helpers.message
// @Failure      401  {object}  helpers.message
// @Failure      500  {object}  helpers.message
// @Router       /sellers/:id [Get]
func (h *SellerHandler) GetSellerByID(c echo.Context) error {
	id := c.Param("id")
	sellerID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}

	in := pb.GetSellerByIDRequest{
		SellerId: int32(sellerID),
	}

	resp, err := h.sellerGRPC.GetSellerByID(context.Background(), &in)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}

	return c.JSON(http.StatusOK, model.SellerDetail{
		ID:         int(resp.SellerId),
		Name:       resp.Name,
		LastActive: resp.LastActive,
		Address: model.Address{
			Id:      int(resp.AddressId),
			Address: resp.AddressName,
			Regency: resp.AddressRegency,
			City:    resp.AddressCity,
		},
	})
}

// @Summary      Get Seller By Name
// @Description  Get seller by seller Name
// @Tags         Seller
// @Accept       json
// @Produce      json
// @Param ID path string true "Seller Name"
// @Success      200  {object}  model.SellerDetail
// @Failure      400  {object}  helpers.message
// @Failure      401  {object}  helpers.message
// @Failure      500  {object}  helpers.message
// @Router       /sellers/name/:name [Get]
func (h *SellerHandler) GetSellerByName(c echo.Context) error {
	name := c.Param("name")

	in := pb.GetSellerByNameRequest{
		Name: name,
	}

	resp, err := h.sellerGRPC.GetSellerByName(context.Background(), &in)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}

	return c.JSON(http.StatusOK, model.SellerDetail{
		ID:         int(resp.SellerId),
		Name:       resp.Name,
		LastActive: resp.LastActive,
		Address: model.Address{
			Id:      int(resp.AddressId),
			Address: resp.AddressName,
			Regency: resp.AddressRegency,
			City:    resp.AddressCity,
		},
	})
}

func (h *SellerHandler) UpdateSellerName(c echo.Context) error {
	id := c.Get("id").(string)
	sellerID, _ := strconv.Atoi(id)

	var input model.SellerName
	if err := c.Bind(&input); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}

	in := pb.UpdateSellerNameRequest{
		SellerId: int32(sellerID),
		Name:     input.Name,
	}

	_, err := h.sellerGRPC.UpdateSellerName(context.Background(), &in)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(400, echo.Map{
			"message": "invalid input", // add custom err later
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Seller name updated to " + input.Name,
	})
}
