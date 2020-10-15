package mappers

import (
	"database/sql"

	"github.com/revel/revel"
)

// MBookStatus маппер стутусов книг
type MBookStatus struct {
	db *sql.DB
}

// Init
func (m *MBookStatus) Init(db *sql.DB) {
	m.db = db
}

// StatusByID получение статуса книги по id
func (m *MBookStatus) StatusByID(id int64) (status string, err error) {
	var (
		query string   // строка запроса
		row   *sql.Row // выборка данных
	)

	// запрос
	query = `
		SELECT
			c_value
		FROM "library".ref_statuses
		WHERE pk_id = $1;
	`

	// выполнение запроса
	row = m.db.QueryRow(query, id)

	// считывание строки выборки
	err = row.Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}

		revel.AppLog.Errorf("MBookStatus.StatusByID : row.Scan, %s\n", err)
		return
	}

	return
}

// IDByStatus получение id статуса по значению
func (m *MBookStatus) IDByStatus(status string) (id int64, err error) {
	var (
		query string   // строка запроса
		row   *sql.Row // выборка данных
	)

	revel.AppLog.Debugf("MBookStatus.IDByStatus, status: %+v\n", status)

	// запрос
	query = `
		SELECT
			pk_id
		FROM "library".ref_statuses
		WHERE c_value = $1;
	`

	// выполнение запроса
	row = m.db.QueryRow(query, status)

	// считывание строки выборки
	err = row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}

		revel.AppLog.Errorf("MBookStatus.IDByStatus : row.Scan, %s\n", err)
		return
	}

	return
}
