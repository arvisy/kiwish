package handler

import (
	"context"
	"ms-gateway/helper"
	"ms-gateway/model"
	pb "ms-gateway/pb"

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
		"id": response.Id,
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
		"message": "user information retrieved successfully",
		"user":    response,
	})
}

// func (u *UserHandler) UpdateUser(c echo.Context) error {
// 	userID := c.Get("id").(string)

// 	var updateRequest pb.UpdateUserRequest
// 	if err := c.Bind(&updateRequest); err != nil {
// 		return c.JSON(400, helper.Response{
// 			Message: "invalid update request payload",
// 		})
// 	}

// 	response, err := u.userGRPC.UpdateUser(context.TODO(), &pb.UpdateUserRequest{
// 		Id:       userID,
// 		Username: updateRequest.Username,
// 		Password: updateRequest.Password,
// 	})

// 	if err != nil {
// 		return c.JSON(500, helper.Response{
// 			Message: "failed to update user",
// 		})
// 	}

// 	return c.JSON(200, echo.Map{
// 		"message": "user information retrieved successfully",
// 		"user":    response,
// 	})
// }

// func (u *UserHandler) DeleteUser(c echo.Context) error {
// 	userID := c.Get("id").(string)

// 	response, err := u.userGRPC.DeleteUser(context.TODO(), &pb.DeleteUserRequest{
// 		Id: userID,
// 	})

// 	if err != nil {
// 		return c.JSON(500, helper.Response{
// 			Message: "failed to delete user",
// 		})
// 	}

// 	return c.JSON(200, echo.Map{
// 		"message": "user deleted successfully",
// 		"user":    response,
// 	})
// }

// func (u *UserHandler) AddTask(c echo.Context) error {
// 	userID := c.Get("id").(string)

// 	var addTaskRequest pb.AddTaskRequest
// 	if err := c.Bind(&addTaskRequest); err != nil {
// 		return c.JSON(400, helper.Response{
// 			Message: "invalid add task request payload",
// 		})
// 	}

// 	response, err := u.userGRPC.AddTask(context.TODO(), &pb.AddTaskRequest{
// 		UserId:      userID,
// 		Title:       addTaskRequest.Title,
// 		Description: addTaskRequest.Description,
// 		DueDate:     addTaskRequest.DueDate,
// 	})

// 	if err != nil {
// 		return c.JSON(500, helper.Response{
// 			Message: "failed to add task",
// 		})
// 	}

// 	return c.JSON(200, echo.Map{
// 		"message": "user information retrieved successfully",
// 		"task":    response,
// 	})
// }
