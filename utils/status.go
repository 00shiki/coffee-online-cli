package utils

import "coffee-online-cli/entity"

func PrintShippingStatus(status entity.ShippingStatus) string {
	switch status {
	case entity.Pending:
		return "Pending"
	case entity.Shipped:
		return "Shipped"
	case entity.Delivered:
		return "Delivered"
	}
	return "Unknown"
}
