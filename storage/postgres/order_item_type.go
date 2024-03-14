package postgres

import (
	"bw-erp/helper"
	"bw-erp/models"
	"errors"
)

func (stg *Postgres) CreateOrderItemTypeModel(id string, entity models.OrderItemTypeModel) error {
	_, err := stg.GetCompanyById(entity.CopmanyID)
	if err != nil {
		return errors.New("selected company not found")
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
					o.id,
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
			&orderItemType.ID,
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

func (stg *Postgres) GetOrderItemTypeById(orderItemTypeId string) (models.OrderItemByCompany, error) {
	var orderItemType models.OrderItemByCompany
	err := stg.db.QueryRow(`select o.id, o.name, o.price, c.name, c.id  from order_item_types o inner join companies c on c.id = o.company_id where o.id = $1`, orderItemTypeId).Scan(
		&orderItemType.ID,
		&orderItemType.Name,
		&orderItemType.Price,
		&orderItemType.CopmanyName,
		&orderItemType.CopmanyID,
	)
	if err != nil {
		return orderItemType, err
	}

	return orderItemType, nil
}

func (stg *Postgres) UpdateOrderItemTypeModel(entity models.EditOrderItemTypeRequest) (rowsAffected int64, err error) {
	query := `UPDATE 
					"order_item_types" 
						SET price = :price,updated_at = now()
					WHERE
		  				id = :id and company_id = :company_id`

	params := map[string]interface{}{
		"id":         entity.ID,
		"price":      entity.Price,
		"company_id": entity.CopmanyID,
	}

	query, arr := helper.ReplaceQueryParams(query, params)
	result, err := stg.db.Exec(query, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
