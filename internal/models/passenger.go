package models

import "github.com/google/uuid"

type Passenger struct {
	UUID       uuid.UUID `json:"uuid"`
	LastName   string    `json:"last_name"`
	FirstName  string    `json:"first_name"`
	MiddleName string    `json:"middle_name"`
}

// PassengerWithDocs Используется для вывода пассажиров с документами
type PassengerWithDocs struct {
	Passenger
	Documents []Document `json:"documents"`
}
