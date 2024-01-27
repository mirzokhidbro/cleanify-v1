package postgres

import (
	"bw-erp/helper"
	"bw-erp/models"
	"errors"
)

func (stg *Postgres) CreateBotUserModel(entity models.CreateBotUserModel) error {
	_, err := stg.db.Exec(`INSERT INTO bot_users(
		chat_id,
		page,
		dialog_step,
		bot_id
	) VALUES (
		$1,
		$2,
		$3,
		$4
	)`,
		entity.ChatID,
		entity.Page,
		entity.DialogStep,
		entity.BotID,
	)

	if err != nil {
		return err
	}
	return nil
}

func (stg *Postgres) UpdateBotUserModel(entity models.BotUser) (rowsAffected int64, err error) {
	query := `UPDATE "bot_users" SET `

	if entity.UserID != nil {
		query += `user_id = :user_id,`
	}
	if entity.Status != nil {
		query += `status = :status,`
	}
	if entity.Page != nil {
		query += `page = :page,`
	}
	if entity.DialogStep != nil {
		query += `dialog_step = :dialog_step,`
	}

	if query[len(query)-1:] == "," {
		query = query[:len(query)-1]
	}

	query += ` WHERE bot_id = :bot_id`

	params := map[string]interface{}{
		"bot_id":      entity.BotID,
		"status":      entity.Status,
		"page":        entity.Page,
		"dialog_step": entity.DialogStep,
		"user_id":     entity.UserID,
	}

	query, arr := helper.ReplaceQueryParams(query, params)
	result, err := stg.db.Exec(query, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (stg *Postgres) GetBotUserByChatIDModel(ChatID int64) (models.BotUser, error) {
	var botUser models.BotUser
	err := stg.db.QueryRow(`select bot_id, user_id, status, page, dialog_step from bot_users where chat_id::bigint = $1`, ChatID).Scan(
		&botUser.BotID,
		&botUser.UserID,
		&botUser.Status,
		&botUser.Page,
		&botUser.DialogStep,
	)
	if err != nil {
		return botUser, err
	}

	return botUser, nil
}

func (stg *Postgres) GetSelectedUser(BotID int64, Phone string) (models.SelectedUser, error) {
	var user models.SelectedUser

	query := `with users as (
		select c.id as company_id, u.phone, u.id from users u
		inner join company_roles cr on cr.id = u.role_id
		inner join companies c on c.id = cr.company_id
	)
	select users.company_id, users.phone, users.id from users
	inner join company_bots cb on cb.company_id = users.company_id
	where users.phone = $1 and cb.bot_id = $2`

	err := stg.db.QueryRow(query, Phone, BotID).Scan(
		&user.CompanyID,
		&user.Phone,
		&user.UserID,
	)
	if err != nil {
		return user, errors.New("Tizimda bunday foydalanuvchi mavjud emas!")
	}

	return user, nil
}
