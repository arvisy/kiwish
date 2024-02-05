package routes

import (
	"ms-gateway/handler"
	"ms-gateway/middleware"

	"github.com/labstack/echo/v4"
)

func ApiRoutes(r *echo.Echo, user *handler.UserHandler, seller *handler.SellerHandler) {

	// public endpoints
	r.POST("/register", user.Register)
	r.POST("/login", user.Login)

	// private
	u := r.Group("/users")
	u.Use(middleware.Authentication)
	{
		u.GET("", user.GetInfoCustomer)
	}

	p := r.Group("/products")
	{
		p.GET("/:id", seller.GetProductByID)
		p.GET("/category/:category", seller.GetProductsByCategory)
		p.GET("/seller/:id", seller.GetProductsBySeller)
	}
	px := p.Group("")
	px.Use(middleware.Authentication, middleware.SellerAuth)
	{
		px.POST("", seller.AddProduct)
		px.DELETE("/:id", seller.DeleteProduct)
	}

	// s := r.Group("/sellers")
	// {
	// 	s.POST("", seller.)
	// }

}
