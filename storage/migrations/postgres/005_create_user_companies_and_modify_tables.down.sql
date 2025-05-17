ALTER TABLE companies ADD COLUMN owner_id VARCHAR(255) NULL;


ALTER TABLE users ADD COLUMN permission_ids JSONB NULL;
ALTER TABLE users DROP COLUMN is_active;


DROP TABLE user_companies;
