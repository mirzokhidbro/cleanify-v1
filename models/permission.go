package models

type Permission struct {
	ID    string `json:"id"`
	Slug  string `json:"slug"`
	Name  string `json:"name"`
	Scope string `json:"scope"`
}

type RoleAndPermission struct {
	RoleID        string
	PermissionIDs string
}
