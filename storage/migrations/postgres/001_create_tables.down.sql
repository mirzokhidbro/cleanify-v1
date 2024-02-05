DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "companies";
DROP TABLE IF EXISTS "roles";
DROP TABLE IF EXISTS "bot_users";

ALTER TABLE "users"
DROP COLUMN IF EXISTS "role_id";

DROP TABLE IF EXISTS "orders";
DROP TABLE IF EXISTS "order_items";
DROP TABLE IF EXISTS "order_item_types";
DROP TABLE IF EXISTS "company_bots";
DROP TABLE IF EXISTS "telegram_sessions";
DROP TABLE IF EXISTS "permissions"
DROP TABLE IF EXISTS "role_and_permissions"