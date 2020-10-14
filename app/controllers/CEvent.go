package controllers

import (
	"sample-project/app/models/entities"
	"sample-project/app/models/providers/employee_provider"

	"github.com/revel/revel"
)

// CEvent
type CEvent struct {
	*revel.Controller
	provider *employee_provider.PEmployee
}

// Before интерцептор контроллера CEvent
func (c *CEvent) Before() (result revel.Result, rc CEvent) {
	return
}

// GetAll получение всех событий
func (c *CEvent) GetAll() revel.Result {
	return nil
}

// GetByBookID получение событий по id книги
func (c *CEvent) GetByBookID() revel.Result {
	return nil
}

// GetByEmployeeID получение событий по id сотрудникаы
func (c *CEvent) GetByEmployeeID() revel.Result {
	return nil
}

// Create создание события
func (c *CEvent) Create() revel.Result {
	return nil
}

// fetchPostEvent метод получения сущности из post параметров
func (c *CEvent) fetchPostEvent() (es entities.Event, err error) {
	return
}
