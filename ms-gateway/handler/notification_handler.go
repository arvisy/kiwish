package handler

import (
	"context"
	"ms-gateway/pb"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type NotificationHandler struct {
	notifGRPC pb.NotificationServiceClient
}

func NewNotificationHandler(notif pb.NotificationServiceClient) *NotificationHandler {
	return &NotificationHandler{
		notifGRPC: notif,
	}
}

// @Summary      Get Notification
// @Description  Get All Notification
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Success      200  {object} 	object{message=string}
// @Failure      400  {object}  helper.Message
// @Failure      401  {object}  helper.Message
// @Failure      500  {object}  helper.Message
// @Router       /order/notification [Get]
func (h NotificationHandler) GetAll(c echo.Context) error {
	userid, _ := strconv.ParseInt(c.Get("id").(string), 10, 64)
	res, err := h.notifGRPC.GetAllNotification(context.Background(), &pb.GetAllNotificationRequest{
		ReceiverId: userid,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary      Mark All Notification
// @Description  Mark All notification as read
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Success      200  {object} 	object{message=string}
// @Failure      400  {object}  helper.Message
// @Failure      401  {object}  helper.Message
// @Failure      500  {object}  helper.Message
// @Router       /order/notification/mark [Put]
func (h NotificationHandler) MarkAllAsRead(c echo.Context) error {
	userid, _ := strconv.ParseInt(c.Get("id").(string), 10, 64)
	res, err := h.notifGRPC.MarkAllAsRead(context.Background(), &pb.MarkAllAsReadRequest{
		ReceiverId: userid,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary      Mark Notification By Id
// @Description  Mark Notification By id
// @Tags         Order
// @Accept       json
// @Produce      json
// @Param Authorization header string true "JWT Token"
// @Param ID path int true "Order ID"
// @Success      200  {object} 	object{message=string}
// @Failure      400  {object}  helper.Message
// @Failure      401  {object}  helper.Message
// @Failure      500  {object}  helper.Message
// @Router       /order/notification/mark/:id [Put]
func (h NotificationHandler) MarkAsRead(c echo.Context) error {
	userid, _ := strconv.ParseInt(c.Get("id").(string), 10, 64)
	notifid := c.Param("id")
	res, err := h.notifGRPC.MarkAsRead(context.Background(), &pb.MarkAsReadRequest{
		Id:         notifid,
		ReceiverId: userid,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}
