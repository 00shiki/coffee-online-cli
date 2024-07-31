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
	order.ID = int(orderID)
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

func (r *Repository) UpdateOrderShippingStatus(orderID int, shippingStatus entity.ShippingStatus) error {
	query, err := r.db.Prepare("UPDATE Orders SET ShippingID = ? WHERE ID = ?")
	if err != nil {
		return fmt.Errorf("could not prepare query: %v", err)
	}
	defer query.Close()
	result, err := query.Exec(shippingStatus, orderID)
	if err != nil {
		return fmt.Errorf("could not execute query: %v", err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get rows affected: %v", err)
	}
	if affected == 0 {
		return fmt.Errorf("order not found")
	}
	return nil
}

func (r *Repository) FetchUserOrders(userID int) ([]entity.Order, error) {
	var orders []entity.Order
	query, err := r.db.Prepare("SELECT OrderID, OrderDate, ShippingID FROM Orders WHERE UserID = ?")
	if err != nil {
		return nil, fmt.Errorf("could not prepare query: %v", err)
	}
	defer query.Close()
	rows, err := query.Query(userID)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var order entity.Order
		err := rows.Scan(&order.ID, &order.Date, &order.ShippingStatus)
		if err != nil {
			return nil, fmt.Errorf("could not scan row: %v", err)
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *Repository) FetchPendingOrders() ([]entity.Order, error) {
	var orders []entity.Order
	query, err := r.db.Prepare("SELECT OrderID, OrderDate, ShippingID FROM Orders")
	if err != nil {
		return nil, fmt.Errorf("could not prepare query: %v", err)
	}
	defer query.Close()
	rows, err := query.Query()
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var order entity.Order
		err := rows.Scan(&order.ID, &order.Date, &order.ShippingStatus)
		if err != nil {
			return nil, fmt.Errorf("could not scan row: %v", err)
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *Repository) GetOrderByID(orderID int) (*entity.Order, error) {
	order := &entity.Order{
		ID: orderID,
	}
	query1, err := r.db.Prepare(`
	SELECT u.UserID, u.Name AS UserName, p.PaymentID, p.PaymentAmount, p.PaymentDate, o.OrderDate, o.ShippingID AS ShippingStatus
	FROM Orders o
			 JOIN Users u ON o.UserID = u.UserID
			 JOIN Payments p ON o.PaymentID = p.PaymentID
	WHERE OrderID = ?
	`)
	if err != nil {
		return nil, fmt.Errorf("could not prepare query: %v", err)
	}
	defer query1.Close()
	rowsOrders, err := query1.Query(orderID)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}
	defer rowsOrders.Close()
	if !rowsOrders.Next() {
		return nil, fmt.Errorf("could not find order")
	}
	err = rowsOrders.Scan(
		&order.User.ID,
		&order.User.Name,
		&order.Payment.ID,
		&order.Payment.PaymentAmount,
		&order.Payment.Date,
		&order.Date,
		&order.ShippingStatus,
	)
	if err != nil {
		return nil, fmt.Errorf("could not scan row: %v", err)
	}
	query2, err := r.db.Prepare(`
	SELECT o.OrderProductID, p.ProductID, p.ProductName, p.Stock, p.Price, o.Quantity
	FROM OrderProduct o
			 JOIN Product p ON o.ProductID = p.ProductID
	WHERE OrderID = ?
	`)
	if err != nil {
		return nil, fmt.Errorf("could not prepare query: %v", err)
	}
	defer query2.Close()
	rowsProduct, err := query2.Query(orderID)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}
	defer rowsProduct.Close()
	var products []entity.OrderProduct
	for rowsProduct.Next() {
		var product entity.OrderProduct
		err := rowsProduct.Scan(
			&product.ID,
			&product.Product.ID,
			&product.Product.Name,
			&product.Product.Price,
			&product.Quantity,
		)
		if err != nil {
			return nil, fmt.Errorf("could not scan row: %v", err)
		}
		products = append(products, product)
	}
	order.OrderProduct = products
	return order, nil
}
