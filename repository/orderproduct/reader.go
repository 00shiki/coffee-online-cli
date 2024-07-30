package orderproduct

import "pair-project/entity"

func (r *Repository) popularProducts() ([]entity.OrderProduct, error) {
	// TODO: add query
	query := `
			SELECT
				p.ProductID,
				p.ProductName,
				SUM(op.Quantity) AS TotalOrdered
			FROM
				OrderProduct op
			JOIN
				Product p ON op.ProductID = p.ProductID
			GROUP BY
				p.ProductID, p.ProductName
			ORDER BY
				TotalOrdered DESC;`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orderproduct []entity.OrderProduct
	for rows.Next() {
		var odProd entity.OrderProduct
		rows.Scan(
			&odProd.Product.ID,
			&odProd.Product.Name,
			&odProd.Quantity
		)
		orderproduct = append(orderproduct, odProd)
	}
	return orderproduct, nil
}
