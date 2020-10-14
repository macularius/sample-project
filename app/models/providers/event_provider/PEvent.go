package event_provider

import (
	"database/sql"
	"errors"
	"sample-project/app/models/entities"
	"sample-project/app/models/mappers"
	"time"

	"github.com/revel/revel"
)

// объявление переменных пакета
var (
	ErrBookIsNotExist     = errors.New("в структуре события отсутствует книга")
	ErrEmployeeIsNotExist = errors.New("в структуре события отсутствует сотрудник")
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

// GetEventByBookID метод создания события
func (p *PEvent) Create(event *entities.Event) (e *entities.Event, err error) {
	var (
		edbt   *mappers.EventDBType
		bdbt   *mappers.BookDBType
		empdbt *mappers.EmployeeDBType
	)

	// преобразование типа сущности к типу бд
	err = edbt.FromType(event)
	if err != nil {
		revel.AppLog.Errorf("PEvent.Create : edbt.FromType, %s\n", err)
		return
	}

	// валидация книги
	if event.Book == nil {
		err = ErrBookIsNotExist
		revel.AppLog.Errorf("PEvent.Create, %s\n", err)
		return
	}
	edbt.Fk_book = event.Book.ID

	// валидация сотрудника
	if event.Employee == nil {
		err = ErrEmployeeIsNotExist
		revel.AppLog.Errorf("PEvent.Create, %s\n", err)
		return
	}
	edbt.Fk_employee = event.Employee.ID

	// фиксация времени
	t := time.Now()
	edbt.C_date = &t

	// добавление события
	edbt.Pk_id, err = p.eventMapper.Insert(edbt)
	if err != nil {
		revel.AppLog.Errorf("PEvent.Create : p.eventMapper.Insert, %s\n", err)
		return
	}

	// получение книги
	bdbt, err = p.bookMapper.SelectByID(edbt.Fk_book)
	if err != nil {
		revel.AppLog.Errorf("PEvent.Create : p.bookMapper.SelectByID, %s\n", err)
		return
	}

	// получение сотрудника
	empdbt, err = p.employeeMapper.SelectByID(edbt.Fk_employee)
	if err != nil {
		revel.AppLog.Errorf("PEvent.Create : p.employeeMapper.SelectByID, %s\n", err)
		return
	}

	// преобразование типа бд к типу сущности
	e, err = edbt.ToType()
	if err != nil {
		revel.AppLog.Errorf("PEvent.Create : edbt.ToType, %s\n", err)
		return
	}
	e.Book, err = bdbt.ToType()
	if err != nil {
		revel.AppLog.Errorf("PEvent.Create : bdbt.ToType, %s\n", err)
		return
	}
	e.Employee, err = empdbt.ToType()
	if err != nil {
		revel.AppLog.Errorf("PEvent.Create : empdbt.ToType, %s\n", err)
		return
	}

	return
}

// GetEventByBookID метод получения события по id книги
func (p *PEvent) GetEventByBookID(id int64) (es []*entities.Event, err error) {
	var (
		edbts []*mappers.EventDBType
		e     *entities.Event
	)

	// получение данных событий
	edbts, err = p.eventMapper.SelectByBookID(id)
	if err != nil {
		revel.AppLog.Errorf("PEvent.GetEvents : p.eventMapper.SelectByBookID, %s\n", err)
		return
	}

	for _, edbt := range edbts {
		// преобразование к типу сущности
		e, err = edbt.ToType()
		if err != nil {
			revel.AppLog.Errorf("PEvent.GetEvents : edbt.ToType, %s\n", err)
			return
		}

		es = append(es, e)
	}

	return
}

// GetEventByEmployeeID метод получения события по id сотрудника
func (p *PEvent) GetEventByEmployeeID(id int64) (es []*entities.Event, err error) {
	var (
		edbts []*mappers.EventDBType
		e     *entities.Event
	)

	// получение данных событий
	edbts, err = p.eventMapper.SelectByEmployeeID(id)
	if err != nil {
		revel.AppLog.Errorf("PEvent.GetEvents : p.eventMapper.SelectByEmployeeID, %s\n", err)
		return
	}

	for _, edbt := range edbts {
		// преобразование к типу сущности
		e, err = edbt.ToType()
		if err != nil {
			revel.AppLog.Errorf("PEvent.GetEvents : edbt.ToType, %s\n", err)
			return
		}

		es = append(es, e)
	}

	return
}

// GetEvents метод получения событий
func (p *PEvent) GetEvents() (es []*entities.Event, err error) {
	var (
		edbts []*mappers.EventDBType
		e     *entities.Event
	)

	// получение данных событий
	edbts, err = p.eventMapper.SelectAll()
	if err != nil {
		revel.AppLog.Errorf("PEvent.GetEvents : p.eventMapper.SelectAll, %s\n", err)
		return
	}

	for _, edbt := range edbts {
		// преобразование к типу сущности
		e, err = edbt.ToType()
		if err != nil {
			revel.AppLog.Errorf("PEvent.GetEvents : edbt.ToType, %s\n", err)
			return
		}

		es = append(es, e)
	}

	return
}
