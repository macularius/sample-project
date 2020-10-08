package controllers

import (
	"sample-project/app/models/providers/user_provider"

	"github.com/revel/revel"
)

// CAuth
type CAuth struct {
	*revel.Controller
	provider *user_provider.PUser
}

// Before интерцептор контроллера CAuth
func (c *CAuth) Before() (result revel.Result, rc CAuth) {
	return
}

// Login
func (c *CAuth) Login() revel.Result {
	return nil
}

// Logout
func (c *CAuth) Logout() revel.Result {
	return nil
}
