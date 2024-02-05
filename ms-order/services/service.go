package services

import (
	"ms-order/client"
	"ms-order/config"
	"ms-order/repository"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	CartService
	OrderService
	cfg config.Config
	log *zap.Logger
}

type NewServiceParam struct {
	Repo   *repository.MongoRepository
	Client *client.Client
	Cfg    config.Config
	Log    *zap.Logger
}

func New(param NewServiceParam) *Service {
	return &Service{
		CartService:  CartService{repo: param.Repo, client: param.Client, log: param.Log},
		OrderService: OrderService{repo: param.Repo, client: param.Client, log: param.Log},
		cfg:          param.Cfg,
		log:          param.Log,
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
