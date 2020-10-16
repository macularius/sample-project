package mappers

import (
	"database/sql"

	"github.com/revel/revel"
)

// MLibraryCard маппер книг
type MLibraryCard struct {
	db *sql.DB
}

// Init
func (m *MLibraryCard) Init(db *sql.DB) {
	m.db = db
}

// GetLibraryCardForEmployee получение всех книг
func (m *MLibraryCard) GetLibraryCardForEmployee(id int64) (bs []*BookDBType, err error) {
	var (
		query string    // строка запроса
		rows  *sql.Rows // выборка данных
	)

	// запрос
	query = `
		SELECT 
			b.pk_id,
			b.fk_status,
			b.c_name,
			b.c_isbn,
			b.c_author,
			b.c_publisher,
			b.c_year,
			b.c_is_archive
		FROM "library".t_books b, "library".toc_books_employees toc 
		WHERE b.c_is_archive = 0 and 
			b.pk_id  = toc.fk_book and
			toc.fk_employee = $1;
	`

	// выполнение запроса
	rows, err = m.db.Query(query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}

		revel.AppLog.Errorf("MLibraryCard.GetLibraryCardForEmployee : m.db.query, %s\n", err)
		return
	}

	// обработка строк выборки
	for rows.Next() {
		// создание экземпляра сущности для считывания строки выборки
		b := new(BookDBType)

		// считывание строки выборки
		err = rows.Scan(&b.Pk_id, &b.Fk_status, &b.C_name, &b.C_isbn, &b.C_author, &b.C_publisher, &b.C_year, &b.C_is_archive)
		if err != nil {
			revel.AppLog.Errorf("MLibraryCard.GetLibraryCardForEmployee : rows.Scan, %s\n", err)
			continue
		}

		// добавление сущности в массив
		bs = append(bs, b)
	}

	return
}

// AddToCard добавление книги в читательский билет
func (m *MLibraryCard) AddToCard(bookID int64, employeeID int64) (err error) {
	var (
		query string // строка запроса
	)

	// запрос
	query = `
		INSERT INTO "library".toc_books_employees
			(fk_book, fk_employee)
		VALUES
			($1, $2);	
	`

	// выполнение запроса
	_, err = m.db.Exec(query, bookID, employeeID)
	if err != nil {
		revel.AppLog.Errorf("MLibraryCard.AddToCard : row.Scan, %s\n", err)
		return
	}

	return
}

// RemoveFromCard удаление книги из читательского билета
func (m *MLibraryCard) RemoveFromCard(bookID int64, employeeID int64) (err error) {
	var (
		query string // строка запроса
	)

	// запрос
	query = `
		DELETE FROM "library".toc_books_employees
		WHERE fk_book = $1 and
			fk_employee = $2;
	`

	// выполнение запроса
	_, err = m.db.Exec(query, bookID, employeeID)
	if err != nil {
		revel.AppLog.Errorf("MLibraryCard.RemoveFromCard : row.Scan, %s\n", err)
		return
	}

	return
}
