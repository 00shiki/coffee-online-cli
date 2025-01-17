package products

import (
	"coffee-online-cli/entity"
	"database/sql"
	"errors"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) FetchProducts() ([]entity.Product, error) {
	var products []entity.Product
	rows, err := r.db.Query("SELECT ProductID, ProductName, Stock, Price FROM Product")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var p entity.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Stock, &p.Price)
		if err != nil {
			return nil, fmt.Errorf("could not scan row: %v", err)
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *Repository) GetProductByID(id int) (*entity.Product, error) {
	product := &entity.Product{
		ID: id,
	}
	query, err := r.db.Prepare("SELECT ProductName, Stock, Price FROM Product WHERE ProductID=?")
	if err != nil {
		return nil, fmt.Errorf("could not prepare query: %v", err)
	}
	defer query.Close()
	rows, err := query.Query(id)
	if err != nil {
		return nil, fmt.Errorf("could not query product: %v", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, errors.New("could not find product")
	}
	err = rows.Scan(&product.Name, &product.Stock, &product.Price)
	if err != nil {
		return nil, fmt.Errorf("could not scan row: %v", err)
	}
	return product, nil
}

func (r *Repository) ProductStockUpdate(id int, newStock int) error {
	query, err := r.db.Prepare("UPDATE Product SET Stock=? WHERE ProductID=?")
	if err != nil {
		return fmt.Errorf("could not prepare query: %v", err)
	}
	defer query.Close()
	result, err := query.Exec(newStock, id)
	if err != nil {
		return fmt.Errorf("could not update product stock: %v", err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("could not find product")
	}
	return nil
}

func (r *Repository) PopularProduct() ([]entity.ProductPopular, error) {
	query := `
			SELECT
				p.ProductID,
				p.ProductName,
				SUM(op.Quantity) AS TotalOrder,
				SUM(op.Quantity*p.Price) AS TotalRevenue
			FROM
				OrderProduct op
			JOIN
				Product p ON op.ProductID = p.ProductID
			GROUP BY
				p.ProductID, p.ProductName
			ORDER BY
				TotalOrder DESC;`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []entity.ProductPopular
	for rows.Next() {
		var product entity.ProductPopular
		rows.Scan(
			&product.ID,
			&product.Name,
			&product.TotalOrder,
			&product.TotalRevenue,
		)
		products = append(products, product)
	}
	return products, nil
}

func (r *Repository) CreateProduct(product entity.Product) error {
	query, err := r.db.Prepare("INSERT INTO Product (ProductName, Stock, Price) VALUES (?, ?, ?)")
	if err != nil {
		return fmt.Errorf("could not prepare query: %v", err)
	}
	defer query.Close()
	_, err = query.Exec(product.Name, product.Stock, product.Price)
	if err != nil {
		return fmt.Errorf("could not insert product: %v", err)
	}
	return nil
}
