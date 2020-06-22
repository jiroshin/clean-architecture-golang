package usecase

import "jiroshin/clean-architecture-golang/entity"

type Book struct {
	bookRepo BookRepository
}

func NewBook(bookRepo BookRepository) *Book {
	return &Book{
		bookRepo: bookRepo,
	}
}

func (uc *Book) CreateBook(book *entity.Book) error {
	return uc.bookRepo.StoreBook(book)
}
