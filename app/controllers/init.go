package controllers

import (
	"github.com/revel/revel"
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	// Инициализация интерцепторов
	revel.InterceptMethod((*CEmployee).Init, revel.BEFORE)
	revel.InterceptMethod((*CBook).Init, revel.BEFORE)
	revel.InterceptMethod((*CEvent).Init, revel.BEFORE)
	revel.InterceptMethod((*CPosition).Init, revel.BEFORE)
	revel.InterceptMethod((*CAuth).Init, revel.BEFORE)
}
