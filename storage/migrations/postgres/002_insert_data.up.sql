INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c07', 'create_order',           'company',     'Buyurtma qo''shish', 'order');
-- INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c08', 'get_orders_list',        'company',     'Buyurtmalar ro''yxatini ko''rish', 'order');
-- INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c09', 'get_order',              'company',     'Buyurtma ma''lumotlarini batafsil ko''rish', 'order');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c10', 'edit_order',             'company',     'Buyurtmani taxrirlash', 'order');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c40', 'delete_order',             'company',     'Buyurtmani o''chirish', 'order');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c41', 'attach_courier_to_order',             'company',     'Buyurtmaga kuryer biriktirish', 'order');

INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c11', 'change_status_to_1',             'company',     'Buyurtma statusini "Olish kerak"ga o''zgartirish', 'order');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c12', 'change_status_to_2',             'company',     'Buyurtma statusini "Olingan"ga o''zgartirish', 'order');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c13', 'change_status_to_3',             'company',     'Buyurtma statusini "Yuvilgan"ga o''zgartirish', 'order');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c14', 'change_status_to_4',             'company',     'Buyurtma statusini "Tayyor"ga o''zgartirish', 'order');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c15', 'change_status_to_5',             'company',     'Buyurtma statusini "Oborish kerak"ga o''zgartirish', 'order');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c16', 'change_status_to_6',             'company',     'Buyurtma statusini "Topshirildi"ga o''zgartirish', 'order');
-- INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c17', 'change_status_to_7',             'company',     'Buyurtma statusini "Qaytarildi"ga o''zgartirish', 'order');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c18', 'change_status_to_8',             'company',     'Buyurtma statusini "Omborda"ga o''zgartirish', 'order');
-- INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c19', 'change_status_to_99',             'company',     'Buyurtma statusini "Bekor qilindi"ga o''zgartirish', 'order');

INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c01', 'user_create',            'company',     'Foydalanuvchi qo''shish', 'user');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c02', 'get_users_list',         'company',     'Foydalanuvchilar ro''yxatini olish', 'user');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c03', 'edit_user',         'company',     'Foydalanuvchi ma''lumotini taxrirlash', 'user');

INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c20', 'create_order_item',      'company',     'Gilam qo''shish', 'order_item');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c21', 'edit_order_item',       'company',      'Gilam ma''lumotlarini taxrirlash', 'order_item');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c22', 'delete_order_item',     'company',      'Gilamni o''chirish', 'order_item');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c23', 'change_status_to_washed',    'company',     'Gilamni statusini "Yuvilgan"ga o''zgartirish', 'order_item');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c24', 'change_status_to_prepare',    'company',     'Gilamni statusini "Tayyor"ga o''zgartirish', 'order_item');


INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c25', 'create_order_item_type', 'company',     'Mahsulot turini qo''shish', 'order_item_type');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c26', 'edit_order_item_type',   'company',     'Mahsulot turini taxrirlash', 'order_item_type');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c27', 'get_order_item_type',   'company',     'Mahsulot turlarini ko''rish', 'order_item_type');

INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c28', 'get_work_volume_by_day', 'company',     'Kunlik ish hajmi statistikasini ko''rish', 'statistics');

INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c29', 'get_clients_list',        'company',     'Klientlar ro''yxatini ko''rish', 'clients');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c30', 'show_client',        'company',     'Klient ma''lumotini batafsil ko''rish', 'clients');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c31', 'edit_client',        'company',     'Klientlar ma''lumotini tahrirlash', 'clients');


INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c32', 'get_telegram_groups_list',        'company',     'Telegram guruhlar ro''yxatini ko''rish', 'telegram_groups');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c33', 'add_telegram_group',              'company',     'Telegram guruh qo''shish', 'telegram_groups');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c34', 'show_telegram_group',             'company',     'Telegram guruh ma''lumotini batafsil ko''rish', 'telegram_groups');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c35', 'edit_telegram_group',        'company',     'Telegram guruhni tahrirlash', 'telegram_groups');

INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c36', 'get_employees_list',          'company',     'Xodimlar ro''yxatini ko''rish', 'Xodimlar');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c37', 'add_employee',                'company',     'Xodim qo''shish', 'Xodimlar');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c38', 'show_employee',               'company',     'Xodim ma''lumotini batafsil ko''rish', 'Xodimlar');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c39', 'add_transaction_to_employee', 'company',     'Xodimga oylik berish', 'Xodimlar');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c45', 'attendance_employee',         'company',     'Xodimlar davomati', 'Xodimlar');

INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c43', 'setting_get_order_statuses',               'company',     'Buyurtma statuslarini ko''rish', '[Sozlama] Buyurtma Statuslari');
INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c44', 'setting_edit_order_status',                'company',     'Buyurtma statuslarini taxrirlash', '[Sozlama] Buyurtma Statuslari');


INSERT INTO "permissions"("id", "slug", "scope", "name", "group") VALUES ('9cbb32da-e473-4312-8413-95524ec08c46', 'get_transactions',                'company',     'Tranzaksiyalar ro''yxatini olish', 'Tranzaksiyalar');

INSERT INTO "order_statuses"("company_id", "name", "number", "slug", "description") 
VALUES 
('8070790f-429b-449e-bfe5-fab0440c518f', 'Olish kerak',    1,  'change_status_to_1', 'Bu status mijoz tomonidan yangi zayavka tushganini anglatadi'),
('8070790f-429b-449e-bfe5-fab0440c518f', 'Olingan',        2,  'change_status_to_2', 'Bu status buyurtma mijozdan kuryer tomonidan olinganini anglatadi.'),
('8070790f-429b-449e-bfe5-fab0440c518f', 'Yuvilgan',       3,  'change_status_to_3', 'Bu status buyurtma yuvilganini bildiradi'),
('8070790f-429b-449e-bfe5-fab0440c518f', 'Tayyor',         4,  'change_status_to_4', 'Bu status buyurtma tayyorligini anglatadi'),
('8070790f-429b-449e-bfe5-fab0440c518f', 'Oborish kerak',  5,  'change_status_to_5', 'Buyurtma oborishga tayyor'),
('8070790f-429b-449e-bfe5-fab0440c518f', 'Topshirildi',    6,  'change_status_to_6', 'Buyurtma topshirildi'),
-- ('8070790f-429b-449e-bfe5-fab0440c518f', 'Qaytarildi',     7,  'change_status_to_7', 'Buyurtma Mijoz tomonidan qaytarildi'),
('8070790f-429b-449e-bfe5-fab0440c518f', 'Omborda',        8,  'change_status_to_8', 'Buyurtma omborga olindi'),
-- ('8070790f-429b-449e-bfe5-fab0440c518f', 'Bekor qilindi',  99, 'change_status_to_99', 'Buyurtma bekor qilindi');

"change_status_to_1"
"change_status_to_2"
"change_status_to_3"
"change_status_to_4"
"change_status_to_5"
"change_status_to_6"
"change_status_to_7"
"change_status_to_8"
"change_status_to_99"

update order_statuses
set slug = 'change_status_to_1'
where number = 1



INSERT INTO "payment_purposes"("name", "type")
VALUES
('from_order', 'income'),
('give_salary_to_worker', 'outcome'),
('debt_collection_from_the_employee', 'income')
('employee_loan', 'outcome');