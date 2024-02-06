package main

import (
	"log"
	"ms-gateway/handler"
	"ms-gateway/middleware"
	pb "ms-gateway/pb"
	routes "ms-gateway/router"

	"github.com/labstack/echo/v4"
	midd "github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	userConn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	sellerConn, err := grpc.Dial(":50002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer sellerConn.Close()

	userService := pb.NewUserServiceClient(userConn)
	u := handler.NewUserHandler(userService)

	sellerService := pb.NewSellerServiceClient(sellerConn)
	s := handler.NewSellerHandler(sellerService)

	e := echo.New()

	routes.ApiRoutes(e, u, s)

	e.Use(midd.Logger())
	e.Use(midd.Recover())

	// public := e.Group("/user")
	// {
	// 	public.POST("/register", u.Register)
	// 	public.POST("/login", u.Login)
	// }

	private := e.Group("/api")
	private.Use(middleware.Authentication, middleware.CustomerAuth)
	{
		private.GET("/user", u.GetInfoCustomer)
		private.PUT("/user", u.UpdateCustomer)
		private.DELETE("/user", u.DeleteCustomer)
		private.POST("/user/address", u.AddAddress)
		private.PUT("/user/address", u.UpdateAddress)
	}

	admin := e.Group("/api/admin")
	admin.Use(middleware.Authentication, middleware.AdminAuth)
	{
		admin.GET("/user/:id", u.GetCustomerAdmin)
		admin.GET("/user", u.GetAllCustomerAdmin)
		admin.PUT("/user/:id", u.UpdateCustomerAdmin)
	}
	// private := e.Group("")
	// // private.Use(middleware.Authentication)
	// {
	// 	private.GET("/users", u.GetInfoCustomer, middleware.Authentication)
	// }

	e.Logger.Fatal(e.Start(":8080"))
}
