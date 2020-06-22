package gateway

import (
	"jiroshin/clean-architecture-golang/entity"
	"jiroshin/clean-architecture-golang/usecase"

	"database/sql"

	_ "github.com/lib/pq"
)

var _ usecase.BookRepository = &Postgres{}

type Postgres struct {
	DB *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{
		DB: db,
	}
}

func (gw *Postgres) StoreBook(book *entity.Book) error {
	const query = `		
		INSERT INTO book (title, author, overview)
		VALUES ($1, $2, $3)
	`

	if _, err := gw.DB.Exec(query, book.Title, book.Author, book.Overview); err != nil {
		return err
	}
	return nil
}
