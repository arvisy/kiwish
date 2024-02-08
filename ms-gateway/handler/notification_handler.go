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
