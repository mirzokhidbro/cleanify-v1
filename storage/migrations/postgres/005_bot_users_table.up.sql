CREATE TABLE IF NOT EXISTS "bot_users" (
    "id" SERIAL PRIMARY KEY,
    "user_id" UUID REFERENCES "users"("id") NULL,
    "chat_id" VARCHAR UNIQUE,
    "status" VARCHAR,
    "role" VARCHAR,
    "page" VARCHAR,
    "dialog_step" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

ALTER TABLE "bot_users"
ADD "bot_id" BIGINT REFERENCES "company_bots"("bot_id") NOT NULL;