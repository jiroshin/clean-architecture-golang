package router

import (
	"jiroshin/clean-architecture-golang/infrastructure/server/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	attachHandlers(router)
	return router
}

func attachHandlers(mux *mux.Router) {
	mux.HandleFunc("/books", handler.CreateBook()).Methods(http.MethodPost)
}
