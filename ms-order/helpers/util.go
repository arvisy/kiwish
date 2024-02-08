package helpers

import "ms-order/model"

func ValidOrderStatus(status string) bool {
	return status == model.ORDER_STATUS_UNPAID || status == model.ORDER_STATUS_PACKED || status == model.ORDER_STATUS_SHIPPED || status == model.ORDER_STATUS_COMPLETE
}
