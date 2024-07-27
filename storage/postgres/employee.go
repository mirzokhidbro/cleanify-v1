package postgres

import (
	"bw-erp/helper"
	"bw-erp/models"
	"bw-erp/storage/repo"
	"math"

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

func (stg *employeeRepo) GetDetailedData(queryParam models.ShowEmployeeRequest) (models.ShowEmployeeResponse, error) {
	var employee models.ShowEmployeeResponse
	err := stg.db.QueryRow(`select id, company_id, phone, firstname, lastname from employees where company_id=$1 and id=$2`, queryParam.CompanyID, queryParam.EmployeeID).Scan(
		&employee.ID,
		&employee.CompanyID,
		&employee.Phone,
		&employee.Firstname,
		&employee.Lastname,
	)
	if err != nil {
		return employee, err
	}

	return employee, nil
}

func (stg *employeeRepo) AddTransaction(entity models.EmployeeTransactionRequest) error {
	difference := entity.Salary - entity.ReceivedMoney

	if difference == 0 {
		_, err := stg.db.Exec(`INSERT INTO transactions(
			company_id,
			payer_id,
			payer_type,
			amount,
			receiver_id,
			receiver_type,
			payment_type,
			payment_purpose_id
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
			entity.CompanyID,
			entity.UserID,
			"users",
			entity.ReceivedMoney,
			entity.EmployeeID,
			"employees",
			"cach",
			models.PaymentPurposeGiveSalaryToWorker,
		)

		if err != nil {
			return err
		}
	} else if math.Signbit(difference) {
		stg.db.Exec(`INSERT INTO transactions(
			company_id,
			payer_id,
			payer_type,
			amount,
			receiver_id,
			receiver_type,
			payment_type,
			payment_purpose_id
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
			entity.CompanyID,
			entity.UserID,
			"users",
			entity.Salary,
			entity.EmployeeID,
			"employees",
			"cach",
			models.PaymentPurposeGiveSalaryToWorker,
		)

		stg.db.Exec(`INSERT INTO transactions(
			company_id,
			payer_id,
			payer_type,
			amount,
			receiver_id,
			receiver_type,
			payment_type,
			payment_purpose_id
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
			entity.CompanyID,
			entity.UserID,
			"users",
			math.Abs(difference),
			entity.EmployeeID,
			"employees",
			"cach",
			models.PaymentPurposeEmployeeLoan,
		)

	} else {
		stg.db.Exec(`INSERT INTO transactions(
			company_id,
			payer_id,
			payer_type,
			amount,
			receiver_id,
			receiver_type,
			payment_type,
			payment_purpose_id
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9
		)`,
			entity.CompanyID,
			entity.UserID,
			"users",
			entity.Salary,
			entity.EmployeeID,
			"employees",
			"cach",
			models.PaymentPurposeGiveSalaryToWorker,
		)

		stg.db.Exec(`INSERT INTO transactions(
			company_id,
			payer_id,
			payer_type,
			amount,
			receiver_id,
			receiver_type,
			payment_type,
			payment_purpose_id
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9
		)`,
			entity.CompanyID,
			entity.EmployeeID,
			"employees",
			difference,
			entity.UserID,
			"users",
			"cach",
			models.PaymentPurposeDebtCollectionFromTheEmployee,
		)
	}

	return nil
}
