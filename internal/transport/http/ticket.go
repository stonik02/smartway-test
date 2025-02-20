package http

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/google/uuid"
	"test-task/internal/models"
	"test-task/internal/storage"
	"test-task/internal/storage/query"
	"test-task/internal/transport"
	"test-task/internal/transport/http/dto"
	"time"
)

type ticketHandler struct {
	storage storage.Ticket
	router  fiber.Router
}

func NewTicketHandler(storage storage.Ticket, router fiber.Router) transport.Handler {
	return &ticketHandler{
		storage: storage,
		router:  router,
	}
}

func (h ticketHandler) Register() {
	h.router.Post("/create", h.Create)
	h.router.Post("/all", h.Get)
	h.router.Put("/", h.Update)
	h.router.Delete("/", h.Delete)
	h.router.Post("/full-info", h.GetTicketFullInfo)
	h.router.Post("/passengers", h.GetPassengers)
}

func (h ticketHandler) Create(ctx fiber.Ctx) error {
	var body *dto.CreateTicketRequest
	err := json.Unmarshal(ctx.Body(), &body)
	if err != nil {
		log.Errorf("Failed to unmarshal CreateTicketRequest: %v", err)
		return fiber.ErrBadRequest
	}

	date, err := time.Parse(time.DateOnly, body.BookingDate)
	if err != nil {
		log.Errorf("Failed to parse date: %v", err)
		return fiber.ErrBadRequest
	}

	ticket := &models.Ticket{
		UUID:          uuid.New(),
		PassengerUUID: body.PassengerUUID,
		Departure:     body.Departure,
		Destination:   body.Destination,
		DepartureDate: body.DepartureDate,
		ArrivalDate:   body.ArrivalDate,
		OrderNumber:   body.OrderNumber,
		Provider:      body.Provider,
		BookingDate:   date,
		FlightNumber:  body.FlightNumber,
	}

	err = h.storage.Create(ctx.Context(), ticket)
	if err != nil {
		log.Errorf("Failed to create Ticket: %v", err)
		return err
	}

	ctx.Status(fiber.StatusOK)
	return ctx.JSON(ticket)
}

func (h ticketHandler) Get(ctx fiber.Ctx) error {
	var body = &query.Limit{}
	err := json.Unmarshal(ctx.Body(), &body)
	if err != nil {
		body.Page = 0
		body.Size = 10
	}

	resp, err := h.storage.Get(ctx.Context(), body)
	if err != nil {
		log.Errorf("Failed getting ticket: %v", err)
		return err
	}
	return ctx.JSON(resp)
}

func (h ticketHandler) Update(ctx fiber.Ctx) error {
	var body *models.Ticket
	err := json.Unmarshal(ctx.Body(), &body)
	if err != nil {
		log.Errorf("Failed to unmarshal ticket: %v", err)
		return fiber.ErrBadRequest
	}

	err = h.storage.Update(ctx.Context(), body)
	if err != nil {
		log.Errorf("Failed to update Ticket: %v", err)
		return err
	}

	ctx.Status(fiber.StatusOK)
	return ctx.JSON(body)
}

func (h ticketHandler) Delete(ctx fiber.Ctx) error {
	var body *dto.UUID
	err := json.Unmarshal(ctx.Body(), &body)
	if err != nil {
		log.Errorf("Failed to unmarshal uuid: %v", err)
		return fiber.ErrBadRequest
	}

	err = h.storage.Delete(ctx.Context(), body.UUID)
	if err != nil {
		log.Errorf("Failed to delete Ticket: %v", err)
		return err
	}

	ctx.Status(fiber.StatusOK)
	return nil
}

func (h ticketHandler) GetTicketFullInfo(ctx fiber.Ctx) error {
	var body *dto.UUID
	err := json.Unmarshal(ctx.Body(), &body)
	if err != nil {
		log.Errorf("Failed to unmarshal uuid: %v", err)
		return fiber.ErrBadRequest
	}

	resp, err := h.storage.GetTicketFullInfo(ctx.Context(), body.UUID)
	if err != nil {
		log.Errorf("Failed to GetTicketFullInfo(): %v", err)
		return err
	}

	ctx.Status(fiber.StatusOK)
	return ctx.JSON(resp)
}

func (h ticketHandler) GetPassengers(ctx fiber.Ctx) error {
	var body *dto.UUID
	err := json.Unmarshal(ctx.Body(), &body)
	if err != nil {
		log.Errorf("Failed to unmarshal uuid: %v", err)
		return fiber.ErrBadRequest
	}

	resp, err := h.storage.GetPassengers(ctx.Context(), body.UUID)
	if err != nil {
		log.Errorf("Failed to GetPassengers(): %v", err)
		return err
	}

	ctx.Status(fiber.StatusOK)
	return ctx.JSON(resp)
}
