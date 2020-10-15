import Model from "../../helpers/model.js";

/// BookModel объект для работы(CRUD) с данными
class BookModel extends Model {
    constructor() {
        super()
    }

    // получение всех книг
    getBooks() {
        return this.get('/book/all')
    }

    // получение книги по ее ID
    getBookByID(id) {
        return this.get(`/book/${id}`)
    }

    // создание книги
    createBook(book) {
        // преобразование года в дату
        let date = new Date(book.year, 0)
        book.year = new Date(date.getTime() - date.getTimezoneOffset() * 60000)

        return this.post('/book/create', book)
    }

    // изменение книги
    updateBook(book) {
        // преобразование года в дату
        let date = new Date(book.year, 0)
        book.year = new Date(date.getTime() - date.getTimezoneOffset() * 60000)

        return this.post('/book/update', book)
    }

    // удаление книги
    deleteBook(book) {
        // преобразование года в дату
        let date = new Date(book.year, 0)
        book.year = new Date(date.getTime() - date.getTimezoneOffset() * 60000)
        
        return this.post('/book/delete', book)
    }
}
const bookModel = new BookModel();
export default bookModel