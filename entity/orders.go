package entity

type Order struct {
	OrderID		int
	UserID		int
	Payment		
	Shipping	
	OrderDate	string
}

type Payment struct {
	PaymentID		int
	PaymentAmount	float64
	PaymentDate		string
}

type Shipping struct {
	ShippingID		int
	ShippingStatus	string
}