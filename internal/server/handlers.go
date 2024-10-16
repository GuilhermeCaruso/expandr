package server

import "github.com/gofiber/fiber/v2"

type Handler interface {
	Routes(dispatcher HandlerDispatcher)
}

type HandlerContainer struct {
	Routes []Handler
}

type HandlerDispatcher struct {
	Public  fiber.Router
	Private fiber.Router
}
