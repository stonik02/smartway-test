package models

import "github.com/google/uuid"

// Document На всякий случай сделал несколько документов на пассажира
type Document struct {
	UUID          uuid.UUID `json:"uuid"`
	Type          string    `json:"type"`
	Number        string    `json:"number"`
	PassengerUUID uuid.UUID `json:"-"`
}
