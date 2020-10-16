package controllers

import "github.com/revel/revel"

// ServerResponse структура ответа сервера
type ServerResponse struct {
	Status       ResponseStatus `json:"status"` // статус ответа
	ErrorMessage string         `json:"error"`  // текст ошибки
	Data         interface{}    `json:"data"`   // данные
}

// ResponseStatus тип статуса ответа сервера
type ResponseStatus string

// перечисление статусов ответа сервера
const (
	RESPONSE_STATUS_SUCCESS ResponseStatus = "succes" // строковое представление статуса ответа "успешно"
	RESPONSE_STATUS_FAILED  ResponseStatus = "failed" // строковое представление статуса ответа "неудачно"
)

// Succes получение структуры ответа при успешном запросе
func Succes(data interface{}) (response ServerResponse) {
	response.Status = RESPONSE_STATUS_SUCCESS
	response.Data = data

	revel.AppLog.Debugf("Succes, response: %+v\n", response)
	revel.AppLog.Debugf("Succes, data: %+v\n", data)

	return response
}

// Failed получение структуры ответа при неудачном запросе
func Failed(err string) (response ServerResponse) {
	response.Status = RESPONSE_STATUS_FAILED
	response.ErrorMessage = err

	revel.AppLog.Debugf("Failed, response: %+v\n", response)
	revel.AppLog.Debugf("Failed, ErrorMessage: %+v\n", err)

	return response
}
