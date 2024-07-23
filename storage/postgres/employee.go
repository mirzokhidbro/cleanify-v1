package postgres

import (
	"bw-erp/helper"
	"bw-erp/models"
	"bw-erp/storage/repo"

	"github.com/jmoiron/sqlx"
)

type employeeRepo struct {
	db *sqlx.DB
}

func NewEmployeeRepo(db *sqlx.DB) repo.EmployeeI {
	return &employeeRepo{db: db}
}

func (stg employeeRepo) Create(entity models.CreateEmployeeRequest) error {
	_, err := stg.db.Exec(`INSERT INTO employees(
		company_id,
		phone,
		firstname,
		lastname
	) VALUES (
		$1,
		$2,
		$3,
		$4
	)`,
		entity.CompanyID,
		entity.Phone,
		entity.Firstname,
		entity.Lastname,
	)

	if err != nil {
		return err
	}

	return nil
}

func (stg *employeeRepo) GetList(companyID string) (res []models.GetEmployeeList, err error) {
	var employees []models.GetEmployeeList
	var arr []interface{}
	params := make(map[string]interface{})

	query := "select id, company_id, phone, firstname, lastname from employees"
	filter := " WHERE true"
	order := " ORDER BY firstname"

	params["company_id"] = companyID
	filter += " AND (company_id = :company_id)"

	q := query + filter + order

	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := stg.db.Query(q, arr...)
	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		var employee models.GetEmployeeList
		err = rows.Scan(
			&employee.ID,
			&employee.CompanyID,
			&employee.Phone,
			&employee.Firstname,
			&employee.Lastname)
		if err != nil {
			return res, err
		}
		employees = append(employees, employee)
	}

	if err = rows.Err(); err != nil {
		return res, err
	}

	return employees, nil
}
