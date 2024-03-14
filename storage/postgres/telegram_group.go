package postgres

import "bw-erp/models"

func (stg *Postgres) CreateTelegramGroupModel(entity models.CreateTelegramGroupRequest) error {
	_, err := stg.db.Exec(`INSERT INTO telegram_groups(
		chat_id,
		code,
		name
	) VALUES (
		$1,
		$2,
		$3
	)`,
		entity.ChatID,
		entity.Code,
		entity.Name,
	)

	if err != nil {
		return err
	}
	return nil
}
