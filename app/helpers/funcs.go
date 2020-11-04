package helpers

import (
	"net/http"

	"github.com/revel/revel"
)

// GetToken функция получения токена
func GetToken(c *revel.Controller) (string, error) {
	// получение токена клиента
	token, err := c.Request.Cookie("auth-token")
	if err != nil {
		if err == http.ErrNoCookie {
			return "", nil
		}

		revel.AppLog.Errorf("helpers.GetToken : c.Request.Cookie, %s\n", err)
		return "", err
	}

	return token.GetValue(), nil
}
