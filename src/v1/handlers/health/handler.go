package health

import (
	"github.com/expandr/expandr/internal/server"
	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct{}

func NewHealthHandle() HealthHandler {
	return HealthHandler{}
}

func (hh HealthHandler) Routes(dispatcher server.HandlerDispatcher) {
	g := dispatcher.Public.Group("/health")
	g.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON("i'm alive")
	})
}
