package orders

import "coffee-online-cli/entity"

type Repo interface {
	Reader
	Writer
}

type Reader interface {
	FetchUserOrders(userID int) ([]entity.Order, error)
	FetchPendingOrders() ([]entity.Order, error)
	GetOrderByID(id int) (*entity.Order, error)
}

type Writer interface {
	OrderPayment(order *entity.Order) error
	CreateOrder(order *entity.Order) error
	UpdateOrderShippingStatus(orderID int, shippingStatus entity.ShippingStatus) error
}
