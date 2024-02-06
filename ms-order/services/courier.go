package services

import (
	"ms-order/exception"
	"ms-order/pb"
	"ms-order/repository"
)

type CourierService struct {
	pb.UnimplementedCartServiceServer
	repo *repository.MongoRepository
	err  *exception.ErrorHandler
}

func NewCourierService(repo *repository.MongoRepository, errh *exception.ErrorHandler) *CourierService {
	return &CourierService{
		repo: repo,
		err:  errh,
	}
}
