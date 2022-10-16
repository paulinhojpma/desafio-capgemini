package web

import (
	"desafio-api/sequence-validator/core"

	mux "github.com/gorilla/mux"
	// newrelic "github.com/newrelic/go-agent"
)

var (
	ExcludeRoutes []string
)

// Router ...
func Router(h *Handler) *mux.Router {
	ExcludeRoutes = make([]string, 0)
	router := core.Router()
	router.HandleFunc("/sequence", h.ValidateSequence).Methods("POST")
	router.HandleFunc("/stats", h.GetInfoSequences).Methods("GET")

	excludeRoutes()

	return router
}

func excludeRoutes() {

}

// RouterTest ...
func RouterTest(h *Handler) *mux.Router {
	router := core.Router()
	router.HandleFunc("/sequence", h.ValidateSequence).Methods("POST")
	router.HandleFunc("/stats", h.GetInfoSequences).Methods("GET")

	return router
}
