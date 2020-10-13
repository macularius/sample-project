package mappers

import (
	"database/sql"
	"sample-project/app/models/entities"
	"time"

	"github.com/revel/revel"
)

// EventDBType тип сущности "событие" бд
type EventDBType struct {
	Pk_id       int64      // идентификатор
	Fk_book     int64      // FK на книгу
	Fk_employee int64      // FK на сотрудника
	Fk_type     int64      // FK на тип события
	C_date      *time.Time // дата события
}

// ToType функция преобразования типа бд к типу сущности
func (dbt *EventDBType) ToType() (e *entities.Event) {
	e = new(entities.Event)

	e.ID = dbt.Pk_id
	e.Date = dbt.C_date

	return
}

// MEvent маппер событий
type MEvent struct {
	db *sql.DB
}

// SelectAll получение всех событий
func (m *MEvent) SelectAll() (edbts []*EventDBType, err error) {
	var (
		query string    // строка запроса
		rows  *sql.Rows // выборка данных
	)

	// запрос
	query = `
		SELECT
			pk_id,
			fk_book,
			fk_employee,
			fk_event_type,
			c_date
		FROM "library".t_event;	
	`

	// выполнение запроса
	rows, err = m.db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}

		revel.AppLog.Errorf("MEvent.SelectAll : m.db.query, %s\n", err)
		return
	}

	// обработка строк выборки
	for rows.Next() {
		// создание экземпляра сущности для считывания строки выборки
		edbt := new(EventDBType)

		// считывание строки выборки
		err = rows.Scan(&edbt.Pk_id, &edbt.Fk_book, &edbt.Fk_employee, &edbt.Fk_type, &edbt.C_date)
		if err != nil {
			revel.AppLog.Errorf("MEvent.SelectAll : rows.Scan, %s\n", err)
			continue
		}

		// добавление сущности в массив
		edbts = append(edbts, edbt)
	}

	return
}

// SelectByBookID получение события по ID книги
func (m *MEvent) SelectByBookID(id int64) (edbts []*EventDBType, err error) {
	var (
		query string    // строка запроса
		rows  *sql.Rows // выборка данных
	)

	// запрос
	query = `
		SELECT
			pk_id,
			fk_book,
			fk_employee,
			fk_event_type,
			c_date
		FROM "library".t_event
		WHERE fk_book = $1;
	`

	// выполнение запроса
	rows, err = m.db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}

		revel.AppLog.Errorf("MEvent.SelectAll : m.db.query, %s\n", err)
		return
	}

	// обработка строк выборки
	for rows.Next() {
		// создание экземпляра сущности для считывания строки выборки
		edbt := new(EventDBType)

		// считывание строки выборки
		err = rows.Scan(&edbt.Pk_id, &edbt.Fk_book, &edbt.Fk_employee, &edbt.Fk_type, &edbt.C_date)
		if err != nil {
			revel.AppLog.Errorf("MEvent.SelectAll : rows.Scan, %s\n", err)
			continue
		}

		// добавление сущности в массив
		edbts = append(edbts, edbt)
	}

	return
}

// SelectByEmployeeID получение события по ID сотрудника
func (m *MEvent) SelectByEmployeeID(id int64) (edbts []*EventDBType, err error) {
	var (
		query string    // строка запроса
		rows  *sql.Rows // выборка данных
	)

	// запрос
	query = `
		SELECT
			pk_id,
			fk_book,
			fk_employee,
			fk_event_type,
			c_date
		FROM "library".t_event
		WHERE fk_employee = $1;
	`

	// выполнение запроса
	rows, err = m.db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}

		revel.AppLog.Errorf("MEvent.SelectAll : m.db.query, %s\n", err)
		return
	}

	// обработка строк выборки
	for rows.Next() {
		// создание экземпляра сущности для считывания строки выборки
		edbt := new(EventDBType)

		// считывание строки выборки
		err = rows.Scan(&edbt.Pk_id, &edbt.Fk_book, &edbt.Fk_employee, &edbt.Fk_type, &edbt.C_date)
		if err != nil {
			revel.AppLog.Errorf("MEvent.SelectAll : rows.Scan, %s\n", err)
			continue
		}

		// добавление сущности в массив
		edbts = append(edbts, edbt)
	}

	return
}

// Insert добавление события
func (m *MEvent) Insert(event EventDBType) (id int64, err error) {
	var (
		query string   // строка запроса
		row   *sql.Row // выборка данных
	)

	// запрос
	query = `
		INSERT INTO "library".t_event(
			fk_book,
			fk_employee,
			fk_event_type,
			c_date
		)
		VALUES(
			$1,	-- fk_book
			$2,	-- fk_employee
			$3,	-- fk_event_type
			$4	-- c_date
		)	
		returning pk_id;
	`

	// выполнение запроса
	row = m.db.QueryRow(query,
		event.Fk_book,     // pk_id
		event.Fk_employee, // fk_user
		event.Fk_type,     // fk_position
		event.C_date,      // c_firstname
	)

	// считывание id
	err = row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}

		revel.AppLog.Errorf("MEvent.Insert : row.Scan, %s\n", err)
		return
	}

	return
}
