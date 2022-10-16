package web

import (
	route "defafio-cap/sequence-validator/routes"
)

// NewRoutes ...
func NewRoutes(h *Handler) *route.Manager {
	manager := route.NewManager()

	manager.AddRoute("POST:SEQUENCE", h.CreateSequence)
	manager.AddRoute("GET:SEQUENCE", h.GetInfoSequences)

	return manager
}
