package controllers

import (
	"sample-project/app/models/entities"
	"sample-project/app/models/providers/book_provider"

	"github.com/revel/revel"
)

// CBook
type CBook struct {
	*revel.Controller
	provider *book_provider.PBook
}

// Before интерцептор контроллера CBook
func (c *CBook) Before() (result revel.Result, rc CBook) {
	return
}

// GetBookByID получение книги по id
func (c *CBook) GetBookByID() revel.Result {
	return nil
}

// GetBooks получение всех книг
func (c *CBook) GetBooks() revel.Result {
	return nil
}

// CreateBook создание книги
func (c *CBook) CreateBook() revel.Result {
	return nil
}

// UpdateBook изменение книги
func (c *CBook) UpdateBook() revel.Result {
	return nil
}

// DeleteBook удаление книги
func (c *CBook) DeleteBook() revel.Result {
	return nil
}

// GiveBook выдача книги
func (c *CBook) GiveBook() revel.Result {
	return nil
}

// fetchPostBook метод получения сущности из post параметров
func (c *CBook) fetchPostBook() (bs entities.Book, err error) {
	return
}
