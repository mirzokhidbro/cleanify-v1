package postgres

import (
	"bw-erp/models"
	"errors"
)

func (stg *Postgres) CreateOrderItemTypeModel(id string, entity models.OrderItemTypeModel) error {
	_, err := stg.GetCompanyById(entity.CopmanyID)
	if err != nil {
		return errors.New("Selected company not found")
	}

	_, err = stg.db.Exec(`INSERT INTO order_item_types(
		id,
		name,
		company_id,
		price
	) VALUES (
		$1,
		$2,
		$3,
		$4
	)`,
		id,
		entity.Name,
		entity.CopmanyID,
		entity.Price,
	)

	if err != nil {
		return err
	}
	return nil
}

func (stg *Postgres) GetOrderItemTypesByCompany(CompanyID string) ([]models.OrderItemByCompany, error) {
	rows, err := stg.db.Query(`select  
					o.name, 
					o.price,
					c.name,
					c.id
					from order_item_types o inner join companies c on c.id = o.company_id 
					where o.company_id = $1`, CompanyID)
	if err != nil {
		return nil, err
	}

	var orderItemTypes []models.OrderItemByCompany
	for rows.Next() {
		var orderItemType models.OrderItemByCompany
		err = rows.Scan(
			&orderItemType.Name,
			&orderItemType.Price,
			&orderItemType.CopmanyName,
			&orderItemType.CopmanyID,
		)
		if err != nil {
			return nil, err
		}
		orderItemTypes = append(orderItemTypes, orderItemType)
	}

	return orderItemTypes, nil
}
