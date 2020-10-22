package book_provider

import (
	"database/sql"
	"sample-project/app/helpers"
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
func (p *PBook) Init() (err error) {
	var db *sql.DB // экземпляр подключения к бд

	// получение экземпляра подключения к бд
	db, err = helpers.GetDBConnection()
	if err != nil {
		revel.AppLog.Errorf("PBook.Init : helpers.GetDBConnection, %s\n", err)
		return err
	}

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
	b, err = bdbt.ToType()
	if err != nil {
		revel.AppLog.Errorf("PBook.GetBookByID : bdbt.ToType, %s\n", err)
		return
	}

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
		b     *entities.Book
	)

	// получение данных книг
	bdbts, err = p.bookMapper.SelectAll()
	if err != nil {
		revel.AppLog.Errorf("PBook.GetBooks : p.bookMapper.SelectAll, %s\n", err)
		return
	}

	for _, bdbt := range bdbts {
		// преобразование к типу сущности
		b, err = bdbt.ToType()
		if err != nil {
			revel.AppLog.Errorf("PBook.GetBooks : bdbt.ToType, %s\n", err)
			return
		}

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

	// инициализация структуры бд из струткуры сущности
	bdbt, err = bdbt.FromType(*book)
	if err != nil {
		revel.AppLog.Errorf("PBook.CreateBook : bdbt.FromType, %s\n", err)
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

	// инициализация структуры бд из струткуры сущности
	bdbt, err = bdbt.FromType(*book)
	if err != nil {
		revel.AppLog.Errorf("PBook.UpdateBook : bdbt.FromType, %s\n", err)
		return
	}

	// получение внешнего ключа на статус
	bdbt.Fk_status, err = p.bookStatusMapper.IDByStatus(book.Status)
	if err != nil {
		revel.AppLog.Errorf("PBook.UpdateBook : p.bookStatusMapper.IDByStatus, %s\n", err)
		return
	}

	// обновление книги
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

	// инициализация структуры бд из струткуры сущности
	bdbt, err = bdbt.FromType(*book)
	if err != nil {
		revel.AppLog.Errorf("PBook.DeleteBook : bdbt.FromType, %s\n", err)
		return
	}

	// удаление книги
	err = p.bookMapper.Delete(bdbt)
	if err != nil {
		revel.AppLog.Errorf("PBook.DeleteBook : p.bookMapper.Delete, %s\n", err)
		return
	}

	return
}
