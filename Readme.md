# Тестовое задание для Smartway

## БД
Диаграмма базы данных

![](/img/diagram.png)

<hr>
В таблице tickets добавил поле flight_number - номер рейса.
В ТЗ написано, что по билету нужно выводить информацию о "пассажирах" - во множественном числе. Возможно это опечатка,
но я на всякий случай добавил это поле, и по uuid билета вывожу всех пассажиров на данном рейсе.
<hr>

**В директории migrates лежат скрипты для создания таблиц, а так же для заполнения тестовыми данными**

## Config
Конфиг парсится из энвов
- PG_HOST
- PG_PORT
- PG_USER
- PG_PASSWORD
- PG_DB_NAME
- APP_PORT

## Handlers
### Document
1. Обновление PUT 127.0.0.1:8081/api/document

body:
```json
    {
        "uuid" : "uuid",
        "type" : "type",
        "value" : "value"
    }
```
2. Удаление DELETE 127.0.0.1:8081/api/document
---
body:
```json
    {
        "uuid" : "uuid"
    }
```
---
3. Получение по uuid пассажира POST 127.0.0.1:8081/api/document/get-by-passenger
---
body:
```json
    {
        "uuid" : "uuid"
    }
```
---

### Passenger
1. Обновление PUT 127.0.0.1:8081/api/passenger

body:
```json
    {
        "uuid" : "uuid",
        "last_name" : "last_name",
        "first_name" : "first_name",
        "middle_name" : "middle_name"
    }
```
2. Удаление DELETE 127.0.0.1:8081/api/passenger
---
body:
```json
    {
        "uuid" : "uuid"
    }
```
---
3. Получение отчета POST 127.0.0.1:8081/api/document/report
---
body с тестовыми данными:
```json
    {
      "passenger_uuid" : "d3febc99-9c0b-4ef8-bb6d-6bb9bd380a15",
      "start_date" : "2020-01-01",
      "end_date" : "2024-12-31"
    }
```
---

### Ticket
1. Обновление PUT 127.0.0.1:8081/api/ticket

body:
```json
    {
        "uuid" : "uuid",
        "departure" : "departure",
        "departure_date" : "departure_date",
        "destination" : "destination",
        "arrival_date" : "arrival_date",
        "order_number" : "order_number",
        "provider" : "provider",
        "booking_date" : "booking_date",
        "flight_number" : "flight_number"
    }
```
2. Удаление DELETE 127.0.0.1:8081/api/ticket
---
body:
```json
    {
        "uuid" : "uuid"
    }
```
---
3. Получение всех билетов Post 127.0.0.1:8081/api/document/all
---
body:
```json
    {
      "size" : 10,
      "page" : 0
    }
```
4. Получение полной информации по билету Post 127.0.0.1:8081/api/document/full-info
---
body:
```json
    {
      "uuid" : "uuid"
    }
```
5. Получение пассажиров Post 127.0.0.1:8081/api/document/passengers
---
body:
```json
    {
      "uuid" : "uuid"
    }
```
---

## Репозитории

### Document
1. Получение документов по uuid пассажира
2. Обновление документа
3. Удаление документа

### Passenger
1. Обновление пассажира
2. Удаление пассажира
3. Получение отчета

### Tickets
1. Получение всех билетов
2. Обновление билета
3. Удаление билета
4. Получение полной информации по билету (вывод всех пассажиров + документы)
5. Получение всех пассажиров на рейс по uuid билета