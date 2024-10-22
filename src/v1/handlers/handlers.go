package handlers

import (
	"github.com/expandr/expandr/internal/server"
	"github.com/expandr/expandr/src/v1/handlers/health"
)

func NewHandlers() server.HandlerContainer {
	return server.HandlerContainer{
		Routes: []server.Handler{
			health.NewHealthHandle(),
		},
	}
}
