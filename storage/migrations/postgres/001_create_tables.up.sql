CREATE TABLE IF NOT EXISTS "users" (
    "id" UUID PRIMARY KEY,
    "phone" VARCHAR UNIQUE,
    "firstname" VARCHAR,
    "lastname" VARCHAR,
    "password" VARCHAR(1000),
    "permission_ids" VARCHAR[],
    "company_id" UUID REFERENCES "companies"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "companies" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR UNIQUE,
    "owner_id" UUID REFERENCES "users"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);


CREATE TABLE IF NOT EXISTS "telegram_bots" (
    "id" UUID PRIMARY KEY,
    "company_id" UUID REFERENCES "companies"("id") NOT NULL,
    "bot_token" VARCHAR NOT NULL,
	"bot_id" BIGINT UNIQUE,
    "type" VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS "bot_users" (
    "id" SERIAL PRIMARY KEY,
    "user_id" VARCHAR,
    "chat_id" BIGINT UNIQUE,
    "status" VARCHAR,
    "role" VARCHAR,
    "page" VARCHAR,
    "dialog_step" VARCHAR,
    "firstname" VARCHAR,
    "lastname" VARCHAR,
    "username" VARCHAR,
    "company_id" UUID REFERENCES "companies"("id"),
    "bot_id" BIGINT REFERENCES "telegram_bots"("bot_id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "orders" (
    "id" SERIAL PRIMARY KEY,
    "company_id" UUID REFERENCES "companies"("id"),
    "chat_id" BIGINT,
    "phone" VARCHAR NULL,
    "count" INTEGER,
    "status" INTEGER,
    "slug" VARCHAR,
    "latitute" FLOAT,
    "longitude" FLOAT,
    "description" VARCHAR,
    "uuid" UUID DEFAULT gen_random_uuid(),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "courier_id" UUID REFERENCES "users"("id"),
    "address" VARCHAR,
    "client_id" INTEGER REFERENCES "clients"("id"),
    "service_price" FLOAT,
    "discount_percentage" FLOAT,
    "payment_status" INTEGER,
    "discounted_price" FLOAT
);


ALTER TABLE "orders"
ALTER COLUMN "chat_id" SET DATA TYPE BIGINT
USING chat_id::BIGINT;

ALTER TABLE "orders"
ALTER COLUMN "status" SET DEFAULT 1;

ALTER TABLE "orders"
ADD COLUMN "courier_id" UUID REFERENCES "users"("id");

CREATE TABLE IF NOT EXISTS "order_items" (
    "id" SERIAL PRIMARY KEY,
    "order_id" INTEGER REFERENCES "orders"("id") NOT NULL,
    "type" VARCHAR NOT NULL,
    "price" FLOAT NOT NULL,
    "width" FLOAT NOT NULL,
    "height" FLOAT NOT NULL,
    "description" VARCHAR,
    "washed_at"  TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "is_countable" BOOLEAN DEFAULT FALSE,
    "status" INTEGER DEFAULT 1,
    "order_item_type_id" UUID REFERENCES "order_item_types"("id")
);

CREATE TABLE IF NOT EXISTS "order_item_types" (
    "id" UUID PRIMARY KEY,
    "company_id" UUID REFERENCES "companies"("id") NOT NULL,
    "name" VARCHAR NOT NULL,
    "price" FLOAT NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "is_countable" BOOLEAN DEFAULT FALSE
);

ALTER TABLE "telegram_bots"
ADD UNIQUE (bot_id);


ALTER TABLE "telegram_bots"
ADD "firstname" VARCHAR,
ADD "lastname" VARCHAR,
ADD "username" VARCHAR;

CREATE TABLE IF NOT EXISTS "telegram_sessions" (
    "id" SERIAL PRIMARY KEY,
    "bot_id" BIGINT REFERENCES "telegram_bots"("bot_id") NOT NULL,
    "order_id" INT,
    "chat_id" BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS "permissions" (
    "id" UUID PRIMARY KEY,
    "slug" VARCHAR UNIQUE,
    "name" VARCHAR UNIQUE,
    "scope" VARCHAR DEFAULT ('company'),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "group" VARCHAR
);

CREATE TABLE IF NOT EXISTS "role_and_permissions" (
    "id" SERIAL PRIMARY KEY,
    "role_id" VARCHAR UNIQUE,
    "permission_ids" VARCHAR[]
);

ALTER TABLE "orders"
ADD "address" VARCHAR;


CREATE TABLE IF NOT EXISTS "clients" (
    "id" SERIAL PRIMARY KEY,
    "company_id" UUID REFERENCES "companies"("id") NOT NULL,
    "address" VARCHAR,
    "full_name" VARCHAR,
    "phone_number" VARCHAR,
    "additional_phone_number" VARCHAR,
    "work_number" VARCHAR,
    "latitute" FLOAT,
    "longitude" FLOAT,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE EXTENSION pg_trgm;
CREATE INDEX idx_clients_phones ON "clients" USING gin (("phone_number" || ' ' || "additional_phone_number" || ' ' || "work_number") gin_trgm_ops);

ALTER TABLE "orders"
ADD "client_id" INTEGER REFERENCES "clients"("id");

CREATE TABLE IF NOT EXISTS "telegram_groups" (
    "id" SERIAL PRIMARY KEY,
    "company_id" UUID REFERENCES "companies"("id"),
    "name"  VARCHAR,
    "notification_statuses" INT[],
    "with_location" BOOLEAN,
    "code" INT,
    "chat_id" BIGINT NOT NULL UNIQUE,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

ALTER TABLE "order_items"
ADD COLUMN "is_countable" BOOLEAN DEFAULT FALSE;

ALTER TABLE "order_items"
ADD COLUMN "status" INTEGER DEFAULT 1;

ALTER TABLE "order_item_types"
ADD COLUMN "is_countable" BOOLEAN DEFAULT FALSE;

CREATE TABLE IF NOT EXISTS "order_statuses" (
    "id" SERIAL PRIMARY KEY,
    "company_id" UUID REFERENCES "companies"("id"),
    "name"  VARCHAR,
    "color" VARCHAR,
    "number" INTEGER,
    "description" TEXT,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "slug" VARCHAR
);


ALTER TABLE "order_items"
ADD COLUMN "order_item_type_id" UUID REFERENCES "order_item_types"("id");

CREATE TABLE IF NOT EXISTS "status_change_histories" (
    "id"                SERIAL PRIMARY KEY,
    "historyable_type"  VARCHAR,
    "historyable_id"    VARCHAR,
    "status"            INTEGER,
    "user_id"           UUID REFERENCES "users"("id"),
    "created_at"        TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at"        TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);


CREATE TABLE IF NOT EXISTS "user_permissions" (
    "id"             SERIAL PRIMARY KEY,
    "permission_ids" VARCHAR[],
    "company_id"     UUID REFERENCES "companies"("id") NOT NULL,
    "user_id"        UUID REFERENCES "users"("id") NOT NULL,
    "created_at"     TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at"     TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "is_courier" BOOLEAN DEFAULT FALSE
);

ALTER TABLE "orders"
ADD "service_price" FLOAT,
ADD "discount_percentage" FLOAT,
ADD "payment_status" INTEGER,
ADD "discounted_price" FLOAT;



CREATE TABLE IF NOT EXISTS "transactions" (
    "id" SERIAL PRIMARY KEY,
    "company_id" UUID REFERENCES "companies"("id") NOT NULL,
    "receiver_id" VARCHAR,
    "receiver_type" VARCHAR,
    "payment_purpose_id" INTEGER,
    "amount" FLOAT,
    "payer_id" VARCHAR,
    "payer_type" VARCHAR,
    "description" TEXT,
    "payment_type" TEXT,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "payment_purposes" (
    "id" SERIAL PRIMARY KEY,
    "type" VARCHAR,
    "name" VARCHAR
)

CREATE TABLE IF NOT EXISTS "employees"(
    "id" SERIAL PRIMARY KEY,
    "company_id" UUID REFERENCES "companies"("id") NOT NULL,
    "phone" VARCHAR,
    "firstname" VARCHAR,
    "lastname" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "balance" FLOAT DEFAULT 0
)

ALTER TABLE "order_statuses"
ADD "slug" VARCHAR;

CREATE TABLE IF NOT EXISTS "attendance"(
    "id" SERIAL PRIMARY KEY,
    "company_id" UUID REFERENCES "companies"("id") NOT NULL,
    "date" DATE,
    "employees" JSONB,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
)


CREATE TABLE IF NOT EXISTS "notification_settings"(
    "id" SERIAL PRIMARY KEY,
    "company_id" UUID REFERENCES "companies"("id") NOT NULL,
    "status" INT,
    "user_ids" VARCHAR[],
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
)

ALTER TABLE notification_settings
ADD CONSTRAINT user_company_unique UNIQUE ("status", "company_id");

CREATE TABLE IF NOT EXISTS "notifications"(
    "id" SERIAL PRIMARY KEY,
    "company_id" UUID REFERENCES "companies"("id") NOT NULL,
    "model_type" VARCHAR(255),
    "model_id" INT,
    "details" JSONB,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
)

CREATE TABLE IF NOT EXISTS "user_notifications"(
    "id" SERIAL PRIMARY KEY,
    "notification_id" BIGINT REFERENCES "notifications"("id") NOT NULL,
    "is_read" BOOLEAN DEFAULT FALSE,
    "user_id" UUID,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
)

CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    model_type VARCHAR(50) NOT NULL,
    model_id INTEGER NOT NULL,
    type VARCHAR(10) NOT NULL,
    message TEXT,
    user_id UUID REFERENCES "users"("id"),
    voice_url TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_comments_model ON comments(model_type, model_id);

CREATE TABLE push_subscriptions (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES "users"("id") NOT NULL,
    data JSONB NOT NULL,
    -- endpoint TEXT NOT NULL,
    -- auth_key TEXT NOT NULL,
    -- p256dh_key TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "order_statuses"
ADD "order" INTEGER;