package controllers

import (
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

// GetEventByBookID получение сотрудника по id
func (c *CEvent) GetEventByBookID() revel.Result {
	return nil
}

// GetEventByEmployeeID получение сотрудника по id
func (c *CEvent) GetEventByEmployeeID() revel.Result {
	return nil
}

// GetEvents получение всех сотрудников
func (c *CEvent) GetEvents() revel.Result {
	return nil
}
