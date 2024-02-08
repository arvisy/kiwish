package main

import (
	"log"
	"ms-gateway/handler"
	pb "ms-gateway/pb"
	routes "ms-gateway/router"

	"github.com/labstack/echo/v4"
	midd "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// @title Kiwish
// @version 1.0
// @description E-commerce API with product management and shipment features.

// @contact.name Kiet Asmara
// @contact.url http://www.swagger.io/support
// @contact.email kiet123pascal@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
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

	notifConn, err := grpc.Dial(":50004", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	defer sellerConn.Close()

	userService := pb.NewUserServiceClient(userConn)
	sellerService := pb.NewSellerServiceClient(sellerConn)
	orderService := pb.NewOrderServiceClient(orderConn)
	notifService := pb.NewNotificationServiceClient(notifConn)

	u := handler.NewUserHandler(userService, sellerService)
	s := handler.NewSellerHandler(sellerService)
	o := handler.NewOrderHandler(userService, sellerService, orderService, notifService)
	nt := handler.NewNotificationHandler(notifService)

	e := echo.New()

	routes.ApiRoutes(e, u, s, o, nt, userService, sellerService, orderService, notifService)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// e.Use(midd.Logger())
	e.Use(midd.Recover())

	e.Logger.Fatal(e.Start(":8080"))
}
