package postgres

import (
	"bw-erp/models"
	"bw-erp/storage/repo"
	"errors"

	"github.com/jmoiron/sqlx"
)

type telegramBotRepo struct {
	db *sqlx.DB
}

func NewTelegramBotRepo(db *sqlx.DB) repo.TelegramBotI {
	return &telegramBotRepo{db: db}
}

func (stg *telegramBotRepo) Create(id string, entity models.CreateCompanyBotModel) error {
	_, err := stg.GetCompanyById(entity.CompanyID)
	if err != nil {
		return errors.New("company not found")
	}

	_, err = stg.db.Exec(`INSERT INTO telegram_bots(
		id,
		company_id,
		bot_token,
		type,
		bot_id,
		username,
		firstname,
		lastname
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8
	)`,
		id,
		entity.CompanyID,
		entity.BotToken,
		"order",
		entity.BotID,
		entity.Username,
		entity.Firstname,
		entity.Lastname,
	)

	if err != nil {
		return err
	}
	return nil
}

func (stg *telegramBotRepo) GetByCompany(companyID string) (models.CompanyTelegramBot, error) {
	var bot models.CompanyTelegramBot
	err := stg.db.QueryRow(`select id, bot_token, company_id from telegram_bots where company_id = $1`, companyID).Scan(
		&bot.ID,
		&bot.BotToken,
		&bot.CompanyID,
	)
	if err != nil {
		return bot, err
	}

	return bot, nil
}

func (stg *telegramBotRepo) GetOrderBot() ([]models.CompanyTelegramBot, error) {
	rows, err := stg.db.Query(`select id, bot_token from telegram_bots where type = 'order'`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bots []models.CompanyTelegramBot
	for rows.Next() {
		var bot models.CompanyTelegramBot
		err = rows.Scan(&bot.ID, &bot.BotToken)
		if err != nil {
			return nil, err
		}
		bots = append(bots, bot)

	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bots, nil
}

func (stg *telegramBotRepo) GetCompanyIDByBot(botID int64) (models.CompanyTelegramBot, error) {
	var bot models.CompanyTelegramBot
	err := stg.db.QueryRow(`select id, bot_token, company_id from telegram_bots where company_id = $1`, botID).Scan(
		&bot.ID,
		&bot.BotToken,
		&bot.CompanyID,
	)
	if err != nil {
		return bot, err
	}

	return bot, nil
}

func (stg *telegramBotRepo) GetCompanyById(id string) (models.Company, error) {
	var company models.Company
	err := stg.db.QueryRow(`select id, name from companies where id = $1`, id).Scan(
		&company.ID,
		&company.Name,
	)
	if err != nil {
		return company, err
	}

	return company, nil
}
