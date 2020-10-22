package employee_provider

import (
	"database/sql"
	"sample-project/app/helpers"
	"sample-project/app/models/entities"
	"sample-project/app/models/mappers"

	"github.com/revel/revel"
)

// PEmployee провайдер контроллера сотрудников
type PEmployee struct {
	employeeMapper    *mappers.MEmployee
	positionsMapper   *mappers.MPosition
	bookStatusMapper  *mappers.MBookStatus
	libraryCardMapper *mappers.MLibraryCard
}

// Init
func (p *PEmployee) Init() (err error) {
	var db *sql.DB // экземпляр подключения к бд

	// получение экземпляра подключения к бд
	db, err = helpers.GetDBConnection()
	if err != nil {
		revel.AppLog.Errorf("PEmployee.Init : helpers.GetDBConnection, %s\n", err)
		return err
	}

	// инициализация маппера сотрудников
	p.employeeMapper = new(mappers.MEmployee)
	p.employeeMapper.Init(db)

	// инициализация маппера должностей
	p.positionsMapper = new(mappers.MPosition)
	p.positionsMapper.Init(db)

	// инициализация маппера статусов книг
	p.bookStatusMapper = new(mappers.MBookStatus)
	p.bookStatusMapper.Init(db)

	// инициализация маппера читательского билета
	p.libraryCardMapper = new(mappers.MLibraryCard)
	p.libraryCardMapper.Init(db)

	return
}

// GetEmployeeByID метод получения сотрудника по id
func (p *PEmployee) GetEmployeeByID(id int64) (e *entities.Employee, err error) {
	var (
		edbt *mappers.EmployeeDBType
	)

	// получение данных сотрудника
	edbt, err = p.employeeMapper.SelectByID(id)
	if err != nil {
		revel.AppLog.Errorf("PEmployee.GetEmployeeByID : p.employeeMapper.SelectByID, %s\n", err)
		return
	}

	// преобразование типа бд к типу сущности
	e, err = edbt.ToType()
	if err != nil {
		revel.AppLog.Errorf("PEmployee.GetEmployeeByID : edbt.ToType, %s\n", err)
		return
	}

	// получение значения должности по ключу
	e.Position, err = p.positionsMapper.PositionNameByID(edbt.Fk_position)
	if err != nil {
		revel.AppLog.Errorf("PEmployee.GetEmployeeByID : p.positionsMapper.PositionNameByID, %s\n", err)
		return
	}

	return
}

// GetEmployees метод получения сотрудников
func (p *PEmployee) GetEmployees() (es []*entities.Employee, err error) {
	var (
		edbts []*mappers.EmployeeDBType
		e     *entities.Employee
	)

	// получение данных книг
	edbts, err = p.employeeMapper.SelectAll()
	if err != nil {
		revel.AppLog.Errorf("PEmployee.GetEmployees : p.employeeMapper.SelectAll, %s\n", err)
		return
	}

	for _, edbt := range edbts {
		// преобразование к типу сущности
		e, err = edbt.ToType()
		if err != nil {
			revel.AppLog.Errorf("PEmployee.GetEmployees : edbt.ToType, %s\n", err)
			return
		}

		// получение значения должности по ключу
		e.Position, err = p.positionsMapper.PositionNameByID(edbt.Fk_position)
		if err != nil {
			revel.AppLog.Errorf("PEmployee.GetEmployees : p.positionsMapper.PositionNameByID, %s\n", err)
			return
		}

		es = append(es, e)
	}

	return
}

// CreateEmployee метод создания сотрудника
func (p *PEmployee) CreateEmployee(employee *entities.Employee) (e *entities.Employee, err error) {
	var (
		edbt *mappers.EmployeeDBType
	)

	// инициализация структур бд из струткур сущности
	edbt, err = edbt.FromType(*employee)
	if err != nil {
		revel.AppLog.Errorf("PEmployee.CreateEmployee : edbt.FromType, %s\n", err)
		return
	}

	// получение внешнего ключа на статус
	edbt.Fk_position, err = p.positionsMapper.IDByPositionName(employee.Position)
	if err != nil {
		revel.AppLog.Errorf("PEmployee.CreateEmployee : p.bookStatusMapper.IDByPositionName, %s\n", err)
		return
	}

	// добавление сотрудника
	employee.ID, err = p.employeeMapper.Insert(edbt)
	if err != nil {
		revel.AppLog.Errorf("PEmployee.CreateEmployee : p.employeeMapper.Create, %s\n", err)
		return
	}

	return employee, nil
}

// UpdateEmployee метод обновления сотрудника
func (p *PEmployee) UpdateEmployee(employee *entities.Employee) (e *entities.Employee, err error) {
	var (
		edbt *mappers.EmployeeDBType
	)

	// инициализация структуры бд из струткуры сущности
	edbt, err = edbt.FromType(*employee)
	if err != nil {
		revel.AppLog.Errorf("PEmployee.UpdateEmployee : edbt.FromType, %s\n", err)
		return
	}

	// получение внешнего ключа на должность
	edbt.Fk_position, err = p.positionsMapper.IDByPositionName(employee.Position)
	if err != nil {
		revel.AppLog.Errorf("PEmployee.UpdateEmployee : p.positionsMapper.IDByPositionName, %s\n", err)
		return
	}

	// обновление сотрудника
	err = p.employeeMapper.Update(edbt)
	if err != nil {
		revel.AppLog.Errorf("PEmployee.UpdateEmployee : p.employeeMapper.Update, %s\n", err)
		return
	}

	return employee, nil
}

// DeleteEmployee метод удаления сотрудника
func (p *PEmployee) DeleteEmployee(employee *entities.Employee) (err error) {
	var (
		edbt *mappers.EmployeeDBType
	)

	// инициализация структуры бд из струткуры сущности
	edbt, err = edbt.FromType(*employee)
	if err != nil {
		revel.AppLog.Errorf("PEmployee.DeleteEmployee : edbt.FromType, %s\n", err)
		return
	}

	// удаление сотрудника
	err = p.employeeMapper.Delete(edbt)
	if err != nil {
		revel.AppLog.Errorf("PEmployee.DeleteEmployee : p.employeeMapper.Update, %s\n", err)
		return
	}

	return
}

// GetEmployeeByID метод получения сотрудника по id
func (p *PEmployee) GetCardBooks(id int64) (bs []*entities.Book, err error) {
	var (
		bdbts []*mappers.BookDBType
		b     *entities.Book
	)

	// получение книг читательского билета сотрудника
	bdbts, err = p.libraryCardMapper.GetLibraryCardForEmployee(id)
	if err != nil {
		revel.AppLog.Errorf("PEmployee.GetCardBooks : p.libraryCardMapper.GetCardForEmployee, %s\n", err)
		return
	}

	for _, bdbt := range bdbts {
		// преобразование к типу сущности
		b, err = bdbt.ToType()
		if err != nil {
			revel.AppLog.Errorf("PEmployee.GetCardBooks : bdbt.ToType, %s\n", err)
			return
		}

		// получение значения статуса по ключу
		b.Status, err = p.bookStatusMapper.StatusByID(bdbt.Fk_status)
		if err != nil {
			revel.AppLog.Errorf("PEmployee.GetCardBooks : p.bookStatusMapper.StatusByID, %s\n", err)
			return
		}

		bs = append(bs, b)
	}

	return
}
