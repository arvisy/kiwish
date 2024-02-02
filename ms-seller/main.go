package main

import (
	"log"
	"ms-seller/config"
	"ms-seller/pb"
	"ms-seller/repository"
	"ms-seller/service"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := config.Open(config.DefaultPostgresConfig())
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	postgresRepo := repository.NewPostgresRepository(db)
	sellerService := service.NewSellerService(postgresRepo)

	grpcServer := grpc.NewServer()
	pb.RegisterSellerServiceServer(grpcServer, sellerService)

	listen, err := net.Listen("tcp", ":50002")
	if err != nil {
		log.Println(err)
	}

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}
}
