package models

type CreateBotUserModel struct {
	BotID      int    `json:"bot_id"`
	ChatID     int    `json:"chat_id"`
	Page       string `json:"page"`
	DialogStep string `json:"dialog_step"`
	Role       string `json:"role"`
}

type BotUser struct {
	BotID      int     `json:"bot_id"`
	UserID     *string `json:"user_id"`
	Status     *string `json:"status"`
	Page       *string `json:"page"`
	DialogStep *string `json:"dialog_step"`
	ChatID     int64   `json:"chat_id"`
	BotToken   string
}

type SelectedUser struct {
	CompanyID   string  `json:"company_id"`
	Phone       string  `json:"phone"`
	UserID      *string `json:"user_id"`
	CompanyName string  `json:"company_name"`
}

type BotUserByCompany struct {
	CompanyID string `json:"company_id"`
	BotID     int64  `json:"bot_id"`
	ChatID    int64  `json:"chat_id"`
}
