package entity

type Product struct {
	ID        int
	Name      string
	Stock     int
	Price     float64
	CreatedAt string
	UpdateAt  string
}

type ProductPopular struct {
	Name          string
	TotalOrder    int
	TotalSpending float64
}
