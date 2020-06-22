package router

import (
	"database/sql"
	"jiroshin/clean-architecture-golang/infrastructure/api/handler"
	"jiroshin/clean-architecture-golang/infrastructure/api/registry"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func NewRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	bookUseCase := registry.NewBookUseCase(db)
	r.HandleFunc("/books", handler.CreateBook(bookUseCase)).Methods(http.MethodPost)

	return r
}
