package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/gofiber/fiber/v2"
)

var (
	once   sync.Once
	server Server
)

type Server struct {
	app *fiber.App
	cfg *config
}

func NewServer(opts ...Option) Server {
	once.Do(func() {
		app := fiber.New()
		cfg := newConfig(opts...)

		server = Server{
			app: app,
			cfg: cfg,
		}
	})
	return server
}

func (s Server) RegisterVersion(version int, handlerContainer HandlerContainer) Server {

	dispatcher := HandlerDispatcher{
		Public:  s.app.Group(fmt.Sprintf("/api/v%v", version)),
		Private: s.app.Group(fmt.Sprintf("/ws/v%v", version)),
	}

	for _, r := range handlerContainer.Routes {
		r.Routes(dispatcher)
	}

	return s
}

func (s Server) Listen() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("G=gracefully shutting down...")
		_ = s.app.Shutdown()
	}()

	if err := s.app.Listen(fmt.Sprintf(":%v", s.cfg.port)); err != nil {
		log.Panic(err)
	}

	fmt.Println("Running cleanup tasks...")

}
