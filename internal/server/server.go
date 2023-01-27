package server

import (
	"fmt"
	"hexagonal-go/internal/application"
	"hexagonal-go/internal/infraestructure"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	httpAddress     string
	engine          *fiber.App
	shutdownTimeout time.Duration

	userCreateService application.UserCreateService
}

func New(appName string, port uint32, shutdownTimeout time.Duration, userCreateService application.UserCreateService) Server {
	server := Server{
		httpAddress: fmt.Sprintf(":%d", port),
		engine: fiber.New(fiber.Config{
			AppName: appName,
		}),
		userCreateService: userCreateService,
		shutdownTimeout:   shutdownTimeout,
	}
	server.registerMiddlewares()
	server.registerRoutes()
	return server
}

func (s *Server) Run() error {

	log.Println("Server running on", s.httpAddress)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	serverShutdown := make(chan struct{})

	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = s.engine.Shutdown()
		serverShutdown <- struct{}{}
	}()

	if err := s.engine.Listen(s.httpAddress); err != nil {
		return err
	}

	<-serverShutdown

	log.Println("Running cleanup tasks...")

	return nil
}
func (s *Server) registerMiddlewares() {
	s.engine.Use(recover.New())
	s.engine.Use(logger.New())
}

func (s *Server) registerRoutes() {
	api := s.engine.Group("/api").Group("/v1")
	api.Get("/health", infraestructure.CheckHandler())
	api.Post("/user", infraestructure.UserCreateHandler(s.userCreateService))
}
