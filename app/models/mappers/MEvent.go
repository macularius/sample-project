package mappers

import (
	"database/sql"
	"sample-project/app/models/entities"
)

// MEvent маппер событий
type MEvent struct {
	db *sql.DB
}

// SelectAll получение всех событий
func (m *MEvent) SelectAll() (bs []*entities.Event, err error) {
	return
}

// SelectByBookID получение события по ID книги
func (m *MEvent) SelectByBookID(id int64) (b *entities.Event, err error) {
	return
}

// SelectByEmployeeID получение события по ID сотрудника
func (m *MEvent) SelectByEmployeeID(id int64) (b *entities.Event, err error) {
	return
}

// Insert добавление события
func (m *MEvent) Insert(event entities.Event) (b *entities.Event, err error) {
	return
}
