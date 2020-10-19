package controllers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sample-project/app/helpers"
	"sample-project/app/models/entities"
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

// Login авторизация пользователя
func (c *CAuth) Login() revel.Result {
	var (
		u   *entities.User
		err error
	)

	// Проверка авторизованности
	if token, ok := helpers.ActualToken[c.Session.ID()]; ok {
		c.SetCookie(&http.Cookie{Name: "auth-token", Value: token, Domain: c.Request.Host, Path: "/"})
		return c.RenderJSON(Succes("Пользователь уже авторизован"))
	}

	// получение пользователя и зpost параметров
	u, err = c.fetchUserData()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}

	// проверка логина и пароля
	if u.Login == "admin" && u.Password == "123" {
		// создание токена
		helpers.ActualToken[c.Session.ID()] = "yep" // TODO
		// установка токена в cookies клиента
		c.SetCookie(&http.Cookie{Name: "auth-token", Value: "yep", Domain: c.Request.Host, Path: "/"}) // TODO
	}

	return c.RenderJSON(Succes("Пользователь авторизован"))
}

// Logout
func (c *CAuth) Logout() revel.Result {
	// удаление токена
	delete(helpers.ActualToken, c.Session.ID())

	return c.Redirect((*CIndex).Index)
}

// Check проверка авторизованности пользователя
func (c *CAuth) Check() revel.Result {
	// получение токена клиента для пользователя
	userToken, err := c.Request.Cookie("auth-token")
	if err != nil {
		revel.AppLog.Errorf("CAuth.Check : c.fetchToken, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// получение токена сервера для пользователя
	token, ok := helpers.ActualToken[c.Session.ID()]
	if !ok {
		return c.RenderJSON(Succes(false))
	}

	// проверка соответствия токена пользователя сервера и клиента
	if token == userToken.GetValue() {
		return c.RenderJSON(Succes(true))
	}

	return c.RenderJSON(Succes(false))
}

// GetCurrentEmployee
func (c *CAuth) GetCurrentEmployee() revel.Result {
	return c.RenderJSON(Succes(nil))
}

// fetchUserData метод получения сущности из post параметров
func (c *CAuth) fetchUserData() (u *entities.User, err error) {
	var (
		rawRequest []byte // байтовое представление тела запроса
	)

	// получение тела запроса
	rawRequest, err = ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		revel.AppLog.Errorf("CAuth.fetchUserData : ioutil.ReadAll, %s\n", err)
		return
	}

	// преобразование тела запроса в структуру сущности
	err = json.Unmarshal(rawRequest, &u)
	if err != nil {
		revel.AppLog.Errorf("CAuth.fetchUserData : json.Unmarshal, %s\n", err)
		return
	}

	revel.AppLog.Debugf("CAuth.fetchUserData, user: %+v\n", u)

	return
}

// fetchToken метод получения token'а из post параметров
func (c *CAuth) fetchToken() (token string, err error) {
	var (
		rawRequest []byte // байтовое представление тела запроса
	)

	// получение тела запроса
	rawRequest, err = ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		revel.AppLog.Errorf("CAuth.fetchToken : ioutil.ReadAll, %s\n", err)
		return
	}

	// преобразование тела запроса в структуру сущности
	err = json.Unmarshal(rawRequest, &token)
	if err != nil {
		revel.AppLog.Errorf("CAuth.fetchToken : json.Unmarshal, %s\n", err)
		return
	}

	revel.AppLog.Debugf("CAuth.fetchToken, token: %+v\n", token)

	return
}
