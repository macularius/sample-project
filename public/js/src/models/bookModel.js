import Model from "../../helpers/model.js";

/// BookModel объект для работы(CRUD) с данными
class BookModel extends Model {
    // поолучение всех книг
    getBooks() {
        return new Promise(function(resolve, reject) {
            resolve({ ID: 1, isbn: 'isbn1', name: 'name', author: 'author', publisher: 'publisher', year: 'year', status: 'status', },)
        })
    }

    // поолучение книги по ее ID
    getBookByID(id) {
        return new Promise(function(resolve, reject) {
            resolve({ ID: 1, isbn: 'isbn1', name: 'name', author: 'author', publisher: 'publisher', year: 'year', status: 'status', },)
        })
    }

    // создание книги
    createBook(book) {
        return new Promise(function(resolve, reject) {
            resolve({ ID: 1, isbn: 'isbn1', name: 'name', author: 'author', publisher: 'publisher', year: 'year', status: 'status', },)
        })
    }

    // изменение книги
    updateBook(book) {
        return new Promise(function(resolve, reject) {
            resolve({ ID: 1, isbn: 'isbn1', name: 'name', author: 'author', publisher: 'publisher', year: 'year', status: 'status', },)
        })
    }

    // удаление книги
    deleteBook(book) {
        return new Promise(function(resolve, reject) {
            resolve({ ID: 1, isbn: 'isbn1', name: 'name', author: 'author', publisher: 'publisher', year: 'year', status: 'status', },)
        })
    }
}
const bookModel  = new BookModel();
export default bookModel