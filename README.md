## Тестовое задание магазин комплектующих для ПК.

В config-default.yaml указать соответствующие параметры.

Для базы выполнить миграцию.
Сам использую https://github.com/golang-migrate/migrate
Для установки на Linux с установленным golang выполнить:
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```
После установки, прогнать миграции(указав свои конфигурации) из корня проекта:
```bash
migrate -path migrations -database "postgres://user:password@localhost:5432/database?sslmode=disable" up
```
Запускаем проект из корня проекта:
```bash
go run -v ./cmd/app
```
### Сurl запросов для проверки

Получить товары с фильтрами:
```bash
curl
--location 'http://localhost:8010/storage/products?property_id=915dc185-0f7d-4139-a6c7-9e8df33efe68&category_id=51f2ee06-e34b-49b9-a9ac-d3a1f5be39ce&property_val=AKM%20AK4396'
```
Получить товар по его идентификатору:
```bash
curl --location 'http://localhost:8010/storage/product?id=672476ca-a87b-4e9b-a368-e014b39b3fd8'
```
Добавить товар:
```bash
curl --location 'http://localhost:8010/storage/product' \
--header 'Content-Type: application/json' \
--data '{
"name": "i9-9900k",
"price": 3000000,
"category_id": "3575a27c-f37b-4ed7-9d27-1a36957db6c3",
"properties": [
{
"name": "Общее количество ядер",
"value": "8"
},
{
"name": "Базовая частота процессора ГГц",
"value": "3.6"
}
]
}'
```

Обновить товар по его идентификатору:
```bash
curl --location --request PUT 'http://localhost:8010/storage/product?id=3de1590b-afdb-4bef-8179-0304969dfe4a' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Процессор Intel Core i9-10940X OEM",
    "price": 7550000,
    "category_id": "3575a27c-f37b-4ed7-9d27-1a36957db6c3",
    "properties": [
        {
            "name": "Общее количество ядер",
            "value": "14"
        },
        {
            "name": "Максимальное число потоков",
            "value": "28"
        }
    ]
}'
```

Удаление товара:
```bash
curl --location --request DELETE 'http://localhost:8010/storage/product?id=3de1590b-afdb-4bef-8179-0304969dfe4a'
```

По заданию не было условий реализации добавления категорий, решил сделать отдельные эндпоинты получения всех категорий и добавления.
Считаю что в момент создания товара категория уже должна существовать, иначе будет лишняя генерация дублирующих категорий.

Список категорий и их идентификатор:
```bash
curl --location 'http://localhost:8010/storage/categories'
```

Добавить категорию:
```bash
curl --location 'http://localhost:8010/storage/category' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Видеокарты",
    "parent_id": "9527bd9a-19a9-4f53-8f2f-9b0cd689fb39"

}'
```

by F7icK™ 2023