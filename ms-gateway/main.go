package main

import (
	"ms-gateway/handler"
	"ms-gateway/middleware"
	pb "ms-gateway/pb"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func main() {
	grpcConn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	grpcService := pb.NewUserServiceClient(grpcConn)
	u := handler.NewUserHandler(grpcService)

	e := echo.New()

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	public := e.Group("/user")
	{
		public.POST("/register", u.Register)
		public.POST("/login", u.Login)
	}

	private := e.Group("/api")
	private.Use(middleware.Authentication)
	{
		private.GET("/user", u.GetInfoCustomer)
		private.PUT("/user", u.UpdateCustomer)
		private.DELETE("/user", u.DeleteCustomer)
		private.POST("/user/address", u.AddAddress)
		// private.PUT("/user/address", )
	}

	e.Logger.Fatal(e.Start(":8080"))
}
