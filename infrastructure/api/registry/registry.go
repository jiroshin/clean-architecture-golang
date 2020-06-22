package registry

import (
	"database/sql"
	"jiroshin/clean-architecture-golang/adapter/gateway"
	"jiroshin/clean-architecture-golang/usecase"
	"os"

	_ "github.com/lib/pq"
)

func NewBookUseCase(db *sql.DB) *usecase.Book {
	return usecase.NewBook(newBookRepository(db))
}

func newBookRepository(db *sql.DB) usecase.BookRepository {
	if os.Getenv("DB_TYPE") == "mysql" {
		return gateway.NewMysql(db)
	} else {
		return gateway.NewPostgres(db)
	}
}
