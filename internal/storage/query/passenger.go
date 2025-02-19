package query

import (
	"github.com/google/uuid"
)

// TODO можно сделать кастомный тип времени для маршалинга формата 2020-12-12
type GetReport struct {
	PassengerUUID uuid.UUID `json:"passenger_uuid"`
	StartDate     string    `json:"start_date"`
	EndDate       string    `json:"end_date"`
}
