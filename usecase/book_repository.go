package usecase

import "jiroshin/clean-architecture-golang/entity"

type BookRepository interface {
	StoreBook(book *entity.Book) error
}
