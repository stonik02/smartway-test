package dto

import "github.com/google/uuid"

type CreateDocumentRequest struct {
	Type          string    `json:"type"`
	Number        string    `json:"number"`
	PassengerUUID uuid.UUID `json:"passenger_uuid"`
}
