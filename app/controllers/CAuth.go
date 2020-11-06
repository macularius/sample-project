package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sample-project/app/helpers"
	"sample-project/app/models/entities"
	"sample-project/app/models/providers/user_provider"

	"github.com/google/uuid"

	"github.com/revel/revel"
)

// CAuth
type CAuth struct {
	*revel.Controller
	provider *user_provider.PUser
	cache    helpers.ICache
}

// Init интерцептор контроллера CAuth
func (c *CAuth) Init() revel.Result {
	var (
		err error // ошибка в ходе выполнения функции
	)

	// инициализация провайдера
	c.provider = new(user_provider.PUser)
	err = c.provider.Init()
	if err != nil {
		revel.AppLog.Errorf("CAuth.Init : c.provider.Init, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// инициализация кэша
	c.cache, err = helpers.GetCache()
	if err != nil {
		revel.AppLog.Errorf("CAuth.Init : helpers.GetCache, %s\n", err)
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
		u       *entities.User
		isValid bool
		err     error
	)

	// получение пользователя из post параметров
	u, err = c.fetchUserData()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	// получение сотруднкиа пользователя
	err = c.provider.AttachEmployee(u)
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}

	// валиадция пользователя
	isValid, err = c.provider.Validate(u)
	if err != nil {
		revel.AppLog.Errorf("CAuth.Login : c.provider.Validate, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}
	if isValid {
		// создание токена
		token := uuid.New().String()

		// установка токена в cache сервера
		err = c.cache.Set(token, u)
		if err != nil {
			revel.AppLog.Errorf("CAuth.Login : c.cache.Set, %s\n", err)
			return c.RenderJSON(Failed(err.Error()))
		}

		// установка токена в cookies клиента
		c.SetCookie(&http.Cookie{Name: "auth-token", Value: token, Domain: c.Request.Host, Path: "/"})
	} else {
		return c.RenderJSON(Succes("Пользователь не прошел валидацию"))
	}

	return c.RenderJSON(Succes(true))
}

// Logout
func (c *CAuth) Logout() revel.Result {
	// получение токена клиента
	token, err := helpers.GetToken(c.Controller)
	if err != nil {
		revel.AppLog.Errorf("CAuth.Check : helpers.GetToken, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// удаление токена
	err = c.cache.Delete(token)
	if err != nil {
		revel.AppLog.Errorf("CAuth.Logout : c.cache.Delete, %s\n", err)
	}

	return c.Redirect((*CIndex).Index)
}

// Check проверка авторизованности пользователя
func (c *CAuth) Check() revel.Result {
	// получение токена клиента
	token, err := helpers.GetToken(c.Controller)
	if err != nil {
		revel.AppLog.Errorf("CAuth.Check : helpers.GetToken, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// проверка токена
	if isExist := c.cache.TokenIsActual(token); !isExist {
		return c.RenderJSON(Succes(false))
	}

	return c.RenderJSON(Succes(true))
}

// GetCurrentEmployee
func (c *CAuth) GetCurrentEmployee() revel.Result {
	// получение токена клиента
	token, err := helpers.GetToken(c.Controller)
	if err != nil {
		revel.AppLog.Errorf("CAuth.Check : helpers.GetToken, %s\n", err)
		return c.Redirect((*CError).Unauthorized)
	}

	// проверка токена
	if isExist := c.cache.TokenIsActual(token); !isExist {
		return c.Redirect((*CError).Unauthorized)
	}

	// получение токена сервера для пользователя
	u, err := c.cache.Get(token)
	if err != nil {
		revel.AppLog.Errorf("CAuth.Check : c.cache.Get, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	return c.RenderJSON(Succes(u.Employee))
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

	return
}
