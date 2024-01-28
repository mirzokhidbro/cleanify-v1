CREATE TABLE IF NOT EXISTS "telegram_sessions" (
    "id" SERIAL PRIMARY KEY,
    "bot_id" BIGINT REFERENCES "company_bots"("bot_id") NOT NULL,
    "order_id" INT REFERENCES "orders"("id") NOT NULL,
	"chat_id" BIGINT NOT NULL
);