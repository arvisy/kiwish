package handler

import (
	"context"
	"errors"
	"log"
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

	intUserConv := strconv.Itoa(user.Id)
	intRoleConv := strconv.Itoa(user.RoleID)

	return &pb.LoginResponse{
		Id:    intUserConv,
		Role:  intRoleConv,
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

func (u *UserHandler) UpdateCustomer(ctx context.Context, in *pb.UpdateCustomerRequest) (*pb.UpdateCustomerResponse, error) {
	strCon, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Id:       strCon,
		Name:     in.Name,
		Email:    in.Email,
		Password: string(hashedPassword),
	}

	err = u.UserRepository.UpdateCustomer(user.Id, user)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateCustomerResponse{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func (u *UserHandler) DeleteCustomer(ctx context.Context, in *pb.DeleteCustomerRequest) (*pb.Empty, error) {
	strCon, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Id: strCon,
	}

	err = u.UserRepository.Delete(user.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (u *UserHandler) AddAddress(ctx context.Context, in *pb.AddAddressRequest) (*pb.AddAddressResponse, error) {
	strCon, err := strconv.Atoi(in.UserId)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Id: strCon,
	}

	address := model.Address{
		Address: in.Address,
		Regency: in.Regency,
		City:    in.City,
	}

	addressID, err := u.UserRepository.AddAddress(address)
	if err != nil {
		panic(err)
	}

	err = u.UserRepository.SetAddressCustomer(user, addressID)
	if err != nil {
		panic(err)
	}

	return &pb.AddAddressResponse{
		Address: address.Address,
		Regency: address.Regency,
		City:    address.City,
	}, nil
}

func (u *UserHandler) UpdateAddress(ctx context.Context, in *pb.UpdateAddressRequest) (*pb.UpdateAddressResponse, error) {
	strCon, err := strconv.Atoi(in.UserId)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Id: strCon,
	}

	address := model.Address{
		Address: in.Address,
		Regency: in.Regency,
		City:    in.City,
	}

	err = u.UserRepository.UpdateAddress(user.Id, address)
	if err != nil {
		log.Fatal(err)
	}

	return &pb.UpdateAddressResponse{
		Address: address.Address,
		Regency: address.Regency,
		City:    address.City,
	}, nil
}
