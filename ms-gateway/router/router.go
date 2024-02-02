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
}
