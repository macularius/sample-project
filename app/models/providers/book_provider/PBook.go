package book_provider

import (
	"database/sql"
	"sample-project/app/models/entities"
	"sample-project/app/models/mappers"

	"github.com/revel/revel"
)

// PBook провайдер контроллера книг
type PBook struct {
	bookMapper       *mappers.MBook
	bookStatusMapper *mappers.MBookStatus
}

// Init
func (p *PBook) Init(db *sql.DB) (err error) {
	// инициализация маппера книг
	p.bookMapper = new(mappers.MBook)
	p.bookMapper.Init(db)

	// инициализация маппера статусов книг
	p.bookStatusMapper = new(mappers.MBookStatus)
	p.bookStatusMapper.Init(db)

	return
}

// GetBookByID метод получения книги по id
func (p *PBook) GetBookByID(id int64) (b *entities.Book, err error) {
	var (
		bdbt *mappers.BookDBType
	)

	// получение данных книги
	bdbt, err = p.bookMapper.SelectByID(id)
	if err != nil {
		revel.AppLog.Errorf("PBook.GetBookByID : p.bookMapper.SelectByID, %s\n", err)
		return
	}

	// преобразование к типу сущности
	b = bdbt.ToType()

	// получение значения статуса по ключу
	b.Status, err = p.bookStatusMapper.StatusByID(bdbt.Fk_status)
	if err != nil {
		revel.AppLog.Errorf("PBook.GetBookByID : p.bookStatusMapper.StatusByID, %s\n", err)
		return
	}

	return
}

// GetBooks метод получения книг
func (p *PBook) GetBooks() (bs []*entities.Book, err error) {
	var (
		bdbts []*mappers.BookDBType
	)

	// получение данных книг
	bdbts, err = p.bookMapper.SelectAll()
	if err != nil {
		revel.AppLog.Errorf("PBook.GetBooks : p.bookMapper.SelectAll, %s\n", err)
		return
	}

	for _, bdbt := range bdbts {
		// преобразование к типу сущности
		b := bdbt.ToType()

		// получение значения статуса по ключу
		b.Status, err = p.bookStatusMapper.StatusByID(bdbt.Fk_status)
		if err != nil {
			revel.AppLog.Errorf("PBook.GetBooks : p.bookStatusMapper.StatusByID, %s\n", err)
			return
		}

		bs = append(bs, b)
	}

	return
}

// CreateBook метод создания книги
func (p *PBook) CreateBook(book *entities.Book) (b *entities.Book, err error) {
	var (
		bdbt *mappers.BookDBType
	)

	// формирование структуры бд
	bdbt = &mappers.BookDBType{
		C_isbn:      book.ISBN,
		C_name:      book.Name,
		C_author:    book.Author,
		C_publisher: book.Publisher,
		C_year:      book.Year,
	}

	// получение внешнего ключа на статус
	bdbt.Fk_status, err = p.bookStatusMapper.IDByStatus(book.Status)
	if err != nil {
		revel.AppLog.Errorf("PBook.CreateBook : p.bookStatusMapper.IDByStatus, %s\n", err)
		return
	}

	// добавление книги
	book.ID, err = p.bookMapper.Insert(bdbt)
	if err != nil {
		revel.AppLog.Errorf("PBook.CreateBook : p.bookMapper.Insert, %s\n", err)
		return
	}

	return book, nil
}

// UpdateBook метод обновления книги
func (p *PBook) UpdateBook(book *entities.Book) (b *entities.Book, err error) {
	var (
		bdbt *mappers.BookDBType
	)

	// формирование структуры бд
	bdbt = &mappers.BookDBType{
		C_isbn:      book.ISBN,
		C_name:      book.Name,
		C_author:    book.Author,
		C_publisher: book.Publisher,
		C_year:      book.Year,
	}

	// получение внешнего ключа на статус
	bdbt.Fk_status, err = p.bookStatusMapper.IDByStatus(book.Status)
	if err != nil {
		revel.AppLog.Errorf("PBook.UpdateBook : p.bookStatusMapper.IDByStatus, %s\n", err)
		return
	}

	// добавление книги
	err = p.bookMapper.Update(bdbt)
	if err != nil {
		revel.AppLog.Errorf("PBook.UpdateBook : p.bookMapper.Update, %s\n", err)
		return
	}

	return book, nil
}

// DeleteBook метод удаления книги
func (p *PBook) DeleteBook(book *entities.Book) (err error) {
	var (
		bdbt *mappers.BookDBType
	)

	// формирование структуры бд
	bdbt = &mappers.BookDBType{
		C_isbn:      book.ISBN,
		C_name:      book.Name,
		C_author:    book.Author,
		C_publisher: book.Publisher,
		C_year:      book.Year,
	}

	// получение внешнего ключа на статус
	bdbt.Fk_status, err = p.bookStatusMapper.IDByStatus(book.Status)
	if err != nil {
		revel.AppLog.Errorf("PBook.DeleteBook : p.bookStatusMapper.IDByStatus, %s\n", err)
		return
	}

	// добавление книги
	err = p.bookMapper.Delete(bdbt)
	if err != nil {
		revel.AppLog.Errorf("PBook.DeleteBook : p.bookMapper.Delete, %s\n", err)
		return
	}

	return
}
