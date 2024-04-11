CREATE TABLE IF NOT EXISTS "users" (
    "id" UUID PRIMARY KEY,
    "phone" VARCHAR UNIQUE,
    "firstname" VARCHAR,
    "lastname" VARCHAR,
    "password" VARCHAR(1000),
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

CREATE TABLE IF NOT EXISTS "roles" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR,
    "company_id" UUID REFERENCES "companies"("id"),
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
    "chat_id" BIGINT NOT NULL,
    "status" VARCHAR,
    "page" VARCHAR,
    "role" VARCHAR,
    "dialog_step" VARCHAR,
    "bot_id" BIGINT REFERENCES "telegram_bots"("bot_id")
);

CREATE TABLE IF NOT EXISTS "orders" (
    "id" SERIAL PRIMARY KEY,
    "company_id" UUID REFERENCES "companies"("id"),
    "chat_id" INTEGER REFERENCES "bot_users"("id"),
    "phone" VARCHAR NULL,
    "count" INTEGER,
    "status" INTEGER,
    "slug" VARCHAR,
    "latitute" FLOAT,
    "longitude" FLOAT,
    "description" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

ALTER TABLE "users"
ADD "role_id" UUID NULL REFERENCES "roles"("id");

ALTER TABLE "orders"
ALTER COLUMN "chat_id" SET DATA TYPE BIGINT
USING chat_id::BIGINT;

ALTER TABLE "orders"
ALTER COLUMN "status" SET DEFAULT 83;

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
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "order_item_types" (
    "id" UUID PRIMARY KEY,
    "company_id" UUID REFERENCES "companies"("id") NOT NULL,
    "name" VARCHAR NOT NULL,
    "price" FLOAT NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

ALTER TABLE "telegram_bots"
ADD UNIQUE (bot_id);

ALTER TABLE "bot_users"
ADD UNIQUE (chat_id);

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
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "role_and_permissions" (
    "id" SERIAL PRIMARY KEY,
    "role_id" VARCHAR UNIQUE,
    "permission_ids" VARCHAR[]
);

ALTER TABLE "orders"
ADD "address" VARCHAR;

ALTER TABLE "bot_users"
ADD "firstname" VARCHAR,
ADD "lastname" VARCHAR,
ADD "username" VARCHAR,
ADD "company_id" UUID REFERENCES "companies"("id");

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

ALTER TABLE "users"
ADD COLUMN "permission_ids" VARCHAR[],
ADD COLUMN "company_id" UUID REFERENCES "companies"("id");

ALTER TABLE "order_items"
ADD COLUMN "is_countable" BOOLEAN DEFAULT FALSE,
ADD COLUMN "count" INTEGER;


ALTER TABLE "order_item_types"
ADD COLUMN "is_countable" BOOLEAN DEFAULT FALSE;