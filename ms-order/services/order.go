package services

import (
	"context"
	"ms-order/client"
	"ms-order/pb"
	orderpb "ms-order/pb"
	"ms-order/repository"

	"go.uber.org/zap"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	repo   *repository.MongoRepository
	client *client.Client
	log    *zap.Logger
}

func NewOrderService(repo *repository.MongoRepository, client *client.Client) *OrderService {
	return &OrderService{
		repo:   repo,
		client: client,
	}
}

func (h OrderService) OrderDirectCreate(ctx context.Context, in *orderpb.OrderDirectCreateRequest) (*orderpb.OrderDirectCreateResponse, error) {
	return nil, nil
}
