package entities

// User структура сущности пользователя
type User struct {
	ID       int64     // идентификатор
	Login    string    // логин
	Password string    // пароль
	Employee *Employee // сотрудник
}
