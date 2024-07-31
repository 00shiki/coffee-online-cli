package orders

import (
	"coffee-online-cli/entity"
	"database/sql"
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

func (r *Repository) OrderPayment(order *entity.Order) error {
	query1, err := r.db.Prepare("INSERT INTO Payments (PaymentAmount) VALUES(?)")
	if err != nil {
		return fmt.Errorf("could not prepare query: %v", err)
	}
	defer query1.Close()
	var totalAmount float64
	for _, product := range order.OrderProduct {
		totalAmount += float64(product.Quantity) * product.Product.Price
	}
	result, err := query1.Exec(totalAmount + 9000)
	if err != nil {
		return fmt.Errorf("could not execute query: %v", err)
	}
	paymentID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("could not get last payment id: %v", err)
	}
	order.Payment.ID = int(paymentID)
	query2, err := r.db.Prepare("SELECT PaymentAmount, PaymentDate FROM Payments WHERE PaymentID = ?")
	if err != nil {
		return fmt.Errorf("could not prepare query: %v", err)
	}
	defer query2.Close()
	rows, err := query2.Query(paymentID)
	if err != nil {
		return fmt.Errorf("could not execute query: %v", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return fmt.Errorf("could not find payment with paymentID: %v", paymentID)
	}
	err = rows.Scan(&order.Payment.PaymentAmount, &order.Payment.Date)
	if err != nil {
		return fmt.Errorf("could not scan row: %v", err)
	}
	return nil
}

func (r *Repository) CreateOrder(order *entity.Order) error {
	query1, err := r.db.Prepare("INSERT INTO Orders (UserID, PaymentID, ShippingID) VALUES(?, ?, 1)")
	if err != nil {
		return fmt.Errorf("could not prepare query: %v", err)
	}
	defer query1.Close()
	result, err := query1.Exec(order.User.ID, order.Payment.ID)
	if err != nil {
		return fmt.Errorf("could not execute query: %v", err)
	}
	orderID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("could not get last orderID: %v", err)
	}
	order.OrderID = int(orderID)
	query2, err := r.db.Prepare("INSERT INTO OrderProduct (OrderID, ProductID, Quantity) VALUES (?, ?, ?)")
	if err != nil {
		return fmt.Errorf("could not prepare query: %v", err)
	}
	defer query2.Close()
	for _, product := range order.OrderProduct {
		_, err := query2.Exec(orderID, product.Product.ID, product.Quantity)
		if err != nil {
			return fmt.Errorf("could not execute query: %v", err)
		}
	}
	query3, err := r.db.Prepare("UPDATE Product SET Stock = ? WHERE ProductID = ?")
	if err != nil {
		return fmt.Errorf("could not prepare query: %v", err)
	}
	defer query3.Close()
	for _, product := range order.OrderProduct {
		_, err := query3.Exec(product.Product.Stock-product.Quantity, product.Product.ID)
		if err != nil {
			return fmt.Errorf("could not execute query: %v", err)
		}
	}
	return nil
}
