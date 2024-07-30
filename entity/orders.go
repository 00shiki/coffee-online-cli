package entity

type Order struct {
	OrderID		int
	User
	Payment		
	Shipping	
	Date		time.Time
}

type Payment struct {
	ID				int
	PaymentAmount	float64
	Date			time.Time
}

type Shipping struct {
	ID		int
	Status	string
}