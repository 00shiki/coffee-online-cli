package entity

import "time"

type Order struct {
	OrderID      int
	OrderProduct []OrderProduct
	Date         time.Time
	User
	Payment
	Shipping
}

type Payment struct {
	ID            int
	PaymentAmount float64
	Date          time.Time
}

type Shipping struct {
	ID     int
	Status string
}
