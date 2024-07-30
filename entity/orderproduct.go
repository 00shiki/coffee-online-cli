package entity

type OrderProduct struct {
	ID int
	Order
	Product
	Quantity int
}
