package models

type CreateCompanyBotModel struct {
	BotId     string `json:"bot_token", binding:"required"`
	CompanyID string `json:"company_id", binding:"required"`
}

type CompanyTelegramBot struct {
	ID        string `jsin:"id"`
	BotToken  string `json:"bot_token"`
	CompanyID string `json:"company_id"`
}
