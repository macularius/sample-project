package controllers

import (
	"database/sql"
	"sample-project/app/helpers"
	"sample-project/app/models/providers/user_provider"

	"github.com/revel/revel"
)

// CAuth
type CAuth struct {
	*revel.Controller
	provider *user_provider.PUser
}

// Init интерцептор контроллера CAuth
func (c *CAuth) Init() revel.Result {
	var (
		connector helpers.IDBConnector // экземпляр коннектора, для получения экземпляра соединения с бд
		db        *sql.DB              // экземпляр соединения с бд
		err       error                // ошибка в ходе выполнения функции
	)

	// получение экземпляра соединения с бд
	connector, err = helpers.GetConnector()
	if err != nil {
		revel.AppLog.Errorf("CAuth.Init : connector.GetDBConnection, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}
	db, err = connector.GetDBConnection()
	if err != nil {
		revel.AppLog.Errorf("CAuth.Init : connector.GetDBConnection, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// инициализация провайдера
	c.provider = new(user_provider.PUser)
	err = c.provider.Init(db)
	if err != nil {
		revel.AppLog.Errorf("CAuth.Init : c.provider.Init, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	return nil
}

// Destroy контроллера CAuth
func (c *CAuth) Destroy() {
	c.Controller.Destroy()

	// удаление ссылки на провайдер
	c.provider = nil
}

// Login
func (c *CAuth) Login() revel.Result {
	return nil
}

// Logout
func (c *CAuth) Logout() revel.Result {
	return nil
}

// GetCurrentEmployee
func (c *CAuth) GetCurrentEmployee() revel.Result {
	return c.RenderJSON(Succes(nil))
}
