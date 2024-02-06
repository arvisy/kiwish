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
	// u := r.Group("/users")
	// u.Use(middleware.Authentication)
	// {
	// 	u.GET("", user.GetInfoCustomer)
	// }

	customer := r.Group("/api")
	customer.Use(middleware.Authentication, middleware.CustomerAuth)
	{
		customer.GET("/user", user.GetInfoCustomer)
		customer.PUT("/user", user.UpdateCustomer)
		customer.DELETE("/user", user.DeleteCustomer)
		customer.POST("/user/address", user.AddAddress)
		customer.PUT("/user/address", user.UpdateAddress)
	}

	admin := r.Group("/api/admin")
	admin.Use(middleware.Authentication, middleware.AdminAuth)
	{
		admin.GET("/user/:id", user.GetCustomerAdmin)
		admin.GET("/user/seller/:id", user.GetSellerAdmin)
		admin.GET("/user", user.GetAllCustomerAdmin)
		admin.GET("/user/seller", user.GetAllSellerAdmin)
		admin.PUT("/user/:id", user.UpdateCustomerAdmin)
		admin.DELETE("/user/:id", user.DeleteCustomerAdmin)
		admin.DELETE("/user/seller/:id", user.DeleteSellerAdmin)
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
