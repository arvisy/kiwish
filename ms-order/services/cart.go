package services

import (
	"context"
	"errors"
	"ms-order/client"
	"ms-order/model"
	"ms-order/pb"
	"ms-order/repository"
	"time"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CartService struct {
	pb.UnimplementedCartServiceServer
	repo   *repository.MongoRepository
	client *client.Client
	log    *zap.Logger
}

func NewCartService(repo *repository.MongoRepository, client *client.Client) *CartService {
	return &CartService{
		repo:   repo,
		client: client,
	}
}

func (s CartService) CartCreate(ctx context.Context, in *pb.CartCreateRequest) (*pb.CartCreateResponse, error) {
	cart := &model.Cart{
		ID:        in.Cart.Id,
		UserID:    in.Cart.UserId,
		ProductID: in.Cart.ProductId,
		Quantity:  in.Cart.Quantity,
		CreatedAt: in.Cart.CreatedAt.AsTime(),
		UpdatedAt: in.Cart.UpdatedAt.AsTime(),
	}

	err := s.repo.Cart.Create(ctx, cart)
	if err != nil {
		return nil, ErrInternal(err, s.log)
	}

	return &pb.CartCreateResponse{
		Cart: &pb.Cart{
			Id:        cart.ID,
			UserId:    cart.UserID,
			ProductId: cart.ProductID,
			Quantity:  cart.Quantity,
			CreatedAt: timestamppb.New(cart.CreatedAt),
			UpdatedAt: timestamppb.New(cart.UpdatedAt),
		},
		Message: "Success",
	}, nil
}

func (s CartService) CartGetAll(ctx context.Context, in *pb.CartGetAllRequest) (*pb.CartGetAllResponse, error) {
	carts, err := s.repo.Cart.GetAll(ctx, in.UserId)
	if err != nil {
		return nil,ErrInternal(err, s.log)
	}

	var response pb.CartGetAllResponse
	for _, cart := range carts {
		response.Carts = append(response.Carts, &pb.Cart{
			Id: cart.ID,
		})
	}
	return &response, nil
}

func (s CartService) CartGetByID(ctx context.Context, in *pb.CartGetByIDRequest) (*pb.CartGetByIDResponse, error) {
	cart, err := s.repo.Cart.GetByID(ctx, in.Id, in.UserId)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			return nil, ErrNotFound("the requested resources could not be found")
		default:
			return nil,ErrInternal(err, s.log)
		}
	}

	return &pb.CartGetByIDResponse{
		Cart: &pb.Cart{
			Id:        cart.ID,
			UserId:    cart.UserID,
			ProductId: cart.ProductID,
			Quantity:  cart.Quantity,
			CreatedAt: timestamppb.New(cart.CreatedAt),
			UpdatedAt: timestamppb.New(cart.UpdatedAt),
		},
	}, nil
}

func (s CartService) CartUpdate(ctx context.Context, in *pb.CartUpdateRequest) (*pb.CartUpdateResponse, error) {
	cart, err := s.repo.Cart.GetByID(ctx, in.Id, in.UserId)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			return nil, ErrNotFound("the requested resources could not be found")
		default:
			return nil,ErrInternal(err, s.log)
		}
	}

	cart.Quantity = in.Quantity
	cart.UpdatedAt = time.Now()

	err = s.repo.Cart.Update(ctx, cart)
	if err != nil {
		return nil,ErrInternal(err, s.log)
	}

	return &pb.CartUpdateResponse{
		Cart: &pb.Cart{
			Id:        cart.ID,
			UserId:    cart.UserID,
			ProductId: cart.ProductID,
			Quantity:  cart.Quantity,
			CreatedAt: timestamppb.New(cart.CreatedAt),
			UpdatedAt: timestamppb.New(cart.UpdatedAt),
		},
		Message: "success",
	}, nil
}

func (s CartService) CartDeleteOne(ctx context.Context, in *pb.CartDeleteOneRequest) (*pb.CartDeleteOneResponse, error) {
	cart, err := s.repo.Cart.GetByID(ctx, in.Id, in.UserId)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			return nil, ErrNotFound("the requested resources could not be found")
		default:
			return nil,ErrInternal(err, s.log)
		}
	}
	err = s.repo.Cart.DeleteOne(ctx, in.Id, in.UserId)
	if err != nil {
		return nil,ErrInternal(err, s.log)
	}

	return &pb.CartDeleteOneResponse{
		Cart: &pb.Cart{
			Id:        cart.ID,
			UserId:    cart.UserID,
			ProductId: cart.ProductID,
			Quantity:  cart.Quantity,
			CreatedAt: timestamppb.New(cart.CreatedAt),
			UpdatedAt: timestamppb.New(cart.UpdatedAt),
		},
		Message: "success",
	}, nil
}

func (s CartService) CartDeleteAll(ctx context.Context, in *pb.CartDeleteAllRequest) (*pb.CartDeleteAllResponse, error) {
	carts, err := s.repo.Cart.GetAll(ctx, in.UserId)
	if err != nil {
		return nil,ErrInternal(err, s.log)
	}

	err = s.repo.Cart.DeleteAll(ctx, in.UserId)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			return nil, ErrNotFound("the requested resources could not be found")
		default:
			return nil, err
		}
	}

	var response = &pb.CartDeleteAllResponse{}
	for _, cart := range carts {
		response.Ids = append(response.Ids, cart.ID)
	}
	response.Message = "success"

	return response, nil
}
