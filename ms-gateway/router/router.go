package routes

import (
	"ms-gateway/handler"
	"ms-gateway/middleware"
	"ms-gateway/pb"

	"github.com/labstack/echo/v4"
)

func ApiRoutes(r *echo.Echo, user *handler.UserHandler, seller *handler.SellerHandler, order *handler.OrderHandler, userGRPC pb.UserServiceClient, sellerGRPC pb.SellerServiceClient, orderGRPC pb.OrderServiceClient) {

	// public endpoints
	r.POST("/register", user.Register) // working
	r.POST("/login", user.Login)       // working

	customer := r.Group("/api")
	customer.Use(middleware.Authentication, middleware.CustomerAuth)
	{
		customer.GET("/user", user.GetInfoCustomer)       // working kalo ganti jadi seller gabisa akses
		customer.PUT("/user", user.UpdateCustomer)        // working
		customer.DELETE("/user", user.DeleteCustomer)     // working
		customer.POST("/user/address", user.AddAddress)   // working
		customer.PUT("/user/address", user.UpdateAddress) // working
		customer.POST("/user/seller", user.CreateSeller)  // working
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
		p.GET("/:id", seller.GetProductByID)                       // working
		p.GET("/category/:category", seller.GetProductsByCategory) // working
		p.GET("/seller/:id", seller.GetProductsBySeller)           // working
	}
	px := p.Group("")
	px.Use(middleware.Authentication, middleware.SellerAuth)
	{
		px.POST("", seller.AddProduct) // working
		px.DELETE("/:id", seller.DeleteProduct)
		px.PUT("/:id", seller.UpdateProduct)
	}

	s := r.Group("/sellers")
	{
		s.GET("", seller.GetAllSellers)
		s.GET("/:id", seller.GetSellerByID)
		s.GET("/name/:name", seller.GetSellerByName)
		s.PUT("", seller.UpdateSellerName, middleware.Authentication, middleware.SellerAuth)
	}

	o := r.Group("/order")
	o.Use(middleware.Authentication, middleware.CheckPayment(userGRPC, sellerGRPC, orderGRPC))
	{
		o.POST("", order.CreateOrder)
		o.GET("/customer", order.GetAllOrderForCustomer, middleware.CustomerAuth)
		o.GET("/seller", order.GetAllOrderForSeller, middleware.SellerAuth)
	}

}
