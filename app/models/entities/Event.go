package entities

import "time"

// Event структура сущности события
type Event struct {
	ID       int64      `json:"ID"`       // идентификатор
	Book     *Book      `json:"book"`     // книга, связанная с событием
	Employee *Employee  `json:"employee"` // сотрудник, связанный с событием
	Type     string     `json:"type"`     // тип события
	Date     *time.Time `json:"date"`     // дата события
}
