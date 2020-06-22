package gateway

import (
	"jiroshin/clean-architecture-golang/entity"
	"jiroshin/clean-architecture-golang/usecase"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var _ usecase.BookRepository = &Mysql{}

type Mysql struct {
	DB *sql.DB
}

func NewMysql(db *sql.DB) *Mysql {
	return &Mysql{
		DB: db,
	}
}

func (gw *Mysql) StoreBook(book *entity.Book) error {
	const query = `		
		INSERT INTO book (title, author, overview)
		VALUES (?, ?, ?)
	`

	if _, err := gw.DB.Exec(query, book.Title, book.Author, book.Overview); err != nil {
		return err
	}
	return nil
}
