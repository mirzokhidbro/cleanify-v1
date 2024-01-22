package postgres

import (
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
