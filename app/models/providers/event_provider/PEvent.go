package event_provider

import (
	"database/sql"
	"sample-project/app/models/entities"
	"sample-project/app/models/mappers"
)

// PEvent провайдер контроллера событий
type PEvent struct {
	eventMapper    *mappers.MEvent
	employeeMapper *mappers.MEmployee
	bookMapper     *mappers.MBook
}

// Init
func (p *PEvent) Init(db *sql.DB) (err error) {
	// инициализация маппера событий
	p.eventMapper = new(mappers.MEvent)
	p.eventMapper.Init(db)

	// инициализация маппера сотрудников
	p.employeeMapper = new(mappers.MEmployee)
	p.employeeMapper.Init(db)

	// инициализация маппера книг
	p.bookMapper = new(mappers.MBook)
	p.bookMapper.Init(db)

	return
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
