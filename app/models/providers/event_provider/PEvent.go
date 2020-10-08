package event_provider

import (
	"sample-project/app/models/entities"
	"sample-project/app/models/mappers"
)

// PEvent провайдер контроллера событий
type PEvent struct {
	eventMapper    *mappers.MEvent
	employeeMapper *mappers.MBook
	bookMapper     *mappers.MEmployee
}

// GetEventByBookID метод получения события по id книги
func (p *PEvent) GetEventByBookID(id int64) (b *entities.Event, err error) {
	return
}

// GetEventByEmployeeID метод получения события по id сотрудника
func (p *PEvent) GetEventByEmployeeID(id int64) (b *entities.Event, err error) {
	return
}

// GetEvents метод получения событий
func (p *PEvent) GetEvents() (bs []*entities.Event, err error) {
	return
}
