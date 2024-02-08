package helpers

import "ms-order/model"

func CalculateSubtotal(products []model.Product) float64 {
	var total float64
	for _, p := range products {
		total += p.Price * float64(p.Quantity)
	}
	return total
}

func CalculateTotal(order *model.Order) float64 {
	return order.Shipment.Price + order.Subtotal
}
