package models

type CreateBotUserModel struct {
	BotID      int    `json:"bot_id"`
	ChatID     int    `json:"chat_id"`
	Page       string `json:"page"`
	DialogStep string `json:"dialog_step"`
	Role       string `json:"role"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Username   string `json:"username"`
}

type BotUser struct {
	BotID      int     `json:"bot_id"`
	UserID     *int64  `json:"user_id"`
	Status     *string `json:"status"`
	Page       *string `json:"page"`
	DialogStep *string `json:"dialog_step"`
	ChatID     int64   `json:"chat_id"`
	BotToken   string
	CompanyID  string
	// [TODO: hamma userlarni datalari yozilganidan keyin bu fieldlarni olib tashlaymiz.]
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
}

type SelectedUser struct {
	CompanyID   string `json:"company_id"`
	Phone       string `json:"phone"`
	UserID      *int64 `json:"user_id"`
	CompanyName string `json:"company_name"`
}

type BotUserByCompany struct {
	CompanyID string `json:"company_id"`
	BotID     int64  `json:"bot_id"`
	ChatID    int64  `json:"chat_id"`
}

type GetNotificationGroup struct {
	CompanyID string
	Role      string
}
