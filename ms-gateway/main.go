package main

import (
	"log"
	"ms-gateway/handler"
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

	orderConn, err := grpc.Dial(":50003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer sellerConn.Close()

	userService := pb.NewUserServiceClient(userConn)
	sellerService := pb.NewSellerServiceClient(sellerConn)
	orderService := pb.NewOrderServiceClient(orderConn)

	u := handler.NewUserHandler(userService, sellerService)
	s := handler.NewSellerHandler(sellerService)
	o := handler.NewOrderHandler(userService, sellerService, orderService)

	e := echo.New()

	routes.ApiRoutes(e, u, s, o, userService, sellerService, orderService)

	// e.Use(midd.Logger())
	e.Use(midd.Recover())

	e.Logger.Fatal(e.Start(":8080"))
}
