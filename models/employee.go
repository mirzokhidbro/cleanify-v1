package models

import "time"

type Employee struct {
	CompanyID string `json:"company_id"`
	Phone     string `json:"phone"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type CreateEmployeeRequest struct {
	CompanyID string `json:"company_id" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
}

type ShowEmployeeRequest struct {
	CompanyID  string `json:"company_id" form:"company_id" binding:"required"`
	EmployeeID int    `json:"id" form:"id" binding:"required"`
}

type GetEmployeeList struct {
	CompanyID string `json:"company_id"`
	Phone     string `json:"phone"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	ID        int    `json:"id"`
}

type GetEmployeeListRequest struct {
	CompanyID string `json:"company_id" form:"company_id" binding:"required"`
}

type ShowEmployeeResponse struct {
	ID          int                    `json:"id"`
	CompanyID   string                 `json:"company_id"`
	Phone       string                 `json:"phone"`
	Firstname   string                 `json:"firstname"`
	Lastname    string                 `json:"lastname"`
	Balance     float64                `json:"balance"`
	Transaction []EmployeeTransactions `json:"transactions"`
	Attendance  []EmployeeAttendance   `json:"employee_attendance"`
}

type EmployeeTransactions struct {
	Amount    float64   `json:"amount"`
	Status    float64   `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type EmployeeTransactionRequest struct {
	Salary        float64 `json:"salary"`
	ReceivedMoney float64 `json:"received_money" binding:"required"`
	EmployeeID    int     `json:"employee_id" binding:"required"`
	CompanyID     string  `json:"company_id" binding:"required"`
	UserID        string
}

type AttendanceEmployeeRequest struct {
	CompanyID string `json:"company_id" binding:"required"`
	Date      string `json:"date" binding:"required"`
	Employees []struct {
		WorkSchedule int8 `json:"work_schedule" binding:"required"`
		EmployeeID   int  `json:"employee_id" binding:"required"`
	} `json:"employees"`
}

type EmployeeAttendance struct {
	Date         time.Time `json:"date"`
	WorkSchedule string    `json:"work_schedule"`
}
