package entities

// Employee структура сущности сотрудника
type Employee struct {
	ID          int64  `json:"ID"`           // идентификатор
	Lastname    string `json:"lastname"`     // фамилия
	Firstname   string `json:"firstname"`    // имя
	Middlename  string `json:"middlename"`   // отчество
	Position    string `json:"position"`     // должность
	PhoneNumber string `json:"phone_number"` // телефонный номер
	Email       string `json:"email"`        // почтовый адрес
}
