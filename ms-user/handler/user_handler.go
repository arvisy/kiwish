package handler

import (
	"context"
	"ms-user/model"
	pb "ms-user/pb"
	"ms-user/repository"

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
		Id:       1,
		Name:     in.Name,
		Password: string(hashedPassword),
		RoleID:   2,
	}

	err = u.UserRepository.AddUser(user)
	if err != nil {
		return &pb.RegisterResponse{}, err
	}

	return &pb.RegisterResponse{
		Id:       int64(user.Id),
		Name:     user.Name,
		Password: user.Password,
	}, nil
}

func (u *UserHandler) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := u.UserRepository.FindByCredentials(in.Email, in.Password)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Id:       int64(user.Id),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
