package products

import "coffee-online-cli/entity"

type Repo interface {
	Reader
	Writer
}

type Reader interface {
	FetchProducts() ([]entity.Product, error)
	GetProductByID(id int) (*entity.Product, error)
}

type Writer interface {
	ProductStockUpdate(id int, newStock int) error
}
