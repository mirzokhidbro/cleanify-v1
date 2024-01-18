package models

type CreateUserModel struct {
	Firstname string `json:"firstname" binding:"required" minLength:"2" maxLength:"255" example:"John"`
	Lastname  string `json:"lastname" binding:"required" minLength:"2" maxLength:"255" example:"Doe"`
	Phone     string `json:"phone" binding:"required" example:"991234567"`
	Password  string `json:"password" binding:"required"`
}

type AuthUserModel struct {
	ID       string `json:"id"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
}
