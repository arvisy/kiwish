package handler

import (
	"context"
	"errors"
	"fmt"
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
		fmt.Println(err)
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
		fmt.Println(err)
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
		return nil, err
	}

	err = u.UserRepository.SetAddressCustomer(user, addressID)
	if err != nil {
		return nil, err
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

	addressID, err := u.UserRepository.GetAddressID(user.Id)
	if err != nil {
		return nil, err
	}

	err = u.UserRepository.UpdateAddress(addressID, address)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateAddressResponse{
		Address: address.Address,
		Regency: address.Regency,
		City:    address.City,
	}, nil
}

func (u *UserHandler) GetCustomerAdmin(ctx context.Context, in *pb.GetCustomerAdminRequest) (*pb.GetCustomerAdminResponse, error) {
	userID, err := strconv.Atoi(in.UserId)
	if err != nil {
		return nil, err
	}

	user, err := u.UserRepository.GetUserAdmin(userID)
	if err != nil {
		return nil, err
	}

	return &pb.GetCustomerAdminResponse{
		UserId: in.UserId,
		Name:   user.Name,
		Email:  user.Email,
	}, nil
}

func (u *UserHandler) GetAllCustomerAdmin(ctx context.Context, in *pb.Empty) (*pb.GetAllCustomerAdminResponse, error) {
	users, err := u.UserRepository.GetAllCustomerAdmin()
	if err != nil {
		return nil, err
	}

	var customerResponses []*pb.CustomerResponse
	for _, user := range users {
		userID := strconv.Itoa(user.Id)

		customerResponses = append(customerResponses, &pb.CustomerResponse{
			UserId: userID,
			Name:   user.Name,
			Email:  user.Email,
		})
	}

	return &pb.GetAllCustomerAdminResponse{
		Customers: customerResponses,
	}, nil
}

func (u *UserHandler) UpdateCustomerAdmin(ctx context.Context, in *pb.UpdateCustomerAdminRequest) (*pb.UpdateCustomerAdminResponse, error) {
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

	return &pb.UpdateCustomerAdminResponse{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func (u *UserHandler) DeleteCustomerAdmin(ctx context.Context, in *pb.DeleteCustomerAdminRequest) (*pb.Empty, error) {
	strCon, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Id: strCon,
	}

	err = u.UserRepository.DeleteCustomer(user.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (u *UserHandler) GetSellerAdmin(ctx context.Context, in *pb.GetSellerAdminRequest) (*pb.GetSellerAdminResponse, error) {
	userID, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}

	seller, err := u.UserRepository.GetSellerAdmin(userID)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &pb.GetSellerAdminResponse{
		Name:  seller.Name,
		Email: seller.Email,
	}, nil
}

func (u *UserHandler) GetAllSellerAdmin(ctx context.Context, in *pb.Empty) (*pb.GetAllSellerAdminResponse, error) {
	users, err := u.UserRepository.GetAllSellerAdmin()
	if err != nil {
		return nil, err
	}

	var sellerResponses []*pb.SellerResponseAdmin
	for _, user := range users {
		userID := strconv.Itoa(user.Id)

		sellerResponses = append(sellerResponses, &pb.SellerResponseAdmin{
			UserId: userID,
			Name:   user.Name,
			Email:  user.Email,
		})
	}

	return &pb.GetAllSellerAdminResponse{
		Sellers: sellerResponses,
	}, nil
}

func (u *UserHandler) DeleteSellerAdmin(ctx context.Context, in *pb.DeleteSellerAdminRequest) (*pb.Empty, error) {
	strCon, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Id: strCon,
	}

	err = u.UserRepository.DeleteSellerAdmin(user.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (u *UserHandler) CreateSeller(ctx context.Context, in *pb.CreateSellerRequest) (*pb.Empty, error) {
	strCon, err := strconv.Atoi(in.Id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	user := model.User{
		Id: strCon,
	}

	err = u.UserRepository.CreateSeller(user.Id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (u *UserHandler) GetUserAddress(ctx context.Context, in *pb.GetUserAddressRequest) (*pb.GetUserAddressResponse, error) {
	strCon, err := strconv.Atoi(in.UserId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	addID, err := u.UserRepository.GetAddressID(strCon)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	address := model.Address{
		Id: addID,
	}

	res, err := u.UserRepository.GetAddress(&address)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	resp := pb.GetUserAddressResponse{
		AddressId: int64(res.Id),
		Address:   res.Address,
		Regency:   res.Regency,
		City:      res.City,
	}

	return &resp, nil
}
