INSERT INTO categories (id, name, parent_id, created_at, updated_at, deleted_at)
VALUES ('9527bd9a-19a9-4f53-8f2f-9b0cd689fb39', 'Основные комплектующие для ПК', null,
        '2023-04-27 12:35:36.017626 +00:00', '2023-04-27 12:35:36.017626 +00:00', null);
INSERT INTO categories (id, name, parent_id, created_at, updated_at, deleted_at)
VALUES ('68be4526-d880-4c89-90f4-60d1063ea021', 'Мониторы', null, '2023-04-27 12:35:36.017626 +00:00',
        '2023-04-27 12:35:36.017626 +00:00', null);
INSERT INTO categories (id, name, parent_id, created_at, updated_at, deleted_at)
VALUES ('6b30fb33-b4da-444c-b833-0752fd2c241b', 'Устройства расширения', null, '2023-04-27 12:35:36.017626 +00:00',
        '2023-04-27 12:35:36.017626 +00:00', null);
INSERT INTO categories (id, name, parent_id, created_at, updated_at, deleted_at)
VALUES ('3575a27c-f37b-4ed7-9d27-1a36957db6c3', 'Процессоры', '9527bd9a-19a9-4f53-8f2f-9b0cd689fb39',
        '2023-04-27 12:35:36.017626 +00:00', '2023-04-27 12:35:36.017626 +00:00', null);
INSERT INTO categories (id, name, parent_id, created_at, updated_at, deleted_at)
VALUES ('51f2ee06-e34b-49b9-a9ac-d3a1f5be39ce', 'Звуковые карты', '6b30fb33-b4da-444c-b833-0752fd2c241b',
        '2023-04-27 12:35:36.017626 +00:00', '2023-04-27 12:35:36.017626 +00:00', null);
INSERT INTO categories (id, name, parent_id, created_at, updated_at, deleted_at)
VALUES ('23117557-dd4f-43a1-b9f1-e7ca6a49e03e', 'Оперативная память', '9527bd9a-19a9-4f53-8f2f-9b0cd689fb39',
        '2023-04-27 20:57:20.780402 +00:00', '2023-04-27 20:57:20.780402 +00:00', null);

INSERT INTO products (id, name, price, category_id, created_at, updated_at, deleted_at)
VALUES ('3de1590b-afdb-4bef-8179-0304969dfe4a', 'Процессор Intel Core i9-10940X OEM', 7500000,
        '3575a27c-f37b-4ed7-9d27-1a36957db6c3', '2023-04-27 12:36:16.493843 +00:00',
        '2023-04-27 12:36:16.493843 +00:00', null);
INSERT INTO products (id, name, price, category_id, created_at, updated_at, deleted_at)
VALUES ('62904807-2dad-4f54-8d5e-f01bad26fff0', 'Процессор AMD Ryzen 9 7950X BOX', 5699900,
        '3575a27c-f37b-4ed7-9d27-1a36957db6c3', '2023-04-27 12:36:16.493843 +00:00',
        '2023-04-27 12:36:16.493843 +00:00', null);
INSERT INTO products (id, name, price, category_id, created_at, updated_at, deleted_at)
VALUES ('921ecb4c-1595-45d2-9174-08133007ec16', 'ASUS ProArt PA27UCX-K', 19699900,
        '68be4526-d880-4c89-90f4-60d1063ea021', '2023-04-28 11:34:46.720668 +00:00',
        '2023-04-28 11:34:46.720668 +00:00', null);
INSERT INTO products (id, name, price, category_id, created_at, updated_at, deleted_at)
VALUES ('6fce43d5-0fba-48ae-8d50-ab36dea3500e', 'Philips 329P9H/00', 7829900, '68be4526-d880-4c89-90f4-60d1063ea021',
        '2023-04-28 12:41:48.149658 +00:00', '2023-04-28 12:41:48.149658 +00:00', null);
INSERT INTO products (id, name, price, category_id, created_at, updated_at, deleted_at)
VALUES ('256337e1-5df6-484a-8b2f-f85c8f4ed869', 'Zoom TAC-2R', 500000, '51f2ee06-e34b-49b9-a9ac-d3a1f5be39ce',
        '2023-04-28 12:43:51.264461 +00:00', '2023-04-28 12:43:51.264461 +00:00', null);
INSERT INTO products (id, name, price, category_id, created_at, updated_at, deleted_at)
VALUES ('e5438b28-124e-4ffa-a463-e32b40e78193', 'PreSonus Studio 68C', 4079900, '51f2ee06-e34b-49b9-a9ac-d3a1f5be39ce',
        '2023-04-28 12:43:51.264461 +00:00', '2023-04-28 12:43:51.264461 +00:00', null);
INSERT INTO products (id, name, price, category_id, created_at, updated_at, deleted_at)
VALUES ('08eba499-cd6b-4a9e-b300-2295b935b4e9', 'Kingston FURY Renegade Silver RGB', 3099900,
        '23117557-dd4f-43a1-b9f1-e7ca6a49e03e', '2023-04-28 12:43:51.264461 +00:00',
        '2023-04-28 12:43:51.264461 +00:00', null);

INSERT INTO properties (id, name, created_at, updated_at, deleted_at)
VALUES ('e243a8bd-7a7b-4175-9821-acef97ba5565', 'Общее количество ядер', '2023-04-27 12:37:35.593857 +00:00',
        '2023-04-27 12:37:35.593857 +00:00', null);
INSERT INTO properties (id, name, created_at, updated_at, deleted_at)
VALUES ('9270bf8d-cdf9-4bab-bfc3-1d52b5fa6a64', 'Базовая частота процессора ГГц', '2023-04-27 12:37:35.593857 +00:00',
        '2023-04-27 12:37:35.593857 +00:00', null);
INSERT INTO properties (id, name, created_at, updated_at, deleted_at)
VALUES ('bdaa7f06-3e7a-4584-9a5b-9d7b5f75f052', 'Диагональ экрана', '2023-04-28 11:35:15.523168 +00:00',
        '2023-04-28 11:35:15.523168 +00:00', null);
INSERT INTO properties (id, name, created_at, updated_at, deleted_at)
VALUES ('b10160e0-3e12-4ab9-8e78-ff6e4386f60e', 'Технология изготовления матрицы', '2023-04-28 11:35:23.895976 +00:00',
        '2023-04-28 11:35:23.895976 +00:00', null);
INSERT INTO properties (id, name, created_at, updated_at, deleted_at)
VALUES ('dafd8351-6562-46d7-aa5c-c09dccef0ce2', 'Тип памяти', '2023-04-28 12:47:03.482386 +00:00',
        '2023-04-28 12:47:03.482386 +00:00', null);
INSERT INTO properties (id, name, created_at, updated_at, deleted_at)
VALUES ('97182043-8783-4948-afba-731be4b8b771', 'Тактовая частота МГц', '2023-04-28 12:47:03.482386 +00:00',
        '2023-04-28 12:47:03.482386 +00:00', null);
INSERT INTO properties (id, name, created_at, updated_at, deleted_at)
VALUES ('915dc185-0f7d-4139-a6c7-9e8df33efe68', 'Модели ЦАП', '2023-04-28 12:48:43.123338 +00:00',
        '2023-04-28 12:48:43.123338 +00:00', null);
INSERT INTO properties (id, name, created_at, updated_at, deleted_at)
VALUES ('e819e64e-6c71-48d4-9cf8-a2705af7dceb', 'Разрядность АЦП', '2023-04-28 12:48:43.123338 +00:00',
        '2023-04-28 12:48:43.123338 +00:00', null);

INSERT INTO product_properties (product_id, property_id, value, created_at, updated_at, deleted_at)
VALUES ('921ecb4c-1595-45d2-9174-08133007ec16', 'bdaa7f06-3e7a-4584-9a5b-9d7b5f75f052', '27',
        '2023-04-28 11:36:19.618498 +00:00', '2023-04-28 11:36:19.618498 +00:00', null);
INSERT INTO product_properties (product_id, property_id, value, created_at, updated_at, deleted_at)
VALUES ('921ecb4c-1595-45d2-9174-08133007ec16', 'b10160e0-3e12-4ab9-8e78-ff6e4386f60e', 'IPS',
        '2023-04-28 11:36:19.618498 +00:00', '2023-04-28 11:36:19.618498 +00:00', null);
INSERT INTO product_properties (product_id, property_id, value, created_at, updated_at, deleted_at)
VALUES ('6fce43d5-0fba-48ae-8d50-ab36dea3500e', 'bdaa7f06-3e7a-4584-9a5b-9d7b5f75f052', '31.5',
        '2023-04-28 12:46:05.449000 +00:00', '2023-04-28 12:46:05.449000 +00:00', null);
INSERT INTO product_properties (product_id, property_id, value, created_at, updated_at, deleted_at)
VALUES ('6fce43d5-0fba-48ae-8d50-ab36dea3500e', 'b10160e0-3e12-4ab9-8e78-ff6e4386f60e', 'IPS',
        '2023-04-28 12:46:36.673012 +00:00', '2023-04-28 12:46:36.673012 +00:00', null);
INSERT INTO product_properties (product_id, property_id, value, created_at, updated_at, deleted_at)
VALUES ('08eba499-cd6b-4a9e-b300-2295b935b4e9', 'dafd8351-6562-46d7-aa5c-c09dccef0ce2', 'DDR5',
        '2023-04-28 12:47:56.082202 +00:00', '2023-04-28 12:47:56.082202 +00:00', null);
INSERT INTO product_properties (product_id, property_id, value, created_at, updated_at, deleted_at)
VALUES ('08eba499-cd6b-4a9e-b300-2295b935b4e9', '97182043-8783-4948-afba-731be4b8b771', '6000',
        '2023-04-28 12:47:56.082202 +00:00', '2023-04-28 12:47:56.082202 +00:00', null);
INSERT INTO product_properties (product_id, property_id, value, created_at, updated_at, deleted_at)
VALUES ('256337e1-5df6-484a-8b2f-f85c8f4ed869', '915dc185-0f7d-4139-a6c7-9e8df33efe68', 'AKM AK4396',
        '2023-04-28 12:49:56.659015 +00:00', '2023-04-28 12:49:56.659015 +00:00', null);
INSERT INTO product_properties (product_id, property_id, value, created_at, updated_at, deleted_at)
VALUES ('256337e1-5df6-484a-8b2f-f85c8f4ed869', 'e819e64e-6c71-48d4-9cf8-a2705af7dceb', '24 бит',
        '2023-04-28 12:49:56.659015 +00:00', '2023-04-28 12:49:56.659015 +00:00', null);
INSERT INTO product_properties (product_id, property_id, value, created_at, updated_at, deleted_at)
VALUES ('e5438b28-124e-4ffa-a463-e32b40e78193', '915dc185-0f7d-4139-a6c7-9e8df33efe68', 'AKM AK4355',
        '2023-04-28 12:49:56.659015 +00:00', '2023-04-28 12:49:56.659015 +00:00', null);
INSERT INTO product_properties (product_id, property_id, value, created_at, updated_at, deleted_at)
VALUES ('e5438b28-124e-4ffa-a463-e32b40e78193', 'e819e64e-6c71-48d4-9cf8-a2705af7dceb', '24 бит',
        '2023-04-28 12:49:56.659015 +00:00', '2023-04-28 12:49:56.659015 +00:00', null);
INSERT INTO product_properties (product_id, property_id, value, created_at, updated_at, deleted_at)
VALUES ('3de1590b-afdb-4bef-8179-0304969dfe4a', 'e243a8bd-7a7b-4175-9821-acef97ba5565', '14',
        '2023-04-27 12:37:49.985469 +00:00', '2023-04-27 12:37:49.985469 +00:00', null);
INSERT INTO product_properties (product_id, property_id, value, created_at, updated_at, deleted_at)
VALUES ('3de1590b-afdb-4bef-8179-0304969dfe4a', '9270bf8d-cdf9-4bab-bfc3-1d52b5fa6a64', '3.3',
        '2023-04-27 12:37:49.985469 +00:00', '2023-04-27 12:37:49.985469 +00:00', null);
INSERT INTO product_properties (product_id, property_id, value, created_at, updated_at, deleted_at)
VALUES ('62904807-2dad-4f54-8d5e-f01bad26fff0', '9270bf8d-cdf9-4bab-bfc3-1d52b5fa6a64', '4.5',
        '2023-04-27 12:37:49.985469 +00:00', '2023-04-27 12:37:49.985469 +00:00', null);
INSERT INTO product_properties (product_id, property_id, value, created_at, updated_at, deleted_at)
VALUES ('62904807-2dad-4f54-8d5e-f01bad26fff0', 'e243a8bd-7a7b-4175-9821-acef97ba5565', '16',
        '2023-04-27 12:37:49.985469 +00:00', '2023-04-27 12:37:49.985469 +00:00', null);
