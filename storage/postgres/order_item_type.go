package postgres

import (
	"bw-erp/helper"
	"bw-erp/models"
	"bw-erp/storage/repo"

	"github.com/jmoiron/sqlx"
)

type orderItemTypeRepo struct {
	db *sqlx.DB
}

func NewOrderItemTypeRepo(db *sqlx.DB) repo.OrderItemTypeI {
	return &orderItemTypeRepo{db: db}
}

func (stg *orderItemTypeRepo) Create(id string, entity models.OrderItemTypeModel) error {
	// [TODO: get by primary key metodini yozish kerak]
	// _, err := stg.Company().GetById(entity.CopmanyID)
	// if err != nil {
	// 	return errors.New("selected company not found")
	// }

	_, err := stg.db.Exec(`INSERT INTO order_item_types(
		id,
		name,
		company_id,
		price,
		is_countable
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5
	)`,
		id,
		entity.Name,
		entity.CompanyID,
		entity.Price,
		entity.IsCountable,
	)

	if err != nil {
		return err
	}
	return nil
}

func (stg *orderItemTypeRepo) GetByCompany(CompanyID string) ([]models.OrderItemByCompany, error) {
	rows, err := stg.db.Query(`select  
					o.id,
					o.name, 
					o.price,
					o.is_countable,
					c.name,
					c.id
					from order_item_types o inner join companies c on c.id = o.company_id 
					where o.company_id = $1 order by o.price`, CompanyID)
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
			&orderItemType.IsCountable,
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

func (stg *orderItemTypeRepo) GetById(orderItemTypeId string) (models.OrderItemByCompany, error) {
	var orderItemType models.OrderItemByCompany
	err := stg.db.QueryRow(`select o.id, o.name, o.price, o.is_countable, c.name, c.id  from order_item_types o inner join companies c on c.id = o.company_id where o.id = $1`, orderItemTypeId).Scan(
		&orderItemType.ID,
		&orderItemType.Name,
		&orderItemType.Price,
		&orderItemType.IsCountable,
		&orderItemType.CopmanyName,
		&orderItemType.CopmanyID,
	)
	if err != nil {
		return orderItemType, err
	}

	return orderItemType, nil
}

func (stg *orderItemTypeRepo) Update(entity models.EditOrderItemTypeRequest) (rowsAffected int64, err error) {

	query := `UPDATE 
					"order_item_types" 
						SET price = :price, is_countable = :is_countable, name = :name updated_at = now()
					WHERE
		  				id = :id and company_id = :company_id`

	params := map[string]interface{}{
		"id":           entity.ID,
		"price":        entity.Price,
		"company_id":   entity.CopmanyID,
		"is_countable": entity.IsCountable,
		"name":         entity.Name,
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
