package postgres

import (
	"bw-erp/helper"
	"bw-erp/models"
	"bw-erp/storage/repo"

	"github.com/jmoiron/sqlx"
)

type orderItemRepo struct {
	db *sqlx.DB
}

func NewOrderItemRepo(db *sqlx.DB) repo.OrderItemI {
	return &orderItemRepo{db: db}
}

func (stg orderItemRepo) Create(entity models.CreateOrderItemModel) error {

	_, err := stg.db.Exec(`INSERT INTO order_items(
		order_id,
		type,
		price,
		width,
		height,
		description,
		is_countable,
		order_item_type_id
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8
	)`,
		entity.OrderID,
		entity.ItemType,
		entity.Price,
		entity.Width,
		entity.Height,
		entity.Description,
		entity.IsCountable,
		entity.OrderItemTypeID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (stg orderItemRepo) Update(entity models.UpdateOrderItemRequest) (rowsAffected int64, err error) {
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
	// if entity.ItemType != "" {
	// 	query += `type = :type,`
	// }

	query += `type = :type, is_countable = :is_countable, updated_at = now()
			  WHERE
					id = :id`

	params := map[string]interface{}{
		"id":           entity.ID,
		"price":        entity.Price,
		"width":        entity.Width,
		"height":       entity.Height,
		"description":  entity.Description,
		"type":         entity.ItemType,
		"is_countable": entity.IsCountable,
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

func (stg *orderItemRepo) DeleteByID(ID int) error {
	_, err := stg.db.Exec(`delete from order_items where id = $1`, ID)
	if err != nil {
		return err
	}

	return nil
}
