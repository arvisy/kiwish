package routes

import (
	"ms-gateway/handler"
	"ms-gateway/middleware"
	"ms-gateway/pb"

	"github.com/labstack/echo/v4"
)

func ApiRoutes(
	r *echo.Echo,
	user *handler.UserHandler,
	seller *handler.SellerHandler,
	order *handler.OrderHandler,
	notif *handler.NotificationHandler,
	userGRPC pb.UserServiceClient,
	sellerGRPC pb.SellerServiceClient,
	orderGRPC pb.OrderServiceClient,
	notifGRPC pb.NotificationServiceClient,
) {

	// public endpoints
	r.POST("/register", user.Register)
	r.POST("/login", user.Login)

	customer := r.Group("/api")
	customer.Use(middleware.Authentication, middleware.CustomerAuth)
	{
		customer.GET("/user", user.GetInfoCustomer)
		customer.PUT("/user", user.UpdateCustomer)
		customer.DELETE("/user", user.DeleteCustomer)
		customer.POST("/user/address", user.AddAddress)
		customer.PUT("/user/address", user.UpdateAddress)
		customer.POST("/user/seller", user.CreateSeller)
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
	o.Use(middleware.Authentication, middleware.CheckPayment(userGRPC, sellerGRPC, orderGRPC, notifGRPC))
	o.Use(middleware.Authentication)
	{
		o.POST("", order.CreateOrder)
		o.GET("", order.GetAllOrder)
		o.GET("/:id", order.GetByIdOrder)
		o.PUT("/accept/:id", order.ConfirmOrder, middleware.SellerAuth)
		o.PUT("/reject/:id", order.RejectOrder, middleware.SellerAuth)
		o.POST("/courier/:id", order.AddCourierinfo, middleware.SellerAuth)
		o.GET("/courier/:id", order.TrackCourierShipment)
		o.PUT("", order.CustomerConfirmOrder, middleware.CustomerAuth)
	}

	nt := o.Group("/notification")
	{
		nt.GET("", notif.GetAll)
		nt.PUT("/mark", notif.MarkAllAsRead)
		nt.PUT("/mark/:id", notif.MarkAsRead)
	}
}
