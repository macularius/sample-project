package controllers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"sample-project/app/helpers"
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
	var (
		connector helpers.IDBConnector // экземпляр коннектора, для получения экземпляра соединения с бд
		db        *sql.DB              // экземпляр соединения с бд
		err       error                // ошибка в ходе выполнения функции
	)

	// получение экземпляра соединения с бд
	connector, err = helpers.GetConnector()
	if err != nil {
		revel.AppLog.Errorf("CEmployee.Before : connector.GetDBConnection, %s\n", err)
		return
	}
	db, err = connector.GetDBConnection()
	if err != nil {
		revel.AppLog.Errorf("CEmployee.Before : connector.GetDBConnection, %s\n", err)
		return
	}

	// инициализация провайдера
	c.provider = new(employee_provider.PEmployee)
	err = c.provider.Init(db)
	if err != nil {
		revel.AppLog.Errorf("CEmployee.Before : c.provider.Init, %s\n", err)
		return
	}

	return
}

// Destroy контроллера CEmployee
func (c *CEmployee) Destroy() {
	c.Controller.Destroy()

	// удаление ссылки на провайдер
	c.provider = nil
}

// GetAll получение всех сотрудников
func (c *CEmployee) GetAll() revel.Result {
	// получение отрудников
	employees, err := c.provider.GetEmployees()
	if err != nil {
		revel.AppLog.Errorf("CEmployee.GetAll : c.provider.GetEmployees, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// рендер положительного результата
	return c.RenderJSON(Succes(employees))
}

// GetByID получение сотрудника по id
func (c *CEmployee) GetByID(id int64) revel.Result {
	// получение сотрудника
	employee, err := c.provider.GetEmployeeByID(id)
	if err != nil {
		revel.AppLog.Errorf("CEmployee.GetByID : c.provider.GetEmployeeByID, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// рендер положительного результата
	return c.RenderJSON(Succes(employee))
}

// Create создание сотрудника
func (c *CEmployee) Create() revel.Result {
	var (
		employee *entities.Employee // экземпляр сущности для создания
		err      error              // ошибка в ходе выполнения функции
	)

	// формирование сущности для создания из post параметров
	employee, err = c.fetchPostEmployee()
	if err != nil {
		revel.AppLog.Errorf("CEmployee.Create : c.fetchPostEmployee, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// создание сущности
	employee, err = c.provider.CreateEmployee(employee)
	if err != nil {
		revel.AppLog.Errorf("CEmployee.Create : c.provider.CreateEmployee, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// рендер положительного результата
	return c.RenderJSON(Succes(employee))
}

// UpdateEmployee изменение сотрудника
func (c *CEmployee) Update() revel.Result {
	var (
		employee *entities.Employee // экземпляр сущности для обновления
		err      error              // ошибка в ходе выполнения функции
	)

	// формирование сущности для обновления из post параметров
	employee, err = c.fetchPostEmployee()
	if err != nil {
		revel.AppLog.Errorf("CEmployee.Create : c.fetchPostEmployee, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// обновление сущности
	employee, err = c.provider.UpdateEmployee(employee)
	if err != nil {
		revel.AppLog.Errorf("CEmployee.Create : c.provider.UpdateEmployee, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// рендер положительного результата
	return c.RenderJSON(Succes(employee))
}

// DeleteEmployee удаление сотрудника
func (c *CEmployee) Delete() revel.Result {
	var (
		employee *entities.Employee // экземпляр сущности для удаления
		err      error              // ошибка в ходе выполнения функции
	)

	// формирование сущности для удаления из post параметров
	employee, err = c.fetchPostEmployee()
	if err != nil {
		revel.AppLog.Errorf("CEmployee.Create : c.fetchPostEmployee, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// удаление сущности
	err = c.provider.DeleteEmployee(employee)
	if err != nil {
		revel.AppLog.Errorf("CEmployee.Create : c.provider.DeleteEmployee, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// рендер положительного результата
	return c.RenderJSON(Succes(nil))
}

// fetchPostEmployee метод получения сущности из post параметров
func (c *CEmployee) fetchPostEmployee() (e *entities.Employee, err error) {
	var (
		rawRequest []byte // байтовое представление тела запроса
	)

	// получение тела запроса
	rawRequest, err = ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		revel.AppLog.Errorf("CEmployee.fetchPostEmployee : ioutil.ReadAll, %s\n", err)
		return
	}

	// преобразование тела запроса в структуру сущности
	err = json.Unmarshal(rawRequest, &e)
	if err != nil {
		revel.AppLog.Errorf("CEmployee.fetchPostEmployee : json.Unmarshal, %s\n", err)
		return
	}

	return
}
