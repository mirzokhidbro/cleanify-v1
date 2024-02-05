package models

import (
	"database/sql"
	"encoding/json"
)

type CreateUserModel struct {
	Firstname            string `json:"firstname" binding:"required" minLength:"2" maxLength:"255" example:"John"`
	Lastname             string `json:"lastname" binding:"required" minLength:"2" maxLength:"255" example:"Doe"`
	RoleID               string `json:"role_id" binding:"required"`
	Phone                string `json:"phone" binding:"required" example:"991234567"`
	Password             string `json:"password" binding:"required"`
	ConfirmationPassword string `json:"confirmation_password" binding:"required"`
}

type AuthUserModel struct {
	ID       string `json:"id"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID          string  `json:"id"`
	Firstname   string  `json:"firstname"`
	Lastname    string  `json:"lastname"`
	Phone       string  `json:"phone"`
	Role        *string `json:"role"`
	Company     *string `json:"company"`
	CompanyID   *string `json:"company_id"`
	RoleID      *string `json:"role_id"`
	Permissions string  `json:"permissions"`
}

type NullString struct {
	sql.NullString
}

func (ns NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return json.Marshal(nil)
}
