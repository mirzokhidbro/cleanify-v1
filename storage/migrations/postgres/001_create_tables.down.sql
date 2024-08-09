DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "companies";
DROP TABLE IF EXISTS "roles";
DROP TABLE IF EXISTS "bot_users";

ALTER TABLE "users"
DROP COLUMN IF EXISTS "role_id";

DROP TABLE IF EXISTS "orders";
DROP TABLE IF EXISTS "order_items";
DROP TABLE IF EXISTS "order_item_types";
DROP TABLE IF EXISTS "telegram_bots";
DROP TABLE IF EXISTS "telegram_sessions";
DROP TABLE IF EXISTS "permissions"
DROP TABLE IF EXISTS "role_and_permissions"
DROP TABLE IF EXISTS "clients"
DROP TABLE IF EXISTS "telegram_groups"
DROP TABLE IF EXISTS "status_change_histories"
DROP TABLE IF EXISTS "user_permissions"
DROP TABLE IF EXISTS "transactions"
DROP TABLE IF EXISTS "payment_purposes"
DROP TABLE IF EXISTS "employees"
DROP TABLE IF EXISTS "attendance"