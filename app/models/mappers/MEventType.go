package mappers

import (
	"database/sql"

	"github.com/revel/revel"
)

// MEventType маппер типов событий
type MEventType struct {
	db *sql.DB
}

// Init
func (m *MEventType) Init(db *sql.DB) {
	m.db = db
}

// EventByID получение типа события по id
func (m *MEventType) EventByID(id int64) (eventType string, err error) {
	var (
		query string   // строка запроса
		row   *sql.Row // выборка данных
	)

	revel.AppLog.Debugf("MEventType.EventByID, id: %+v\n", id)

	// запрос
	query = `
		SELECT
			c_value
		FROM "library".ref_event_types
		WHERE pk_id = $1
		ORDER BY pk_id;
	`

	// выполнение запроса
	row = m.db.QueryRow(query, id)

	// считывание строки выборки
	err = row.Scan(&eventType)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}

		revel.AppLog.Errorf("MEventType.EventByID : row.Scan, %s\n", err)
		return
	}

	return
}

// IDByEventType получение id статуса по значению
func (m *MEventType) IDByEventType(eventType string) (id int64, err error) {
	var (
		query string   // строка запроса
		row   *sql.Row // выборка данных
	)

	// запрос
	query = `
		SELECT
			pk_id
		FROM "library".ref_event_types
		WHERE c_value = $1
		ORDER BY pk_id;
	`

	// выполнение запроса
	row = m.db.QueryRow(query, eventType)

	// считывание строки выборки
	err = row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}

		revel.AppLog.Errorf("MEventType.IDByEventType : row.Scan, %s\n", err)
		return
	}

	return
}
