package main

import (
	"log"
	"ms-seller/config"
	pb "ms-seller/proto"
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
	productService := service.NewProductService(postgresRepo)

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, productService)

	listen, err := net.Listen("tcp", ":50002")
	if err != nil {
		log.Println(err)
	}

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}
}
