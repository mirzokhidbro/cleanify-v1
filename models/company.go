package models

type CreateCompanyModel struct {
	Name    string `json:"name" binding:"required" minLength:"2" maxLength:"255"`
	OwnerId int64  `json:"owner_id" binding:"required"`
}

type Company struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
