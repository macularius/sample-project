package entities

import "time"

// Book структура сущности книги
type Book struct {
	ID        int64      `json:"ID"`        // идентификатор
	ISBN      string     `json:"ISBN"`      // уникальный идентификатор книги
	Name      string     `json:"name"`      // название книги
	Author    *string    `json:"author"`    // автор
	Publisher *string    `json:"publisher"` // издательство
	Year      *time.Time `json:"year"`      // год изданиия
	Status    BookStatus `json:"status"`    // статус книги
}

// BookStatus статус книги
type BookStatus string

const (
	BOOK_STATUS_AVAILABLE     BookStatus = "доступна"
	BOOK_STATUS_NOT_AVAILABLE BookStatus = "не доступна"
)
