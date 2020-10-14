package controllers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"sample-project/app/helpers"
	"sample-project/app/models/entities"
	"sample-project/app/models/providers/book_provider"

	"github.com/revel/revel"
)

// CBook
type CBook struct {
	*revel.Controller
	provider *book_provider.PBook
}

// Before интерцептор контроллера CBook
func (c *CBook) Before() (result revel.Result, rc CBook) {
	var (
		connector helpers.IDBConnector // экземпляр коннектора, для получения экземпляра соединения с бд
		db        *sql.DB              // экземпляр соединения с бд
		err       error                // ошибка в ходе выполнения функции
	)

	// получение экземпляра соединения с бд
	connector, err = helpers.GetConnector()
	if err != nil {
		revel.AppLog.Errorf("CBook.Before : connector.GetDBConnection, %s\n", err)
		return
	}
	db, err = connector.GetDBConnection()
	if err != nil {
		revel.AppLog.Errorf("CBook.Before : connector.GetDBConnection, %s\n", err)
		return
	}

	// инициализация провайдера
	c.provider = new(book_provider.PBook)
	err = c.provider.Init(db)
	if err != nil {
		revel.AppLog.Errorf("CBook.Before : c.provider.Init, %s\n", err)
		return
	}

	return
}

// Destroy контроллера CBook
func (c *CBook) Destroy() {
	c.Controller.Destroy()

	// удаление ссылки на провайдер
	c.provider = nil
}

// GetAll получение всех книг
func (c *CBook) GetAll() revel.Result {
	// получение книг
	books, err := c.provider.GetBooks()
	if err != nil {
		revel.AppLog.Errorf("CBook.GetAll : c.provider.GetBooks, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// рендер положительного результата
	return c.RenderJSON(Succes(books))
}

// GetByID получение книги по id
func (c *CBook) GetByID(id int64) revel.Result {
	// получение книг
	book, err := c.provider.GetBookByID(id)
	if err != nil {
		revel.AppLog.Errorf("CBook.GetByID : c.provider.GetBookByID, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// рендер положительного результата
	return c.RenderJSON(Succes(book))
}

// Create создание книги
func (c *CBook) Create() revel.Result {
	var (
		book *entities.Book // экземпляр сущности для создания
		err  error          // ошибка в ходе выполнения функции
	)

	// формирование сущности для создания из post параметров
	book, err = c.fetchPostBook()
	if err != nil {
		revel.AppLog.Errorf("CBook.Create : c.fetchPostBook, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// создание сущности
	book, err = c.provider.CreateBook(book)
	if err != nil {
		revel.AppLog.Errorf("CBook.Create : c.provider.CreateBook, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// рендер положительного результата
	return c.RenderJSON(Succes(book))
}

// Update обновление книги
func (c *CBook) Update() revel.Result {
	var (
		book *entities.Book // экземпляр сущности для обновления
		err  error          // ошибка в ходе выполнения функции
	)

	// формирование сущности для обновления из post параметров
	book, err = c.fetchPostBook()
	if err != nil {
		revel.AppLog.Errorf("CBook.Create : c.fetchPostBook, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// обновление сущности
	book, err = c.provider.UpdateBook(book)
	if err != nil {
		revel.AppLog.Errorf("CBook.Create : c.provider.UpdateBook, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// рендер положительного результата
	return c.RenderJSON(Succes(book))
}

// Delete удаление книги
func (c *CBook) Delete() revel.Result {
	var (
		book *entities.Book // экземпляр сущности для удаления
		err  error          // ошибка в ходе выполнения функции
	)

	// формирование сущности для удаления из post параметров
	book, err = c.fetchPostBook()
	if err != nil {
		revel.AppLog.Errorf("CBook.Create : c.fetchPostBook, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// удаление сущности
	err = c.provider.DeleteBook(book)
	if err != nil {
		revel.AppLog.Errorf("CBook.Create : c.provider.DeleteBook, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// рендер положительного результата
	return c.RenderJSON(Succes(nil))
}

// fetchPostBook метод получения сущности из post параметров
func (c *CBook) fetchPostBook() (b *entities.Book, err error) {
	var (
		rawRequest []byte // байтовое представление тела запроса
	)

	// получение тела запроса
	rawRequest, err = ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		revel.AppLog.Errorf("CBook.fetchPostBook : ioutil.ReadAll, %s\n", err)
		return
	}

	// преобразование тела запроса в структуру сущности
	err = json.Unmarshal(rawRequest, &b)
	if err != nil {
		revel.AppLog.Errorf("CBook.fetchPostBook : json.Unmarshal, %s\n", err)
		return
	}

	return
}
