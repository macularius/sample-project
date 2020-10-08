package entities

import "time"

// Book структура сущности книги
type Book struct {
	ID        int64     `json:"ID"`        // идентификатор
	ISBN      int64     `json:"ISBN"`      // уникальный идентификатор книги
	Name      string    `json:"name"`      // название книги
	Author    string    `json:"author"`    // автор
	Publisher string    `json:"publisher"` // издательство
	Year      time.Time `json:"year"`      // год изданиия
	Status    string    `json:"status"`    // статус книги
}
