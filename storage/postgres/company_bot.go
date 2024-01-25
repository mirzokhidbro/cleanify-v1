package postgres

import (
	"bw-erp/models"
	"errors"
)

func (stg *Postgres) CreateCompanyBotModel(id string, entity models.CreateCompanyBotModel) error {
	_, err := stg.GetCompanyById(entity.CompanyID)
	if err != nil {
		return errors.New("company not found")
	}

	_, err = stg.db.Exec(`INSERT INTO company_bots(
		id,
		company_id,
		bot_token,
		type,
		bot_id
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5
	)`,
		id,
		entity.CompanyID,
		entity.BotToken,
		"order",
		entity.BotID,
	)

	if err != nil {
		return err
	}
	return nil
}

func (stg *Postgres) GetTelegramBotByCompany(companyID string) (models.CompanyTelegramBot, error) {
	var bot models.CompanyTelegramBot
	err := stg.db.QueryRow(`select id, bot_token, company_id from company_bots where company_id = $1`, companyID).Scan(
		&bot.ID,
		&bot.BotToken,
		&bot.CompanyID,
	)
	if err != nil {
		return bot, err
	}

	return bot, nil
}
