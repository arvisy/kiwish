package cmd

import (
	"log"
	"ms-user/handler"
	"ms-user/pb"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func InitGrpc(UserHandler handler.UserHandler) {
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &UserHandler)

	godotenv.Load()
	listen, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Println(err)
	}

	log.Println("Server listening on :50001")
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}
}
