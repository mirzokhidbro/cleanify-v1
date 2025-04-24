-- Users jadvalini qaytarish
ALTER TABLE "users"
ADD COLUMN "firstname" VARCHAR,
ADD COLUMN "lastname" VARCHAR;

-- Mavjud foydalanuvchilar uchun firstname va lastname'ni ajratib olish
UPDATE "users"
SET 
    "firstname" = SPLIT_PART("fullname", ' ', 1),
    "lastname" = SUBSTRING("fullname" FROM POSITION(' ' IN "fullname") + 1);

-- fullname ustunini o'chirish
ALTER TABLE "users"
DROP COLUMN "fullname";

-- Employees jadvalini qaytarish
ALTER TABLE "employees"
ADD COLUMN "firstname" VARCHAR,
ADD COLUMN "lastname" VARCHAR;

-- Mavjud xodimlar uchun firstname va lastname'ni ajratib olish
UPDATE "employees"
SET 
    "firstname" = SPLIT_PART("fullname", ' ', 1),
    "lastname" = SUBSTRING("fullname" FROM POSITION(' ' IN "fullname") + 1);

-- fullname ustunini o'chirish
ALTER TABLE "employees"
DROP COLUMN "fullname";
