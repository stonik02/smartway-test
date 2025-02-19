CREATE TABLE passengers (
    uuid UUID PRIMARY KEY,
    last_name VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    middle_name VARCHAR(255)
);

CREATE TABLE tickets (
    uuid UUID PRIMARY KEY,
    passenger_uuid UUID NOT NULL REFERENCES passengers(uuid) ON DELETE CASCADE,
    departure VARCHAR(255) NOT NULL,
    destination VARCHAR(255) NOT NULL,
    departure_date TIMESTAMP NOT NULL,
    arrival_date TIMESTAMP NOT NULL,
    order_number VARCHAR(255) NOT NULL UNIQUE,
    provider VARCHAR(255) NOT NULL,
    booking_date DATE NOT NULL,
    flight_number VARCHAR(50) NOT NULL -- Номер рейса для выборки всех пассажиров по билету
);

CREATE INDEX idx_tickets_passenger_uuid ON tickets (passenger_uuid);


CREATE TABLE documents (
    uuid UUID PRIMARY KEY,
    passenger_uuid UUID NOT NULL REFERENCES passengers(uuid) ON DELETE CASCADE,
    type VARCHAR(255) NOT NULL,
    number VARCHAR(255) NOT NULL
);

CREATE INDEX idx_documents_passenger_uuid ON documents (passenger_uuid);


-- Получить юзера со всеми документами (на будущее)
CREATE VIEW passengers_with_docs AS
SELECT p.*,
       json_agg(
               json_build_object(
                       'uuid', doc.uuid,
                       'type', doc.type,
                       'number', doc.number
               )
       ) as documents
FROM passengers AS p
         LEFT JOIN documents AS doc
                   ON doc.passenger_uuid = p.uuid
GROUP BY p.uuid;

-- View для отчета
CREATE VIEW passenger_report AS
SELECT
    t.passenger_uuid AS passenger_uuid,
    t.booking_date,
    t.departure_date,
    t.order_number,
    t.departure,
    t.destination,
    CASE
        WHEN t.departure_date < CURRENT_TIMESTAMP THEN TRUE
        ELSE FALSE
        END AS "service_rendered"
FROM tickets t;

