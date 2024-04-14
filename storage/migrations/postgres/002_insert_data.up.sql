INSERT INTO "permissions"("id", "slug", "scope", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c01', 'user_create',            'company',     'Foydalanuvchi qo''shish');
INSERT INTO "permissions"("id", "slug", "scope", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c02', 'get_users_list',         'company',     'Foydalanuvchilar ro''yxatini olish');
INSERT INTO "permissions"("id", "slug", "scope", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c03', 'create_company',         'super-admin', 'Korxona qo''shish');
INSERT INTO "permissions"("id", "slug", "scope", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c04', 'get_company_by_owner',   'super-admin', 'Korxona egasi');
INSERT INTO "permissions"("id", "slug", "scope", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c05', 'create_role',            'company',     'Rol yaratish');
INSERT INTO "permissions"("id", "slug", "scope", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c06', 'get_roles_list',         'company',     'Rollar ro''yxatini ko''rish');
INSERT INTO "permissions"("id", "slug", "scope", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c07', 'create_order',           'company',     'Buyurtma qo''shish');
INSERT INTO "permissions"("id", "slug", "scope", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c08', 'get_orders_list',        'company',     'Buyurtmalar ro''yxatini ko''rish');
INSERT INTO "permissions"("id", "slug", "scope", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c09', 'get_order',              'company',     'Buyurtmani ko''rish');
INSERT INTO "permissions"("id", "slug", "scope", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c10', 'edit_order',             'company',     'Buyurtmani taxrirlash');
INSERT INTO "permissions"("id", "slug", "scope", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c11', 'send_location',          'company',     'Lokatsiyani olish');
INSERT INTO "permissions"("id", "slug", "scope", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c12', 'create_order_item',      'company',     'Buyurtma elementini qo''shish');
INSERT INTO "permissions"("id", "slug", "scope", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c13', 'edit_order_item,',       'company',     'Buyurtma elementini taxrirlash');
INSERT INTO "permissions"("id", "slug", "scope", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c14', 'create_order_item_type', 'company',     'Buyurtma turini qo''shish');
INSERT INTO "permissions"("id", "slug", "scope", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c15', 'edit_order_item_type',   'company',     'Buyurtma turini taxrirlash');
INSERT INTO "permissions"("id", "slug", "scope", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c16', 'create_telegram_bot',     'super-admin', 'Korxona botini qo''shish');
INSERT INTO "permissions"("id", "slug", "scope", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c17', 'start_telegram_bot',      'company',     'Korxona botini ishga tushirish');
INSERT INTO "permissions"("id", "slug", "scope", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c18', 'get_work_volume_by_day', 'company',     'Kunlik ish hajmini ko''rish statistikasini ko''rish');
INSERT INTO "permissions"("id", "slug", "scope", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c19', 'get_permissions_list',   'company',     'Ruxsatlar ro''yxatini ko''rish');
INSERT INTO "permissions"("id", "slug", "scope", "name") VALUES ('9cbb32da-e473-4312-8413-95524ec08c20', 'get_clients_list',        'company',     'Klientlar ro''yxatini ko''rish');


INSERT INTO "order_statuses"("company_id", "name", "number", "description") 
VALUES 
('2287b482-3450-44aa-aa43-8783d016d79b', 'Yangi Buyurtma', 1,  'Bu status mijoz tomonidan yangi zayavka tushganini anglatadi'),
('2287b482-3450-44aa-aa43-8783d016d79b', 'Olingan',        2,  'Bu status buyurtma mijozdan kuryer tomonidan olinganini anglatadi.'),
('2287b482-3450-44aa-aa43-8783d016d79b', 'Yuvilgan',       3,  'Bu status buyurtma yuvilganini bildiradi'),
('2287b482-3450-44aa-aa43-8783d016d79b', 'Tayyor',         4,  'Bu status buyurtma tayyorligini anglatadi'),
('2287b482-3450-44aa-aa43-8783d016d79b', 'Oborish kerak',  5,  'Buyurtma oborishga tayyor'),
('2287b482-3450-44aa-aa43-8783d016d79b', 'Topshirildi',    6,  'Buyurtma topshirildi'),
('2287b482-3450-44aa-aa43-8783d016d79b', 'Qaytarildi',     7,  'Buyurtma Mijoz tomonidan qaytarildi'),
('2287b482-3450-44aa-aa43-8783d016d79b', 'Omborda',        8,  'Buyurtma omborga olindi'),
('2287b482-3450-44aa-aa43-8783d016d79b', 'Bekor qilindi',  99, 'Buyurtma bekor qilindi');

