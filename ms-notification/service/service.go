package services

import (
	"ms-notification/config"
	"ms-notification/pb"
	"ms-notification/repository"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	pb.UnimplementedNotificationServiceServer
	repo *repository.MongoRepository
	cfg  config.Config
	log  *zap.Logger
}

type NewServiceParam struct {
	Repo *repository.MongoRepository
	Cfg  config.Config
	Log  *zap.Logger
}

func New(param NewServiceParam) *Service {
	log := param.Log.WithOptions(zap.AddCallerSkip(1))

	return &Service{
		repo: param.Repo,
		cfg:  param.Cfg,
		log:  log,
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
