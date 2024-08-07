package postgres

import (
	"bw-erp/models"
	"bw-erp/storage/repo"
	"encoding/json"

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

func (stg *employeeRepo) GetList(entity models.GetEmployeeListRequest) (res []models.GetEmployeeList, err error) {
	var employees []models.GetEmployeeList

	query := `with attendance as (SELECT attendance_record->>'work_schedule' AS work_schedule, cast(attendance_record->>'employee_id' as integer) as employee_id, date 
			FROM attendance, jsonb_array_elements(employees) AS attendance_record
			where company_id = $1 and "date" = $2)
			select e.id, e.company_id, e.phone, e.firstname, e.lastname, coalesce(a.work_schedule, '0') as work_schedule, a.date from employees e
			left join attendance a on e.id = a.employee_id where company_id = $1`

	order := " ORDER BY e.firstname"

	q := query + order
	var date string
	// [TODO: must to fix this sh*t]
	date = entity.Date
	if entity.Date == "" {
		date = "2006-12-12"
	}

	rows, err := stg.db.Query(q, entity.CompanyID, date)
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
			&employee.Lastname,
			&employee.WorkSchedule,
			&employee.Date)
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
	err := stg.db.QueryRow(`select id, company_id, phone, firstname, lastname, balance from employees where company_id=$1 and id=$2`, queryParam.CompanyID, queryParam.EmployeeID).Scan(
		&employee.ID,
		&employee.CompanyID,
		&employee.Phone,
		&employee.Firstname,
		&employee.Lastname,
		&employee.Balance,
	)
	if err != nil {
		return employee, err
	}

	rows, err := stg.db.Query(`select id, amount, payment_purpose_id, created_at from transactions where receiver_type = 'employees' and receiver_id = $1 order by created_at desc`, employee.ID)

	if err != nil {
		return employee, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction models.EmployeeTransactions
		if err := rows.Scan(&transaction.ID, &transaction.Amount, &transaction.Status, &transaction.CreatedAt); err != nil {
			return employee, err
		}

		employee.Transaction = append(employee.Transaction, transaction)
	}

	attendances, err := stg.db.Query(`SELECT attendance_record->>'work_schedule' AS work_schedule, date FROM attendance, jsonb_array_elements(employees) AS attendance_record where company_id = $1 and attendance_record->>'employee_id' = $2;`, queryParam.CompanyID, queryParam.EmployeeID)

	if err != nil {
		return employee, err
	}
	defer rows.Close()
	for attendances.Next() {
		var attendance models.EmployeeAttendance
		if err := attendances.Scan(&attendance.WorkSchedule, &attendance.Date); err != nil {
			return employee, err
		}

		employee.Attendance = append(employee.Attendance, attendance)
	}

	return employee, nil
}

func (stg *employeeRepo) AddTransaction(entity models.EmployeeTransactionRequest) error {
	difference := entity.Salary - entity.ReceivedMoney

	if entity.Salary != 0 {
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
			entity.Salary,
			entity.EmployeeID,
			"employees",
			"cach",
			models.PaymentPurposeSalaryOfEmployee,
		)

		if err != nil {
			return err
		}
	}

	if entity.ReceivedMoney != 0 {
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
			models.PaymentPurposeMoneyReceiverByWorker,
		)

		if err != nil {
			return err
		}
	}

	_, err := stg.db.Exec(`UPDATE "employees" SET balance = $1 where id = $2`, difference, entity.EmployeeID)
	if err != nil {
		return err
	}

	return nil
}

func (stg *employeeRepo) Attendance(entity models.AttendanceEmployeeRequest) error {

	employeesJSON, err := json.Marshal(entity.Employees)

	if err != nil {
		return err
	}

	_, err = stg.db.Exec(`INSERT INTO attendance(
		company_id,
		date,
		employees
	) VALUES (
		$1,
		$2,
		$3
	)`,
		entity.CompanyID,
		entity.Date,
		employeesJSON,
	)

	if err != nil {
		return err
	}

	return nil
}
