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

// GetEmployeeByID получение сотрудника по id
func (c *CEmployee) GetEmployeeByID() revel.Result {
	return nil
}

// GetEmployees получение всех сотрудников
func (c *CEmployee) GetEmployees() revel.Result {
	return nil
}

// CreateEmployee создание сотрудника
func (c *CEmployee) CreateEmployee() revel.Result {
	return nil
}

// UpdateEmployee изменение сотрудника
func (c *CEmployee) UpdateEmployee() revel.Result {
	return nil
}

// DeleteEmployee удаление сотрудника
func (c *CEmployee) DeleteEmployee() revel.Result {
	return nil
}

// GetPositions получение должностей
func (c *CEmployee) GetPositions() revel.Result {
	return nil
}

// fetchPostEmployee метод получения сущности из post параметров
func (c *CEmployee) fetchPostEmployee() (es entities.Employee, err error) {
	return
}
