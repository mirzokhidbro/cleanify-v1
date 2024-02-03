package postgres

import (
	"bw-erp/helper"
	"bw-erp/models"
	"errors"
	"fmt"
)

func (stg Postgres) CreateOrderItemModel(entity models.CreateOrderItemModel) error {

	_, err := stg.GetOrderByPrimaryKey(entity.OrderID)
	if err != nil {
		return errors.New("order not found")
	}

	orderItemType, err := stg.GetOrderItemTypeById(entity.OrderItemTypeID)
	if err != nil {
		fmt.Print(err.Error())
		return errors.New("Order item type not found")
	}

	itemType := orderItemType.Name
	price := orderItemType.Price

	_, err = stg.db.Exec(`INSERT INTO order_items(
		order_id,
		type,
		price,
		width,
		height,
		description
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6
	)`,
		entity.OrderID,
		itemType,
		price,
		entity.Width,
		entity.Height,
		entity.Description,
	)

	if err != nil {
		return err
	}

	return nil
}

func (stg *Postgres) UpdateOrderItemModel(entity models.UpdateOrderItemRequest) (rowsAffected int64, err error) {
	query := `UPDATE "order_items" SET `

	if entity.Price != 0 {
		query += `price = :price,`
	}
	if entity.Width != 0 {
		query += `width = :width,`
	}
	if entity.Height != 0 {
		query += `height = :height,`
	}
	if entity.Description != "" {
		query += `description = :description,`
	}
	if entity.Type != "" {
		query += `type = :type,`
	}

	query += `updated_at = now()
			  WHERE
					id = :id`

	params := map[string]interface{}{
		"id":          entity.ID,
		"price":       entity.Price,
		"width":       entity.Width,
		"height":      entity.Height,
		"description": entity.Description,
		"type":        entity.Type,
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
