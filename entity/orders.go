package entity

import "time"

type Order struct {
	ID           int
	OrderProduct []OrderProduct
	Date         time.Time
	User
	Payment
	ShippingStatus
}

type Payment struct {
	ID            int
	PaymentAmount float64
	Date          time.Time
}

type ShippingStatus int

const (
	Pending ShippingStatus = iota + 1
	Shipped
	Delivered
)
