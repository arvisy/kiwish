package services

import (
	"ms-order/client"
	"ms-order/config"
	"ms-order/pb"
	"ms-order/repository"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	pb.UnimplementedOrderServiceServer
	repo   *repository.MongoRepository
	client *client.Client
	cfg    config.Config
	log    *zap.Logger
}

type NewServiceParam struct {
	Repo   *repository.MongoRepository
	Client *client.Client
	Cfg    config.Config
	Log    *zap.Logger
}

func New(param NewServiceParam) *Service {
	log := param.Log.WithOptions(zap.AddCallerSkip(1))

	return &Service{
		repo:   param.Repo,
		client: param.Client,
		cfg:    param.Cfg,
		log:    log,
	}
}

func ErrInternal(err error, log *zap.Logger) error {
	log.Error("server error", zap.Error(err))
	return status.Error(codes.Internal, err.Error())
}

func ErrInvalidArgument(err error) error {
	return status.Error(codes.InvalidArgument, err.Error())
}

func ErrNotFound(msg string) error {
	return status.Error(codes.NotFound, msg)
}
