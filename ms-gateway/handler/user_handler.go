package handler

import (
	"context"
	"fmt"
	"ms-gateway/helper"
	"ms-gateway/model"
	pb "ms-gateway/pb"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userGRPC pb.UserServiceClient
}

func NewUserHandler(userGRPC pb.UserServiceClient) *UserHandler {
	return &UserHandler{userGRPC: userGRPC}
}

func (u *UserHandler) Register(c echo.Context) error {
	var payload model.User

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid request payload",
		})
	}

	if payload.Name == "" || payload.Email == "" || payload.Password == "" {
		return c.JSON(400, helper.Response{
			Message: "name, email, and password are required",
		})
	}

	in := pb.RegisterRequest{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}

	response, err := u.userGRPC.Register(context.TODO(), &in)
	if err != nil {
		return c.JSON(400, helper.Response{
			Message: "failed to register user",
		})
	}

	return c.JSON(200, response)
}

func (u *UserHandler) Login(c echo.Context) error {
	var loginRequest model.User
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid login request payload",
		})
	}

	if loginRequest.Email == "" || loginRequest.Password == "" {
		return c.JSON(400, helper.Response{
			Message: "email and password are required",
		})
	}

	response, err := u.userGRPC.Login(context.TODO(), &pb.LoginRequest{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	})

	if err != nil {
		return c.JSON(401, helper.Response{
			Message: "invalid email or password",
		})
	}

	claims := jwt.MapClaims{
		"id":   response.Id,
		"role": response.Role,
	}

	fmt.Println("id: ", response.Id)
	fmt.Println("role: ", response.Role)

	token, err := helper.GenerateJWTTokenWithClaims(claims)
	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed to generate JWT token",
		})
	}

	return c.JSON(200, echo.Map{
		"message": "login success",
		"token":   token,
	})
}

func (u *UserHandler) GetInfoCustomer(c echo.Context) error {
	userID := c.Get("id").(string)

	response, err := u.userGRPC.GetCustomer(context.TODO(), &pb.GetCustomerRequest{
		Id: userID,
	})

	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed to get user information",
		})
	}

	return c.JSON(200, echo.Map{
		"user": response,
	})
}

func (u *UserHandler) UpdateCustomer(c echo.Context) error {
	userID := c.Get("id").(string)

	var updateRequest pb.UpdateCustomerRequest
	if err := c.Bind(&updateRequest); err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid update request payload",
		})
	}

	response, err := u.userGRPC.UpdateCustomer(context.TODO(), &pb.UpdateCustomerRequest{
		Id:       userID,
		Name:     updateRequest.Name,
		Email:    updateRequest.Email,
		Password: updateRequest.Password,
	})

	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed to update user",
		})
	}

	return c.JSON(200, echo.Map{
		"message": "customer updated successfully",
		"user":    response,
	})
}

func (u *UserHandler) DeleteCustomer(c echo.Context) error {
	userID := c.Get("id").(string)

	response, err := u.userGRPC.DeleteCustomer(context.TODO(), &pb.DeleteCustomerRequest{
		Id: userID,
	})

	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed to delete user",
		})
	}

	return c.JSON(200, echo.Map{
		"message": "user deleted successfully",
		"user":    response,
	})
}

func (u *UserHandler) AddAddress(c echo.Context) error {
	userID := c.Get("id").(string)

	uID, _ := strconv.Atoi(userID)

	user := model.User{
		Id: uID,
	}

	if user.AddressID != 0 {
		return c.JSON(400, helper.Response{
			Message: "address already exist",
		})
	}

	var payload model.Address
	if err := c.Bind(&payload); err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid update request payload",
		})
	}

	in := pb.AddAddressRequest{
		UserId:  userID,
		Address: payload.Address,
		Regency: payload.Regency,
		City:    payload.City,
	}

	if payload.Address == "" || payload.Regency == "" || payload.City == "" {
		return c.JSON(400, helper.Response{
			Message: "address, regency, and city are required",
		})
	}

	response, err := u.userGRPC.AddAddress(context.TODO(), &in)

	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed to add address",
			Detail:  err.Error(),
		})
	}

	return c.JSON(201, echo.Map{
		"message": "address created successfully",
		"address": response,
	})
}

func (u *UserHandler) UpdateAddress(c echo.Context) error {
	userID := c.Get("id").(string)

	uID, _ := strconv.Atoi(userID)

	user := model.User{
		Id: uID,
	}

	if user.AddressID != 0 {
		return c.JSON(400, helper.Response{
			Message: "address not found",
		})
	}

	var updateRequest pb.UpdateAddressRequest
	if err := c.Bind(&updateRequest); err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid update request payload",
		})
	}

	if updateRequest.Address == "" || updateRequest.Regency == "" || updateRequest.City == "" {
		return c.JSON(400, helper.Response{
			Message: "address, regency, and city are required",
		})
	}

	response, err := u.userGRPC.UpdateAddress(context.TODO(), &pb.UpdateAddressRequest{
		UserId:  userID,
		Address: updateRequest.Address,
		Regency: updateRequest.Regency,
		City:    updateRequest.City,
	})

	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed to update address",
			Detail:  err.Error(),
		})
	}

	return c.JSON(200, helper.Response{
		Message: "address updated successfully",
		Detail:  response,
	})
}

func (u *UserHandler) GetCustomerAdmin(c echo.Context) error {
	userID := c.Param("id")

	response, err := u.userGRPC.GetCustomerAdmin(context.TODO(), &pb.GetCustomerAdminRequest{
		UserId: userID,
	})

	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed to get user",
		})
	}

	return c.JSON(200, helper.Response{
		Detail: response,
	})
}

func (u *UserHandler) GetAllCustomerAdmin(c echo.Context) error {
	response, err := u.userGRPC.GetAllCustomerAdmin(context.TODO(), &pb.Empty{})

	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed to get all user",
		})
	}

	return c.JSON(200, helper.Response{
		Detail: response,
	})
}

func (u *UserHandler) UpdateCustomerAdmin(c echo.Context) error {
	userID := c.Param("id")

	var updateRequest pb.UpdateCustomerRequest
	if err := c.Bind(&updateRequest); err != nil {
		return c.JSON(400, helper.Response{
			Message: "invalid update request payload",
		})
	}

	response, err := u.userGRPC.UpdateCustomer(context.TODO(), &pb.UpdateCustomerRequest{
		Id:       userID,
		Name:     updateRequest.Name,
		Email:    updateRequest.Email,
		Password: updateRequest.Password,
	})

	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed to update user",
		})
	}

	return c.JSON(200, echo.Map{
		"message": "customer updated successfully",
		"user":    response,
	})
}

func (u *UserHandler) DeleteCustomerAdmin(c echo.Context) error {
	userID := c.Param("id")

	response, err := u.userGRPC.DeleteCustomer(context.TODO(), &pb.DeleteCustomerRequest{
		Id: userID,
	})

	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed to delete user",
		})
	}

	return c.JSON(200, echo.Map{
		"message": "user deleted successfully",
		"user":    response,
	})
}

func (u *UserHandler) GetSellerAdmin(c echo.Context) error {
	userID := c.Param("id")

	response, err := u.userGRPC.GetSellerAdmin(context.TODO(), &pb.GetSellerAdminRequest{
		Id: userID,
	})

	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed to get seller",
		})
	}

	return c.JSON(200, helper.Response{
		Detail: response,
	})
}

func (u *UserHandler) GetAllSellerAdmin(c echo.Context) error {
	response, err := u.userGRPC.GetAllSellerAdmin(context.TODO(), &pb.Empty{})

	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed to get all seller",
		})
	}

	return c.JSON(200, helper.Response{
		Detail: response,
	})
}

func (u *UserHandler) DeleteSellerAdmin(c echo.Context) error {
	userID := c.Param("id")

	response, err := u.userGRPC.DeleteSellerAdmin(context.TODO(), &pb.DeleteSellerAdminRequest{
		Id: userID,
	})

	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed to delete seller",
		})
	}

	return c.JSON(200, echo.Map{
		"message": "seller deleted successfully",
		"user":    response,
	})
}
