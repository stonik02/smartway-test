package storage

import (
	"context"
	"errors"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"test-task/internal/models"
	"test-task/internal/storage/query"
)

type Ticket interface {
	Create(ctx context.Context, args *models.Ticket) error
	Get(ctx context.Context, query *query.Limit) ([]*models.Ticket, error)
	Update(ctx context.Context, ticket *models.Ticket) error
	Delete(ctx context.Context, uid uuid.UUID) error

	GetTicketFullInfo(ctx context.Context, uid uuid.UUID) (*models.TicketsFullInfo, error)
	GetPassengers(ctx context.Context, uid uuid.UUID) ([]*models.PassengerWithDocs, error)
}

type ticket struct {
	client *pgxpool.Pool
}

func NewTicket(client *pgxpool.Pool) Ticket {
	return &ticket{
		client: client,
	}
}

var createTicket = `
INSERT INTO tickets (uuid, passenger_uuid, departure,
                     destination, departure_date, arrival_date,
                     order_number, provider, booking_date,
                     flight_number)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
`

func (r ticket) Create(ctx context.Context, args *models.Ticket) error {
	_, err := r.client.Exec(ctx, createTicket,
		args.UUID, args.PassengerUUID, args.Departure,
		args.Destination, args.DepartureDate, args.ArrivalDate,
		args.OrderNumber, args.Provider, args.BookingDate,
		args.FlightNumber,
	)
	return err
}

var getTickets = `SELECT * FROM tickets LIMIT $1 OFFSET $2;`

func (t ticket) Get(ctx context.Context, args *query.Limit) (tickets []*models.Ticket, err error) {
	args.Validate()
	err = pgxscan.Select(ctx, t.client, &tickets, getTickets, args.Size, args.Size*args.Page)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	return
}

var updateTickets = `
UPDATE tickets
SET
    departure = $2,
    destination = $3,
    departure_date = $4,
    flight_number = $5,
    arrival_date = $6,
    order_number = $7,
    provider = $8,
    booking_date = $9,
    passenger_uuid = $10
WHERE uuid = $1;
`

func (t ticket) Update(ctx context.Context, ticket *models.Ticket) error {
	res, err := t.client.Exec(ctx, updateTickets,
		ticket.UUID, ticket.Departure, ticket.Destination,
		ticket.DepartureDate, ticket.FlightNumber, ticket.ArrivalDate,
		ticket.OrderNumber, ticket.Provider, ticket.BookingDate)

	if err != nil {
		return err // Ошибка выполнения запроса
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("ticket not found") // Запись не найдена
	}

	return err
}

var deleteTickets = `DELETE FROM tickets WHERE uuid = $1;`

func (t ticket) Delete(ctx context.Context, uid uuid.UUID) error {
	res, err := t.client.Exec(ctx, deleteTickets, uid)
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("ticket not found")
	}

	return nil
}

var getTicketFullInfo = `
SELECT
    t.flight_number,
    t.departure,
    t.destination,
    t.departure_date,
    t.arrival_date,
    json_agg(
            json_build_object(
                    'uuid', p.uuid,
                    'first_name', p.first_name,
                    'last_name', p.last_name,
                    'middle_name', p.middle_name,
                    'documents', p.documents
            )
    ) AS passengers
FROM tickets t
         JOIN passengers_with_docs p ON t.passenger_uuid = p.uuid
WHERE t.flight_number = (SELECT flight_number FROM tickets WHERE uuid = $1)
GROUP BY
    t.flight_number,
    t.departure,
    t.destination,
    t.departure_date,
    t.arrival_date;
`

func (t ticket) GetTicketFullInfo(ctx context.Context, uid uuid.UUID) (ticketsFull *models.TicketsFullInfo, err error) {
	var response []*models.TicketsFullInfo
	err = pgxscan.Select(ctx, t.client, &response, getTicketFullInfo, uid)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if len(response) == 0 {
		return nil, nil
	}

	return response[0], nil
}

var getPassengers = `
SELECT p.* FROM tickets AS t
LEFT JOIN passengers_with_docs as p ON t.passenger_uuid = p.uuid
WHERE t.flight_number = (SELECT flight_number FROM tickets WHERE uuid = $1);
`

func (t ticket) GetPassengers(ctx context.Context, uid uuid.UUID) (passengers []*models.PassengerWithDocs, err error) {
	err = pgxscan.Select(ctx, t.client, &passengers, getPassengers, uid)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	return
}
