package controllers

import (
	"database/sql"
	"sample-project/app/helpers"
	"sample-project/app/models/providers/position_provider"

	"github.com/revel/revel"
)

// CPosition контроллер запросов для сущности "Должность"
type CPosition struct {
	*revel.Controller
	provider *position_provider.PPosition
}

// Init интерцептор контроллера CPosition
func (c *CPosition) Init() revel.Result {
	var (
		connector helpers.IDBConnector // экземпляр коннектора, для получения экземпляра соединения с бд
		db        *sql.DB              // экземпляр соединения с бд
		err       error                // ошибка в ходе выполнения функции
	)

	// получение экземпляра соединения с бд
	connector, err = helpers.GetConnector()
	if err != nil {
		revel.AppLog.Errorf("CPosition.Init : connector.GetDBConnection, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}
	db, err = connector.GetDBConnection()
	if err != nil {
		revel.AppLog.Errorf("CPosition.Init : connector.GetDBConnection, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// инициализация провайдера
	c.provider = new(position_provider.PPosition)
	err = c.provider.Init(db)
	if err != nil {
		revel.AppLog.Errorf("CPosition.Init : c.provider.Init, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	return nil
}

// Destroy контроллера CPosition
func (c *CPosition) Destroy() {
	c.Controller.Destroy()

	// удаление ссылки на провайдер
	c.provider = nil
}

// GetAll получение всех должностей
func (c *CPosition) GetAll() revel.Result {
	// получение должностей
	positions, err := c.provider.GetPositions()
	if err != nil {
		revel.AppLog.Errorf("CPosition.GetAll : c.provider.GetPositions, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}
	revel.AppLog.Debugf("CPosition.GetAll : c.provider.GetPositions, positions: %v\n", positions)

	// рендер положительного результата
	return c.RenderJSON(Succes(positions))
}
