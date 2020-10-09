package controllers

// ServerResponse структура ответа сервера
type ServerResponse struct {
	Status       ResponseStatus
	ErrorMessage string
	Data         interface{}
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

	return response
}

// Failed получение структуры ответа при неудачном запросе
func Failed(err string) (response ServerResponse) {
	response.Status = RESPONSE_STATUS_FAILED
	response.ErrorMessage = err

	return response
}
