package http

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"test-task/internal/models"
	"test-task/internal/storage"
	"test-task/internal/storage/query"
	"test-task/internal/transport"
	"test-task/internal/transport/http/dto"
)

type passengerHandler struct {
	storage storage.Passenger
	router  fiber.Router
}

func NewPassengerHandler(storage storage.Passenger, router fiber.Router) transport.Handler {
	return &passengerHandler{
		storage: storage,
		router:  router,
	}
}

func (h passengerHandler) Register() {
	h.router.Put("/", h.Update)
	h.router.Delete("/", h.Delete)
	h.router.Post("/report", h.GetReport)
}

func (h passengerHandler) Update(ctx fiber.Ctx) error {
	var body *models.Passenger
	err := json.Unmarshal(ctx.Body(), &body)
	if err != nil {
		log.Errorf("Failed to unmarshal passenger: %v", err)
		return fiber.ErrBadRequest
	}

	err = h.storage.Update(ctx.Context(), body)
	if err != nil {
		log.Errorf("Failed to update Passenger: %v", err)
		return err
	}

	ctx.Status(fiber.StatusOK)
	return ctx.JSON(body)
}

func (h passengerHandler) Delete(ctx fiber.Ctx) error {
	var body *dto.UUID
	err := json.Unmarshal(ctx.Body(), &body)
	if err != nil {
		log.Errorf("Failed to unmarshal uuid: %v", err)
		return fiber.ErrBadRequest
	}

	err = h.storage.Delete(ctx.Context(), body.UUID)
	if err != nil {
		log.Errorf("Failed to delete Passenger: %v", err)
		return err
	}

	ctx.Status(fiber.StatusOK)
	return nil
}

func (h passengerHandler) GetReport(ctx fiber.Ctx) error {
	var body *query.GetReport
	err := json.Unmarshal(ctx.Body(), &body)
	if err != nil {
		log.Errorf("Failed to unmarshal GetReport: %v", err)
		return fiber.ErrBadRequest
	}

	resp, err := h.storage.GetReport(ctx.Context(), body)
	if err != nil {
		log.Errorf("Failed to GetReport() Passenger: %v", err)
		return err
	}

	ctx.Status(fiber.StatusOK)
	return ctx.JSON(resp)
}
