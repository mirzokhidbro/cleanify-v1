package postgres

import (
	"bw-erp/models"
	"errors"
)

func (stg *Postgres) CreateOrderModel(entity models.CreateOrderModel) error {
	_, err := stg.GetCompanyById(entity.CompanyID)
	if err != nil {
		return errors.New("Company not found")
	}

	_, err = stg.db.Exec(`INSERT INTO orders(
		company_id,
		phone,
		count,
		slug,
		description
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5
	)`,
		entity.CompanyID,
		entity.Phone,
		entity.Count,
		entity.Slug,
		entity.Description,
	)

	if err != nil {
		return err
	}

	return nil
}

func (stg *Postgres) GetOrdersList(companyID string) ([]models.OrderList, error) {
	rows, err := stg.db.Query(`select id, slug, status from orders where company_id = $1`, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.OrderList
	for rows.Next() {
		var order models.OrderList
		err = rows.Scan(
			&order.ID,
			&order.Slug,
			&order.Status)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
