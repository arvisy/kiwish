package handler

import (
	"context"
	"errors"
	"ms-user/model"
	pb "ms-user/pb"
	"ms-user/repository"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	UserRepository repository.UserRepository
}

func NewUserHandler(UserRepository repository.UserRepository) *UserHandler {
	return &UserHandler{UserRepository: UserRepository}
}

func (u *UserHandler) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Name:     in.Name,
		Email:    in.Email,
		Password: string(hashedPassword),
		RoleID:   2,
	}

	err = u.UserRepository.AddUser(user)
	if err != nil {
		return &pb.RegisterResponse{}, err
	}

	return &pb.RegisterResponse{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (u *UserHandler) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := u.UserRepository.FindByCredentials(in.Email, in.Password)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	intConv := strconv.Itoa(user.Id)

	return &pb.LoginResponse{
		Id:    intConv,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (u *UserHandler) GetCustomer(ctx context.Context, in *pb.GetCustomerRequest) (*pb.GetCustomerResponse, error) {
	strCon, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Id: strCon,
	}

	err = u.UserRepository.GetCustomer(&user)
	if err != nil {
		return nil, err
	}

	if user.Id == 0 {
		return nil, errors.New("user not found")
	}

	return &pb.GetCustomerResponse{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

// func (u *UserHandler) UpdateCustomer(ctx context.Context, in *pb.UpdateCustomerRequest) (*pb.UpdateCustomerResponse, error) {

// }
