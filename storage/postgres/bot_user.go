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
		bot_id,
		firstname,
		lastname,
		username
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7
	)`,
		entity.ChatID,
		entity.Page,
		entity.DialogStep,
		entity.BotID,
		entity.Firstname,
		entity.Lastname,
		entity.Username,
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
	if entity.Firstname != "" {
		query += `firstname = :firstname,`
	}
	if entity.Lastname != "" {
		query += `lastname = :lastname,`
	}
	if entity.Username != "" {
		query += `username = :username,`
	}

	if query[len(query)-1:] == "," {
		query = query[:len(query)-1]
	}

	query += ` WHERE bot_id = :bot_id and chat_id = :chat_id`

	params := map[string]interface{}{
		"bot_id":      entity.BotID,
		"status":      entity.Status,
		"page":        entity.Page,
		"dialog_step": entity.DialogStep,
		"user_id":     entity.UserID,
		"chat_id":     entity.ChatID,
		"firstname":   entity.Firstname,
		"lastname":    entity.Lastname,
		"username":    entity.Username,
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

func (stg *Postgres) GetBotUserByChatIDModel(ChatID int64, BotID int64) (models.BotUser, error) {
	var botUser models.BotUser

	err := stg.db.QueryRow(`with users as (
		select c.id as company_id, c.name as company_name, u.phone, u.id from users u
		inner join roles cr on cr.id = u.role_id
		inner join companies c on c.id = cr.company_id
	)
	select bot_id, user_id, status, page, dialog_step, u.company_id from bot_users bu 
	inner join users u on (bu.user_id = u.id or bu.user_id is null) where chat_id::bigint = $1 and bot_id = $2`, ChatID, BotID).Scan(
		&botUser.BotID,
		&botUser.UserID,
		&botUser.Status,
		&botUser.Page,
		&botUser.DialogStep,
		&botUser.CompanyID,
	)
	if err != nil {
		return botUser, errors.New("Bu botdan foydalanish uchun avtorizatsiyadan o'tish kerak!")
	}

	return botUser, nil
}

func (stg *Postgres) GetBotUserByUserID(UserID string) (models.BotUser, error) {
	var botUser models.BotUser
	err := stg.db.QueryRow(`select tb.bot_id, user_id, status, page, dialog_step, chat_id, tb.bot_token from bot_users bu inner join telegram_bots tb on tb.bot_id = bu.bot_id where user_id = $1`, UserID).Scan(
		&botUser.BotID,
		&botUser.UserID,
		&botUser.Status,
		&botUser.Page,
		&botUser.DialogStep,
		&botUser.ChatID,
		&botUser.BotToken,
	)
	if err != nil {
		return botUser, err
	}

	return botUser, nil
}

func (stg *Postgres) GetSelectedUser(BotID int64, Phone string) (models.SelectedUser, error) {
	var user models.SelectedUser

	query := `select c.id as company_id, c.name as company_name, u.phone, u.id from users u
	inner join roles cr on cr.id = u.role_id
	inner join companies c on c.id = cr.company_id
	where u.phone = $1 `

	err := stg.db.QueryRow(query, Phone).Scan(
		&user.CompanyID,
		&user.CompanyName,
		&user.Phone,
		&user.UserID,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (stg *Postgres) GetBotUserByCompany(BotID int64, ChatID int64) (botUser models.BotUserByCompany, err error) {
	var user models.BotUserByCompany

	query := `select cb.company_id, bu.bot_id, bu.chat_id from bot_users bu inner join telegram_bots cb on bu.bot_id = cb.bot_id where bu.bot_id = $1 and bu.chat_id = $2`

	err = stg.db.QueryRow(query, BotID, ChatID).Scan(
		&user.CompanyID,
		&user.BotID,
		&user.ChatID,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}
