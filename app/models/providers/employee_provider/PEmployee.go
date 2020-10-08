package employee_provider

import (
	"sample-project/app/models/entities"
	"sample-project/app/models/mappers"
)

// PEmployee провайдер контроллера сотрудников
type PEmployee struct {
	employeeMapper *mappers.MEmployee
}

// GetEmployeeByID метод получения сотрудника по id
func (p *PEmployee) GetEmployeeByID(id int64) (b *entities.Employee, err error) {
	return
}

// GetEmployees метод получения сотрудников
func (p *PEmployee) GetEmployees() (bs []*entities.Employee, err error) {
	return
}

// CreateEmployee метод создания сотрудника
func (p *PEmployee) CreateEmployee(employee *entities.Employee) (b *entities.Employee, err error) {
	return
}

// UpdateEmployee метод обновления сотрудника
func (p *PEmployee) UpdateEmployee(employee *entities.Employee) (b *entities.Employee, err error) {
	return
}

// DeleteEmployee метод удаления сотрудника
func (p *PEmployee) DeleteEmployee(employee *entities.Employee) (err error) {
	return
}

// GetPositions метод удаления сотрудника
func (p *PEmployee) GetPositions() (ps []*entities.Position, err error) {
	return
}
