package http

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"test-task/internal/config"
	"test-task/internal/storage"
	"test-task/internal/transport/http"
)

type App struct {
	pgClient *pgxpool.Pool
	cfg      *config.Config
}

func NewApp(client *pgxpool.Pool, cfg *config.Config) *App {
	return &App{
		pgClient: client,
		cfg:      cfg,
	}
}

func (a *App) Start() {

	router := fiber.New(fiber.Config{})
	api := router.Group("/api")

	documentStorage := storage.NewDocument(a.pgClient)
	passengerStorage := storage.NewPassenger(a.pgClient)
	ticketStorage := storage.NewTicket(a.pgClient)

	documentRouter := api.Group("/document")
	passengerRouter := api.Group("/passenger")
	ticketRouter := api.Group("/ticket")

	documentHandler := http.NewDocumentHandler(documentStorage, documentRouter)
	documentHandler.Register()
	passengerHandler := http.NewPassengerHandler(passengerStorage, passengerRouter)
	passengerHandler.Register()
	ticketHandler := http.NewTicketHandler(ticketStorage, ticketRouter)
	ticketHandler.Register()

	log.Fatal(router.Listen(":" + a.cfg.App.Port))
}
