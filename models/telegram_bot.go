package models

type CreateCompanyBotModel struct {
	BotToken  string `json:"bot_token" binding:"required"`
	CompanyID string `json:"company_id" binding:"required"`
	BotID     int    `json:"bot_id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
}

type CompanyTelegramBot struct {
	ID        string `jsin:"id"`
	BotToken  string `json:"bot_token"`
	CompanyID string `json:"company_id"`
}
