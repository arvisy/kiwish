package main

import (
	"fmt"
	"ms-order/client"
	"ms-order/config"
	"ms-order/pb"
	"ms-order/repository"
	"ms-order/services"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	log := config.Newlogger("ms-order", zap.InfoLevel)
	cfg, err := config.New()
	if err != nil {
		log.Fatal("failed initialize config", zap.Error(err))
	}

	db, err := cfg.OpenDB()
	if err != nil {
		log.Fatal("failed open postgres connection", zap.Error(err))
	}

	repo := repository.NewMongo(db)
	client, close, err := client.New(cfg)
	defer close()

	service := services.New(services.NewServiceParam{
		Repo:   repo,
		Client: client,
		Cfg:    cfg,
		Log:    log,
	})

	grpcServer := grpc.NewServer()
	pb.RegisterCartServiceServer(grpcServer, service)
	pb.RegisterOrderServiceServer(grpcServer, service)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Port))
	if err != nil {
		log.Fatal("error create listener", zap.Error(err))
	}

	log.Info("starting order server", zap.Int("port", cfg.Port))
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatal("error serve listener", zap.Error(err))
	}
}
