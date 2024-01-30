package main

import (
	"log"
	"ms-seller/config"
	"net"

	"github.com/joho/godotenv"
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

	// postgresRepo := repository.NewPostgresRepository(db)
	// taskHandler := handlers.NewTaskHandler(postgresRepo)

	// grpcServer := grpc.NewServer()
	// pb.RegisterTaskServiceServer(grpcServer, taskHandler)

	// start gRPC server
	listen, err := net.Listen("tcp", ":50002")
	if err != nil {
		log.Println(err)
	}

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}
}
