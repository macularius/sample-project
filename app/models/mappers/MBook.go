package mappers

import (
	"database/sql"
	"sample-project/app/models/entities"
)

// MBook маппер книг
type MBook struct {
	db *sql.DB
}

// SelectAll получение всех книг
func (m *MBook) SelectAll() (bs []*entities.Book, err error) {
	return
}

// SelectByID получение книги по ID
func (m *MBook) SelectByID(id int64) (b *entities.Book, err error) {
	return
}

// Insert добавление книги
func (m *MBook) Insert(book entities.Book) (b *entities.Book, err error) {
	return
}

// Update изменение книги
func (m *MBook) Update(book entities.Book) (b *entities.Book, err error) {
	return
}

// Delete удаление книги
func (m *MBook) Delete(book entities.Book) (err error) {
	return
}
