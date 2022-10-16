package core

import (
	"net/http"

	mux "github.com/gorilla/mux"
)

//Router ...
func Router() *mux.Router {
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(handleNotFound)

	return router
}
