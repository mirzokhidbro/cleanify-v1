CREATE TABLE user_companies (
    id BIGSERIAL PRIMARY KEY,
    permission_ids INTEGER[] NULL,
    company_id VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    is_courier BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL
);


ALTER TABLE users ADD COLUMN is_active BOOLEAN DEFAULT TRUE;
ALTER TABLE users DROP COLUMN permission_ids;


ALTER TABLE companies DROP COLUMN owner_id;
