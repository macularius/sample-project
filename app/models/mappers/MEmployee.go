package mappers

import (
	"database/sql"
	"sample-project/app/models/entities"
)

// MEmployee маппер сотрудников
type MEmployee struct {
	db *sql.DB
}

// SelectAll получение всех сотрудников
func (m *MEmployee) SelectAll() (bs []*entities.Employee, err error) {
	return
}

// SelectByID получение сотрудника по ID
func (m *MEmployee) SelectByID(id int64) (b *entities.Employee, err error) {
	return
}

// Insert добавление сотрудника
func (m *MEmployee) Insert(employee entities.Employee) (b *entities.Employee, err error) {
	return
}

// Update изменение сотрудника
func (m *MEmployee) Update(employee entities.Employee) (b *entities.Employee, err error) {
	return
}

// Delete удаление сотрудника
func (m *MEmployee) Delete(employee entities.Employee) (err error) {
	return
}
