package mappers

import (
	"database/sql"
	"sample-project/app/models/entities"
	"time"

	"github.com/revel/revel"
)

// BookDBType тип сущности "книга" бд
type BookDBType struct {
	Pk_id        int64      // идентификатор
	Fk_status    int64      // статус книги
	C_isbn       string     // уникальный идентификатор книги
	C_name       string     // название книги
	C_author     *string    // автор
	C_publisher  *string    // издательство
	C_year       *time.Time // год изданиия
	C_is_archive int64      // признак архивности
}

// ToType функция преобразования типа бд к типу сущности
func (dbt *BookDBType) ToType() (b *entities.Book, err error) {
	b = new(entities.Book)

	b.ID = dbt.Pk_id
	b.ISBN = dbt.C_isbn
	b.Name = dbt.C_name
	b.Author = dbt.C_author
	b.Publisher = dbt.C_publisher
	b.Year = dbt.C_year

	return
}

// FromType функция преобразования типа сущности к типу бд
// допускается, что dbt is nil
func (_ *BookDBType) FromType(b entities.Book) (dbt *BookDBType, err error) {
	dbt = &BookDBType{
		Pk_id:       b.ID,
		C_isbn:      b.ISBN,
		C_name:      b.Name,
		C_author:    b.Author,
		C_publisher: b.Publisher,
		C_year:      b.Year,
	}

	return
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
		FROM "library".t_books
		WHERE c_is_archive = 0;
	`

	// выполнение запроса
	rows, err = m.db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
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
		WHERE pk_id = $1 and
			  c_is_archive = 0;
	`

	// выполнение запроса
	row = m.db.QueryRow(query, id)

	// считывание строки выборки
	err = row.Scan(&b.Pk_id, &b.Fk_status, &b.C_name, &b.C_isbn, &b.C_author, &b.C_publisher, &b.C_year, &b.C_is_archive)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}

		revel.AppLog.Errorf("MBook.SelectByID : row.Scan, %s\n", err)
		return
	}

	return
}

// Insert добавление книги
func (m *MBook) Insert(book *BookDBType) (id int64, err error) {
	var (
		query string   // строка запроса
		row   *sql.Row // выборка данных
	)

	revel.AppLog.Debugf("MBook.Insert, book: %+v\n", book)

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
			1,	-- fk_status
			$1,	-- c_name
			$2,	-- c_isbn
			$3,	-- c_author
			$4,	-- c_publisher
			$5,	-- c_year
			$6	-- c_is_archive
		)
		returning pk_id;
	`

	// выполнение запроса
	row = m.db.QueryRow(query, book.C_name, book.C_isbn, book.C_author, book.C_publisher, book.C_year, book.C_is_archive)

	// считывание id
	err = row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}

		revel.AppLog.Errorf("MBook.Insert : row.Scan, %s\n", err)
		return
	}

	return
}

// Update изменение книги
func (m *MBook) Update(book *BookDBType) (err error) {
	var (
		query string // строка запроса
	)

	revel.AppLog.Debugf("MBook.Update, book: %+v\n", book)

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
			err = nil
			return
		}

		revel.AppLog.Errorf("MBook.Update : m.db.Exec, %s\n", err)
		return
	}

	return
}

// Delete удаление книги (архивирование)
func (m *MBook) Delete(book *BookDBType) (err error) {
	var (
		query string // строка запроса
	)

	revel.AppLog.Debugf("MBook.Delete, book: %+v\n", book)

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
			err = nil
			return
		}

		revel.AppLog.Errorf("MBook.Delete : m.db.Exec, %s\n", err)
		return
	}

	return
}
