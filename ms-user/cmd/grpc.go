package cmd

import (
	"log"
	"ms-user/handler"
	pb "ms-user/pb"
	"net"

	"google.golang.org/grpc"
)

func InitGrpc(UserHandler handler.UserHandler) {
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &UserHandler)

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
