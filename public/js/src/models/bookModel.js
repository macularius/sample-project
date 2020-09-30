import Model from "../../helpers/model.js";

/// BookModel объект для работы(CRUD) с данными
class BookModel extends Model {
    constructor() {}

    // поолучение всех книг
    getBooks() {
        return new Promise(function(resolve, reject) {

        })
    }

    // поолучение книги по ее ID
    getBookByID(id) {
        return new Promise(function(resolve, reject) {

        })
    }

    // создание книги
    createBook(book) {
        return new Promise(function(resolve, reject) {

        })
    }

    // удаление книги
    deleteBook(book) {
        return new Promise(function(resolve, reject) {

        })
    }
}
const bookModel  = new BookModel();
export default bookModel