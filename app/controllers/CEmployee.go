package controllers

import (
	"sample-project/app/models/entities"
	"sample-project/app/models/providers/employee_provider"

	"github.com/revel/revel"
)

// CEmployee
type CEmployee struct {
	*revel.Controller
	provider *employee_provider.PEmployee
}

// Before интерцептор контроллера CEmployee
func (c *CEmployee) Before() (result revel.Result, rc CEmployee) {
	return
}

// GetAll получение всех сотрудников
func (c *CEmployee) GetAll() revel.Result {
	return nil
}

// GetByID получение сотрудника по id
func (c *CEmployee) GetByID() revel.Result {
	return nil
}

// Create создание сотрудника
func (c *CEmployee) Create() revel.Result {
	return nil
}

// UpdateEmployee изменение сотрудника
func (c *CEmployee) Update() revel.Result {
	return nil
}

// DeleteEmployee удаление сотрудника
func (c *CEmployee) Delete() revel.Result {
	return nil
}

// fetchPostEmployee метод получения сущности из post параметров
func (c *CEmployee) fetchPostEmployee() (es entities.Employee, err error) {
	return
}
