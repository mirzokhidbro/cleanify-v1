CREATE TABLE IF NOT EXISTS "company_bots" (
    "id" UUID PRIMARY KEY,
    "company_id" UUID REFERENCES "companies"("id") NOT NULL,
    "bot_token" VARCHAR NOT NULL,
    "type" VARCHAR NOT NULL
);