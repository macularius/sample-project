// Book класс для представления сущности книги
export class Book {
    constructor(id, isbn, name, author, publisher, year, status) {
        this.ID = id
        this.ISBN = isbn
        this.name = name
        this.author = author
        this.publisher = publisher
        this.year = year
        this.status = status
    }
}

// допустимые статусы книги
export let BOOK_STATUS = {
    available: 'доступна',
    notAvailable: 'не доступна',
}