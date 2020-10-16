package controllers

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"sample-project/app/helpers"
	"sample-project/app/models/providers/user_provider"
	"strings"

	"github.com/revel/revel"
)

// CAuth
type CAuth struct {
	*revel.Controller
	provider    *user_provider.PUser
	actualNonce map[string]string // [:Session ID:]:nonce:
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
	// Проверка авторизованности
	authorize := c.Session.GetDefault("authorize", nil, false)
	if authorize.(bool) {
		return c.Redirect((*CIndex).Index)
	}

	if c.Request.GetHttpHeader("Authorization") != "" {
		username, realmVal, nonceVal, digestURIVal, responseVal := helpers.GetDigestHeaders(c.Request.GetHttpHeader("Authorization"))
		method := c.Request.Method
		var password string = "123"

		if c.actualNonce[c.Session.ID()] == nonceVal {
			ha1 := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s:%s:%s", username, realmVal, password))))
			ha2 := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s:%s", method, digestURIVal))))
			serverResp := fmt.Sprintf("%x", md5.Sum([]byte(strings.Join([]string{ha1, nonceVal, ha2}, ":"))))

			if serverResp == responseVal {
				c.Session.Set("authorize", true)
				return c.Redirect((*CIndex).Index)
			}

		}
	}

	c.Response.Status = 401
	nonce, digestString, err := helpers.GetDigestString("users@library", c.ClientIP)
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	c.Response.Out.Header().Add("WWW-Authenticate", digestString)
	if c.actualNonce == nil {
		c.actualNonce = make(map[string]string, 0)
	}
	c.actualNonce[c.Session.ID()] = nonce

	return c.Redirect((*CAuth).Login)
}

// Logout
func (c *CAuth) Logout() revel.Result {
	c.Session.Set("authorize", false)
	delete(c.actualNonce, c.Session.ID())

	return c.Redirect((*CAuth).Login)
}

// GetCurrentEmployee
func (c *CAuth) GetCurrentEmployee() revel.Result {
	return c.RenderJSON(Succes(nil))
}
