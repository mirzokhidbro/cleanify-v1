ALTER TABLE "users"
ADD COLUMN "is_active" BOOLEAN DEFAULT TRUE;

UPDATE "users" SET "is_active" = TRUE WHERE "is_active" IS NULL;
