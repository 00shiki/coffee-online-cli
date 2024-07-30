package orderproduct

type OrderProduct struct {
	OrderProductID	int                        
	Order
	Product
	Quantity		int
}