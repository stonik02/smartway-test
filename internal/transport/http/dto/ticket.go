package dto

import (
	"github.com/google/uuid"
	"time"
)

type CreateTicketRequest struct {
	PassengerUUID uuid.UUID `json:"passenger_uuid"`
	Departure     string    `json:"departure"`
	DepartureDate time.Time `json:"departure_date"`
	Destination   string    `json:"destination"`
	ArrivalDate   time.Time `json:"arrival_date"`
	OrderNumber   string    `json:"order_number"`
	Provider      string    `json:"provider"`
	BookingDate   string    `json:"booking_date"`
	FlightNumber  string    `json:"flight_number"`
}
