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