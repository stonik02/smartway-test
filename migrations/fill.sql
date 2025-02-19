-- Заполняем пассажиров
INSERT INTO passengers (uuid, last_name, first_name, middle_name)
VALUES
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'Иванов', 'Иван', 'Иванович'),
    ('b1febc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'Петрова', 'Анна', 'Сергеевна'),
    ('c2febc99-9c0b-4ef8-bb6d-6bb9bd380a13', 'Сидоров', 'Алексей', 'Николаевич'),
    ('d3febc99-9c0b-4ef8-bb6d-6bb9bd380a14', 'Махмуд', 'Третий', ''),
    ('d3febc99-9c0b-4ef8-bb6d-6bb9bd380a15', 'Тестовый', 'Пассажир', 'Летит');

-- Заполняем документы
INSERT INTO documents (uuid, passenger_uuid, type, number) VALUES
    (gen_random_uuid(), 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'PASSPORT', '1234 567890'),
    (gen_random_uuid(), 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'INTERNATIONAL_PASSPORT', '0987654321'),
    (gen_random_uuid(), 'b1febc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'PASSPORT', '1122 334455'),
    (gen_random_uuid(), 'c2febc99-9c0b-4ef8-bb6d-6bb9bd380a13', 'PASSPORT', '5566 778899'),
    (gen_random_uuid(), 'd3febc99-9c0b-4ef8-bb6d-6bb9bd380a14', 'PASSPORT', '5512 778821'),
    (gen_random_uuid(), 'd3febc99-9c0b-4ef8-bb6d-6bb9bd380a15', 'INTERNATIONAL_PASSPORT', 'AB1234567');

INSERT INTO tickets (uuid, passenger_uuid, departure, destination, departure_date, arrival_date, order_number, provider, booking_date, flight_number)
VALUES
-- Рейс SU100: Москва -> Лондон (2 билета)
    (gen_random_uuid(), 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'Moscow', 'London', '2023-10-10 10:00:00', '2023-10-10 13:00:00', 'ORD001', 'Aeroflot', '2023-09-01', 'SU100'),
    (gen_random_uuid(), 'b1febc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'Moscow', 'London', '2023-10-10 10:00:00', '2023-10-10 13:00:00', 'ORD002', 'Aeroflot', '2023-09-02', 'SU100'),

-- Рейс LH200: Берлин -> Париж (2 билета)
    (gen_random_uuid(), 'c2febc99-9c0b-4ef8-bb6d-6bb9bd380a13', 'Berlin', 'Paris', '2023-10-11 09:00:00', '2023-10-11 11:00:00', 'ORD003', 'Lufthansa', '2023-09-03', 'LH200'),
    (gen_random_uuid(), 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'Berlin', 'Paris', '2023-10-11 09:00:00', '2023-10-11 11:00:00', 'ORD004', 'Lufthansa', '2023-09-04', 'LH200'),

-- Рейс EK300: Дубай -> Сингапур (2 билета)
    (gen_random_uuid(), 'd3febc99-9c0b-4ef8-bb6d-6bb9bd380a14', 'Dubai', 'Singapore', '2023-10-12 08:30:00', '2023-10-12 14:00:00', 'ORD005', 'Emirates', '2023-09-05', 'EK300'),
    (gen_random_uuid(), 'b1febc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'Dubai', 'Singapore', '2023-10-12 08:30:00', '2023-10-12 14:00:00', 'ORD006', 'Emirates', '2023-09-06', 'EK300'),

-- Тестовые данные для отчета
    (gen_random_uuid(), 'd3febc99-9c0b-4ef8-bb6d-6bb9bd380a15', 'Москва', 'Питер', '2021-10-12 08:30:00', '2021-10-12 14:00:00', 'ORD007', 'Emirates', '2019-09-06', 'JS100'),
    (gen_random_uuid(), 'd3febc99-9c0b-4ef8-bb6d-6bb9bd380a15', 'Абу-Даби', 'Севастополь', '2025-10-12 08:30:00', '2025-10-12 14:00:00', 'ORD008', 'Emirates', '2024-09-06', 'JS100'),
    (gen_random_uuid(), 'd3febc99-9c0b-4ef8-bb6d-6bb9bd380a15', 'Чебоксары', 'Питер', '2024-10-12 08:30:00', '2023-10-12 14:00:00', 'ORD009', 'Emirates', '2023-09-06', 'JS100'),
    (gen_random_uuid(), 'd3febc99-9c0b-4ef8-bb6d-6bb9bd380a15', 'Марс', 'Юпитер', '2015-10-12 08:30:00', '2015-10-12 14:00:00', 'ORD010', 'Emirates', '2015-09-06', 'JS100');

--  Берем отчет с 2020 - 2024
-- 1. Перелёты заказанные ранее, но совершенные в этот период заказан - 2019 вылет 2021 Москва - Питер
-- 2. Перелёты заказанные в этот период и не совершенные заказан - 2024 вылет 2025 Абу-Даби - Севастополь
-- 3. Перелёты заказанные в этот период и совершенные заказан - 2022 вылет 2023 Чебоксары - Питер
-- 4. Не должен пройти заказан 2015 вылет 2015 Марс - Юпитер