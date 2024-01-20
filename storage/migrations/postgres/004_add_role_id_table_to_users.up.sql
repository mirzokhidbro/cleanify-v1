ALTER TABLE "users"
ADD "role_id" UUID NULL REFERENCES "company_roles"("id");
