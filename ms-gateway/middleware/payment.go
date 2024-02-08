package middleware

import (
	"context"
	"fmt"
	"ms-gateway/helper"
	"ms-gateway/model"
	"ms-gateway/pb"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CheckPayment(userGRPC pb.UserServiceClient, sellerGRPC pb.SellerServiceClient, orderGRPC pb.OrderServiceClient, notifGRPC pb.NotificationServiceClient) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userid, _ := strconv.ParseInt(c.Get("id").(string), 10, 64)
			role := c.Get("role").(string)

			switch {
			case role == "2": // customer
				res, err := orderGRPC.OrderGetAll(context.Background(), &pb.OrderGetAllRequest{
					Userid: userid,
					Role:   role,
					Status: model.ORDER_STATUS_UNPAID,
				})
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, err)
				}

				for _, o := range res.Orders {
					inv, errxendit := helper.GetInvoice(o.Payment.InvoiceId)
					if errxendit != nil {
						switch {
						case errxendit.Status() == "404":
							return next(c)
						default:
							return echo.NewHTTPError(http.StatusInternalServerError, errxendit)
						}
					}

					if inv.GetStatus() == "PAID" {
						// update order
						res, err := orderGRPC.OrderUpdate(context.Background(), &pb.OrderUpdateRequest{
							Id:            o.Id,
							Userid:        userid,
							Role:          role,
							OrderStatus:   model.ORDER_STATUS_PACKED,
							PaymentStatus: "PAID",
						})
						if err != nil {
							return echo.NewHTTPError(http.StatusInternalServerError, err)
						}

						// update stock
						for _, p := range res.Order.Products {
							res, err := sellerGRPC.GetProductByID(context.Background(), &pb.GetProductByIDRequest{
								ProductId: int32(p.Id),
							})
							if err != nil {
								return echo.NewHTTPError(http.StatusInternalServerError, errxendit)
							}

							res.Stock -= int32(p.Quantity)

							_, err = sellerGRPC.UpdateProduct(context.Background(), &pb.UpdateProductRequest{
								Productid:  res.Productid,
								SellerId:   res.SellerId,
								Name:       res.Name,
								Price:      res.Price,
								Stock:      res.Stock,
								CategoryId: res.CategoryId,
								Discount:   res.Discount,
							})
							if err != nil {
								return echo.NewHTTPError(http.StatusInternalServerError, errxendit)
							}
						}

						// send notif seller
						notifGRPC.CreateNotification(context.Background(), &pb.CreateNotificationRequest{
							ReceiverId:  res.Order.Seller.Id,
							Subject:     "Segera Proses Orderanmu",
							Description: fmt.Sprintf("Hai %s anda telah menerima orderan #ID %v. Segera proses orderanmu", res.Order.Seller.Name, res.Order.Id),
						})
					}
				}

			case role == "3": // seller
				res, err := orderGRPC.OrderGetAll(context.Background(), &pb.OrderGetAllRequest{
					Userid: userid,
					Role:   role,
					Status: model.ORDER_STATUS_UNPAID,
				})
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, err)
				}

				for _, o := range res.Orders {
					inv, errxendit := helper.GetInvoice(o.Payment.InvoiceId)
					if errxendit != nil {
						switch {
						case errxendit.Status() == "404":
							return next(c)
						default:
							return echo.NewHTTPError(http.StatusInternalServerError, errxendit)
						}
					}

					if inv.GetStatus() == "PAID" {
						// update order
						res, err := orderGRPC.OrderUpdate(context.Background(), &pb.OrderUpdateRequest{
							Id:            o.Id,
							Userid:        userid,
							Role:          role,
							OrderStatus:   model.ORDER_STATUS_PACKED,
							PaymentStatus: "PAID",
						})
						if err != nil {
							return echo.NewHTTPError(http.StatusInternalServerError, err)
						}

						// update stock
						for _, p := range res.Order.Products {
							res, err := sellerGRPC.GetProductByID(context.Background(), &pb.GetProductByIDRequest{
								ProductId: int32(p.Id),
							})
							if err != nil {
								return echo.NewHTTPError(http.StatusInternalServerError, errxendit)
							}

							res.Stock -= int32(p.Quantity)

							_, err = sellerGRPC.UpdateProduct(context.Background(), &pb.UpdateProductRequest{
								Productid:  res.Productid,
								SellerId:   res.SellerId,
								Name:       res.Name,
								Price:      res.Price,
								Stock:      res.Stock,
								CategoryId: res.CategoryId,
								Discount:   res.Discount,
							})
							if err != nil {
								return echo.NewHTTPError(http.StatusInternalServerError, errxendit)
							}
						}

						// send notif seller
						notifGRPC.CreateNotification(context.Background(), &pb.CreateNotificationRequest{
							ReceiverId:  res.Order.Seller.Id,
							Subject:     "Segera Proses Orderanmu",
							Description: fmt.Sprintf("Hai %s anda telah menerima orderan #ID %v. Segera proses orderanmu", res.Order.Seller.Name, res.Order.Id),
						})
					}
				}
			}

			return next(c)
		}
	}
}
