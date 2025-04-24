-- Users jadvalini yangilash
ALTER TABLE "users" 
ADD COLUMN "fullname" VARCHAR;

-- Mavjud foydalanuvchilar uchun fullname'ni to'ldirish
UPDATE "users" 
SET "fullname" = TRIM(COALESCE("firstname", '') || ' ' || COALESCE("lastname", ''));

-- Eski ustunlarni o'chirish
ALTER TABLE "users"
DROP COLUMN "firstname",
DROP COLUMN "lastname";

-- Employees jadvalini yangilash
ALTER TABLE "employees"
ADD COLUMN "fullname" VARCHAR;

-- Mavjud xodimlar uchun fullname'ni to'ldirish
UPDATE "employees"
SET "fullname" = TRIM(COALESCE("firstname", '') || ' ' || COALESCE("lastname", ''));

-- Eski ustunlarni o'chirish
ALTER TABLE "employees"
DROP COLUMN "firstname",
DROP COLUMN "lastname";
