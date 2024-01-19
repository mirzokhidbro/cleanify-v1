CREATE TABLE IF NOT EXISTS "company_roles" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR,
    "company_id" UUID REFERENCES "companies"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);