package models

import (
	"database/sql"
	"encoding/json"
)

type CreateUserModel struct {
	Fullname string `json:"fullname" binding:"required" minLength:"2" maxLength:"255" example:"John Doe"`
	Phone    string `json:"phone" binding:"required" example:"991234567"`
	// Password  string `json:"password" binding:"required"`
	CompanyID string `json:"company_id" binding:"required"`
	// PermissionIDs        []string `json:"permission_ids" binding:"required"`
	// ConfirmationPassword string `json:"confirmation_password" binding:"required"`
	Permissions []struct {
		CompanyID     string   `json:"company_id"`
		PermissionIDs []string `json:"permission_ids"`
	} `json:"permissions"`
}

type AuthUserModel struct {
	ID       string `json:"id"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID                      string                    `json:"id"`
	Fullname                string                    `json:"fullname"`
	Phone                   string                    `json:"phone"`
	CompanyID               *string                   `json:"company_id"`
	UserPermissionByCompany []UserPermissionByCompany `json:"user_permissions_by_company"`
}

type GetCouriesResponse struct {
	ID       string `json:"id"`
	Fullname string `json:"fullname"`
}

type GetCouriesListRequest struct {
	CompanyID string `json:"company_id" form:"company_id" binding:"required"`
}

type UserPermissionByCompany struct {
	CompanyID     string   `json:"company_id"`
	CompanyName   string   `json:"company_name"`
	PermissionIDs []string `json:"permission_ids"`
	Can           string   `json:"can"`
	IsCourier     bool     `json:"is_courier"`
}

type UpdateUserRequest struct {
	ID string `json:"id" binding:"required"`
	// CompanyID   string `json:"company_id" binding:"required"`
	Fullname    string `json:"fullname"`
	Permissions []struct {
		CompanyID     string   `json:"company_id"`
		PermissionIDs []string `json:"permission_ids"`
	} `json:"permissions"`
}

type GetUserListRequest struct {
	CompanyID string `json:"company_id" form:"company_id" binding:"required"`
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
