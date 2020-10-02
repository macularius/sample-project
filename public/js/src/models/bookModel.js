import Model from "../../helpers/model.js";
import { Book } from "./entities/book.js";

/// BookModel объект для работы(CRUD) с данными
class BookModel extends Model {
    constructor() {
        super()

        this.data = new Map();
        this.data.set(1, new Book(1, 'isbn1', 'name1', 'author1', 'publisher1', '2020', 'status'));
        this.data.set(2, new Book(2, 'isbn2', 'name2', 'author2', 'publisher2', '2020', 'status'));
        this.data.set(3, new Book(3, 'isbn3', 'name3', 'author3', 'publisher3', '2020', 'status'));
        this.data.set(4, new Book(4, 'isbn4', 'name4', 'author4', 'publisher4', '2020', 'status'));
    }

    // получение всех книг
    getBooks() {
        return new Promise((resolve, reject) => {
            let books = []

            for (let book of this.data.values()) {
                books.push(book)
            }

            resolve(books)
        })
    }

    // получение книги по ее ID
    getBookByID(id) {
        return new Promise((resolve, reject) => {
            resolve(this.data.get(id))
        })
    }

    // создание книги
    createBook(book) {
        return new Promise((resolve, reject) => {
            let id

            for (let key of this.data.keys()) {
                id = key
            }
            id++

            book.id = id
            this.data.set(id, book)
            resolve(this.data.get(book.id))
        })
    }

    // изменение книги
    updateBook(book) {
        return new Promise((resolve, reject) => {
            this.data.set(book.id, book)
            resolve(this.data.get(book.id))
        })
    }

    // удаление книги
    deleteBook(book) {
        return new Promise((resolve, reject) => {
            this.data.delete(book.id)
            resolve()
        })
    }

    // выдача книги
    giveBook(bookID, employeeID) {
        return new Promise((resolve, reject) => {
            resolve()
        })
    }
}
const bookModel = new BookModel();
export default bookModel