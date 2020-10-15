package controllers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"sample-project/app/helpers"
	"sample-project/app/models/entities"
	"sample-project/app/models/providers/event_provider"

	"github.com/revel/revel"
)

// CEvent
type CEvent struct {
	*revel.Controller
	provider *event_provider.PEvent
}

// Init интерцептор контроллера CEvent
func (c *CEvent) Init() revel.Result {
	var (
		connector helpers.IDBConnector // экземпляр коннектора, для получения экземпляра соединения с бд
		db        *sql.DB              // экземпляр соединения с бд
		err       error                // ошибка в ходе выполнения функции
	)

	// получение экземпляра соединения с бд
	connector, err = helpers.GetConnector()
	if err != nil {
		revel.AppLog.Errorf("CEvent.Init : connector.GetDBConnection, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}
	db, err = connector.GetDBConnection()
	if err != nil {
		revel.AppLog.Errorf("CEvent.Init : connector.GetDBConnection, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// инициализация провайдера
	c.provider = new(event_provider.PEvent)
	err = c.provider.Init(db)
	if err != nil {
		revel.AppLog.Errorf("CEvent.Init : c.provider.Init, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	return nil
}

// Destroy контроллера CEvent
func (c *CEvent) Destroy() {
	c.Controller.Destroy()

	// удаление ссылки на провайдер
	c.provider = nil
}

// GetAll получение всех событий
func (c *CEvent) GetAll() revel.Result {
	// получение событий
	events, err := c.provider.GetEvents()
	if err != nil {
		revel.AppLog.Errorf("CEvent.GetAll : c.provider.GetEvents, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}
	revel.AppLog.Debugf("CEvent.GetAll : c.provider.GetEvents, events: %+v\n", events)

	// рендер положительного результата
	return c.RenderJSON(Succes(events))
}

// GetByBookID получение событий по id книги
func (c *CEvent) GetByBookID(id int64) revel.Result {
	// получение событий по книге
	events, err := c.provider.GetEventByBookID(id)
	if err != nil {
		revel.AppLog.Errorf("CEvent.GetByBookID : c.provider.GetEventByBookID, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}
	revel.AppLog.Debugf("CEvent.GetByBookID : c.provider.GetEventByBookID, events: %+v\n", events)

	// рендер положительного результата
	return c.RenderJSON(Succes(events))
}

// GetByEmployeeID получение событий по id сотрудникаы
func (c *CEvent) GetByEmployeeID(id int64) revel.Result {
	// получение событий по сотруднику
	events, err := c.provider.GetEventByEmployeeID(id)
	if err != nil {
		revel.AppLog.Errorf("CEvent.GetByEmployeeID : c.provider.GetEventByEmployeeID, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}
	revel.AppLog.Debugf("CEvent.GetByEmployeeID : c.provider.GetEventByEmployeeID, events: %+v\n", events)

	// рендер положительного результата
	return c.RenderJSON(Succes(events))
}

// Create создание события
func (c *CEvent) Create() revel.Result {
	var (
		event *entities.Event // экземпляр сущности для создания
		err   error           // ошибка в ходе выполнения функции
	)

	// формирование сущности для создания из post параметров
	event, err = c.fetchPostEvent()
	if err != nil {
		revel.AppLog.Errorf("CEvent.Create : c.fetchPostEvent, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// создание сущности
	event, err = c.provider.Create(event)
	if err != nil {
		revel.AppLog.Errorf("CEvent.Create : c.provider.Create, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}
	revel.AppLog.Debugf("CEvent.Create : c.provider.Create, event: %+v\n", event)

	// рендер положительного результата
	return c.RenderJSON(Succes(event))
}

// fetchPostEvent метод получения сущности из post параметров
func (c *CEvent) fetchPostEvent() (e *entities.Event, err error) {
	var (
		rawRequest []byte // байтовое представление тела запроса
	)

	// получение тела запроса
	rawRequest, err = ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		revel.AppLog.Errorf("CEvent.fetchPostEmployee : ioutil.ReadAll, %s\n", err)
		return
	}

	// преобразование тела запроса в структуру сущности
	err = json.Unmarshal(rawRequest, &e)
	if err != nil {
		revel.AppLog.Errorf("CEvent.fetchPostEmployee : json.Unmarshal, %s\n", err)
		return
	}

	revel.AppLog.Debugf("CEvent.fetchPostEvent, event: %+v\n", e)

	return
}
