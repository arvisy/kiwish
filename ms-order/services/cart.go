package services

import (
	"context"
	"errors"
	"ms-order/exception"
	"ms-order/model"
	"ms-order/pb"
	"ms-order/repository"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type CartService struct {
	pb.UnimplementedCartServiceServer
	repo *repository.MongoRepository
	err  *exception.ErrorHandler
}

func NewCartService(repo *repository.MongoRepository, errh *exception.ErrorHandler) *CartService {
	return &CartService{
		repo: repo,
		err:  errh,
	}
}

func (h CartService) CartCreate(ctx context.Context, in *pb.CartCreateRequest) (*pb.CartCreateResponse, error) {
	cart := &model.Cart{
		ID:        in.Cart.Id,
		UserID:    in.Cart.UserId,
		ProductID: in.Cart.ProductId,
		Quantity:  in.Cart.Quantity,
		CreatedAt: in.Cart.CreatedAt.AsTime(),
		UpdatedAt: in.Cart.UpdatedAt.AsTime(),
	}

	err := h.repo.Cart.Create(ctx, cart)
	if err != nil {
		return nil, h.err.ErrInternal(err)
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

func (h CartService) CartGetAll(ctx context.Context, in *pb.CartGetAllRequest) (*pb.CartGetAllResponse, error) {
	carts, err := h.repo.Cart.GetAll(ctx, in.UserId)
	if err != nil {
		return nil, h.err.ErrInternal(err)
	}

	var response pb.CartGetAllResponse
	for _, cart := range carts {
		response.Carts = append(response.Carts, &pb.Cart{
			Id: cart.ID,
		})
	}
	return &response, nil
}

func (h CartService) CartGetByID(ctx context.Context, in *pb.CartGetByIDRequest) (*pb.CartGetByIDResponse, error) {
	cart, err := h.repo.Cart.GetByID(ctx, in.Id, in.UserId)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			return nil, h.err.ErrNotFound()
		default:
			return nil, h.err.ErrInternal(err)
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

func (h CartService) CartUpdate(ctx context.Context, in *pb.CartUpdateRequest) (*pb.CartUpdateResponse, error) {
	cart, err := h.repo.Cart.GetByID(ctx, in.Id, in.UserId)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			return nil, h.err.ErrNotFound()
		default:
			return nil, h.err.ErrInternal(err)
		}
	}

	cart.Quantity = in.Quantity
	cart.UpdatedAt = time.Now()

	err = h.repo.Cart.Update(ctx, cart)
	if err != nil {
		return nil, h.err.ErrInternal(err)
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

func (h CartService) CartDeleteOne(ctx context.Context, in *pb.CartDeleteOneRequest) (*pb.CartDeleteOneResponse, error) {
	cart, err := h.repo.Cart.GetByID(ctx, in.Id, in.UserId)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			return nil, h.err.ErrNotFound()
		default:
			return nil, h.err.ErrInternal(err)
		}
	}
	err = h.repo.Cart.DeleteOne(ctx, in.Id, in.UserId)
	if err != nil {
		return nil, h.err.ErrInternal(err)
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

func (h CartService) CartDeleteAll(ctx context.Context, in *pb.CartDeleteAllRequest) (*pb.CartDeleteAllResponse, error) {
	carts, err := h.repo.Cart.GetAll(ctx, in.UserId)
	if err != nil {
		return nil, h.err.ErrInternal(err)
	}

	err = h.repo.Cart.DeleteAll(ctx, in.UserId)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			return nil, h.err.ErrNotFound()
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
