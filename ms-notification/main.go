package main

import (
	"fmt"
	"ms-notification/config"
	"ms-notification/pb"
	"ms-notification/repository"
	services "ms-notification/service"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := config.Newlogger("ms-notification", zap.InfoLevel)
	cfg, err := config.New()
	if err != nil {
		log.Fatal("failed initialize config", zap.Error(err))
	}

	db, err := cfg.OpenDB()
	if err != nil {
		log.Fatal("failed open postgres connection", zap.Error(err))
	}

	repo := repository.NewMongo(db)
	service := services.New(services.NewServiceParam{
		Repo: repo,
		Cfg:  cfg,
		Log:  log,
	})

	grpcServer := grpc.NewServer()
	pb.RegisterNotificationServiceServer(grpcServer, service)
	reflection.Register(grpcServer)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Port))
	if err != nil {
		log.Fatal("error create listener", zap.Error(err))
	}

	log.Info("starting order server", zap.Int("port", cfg.Port))
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatal("error serve listener", zap.Error(err))
	}
	//
}
