CREATE TABLE IF NOT EXISTS "bot_users" (
    "id" SERIAL PRIMARY KEY,
    "chat_id" VARCHAR UNIQUE,
    "status" VARCHAR,
    "type" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);
