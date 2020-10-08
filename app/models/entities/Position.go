package entities

// Position структура сущности должности
type Position struct {
	ID   int64  `json:"ID"`   // идентификатор
	Name string `json:"name"` // название должности
}
