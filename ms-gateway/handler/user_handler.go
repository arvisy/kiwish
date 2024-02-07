package handler

import (
	"context"
	"fmt"
	"ms-gateway/helper"
	"ms-gateway/model"
	pb "ms-gateway/pb"
	"regexp"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userGRPC   pb.UserServiceClient
	sellerGRPC pb.SellerServiceClient
}

func NewUserHandler(userGRPC pb.UserServiceClient, sellerGRPC pb.SellerServiceClient) *UserHandler {
	return &UserHandler{
		userGRPC:   userGRPC,
		sellerGRPC: sellerGRPC,
	}
}

// @Summary      Register
// @Description  Register user name, email, and password
// @Tags         User
// @Accept       json
// @Produce      json
// @Param		 data body model.User true "The input user struct"
// @Success      201  {object}  model.RegisterResponse
// @Failure      400  {object}  helper.Message
// @Failure      401  {object}  helper.Message
// @Failure      404  {object}  helper.Message
// @Failure      500  {object}  helper.Message
// @Router       /register [Post]
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

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(payload.Email) {
		return c.JSON(400, helper.Response{
			Message: "invalid email format",
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

	return c.JSON(201, response)
}

// @Summary      Login
// @Description  Login user with email and password
// @Tags         User
// @Accept       json
// @Produce      json
// @Param		 data body model.LoginRequest true "The input login struct"
// @Success      200  {object}  model.LoginResponse
// @Failure      400  {object}  helper.Message
// @Failure      401  {object}  helper.Message
// @Failure      404  {object}  helper.Message
// @Failure      500  {object}  helper.Message
// @Router       /login [Post]
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

// @Summary      Get User Info
// @Description  Get a user's info
// @Tags         Seller
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Success      201  {object}  model.User
// @Failure      401  {object}  helper.Message
// @Failure      404  {object}  helper.Message
// @Failure      500  {object}  helper.Message
// @Router       /api/user [Get]
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

	getadd := pb.GetUserAddressRequest{
		UserId: userID,
	}

	_, err := u.userGRPC.GetUserAddress(context.TODO(), &getadd)
	if err == nil {
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

	role := c.Get("role").(string)

	if role == "3" {
		in2 := pb.AddSellerAddressRequest{
			SellerId:       int32(uID),
			AddressName:    response.Address,
			AddressRegency: response.Regency,
			AddressCity:    response.City,
		}

		_, err := u.sellerGRPC.AddSellerAddress(context.TODO(), &in2)
		if err != nil {
			return c.JSON(500, helper.Response{
				Message: "failed to add seller address",
				Detail:  err.Error(),
			})
		}
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

	// update seller address
	role := c.Get("role").(string)

	if role == "3" {
		getadd := pb.GetUserAddressRequest{
			UserId: userID,
		}

		address, err := u.userGRPC.GetUserAddress(context.TODO(), &getadd)
		if err != nil {
			return c.JSON(500, helper.Response{
				Message: "please add address before creating seller",
			})
		}

		addressCon, _ := strconv.Atoi(address.AddressId)

		in2 := pb.UpdateSellerAddressRequest{
			AddressId:      int32(addressCon), // todo
			AddressName:    response.Address,
			AddressRegency: response.Regency,
			AddressCity:    response.City,
		}

		_, err = u.sellerGRPC.UpdateAddress(context.TODO(), &in2)
		if err != nil {
			return c.JSON(500, helper.Response{
				Message: "failed to update seller address",
				Detail:  err.Error(),
			})
		}
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
			Message: "user not found",
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
			Message: "seller not found",
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

func (u *UserHandler) CreateSeller(c echo.Context) error {
	userID := c.Get("id").(string)

	fmt.Println(userID)

	strConUser, err := strconv.Atoi(userID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	getadd := pb.GetUserAddressRequest{
		UserId: userID,
	}

	address, err := u.userGRPC.GetUserAddress(context.TODO(), &getadd)
	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "please add address before creating seller",
		})
	}

	addselladd := pb.AddSellerAddressRequest{
		SellerId:       int32(strConUser),
		AddressName:    address.Address,
		AddressRegency: address.Regency,
		AddressCity:    address.City,
	}

	addressSeller, err := u.sellerGRPC.AddSellerAddress(context.TODO(), &addselladd)
	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed " + err.Error(),
		})
	}

	var payload model.SellerIDName
	err = c.Bind(&payload)
	if err != nil {
		fmt.Println(err)
		return c.JSON(400, helper.Response{
			Message: "invalid payload request",
		})
	}

	if payload.Name == "" {
		return c.JSON(400, helper.Response{
			Message: "name is required",
		})
	}

	seller, err := u.sellerGRPC.AddSeller(context.TODO(), &pb.AddSellerRequest{
		SellerId:  int32(strConUser),
		Name:      payload.Name,
		AddressId: addressSeller.AddressId,
	})

	if err != nil {
		fmt.Println(err)
		return c.JSON(500, helper.Response{
			Message: "failed to create seller",
		})
	}

	_, err = u.userGRPC.CreateSeller(context.TODO(), &pb.CreateSellerRequest{
		Id: userID,
	})
	if err != nil {
		return c.JSON(500, helper.Response{
			Message: "failed to change role",
		})
	}

	return c.JSON(200, echo.Map{
		"message": "seller created successfully",
		"seller":  seller,
	})
}
