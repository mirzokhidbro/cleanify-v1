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
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

ALTER TABLE orders
ALTER COLUMN chat_id SET DATA TYPE BIGINT
USING chat_id::BIGINT;