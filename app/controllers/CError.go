package controllers

import (
	"net/http"

	"github.com/revel/revel"
)

// CError
type CError struct {
	*revel.Controller
}

// Unauthorized ошибка 401
func (c *CError) Unauthorized() revel.Result {
	c.Response.Status = http.StatusUnauthorized
	c.ViewArgs["Title"] = "Unauthorized"
	c.ViewArgs["Description"] = "The request has not been applied because it lacks valid authentication credentials for the target resource"
	return c.RenderTemplate("errors/401.html")
}
