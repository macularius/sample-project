package controllers

import (
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

// Init интерцептор контроллера CBook
func (c *CBook) Init() revel.Result {
	var (
		cache helpers.ICache // экземпляр кэша
		err   error          // ошибка в ходе выполнения функции
	)

	// инициализация кэша
	cache, err = helpers.GetCache()
	if err != nil {
		revel.AppLog.Errorf("CBook.Init : helpers.GetCache, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	// получение токена клиента
	token, err := helpers.GetToken(c.Controller)
	if err != nil {
		revel.AppLog.Errorf("CAuth.Check : helpers.GetToken, %s\n", err)
		return c.Redirect((*CError).Unauthorized)
	}

	// проверка токена
	if isExist := cache.TokenIsActual(token); !isExist {
		return c.Redirect((*CError).Unauthorized)
	}

	// инициализация провайдера
	c.provider = new(book_provider.PBook)
	err = c.provider.Init()
	if err != nil {
		revel.AppLog.Errorf("CBook.Init : c.provider.Init, %s\n", err)
		return c.RenderJSON(Failed(err.Error()))
	}

	return nil
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
	revel.AppLog.Debugf("CBook.GetAll : c.provider.GetBooks, books: %+v\n", books)

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
	revel.AppLog.Debugf("CBook.GetByID : c.provider.GetBookByID, book: %+v\n", book)

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
	revel.AppLog.Debugf("CBook.Create : c.provider.CreateBook, book: %+v\n", book)

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
	revel.AppLog.Debugf("CBook.Update : c.provider.UpdateBook, book: %+v\n", book)

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
	revel.AppLog.Debugf("CBook.Delete : c.provider.DeleteBook, book: %+v\n", book)

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

	revel.AppLog.Debugf("CBook.fetchPostBook, book: %+v\n", b)

	return
}
