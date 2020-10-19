package controllers

import (
	"github.com/revel/revel"
)

// CIndex
type CIndex struct {
	*revel.Controller
}

// Init интерцептор контроллера CIndex
func (c CIndex) Init() revel.Result {
	return nil
}

// Index метод, возвращающий html
func (c CIndex) Index() revel.Result {
	return c.Render()
}
