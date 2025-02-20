package storage

import (
	"context"
	"errors"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"test-task/internal/models"
)

type Document interface {
	Create(ctx context.Context, doc *models.Document) error
	GetByPassenger(ctx context.Context, passenger uuid.UUID) ([]*models.Document, error)
	Update(ctx context.Context, document *models.Document) error
	Delete(ctx context.Context, uid uuid.UUID) error
}

type document struct {
	client *pgxpool.Pool
}

func NewDocument(client *pgxpool.Pool) Document {
	return &document{
		client: client,
	}
}

var createDocument = `
INSERT INTO documents (uuid, passenger_uuid, type, number)
VALUES ($1, $2, $3, $4);
`

func (r document) Create(ctx context.Context, doc *models.Document) error {
	if doc.UUID == uuid.Nil {
		doc.UUID = uuid.New()
	}
	_, err := r.client.Exec(ctx, createDocument, doc.UUID, doc.PassengerUUID, doc.Type, doc.Number)
	return err
}

var getByPassenger = `
	SELECT * FROM documents
	WHERE passenger_uuid = $1
`

func (r document) GetByPassenger(ctx context.Context, passengerUUID uuid.UUID) (documents []*models.Document, err error) {
	err = pgxscan.Select(ctx, r.client, &documents, getByPassenger, passengerUUID)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	return
}

var updateDocument = `
UPDATE documents
SET
    type = $2,
    number = $3
WHERE uuid = $1;
`

func (r document) Update(ctx context.Context, document *models.Document) error {
	res, err := r.client.Exec(ctx, updateDocument, document.UUID, document.Type, document.Number)
	if err != nil {
		return err // Ошибка выполнения запроса
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("document not found") // Запись не найдена
	}

	return err
}

var deleteDocument = `
DELETE FROM documents
WHERE uuid = $1;
`

func (r document) Delete(ctx context.Context, uid uuid.UUID) error {
	res, err := r.client.Exec(ctx, deleteDocument, uid)
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("document not found")
	}
	return nil
}
