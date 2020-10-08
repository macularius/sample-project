package book_provider

import (
	"sample-project/app/models/entities"
	"sample-project/app/models/mappers"
)

// PBook провайдер контроллера книг
type PBook struct {
	bookMapper *mappers.MBook
}

// GetBookByID метод получения книги по id
func (p *PBook) GetBookByID(id int64) (b *entities.Book, err error) {
	return
}

// GetBooks метод получения книг
func (p *PBook) GetBooks() (bs []*entities.Book, err error) {
	return
}

// CreateBook метод создания книги
func (p *PBook) CreateBook(book *entities.Book) (b *entities.Book, err error) {
	return
}

// UpdateBook метод обновления книги
func (p *PBook) UpdateBook(book *entities.Book) (b *entities.Book, err error) {
	return
}

// DeleteBook метод удаления книги
func (p *PBook) DeleteBook(book *entities.Book) (err error) {
	return
}
