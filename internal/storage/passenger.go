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

type Passenger interface {
	Update(ctx context.Context, passenger *models.Passenger) error
	Delete(ctx context.Context, uid uuid.UUID) error
	GetReport(ctx context.Context, report *query.GetReport) ([]*models.PassengerReport, error)
}

type passenger struct {
	client *pgxpool.Pool
}

func NewPassenger(client *pgxpool.Pool) Passenger {
	return &passenger{
		client: client,
	}
}

var updatePassenger = `
UPDATE passengers
SET
    last_name = $2,
    first_name = $3,
    middle_name = $4
WHERE uuid = $1;
`

func (r passenger) Update(ctx context.Context, pass *models.Passenger) error {
	res, err := r.client.Exec(ctx, updatePassenger, pass.UUID, pass.LastName, pass.FirstName, pass.MiddleName)
	if err != nil {
		return err // Ошибка выполнения запроса
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("passenger not found") // Запись не найдена
	}

	return err
}

var deletePassenger = `DELETE FROM passengers WHERE uuid = $1;`

func (r passenger) Delete(ctx context.Context, uid uuid.UUID) error {
	res, err := r.client.Exec(ctx, deletePassenger, uid)
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("passenger not found")
	}

	return nil
}

var getReport = `
SELECT * FROM passenger_report
         WHERE (
             (booking_date < $1 AND departure_date BETWEEN $1 AND $2) OR
             (booking_date BETWEEN $1 AND $2 AND service_rendered = FALSE) OR
             (booking_date BETWEEN $1 AND $2 AND service_rendered = TRUE) 
                )AND
             passenger_uuid = $3;
`

func (r passenger) GetReport(ctx context.Context, report *query.GetReport) (passReports []*models.PassengerReport, err error) {
	err = pgxscan.Select(ctx, r.client, &passReports, getReport, report.StartDate, report.EndDate, report.PassengerUUID)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	return
}
