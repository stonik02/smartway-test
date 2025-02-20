package http

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/google/uuid"
	"test-task/internal/models"
	"test-task/internal/storage"
	"test-task/internal/transport"
	"test-task/internal/transport/http/dto"
)

type documentHandler struct {
	storage storage.Document
	router  fiber.Router
}

func NewDocumentHandler(storage storage.Document, router fiber.Router) transport.Handler {
	return &documentHandler{
		storage: storage,
		router:  router,
	}
}

func (h documentHandler) Register() {
	h.router.Post("/create", h.Create)
	h.router.Put("/", h.Update)
	h.router.Delete("/", h.Delete)
	h.router.Post("/get-by-passenger", h.GetByPassenger)
}

func (h documentHandler) Create(ctx fiber.Ctx) error {
	var body *dto.CreateDocumentRequest
	err := json.Unmarshal(ctx.Body(), &body)
	if err != nil {
		log.Errorf("Failed to unmarshal CreateDocumentRequest: %v", err)
		return fiber.ErrBadRequest
	}

	doc := &models.Document{
		UUID:          uuid.New(),
		Type:          body.Type,
		Number:        body.Number,
		PassengerUUID: body.PassengerUUID,
	}

	err = h.storage.Create(ctx.Context(), doc)
	if err != nil {
		log.Errorf("Failed to create Document: %v", err)
		return err
	}

	ctx.Status(fiber.StatusCreated)
	return ctx.JSON(doc)
}

func (h documentHandler) Update(ctx fiber.Ctx) error {
	var body *models.Document
	err := json.Unmarshal(ctx.Body(), &body)
	if err != nil {
		log.Errorf("Failed to unmarshal document: %v", err)
		return fiber.ErrBadRequest
	}

	err = h.storage.Update(ctx.Context(), body)
	if err != nil {
		log.Errorf("Failed to update Document: %v", err)
		return err
	}

	ctx.Status(fiber.StatusOK)
	return ctx.JSON(body)
}

func (h documentHandler) Delete(ctx fiber.Ctx) error {
	var body *dto.UUID
	err := json.Unmarshal(ctx.Body(), &body)
	if err != nil {
		log.Errorf("Failed to unmarshal uuid: %v", err)
		return fiber.ErrBadRequest
	}

	err = h.storage.Delete(ctx.Context(), body.UUID)
	if err != nil {
		log.Errorf("Failed to delete Document: %v", err)
		return err
	}

	ctx.Status(fiber.StatusOK)
	return nil
}

func (h documentHandler) GetByPassenger(ctx fiber.Ctx) error {
	var body *dto.UUID
	err := json.Unmarshal(ctx.Body(), &body)
	if err != nil {
		log.Errorf("Failed to unmarshal uuid: %v", err)
		return fiber.ErrBadRequest
	}

	resp, err := h.storage.GetByPassenger(ctx.Context(), body.UUID)
	if err != nil {
		log.Errorf("Failed to GetByPassenger: %v", err)
		return err
	}

	ctx.Status(fiber.StatusOK)
	return ctx.JSON(resp)
}
