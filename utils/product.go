package utils

import "coffee-online-cli/entity"

func CheckProductExist(cart []entity.OrderProduct, product entity.Product) int {
	for i, p := range cart {
		if p.Name == product.Name {
			return i
		}
	}
	return -1
}
