package models

type CreateUserModel struct {
	Firstname string `json:"firstname" binding:"required" minLength:"2" maxLength:"255" example:"John"`
	Lastname  string `json:"lastname" binding:"required" minLength:"2" maxLength:"255" example:"Doe"`
	Phone     string `json:"phone" binding:"required" example:"991234567"`
	Password  string `json:"password" binding:"required"`
}
