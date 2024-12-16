package postgres

import (
	"bw-erp/models"
	"bw-erp/storage/repo"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type webPushRepo struct {
	db *sqlx.DB
}

func NewWebPushRepo(db *sqlx.DB) repo.WebPushStorageI {
	return &webPushRepo{
		db: db,
	}
}

func (r *webPushRepo) CreatePushSubscription(subscription models.CreatePushSubscriptionRequest) (int64, error) {
	query := `
		INSERT INTO push_subscriptions (
			user_id,
			endpoint,
			auth_key,
			p256dh_key
		) VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var id int64
	err := r.db.QueryRow(
		query,
		subscription.UserID,
		subscription.Endpoint,
		subscription.AuthKey,
		subscription.P256dhKey,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *webPushRepo) GetPushSubscription(userID string) (*models.PushSubscription, error) {
	var subscription models.PushSubscription

	query := `
		SELECT 
			id,
			user_id,
			endpoint,
			auth_key,
			p256dh_key,
			created_at
		FROM push_subscriptions
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT 1
	`

	err := r.db.QueryRow(query, userID).Scan(
		&subscription.ID,
		&subscription.UserID,
		&subscription.Endpoint,
		&subscription.AuthKey,
		&subscription.P256dhKey,
		&subscription.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &subscription, nil
}

func (r *webPushRepo) DeletePushSubscription(userID string) error {
	query := `DELETE FROM push_subscriptions WHERE user_id = $1`
	result, err := r.db.Exec(query, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("subscription not found for user_id: %s", userID)
	}

	return nil
}

func (r *webPushRepo) GetAllPushSubscriptions(params *models.GetPushSubscriptionResponse) (*models.GetPushSubscriptionResponse, error) {
	result := &models.GetPushSubscriptionResponse{
		Subscriptions: make([]models.PushSubscription, 0),
	}

	query := `
		SELECT 
			id,
			user_id,
			endpoint,
			auth_key,
			p256dh_key,
			created_at
		FROM push_subscriptions
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var subscription models.PushSubscription
		err := rows.Scan(
			&subscription.ID,
			&subscription.UserID,
			&subscription.Endpoint,
			&subscription.AuthKey,
			&subscription.P256dhKey,
			&subscription.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		result.Subscriptions = append(result.Subscriptions, subscription)
	}

	countQuery := `SELECT COUNT(*) FROM push_subscriptions`
	err = r.db.QueryRow(countQuery).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return result, nil
}
