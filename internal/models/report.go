package models

import (
	"github.com/google/uuid"
	"time"
)

type PassengerReport struct {
	PassengerUUID   uuid.UUID `json:"passenger_uuid"`
	BookingDate     time.Time `json:"booking_date"`
	DepartureDate   time.Time `json:"departure_date"`
	Departure       string    `json:"departure"`
	Destination     string    `json:"destination"`
	OrderNumber     string    `json:"order_number"`
	ServiceRendered bool      `json:"service_rendered"`
}
