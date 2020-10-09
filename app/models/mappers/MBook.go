package mappers

import (
	"database/sql"
	"time"

	"github.com/revel/revel"
)

// BookDBType тип сущности "книга" бд
type BookDBType struct {
	Pk_id        int64      // идентификатор
	Fk_status    string     // статус книги
	C_isbn       int64      // уникальный идентификатор книги
	C_name       string     // название книги
	C_author     *string    // автор
	C_publisher  *string    // издательство
	C_year       *time.Time // год изданиия
	C_is_archive int64      // признак архивности
}

// MBook маппер книг
type MBook struct {
	db *sql.DB
}

// Init
func (m *MBook) Init(db *sql.DB) {
	m.db = db
}

// SelectAll получение всех книг
func (m *MBook) SelectAll() (bs []*BookDBType, err error) {
	var (
		query string    // строка запроса
		rows  *sql.Rows // выборка данных
	)

	// запрос
	query = `
		SELECT 
			pk_id,
			fk_status,
			c_name,
			c_isbn,
			c_author,
			c_publisher,
			c_year,
			c_is_archive
		FROM "library".t_books;
	`

	// выполнение запроса
	rows, err = m.db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}

		revel.AppLog.Errorf("MBook.SelectAll : m.db.query, %s\n", err)
		return
	}

	// обработка строк выборки
	for rows.Next() {
		// создание экземпляра сущности для считывания строки выборки
		b := new(BookDBType)

		// считывание строки выборки
		err = rows.Scan(&b.Pk_id, &b.Fk_status, &b.C_name, &b.C_isbn, &b.C_author, &b.C_publisher, &b.C_year, &b.C_is_archive)
		if err != nil {
			revel.AppLog.Errorf("MBook.SelectAll : rows.Scan, %s\n", err)
			continue
		}

		// добавление сущности в массив
		bs = append(bs, b)
	}

	return
}

// SelectByID получение книги по ID
func (m *MBook) SelectByID(id int64) (b *BookDBType, err error) {
	var (
		query string   // строка запроса
		row   *sql.Row // выборка данных
	)

	b = new(BookDBType)

	// запрос
	query = `
		SELECT 
			pk_id,
			fk_status,
			c_name,
			c_isbn,
			c_author,
			c_publisher,
			c_year,
			c_is_archive
		FROM "library".t_books
		WHERE pk_id = $1;
	`

	// выполнение запроса
	row = m.db.QueryRow(query, id)

	// считывание строки выборки
	err = row.Scan(&b.Pk_id, &b.Fk_status, &b.C_name, &b.C_isbn, &b.C_author, &b.C_publisher, &b.C_year, &b.C_is_archive)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}

		revel.AppLog.Errorf("MBook.SelectByID : row.Scan, %s\n", err)
		return
	}

	return
}

// Insert добавление книги
func (m *MBook) Insert(book BookDBType) (id int64, err error) {
	var (
		query string   // строка запроса
		row   *sql.Row // выборка данных
	)

	// запрос
	query = `
		INSERT INTO "library".t_books(
			fk_status,
			c_name,
			c_isbn,
			c_author,
			c_publisher,
			c_year,
			c_is_archive
		)
		values(
			$1,	-- fk_status
			$2,	-- c_name
			$3,	-- c_isbn
			$4,	-- c_author
			$5,	-- c_publisher
			$6,	-- c_year
			$7	-- c_is_archive
		)
		returning pk_id;
	`

	// выполнение запроса
	row = m.db.QueryRow(query, book.Fk_status, book.C_name, book.C_isbn, book.C_author, book.C_publisher, book.C_year, book.C_is_archive)

	// считывание id
	err = row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}

		revel.AppLog.Errorf("MBook.Insert : row.Scan, %s\n", err)
		return
	}

	return
}

// Update изменение книги
func (m *MBook) Update(book BookDBType) (err error) {
	var (
		query string // строка запроса
	)

	// запрос
	query = `
		UPDATE "library".t_books
			SET 
				fk_status = $2,
				c_name = $3,
				c_isbn = $4,
				c_author = $5,
				c_publisher = $6,
				c_year = $7
		WHERE pk_id = $1;
	`

	// выполнение запроса
	_, err = m.db.Exec(query, book.Pk_id, book.Fk_status, book.C_name, book.C_isbn, book.C_author, book.C_publisher, book.C_year)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}

		revel.AppLog.Errorf("MBook.Update : m.db.Exec, %s\n", err)
		return
	}

	return
}

// Delete удаление книги (архивирование)
func (m *MBook) Delete(book BookDBType) (err error) {
	var (
		query string // строка запроса
	)

	// запрос
	query = `
		UPDATE "library".t_books
		SET c_is_archive = 1
		WHERE pk_id = $1;
	`

	// выполнение запроса
	_, err = m.db.Exec(query, book.Pk_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}

		revel.AppLog.Errorf("MBook.Delete : m.db.Exec, %s\n", err)
		return
	}

	return
}
