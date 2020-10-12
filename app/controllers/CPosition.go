package controllers

import (
	"sample-project/app/models/providers/position_provider"

	"github.com/revel/revel"
)

// CPosition контроллер запросов для сущности "Должность"
type CPosition struct {
	*revel.Controller
	provider *position_provider.PPosition
}

// Before интерцептор контроллера CPosition
func (c *CPosition) Before() (result revel.Result, rc CPosition) {
	return
}

// GetEmployees получение всех должностей
func (c *CPosition) GetEmployees() revel.Result {
	return nil
}
