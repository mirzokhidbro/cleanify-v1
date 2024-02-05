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

ALTER TABLE "users"
ADD "role_id" UUID NULL REFERENCES "roles"("id");

ALTER TABLE "orders"
ALTER COLUMN "chat_id" SET DATA TYPE BIGINT
USING chat_id::BIGINT;

ALTER TABLE "orders"
ALTER COLUMN "status" SET DEFAULT 1;

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

CREATE TABLE IF NOT EXISTS "company_bots" (
    "id" UUID PRIMARY KEY,
    "company_id" UUID REFERENCES "companies"("id") NOT NULL,
    "bot_token" VARCHAR NOT NULL,
	"bot_id" BIGINT,
    "type" VARCHAR NOT NULL
);

ALTER TABLE "company_bots"
ADD UNIQUE (bot_id);

ALTER TABLE "company_bots"
ADD "firstname" VARCHAR,
ADD "lastname" VARCHAR,
ADD "username" VARCHAR;

CREATE TABLE IF NOT EXISTS "telegram_sessions" (
    "id" SERIAL PRIMARY KEY,
    "bot_id" BIGINT REFERENCES "company_bots"("bot_id") NOT NULL,
    "order_id" INT REFERENCES "orders"("id") NOT NULL,
	"chat_id" BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS "permissions" (
    "id" UUID PRIMARY KEY,
    "slug" VARCHAR UNIQUE,
    "name" VARCHAR UNIQUE,
    "scope" VARCHAR DEFAULT ("company"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "role_and_permissions" (
    "id" SERIAL PRIMARY KEY,
    "role_id" VARCHAR UNIQUE,
    "permission_ids" VARCHAR[]
);
